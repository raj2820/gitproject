package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"gitlab-tool/internal/git"
	"gitlab-tool/internal/models"
	"gitlab-tool/internal/repository"

	"github.com/gin-gonic/gin"
)

type RepositoryHandler struct {
	repoRepo   *repository.RepositoryRepository
	gitService *git.Service
	reposPath  string
}

func NewRepositoryHandler(repoRepo *repository.RepositoryRepository, gitService *git.Service, reposPath string) *RepositoryHandler {
	return &RepositoryHandler{
		repoRepo:   repoRepo,
		gitService: gitService,
		reposPath:  reposPath,
	}
}

type CreateRepositoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Visibility  string `json:"visibility" binding:"oneof=public private"`
	// README options
	AddReadme   bool   `json:"add_readme"`
	ReadmeType  string `json:"readme_type" binding:"omitempty,oneof=markdown text"`
	ReadmeTitle string `json:"readme_title"`
	// Default branch options
	CustomBranch  bool   `json:"custom_branch"`
	DefaultBranch string `json:"default_branch"`
}

func (h *RepositoryHandler) CreateRepository(c *gin.Context) {
	var req CreateRepositoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")
	username := c.GetString("username")

	// Check if repository already exists for this user
	if _, err := h.repoRepo.FindByUsernameAndName(username, req.Name); err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Repository already exists"})
		return
	}

	// Create repository in database
	repo := &models.Repository{
		Name:        req.Name,
		Description: req.Description,
		Visibility:  req.Visibility,
		OwnerID:     userID,
	}

	if err := h.repoRepo.Create(repo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create repository"})
		return
	}

	// Initialize git repository
	if err := h.gitService.InitBareRepository(username, req.Name); err != nil {
		// Clean up database entry if git init fails
		h.repoRepo.Delete(repo.ID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize git repository"})
		return
	}

	// Set custom default branch if requested
	if req.CustomBranch && req.DefaultBranch != "" {
		if err := h.setDefaultBranch(username, req.Name, req.DefaultBranch); err != nil {
			fmt.Printf("Warning: Failed to set custom default branch: %v\n", err)
			// Don't fail the request if branch setting fails
		}
	}

	// Create README file if requested
	if req.AddReadme {
		branchName := "main"
		if req.CustomBranch && req.DefaultBranch != "" {
			branchName = req.DefaultBranch
		}
		if err := h.createReadmeFileWithBranch(username, req.Name, req.ReadmeType, req.ReadmeTitle, branchName); err != nil {
			fmt.Printf("Warning: Failed to create README file: %v\n", err)
			// Don't fail the request if README creation fails
		}
	}

	c.JSON(http.StatusCreated, repo)
}

func (h *RepositoryHandler) ListRepositories(c *gin.Context) {
	userID := c.GetUint("user_id")

	repos, err := h.repoRepo.FindByOwnerID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch repositories"})
		return
	}

	c.JSON(http.StatusOK, repos)
}

func (h *RepositoryHandler) GetRepository(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid repository ID"})
		return
	}

	repo, err := h.repoRepo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Repository not found"})
		return
	}

	// Check if user has access to this repository
	userID := c.GetUint("user_id")
	if repo.OwnerID != userID && repo.Visibility == "private" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	c.JSON(http.StatusOK, repo)
}

func (h *RepositoryHandler) DeleteRepository(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid repository ID"})
		return
	}

	repo, err := h.repoRepo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Repository not found"})
		return
	}

	// Check if user owns this repository
	userID := c.GetUint("user_id")
	if repo.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Delete from database
	if err := h.repoRepo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete repository"})
		return
	}

	// Delete git repository files
	username := c.GetString("username")
	repoPath := h.gitService.GetRepositoryPath(username, repo.Name)
	if err := exec.Command("rm", "-rf", repoPath).Run(); err != nil {
		// Log error but don't fail the request
		fmt.Printf("Warning: Failed to delete git repository files: %v\n", err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Repository deleted successfully"})
}

func (h *RepositoryHandler) CloneRepository(c *gin.Context) {
	// This would typically be handled by the git HTTP backend
	c.JSON(http.StatusOK, gin.H{"message": "Use git clone command with the repository URL"})
}

func (h *RepositoryHandler) PushToRepository(c *gin.Context) {
	// This would typically be handled by the git HTTP backend
	c.JSON(http.StatusOK, gin.H{"message": "Use git push command with the repository URL"})
}

func (h *RepositoryHandler) PullFromRepository(c *gin.Context) {
	// This would typically be handled by the git HTTP backend
	c.JSON(http.StatusOK, gin.H{"message": "Use git pull command with the repository URL"})
}

func (h *RepositoryHandler) GitHTTPBackend(c *gin.Context) {
	username := c.Param("username")
	repoName := c.Param("repo")
	action := c.Param("action")

	// Remove leading slash from action if present
	if strings.HasPrefix(action, "/") {
		action = strings.TrimPrefix(action, "/")
	}

	// Debug logging
	fmt.Printf("GitHTTPBackend called: username=%s, repo=%s, action=%s\n", username, repoName, action)

	// Remove .git suffix if present
	if strings.HasSuffix(repoName, ".git") {
		repoName = strings.TrimSuffix(repoName, ".git")
	}

	// Check if repository exists in database first
	_, err := h.repoRepo.FindByUsernameAndName(username, repoName)
	if err != nil {
		fmt.Printf("Repository not found in database: %v\n", err)
		// For Git protocol, return minimal error without HTTP headers
		c.Data(http.StatusNotFound, "text/plain", []byte("Repository not found"))
		return
	}

	// Check if repository exists on filesystem
	if !h.gitService.RepositoryExists(username, repoName) {
		fmt.Printf("Repository not found on filesystem, initializing...\n")
		// Repository doesn't exist on filesystem, try to initialize it
		if err := h.gitService.InitBareRepository(username, repoName); err != nil {
			fmt.Printf("Failed to initialize repository: %v\n", err)
			// For Git protocol, return minimal error without HTTP headers
			c.Data(http.StatusInternalServerError, "text/plain", []byte("Failed to initialize repository"))
			return
		}
		fmt.Printf("Repository initialized successfully\n")
	}

	// Ensure repository has a proper HEAD reference
	repoPath := h.gitService.GetRepositoryPath(username, repoName)
	headPath := filepath.Join(repoPath, "HEAD")
	if _, err := os.Stat(headPath); err != nil {
		fmt.Printf("HEAD file not found, creating default branch...\n")
		// Create a default branch to fix the HEAD reference
		if err := h.ensureDefaultBranch(repoPath); err != nil {
			fmt.Printf("Warning: Failed to create default branch: %v\n", err)
		}
	}

	// Set environment variables for git-http-backend
	// CRITICAL: git-http-backend expects CGI-style environment variables
	env := []string{
		fmt.Sprintf("GIT_PROJECT_ROOT=%s", h.reposPath), // Points to /tmp/repos
		fmt.Sprintf("GIT_HTTP_EXPORT_ALL=1"),
		fmt.Sprintf("REMOTE_USER=%s", username),
		fmt.Sprintf("PATH=%s", os.Getenv("PATH")),
		fmt.Sprintf("REQUEST_METHOD=%s", c.Request.Method),
		fmt.Sprintf("QUERY_STRING=%s", c.Request.URL.RawQuery),
		fmt.Sprintf("CONTENT_TYPE=%s", c.GetHeader("Content-Type")),
		fmt.Sprintf("CONTENT_LENGTH=%s", c.GetHeader("Content-Length")),
		fmt.Sprintf("HTTP_ACCEPT=%s", c.GetHeader("Accept")),
		fmt.Sprintf("HTTP_USER_AGENT=%s", c.GetHeader("User-Agent")),
		// PATH_INFO should be the repository path relative to GIT_PROJECT_ROOT
		fmt.Sprintf("PATH_INFO=/%s/%s.git", username, repoName),
		fmt.Sprintf("SCRIPT_NAME=/git"),
		// Add server info
		fmt.Sprintf("SERVER_NAME=localhost"),
		fmt.Sprintf("SERVER_PORT=8080"),
		// Add request info
		fmt.Sprintf("REQUEST_URI=%s", c.Request.URL.RequestURI()),
	}

	// Execute git-http-backend (NO EXTRA ARGUMENTS!)
	cmd := exec.Command("git", "http-backend")
	cmd.Env = env
	cmd.Dir = h.reposPath
	cmd.Stdin = c.Request.Body
	cmd.Stdout = c.Writer

	// Capture stderr for debugging
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	fmt.Printf("Executing git http-backend\n")
	fmt.Printf("Working directory: %s\n", h.reposPath)
	fmt.Printf("Environment: %v\n", env)

	if err := cmd.Run(); err != nil {
		fmt.Printf("Git operation failed: %v\n", err)
		fmt.Printf("Git stderr output: %s\n", stderr.String())
		c.Data(http.StatusInternalServerError, "text/plain", []byte(fmt.Sprintf("Git operation failed: %v", err)))
		return
	}

	fmt.Printf("Git operation completed successfully\n")
	if stderr.Len() > 0 {
		fmt.Printf("Git stderr output: %s\n", stderr.String())
	}
}

// createReadmeFile creates a README file and makes an initial commit
func (h *RepositoryHandler) createReadmeFile(username, repoName, readmeType, readmeTitle string) error {
	return h.createReadmeFileWithBranch(username, repoName, readmeType, readmeTitle, "main")
}

// createReadmeFileWithBranch creates a README file and makes an initial commit with a specific branch
func (h *RepositoryHandler) createReadmeFileWithBranch(username, repoName, readmeType, readmeTitle, branchName string) error {
	repoPath := h.gitService.GetRepositoryPath(username, repoName)

	// Set default values if not provided
	if readmeType == "" {
		readmeType = "markdown"
	}
	if readmeTitle == "" {
		readmeTitle = repoName
	}

	// Create README content based on type
	var readmeContent string
	var filename string

	switch readmeType {
	case "markdown":
		filename = "README.md"
		readmeContent = fmt.Sprintf(`# %s

%s

## Description

This repository was created using GitLab-like Tool.

## Getting Started

`+"```"+`bash
# Clone the repository
git clone http://localhost:8080/git/%s/%s.git

# Navigate to the project
cd %s

# Start developing!
`+"```"+`

## Contributing

1. Fork the repository
2. Create your feature branch (`+"`"+`git checkout -b feature/amazing-feature`+"`"+`)
3. Commit your changes (`+"`"+`git commit -m 'Add some amazing feature'`+"`"+`)
4. Push to the branch (`+"`"+`git push origin feature/amazing-feature`+"`"+`)
5. Open a Pull Request

## License

This project is open source and available under the MIT License.
`, readmeTitle, repoName, username, repoName, repoName, repoName)

	case "text":
		filename = "README.txt"
		readmeContent = fmt.Sprintf(`%s

%s

Description:
This repository was created using GitLab-like Tool.

Getting Started:
1. Clone the repository: git clone http://localhost:8080/git/%s/%s.git
2. Navigate to the project: cd %s
3. Start developing!

Contributing:
1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Open a Pull Request

License:
This project is open source and available under the MIT License.
`, readmeTitle, repoName, username, repoName, repoName)

	default:
		return fmt.Errorf("unsupported readme type: %s", readmeType)
	}

	// Create the README file
	readmePath := filepath.Join(repoPath, filename)
	if err := os.WriteFile(readmePath, []byte(readmeContent), 0644); err != nil {
		return fmt.Errorf("failed to create README file: %w", err)
	}

	// Initialize git repository and make initial commit
	if err := h.initializeGitRepositoryWithBranch(repoPath, filename, branchName); err != nil {
		return fmt.Errorf("failed to initialize git repository: %w", err)
	}

	fmt.Printf("README file created successfully: %s\n", filename)
	return nil
}

// initializeGitRepository initializes the git repository and makes the initial commit
func (h *RepositoryHandler) initializeGitRepository(repoPath, filename string) error {
	return h.initializeGitRepositoryWithBranch(repoPath, filename, "main")
}

// initializeGitRepositoryWithBranch initializes the git repository with a specific default branch
func (h *RepositoryHandler) initializeGitRepositoryWithBranch(repoPath, filename, branchName string) error {
	// For bare repositories, we need to use a different approach
	// Create a temporary directory, clone the bare repo, add README, and push back

	tempDir, err := os.MkdirTemp("", "git-temp-*")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	// Clone the bare repository to the temp directory
	cloneCmd := exec.Command("git", "clone", repoPath, tempDir)
	if err := cloneCmd.Run(); err != nil {
		// If clone fails, the repository might be empty, so we initialize it
		initCmd := exec.Command("git", "init", "-b", branchName)
		initCmd.Dir = tempDir
		if err := initCmd.Run(); err != nil {
			return fmt.Errorf("failed to initialize temp repository: %w", err)
		}

		// Add the bare repository as origin
		remoteCmd := exec.Command("git", "remote", "add", "origin", repoPath)
		remoteCmd.Dir = tempDir
		if err := remoteCmd.Run(); err != nil {
			return fmt.Errorf("failed to add remote origin: %w", err)
		}
	}

	// Change to temp directory
	originalDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}
	defer os.Chdir(originalDir)

	if err := os.Chdir(tempDir); err != nil {
		return fmt.Errorf("failed to change to temp directory: %w", err)
	}

	// Create the README file in the temp directory
	readmePath := filepath.Join(tempDir, filename)
	readmeContent, err := os.ReadFile(filepath.Join(repoPath, filename))
	if err != nil {
		return fmt.Errorf("failed to read README content from bare repo: %w", err)
	}

	if err := os.WriteFile(readmePath, readmeContent, 0644); err != nil {
		return fmt.Errorf("failed to write README file in temp dir: %w", err)
	}

	// Add the README file
	addCmd := exec.Command("git", "add", filename)
	addCmd.Dir = tempDir
	if err := addCmd.Run(); err != nil {
		return fmt.Errorf("failed to add README file: %w", err)
	}

	// Make initial commit
	commitCmd := exec.Command("git", "commit", "-m", "Initial commit: Add README file")
	commitCmd.Dir = tempDir
	commitCmd.Env = append(os.Environ(), "GIT_AUTHOR_NAME=GitLab Tool", "GIT_AUTHOR_EMAIL=system@gitlab-tool.local", "GIT_COMMITTER_NAME=GitLab Tool", "GIT_COMMITTER_EMAIL=system@gitlab-tool.local")
	if err := commitCmd.Run(); err != nil {
		return fmt.Errorf("failed to make initial commit: %w", err)
	}

	// Push to the bare repository
	pushCmd := exec.Command("git", "push", "origin", branchName)
	pushCmd.Dir = tempDir
	if err := pushCmd.Run(); err != nil {
		return fmt.Errorf("failed to push to bare repository: %w", err)
	}

	return nil
}

// setDefaultBranch sets the default branch for a bare repository
func (h *RepositoryHandler) setDefaultBranch(username, repoName, branchName string) error {
	repoPath := h.gitService.GetRepositoryPath(username, repoName)

	// For bare repositories, we need to create the branch and set it as default
	// This is done by creating a symbolic ref HEAD pointing to the branch
	cmd := exec.Command("git", "symbolic-ref", "HEAD", fmt.Sprintf("refs/heads/%s", branchName))
	cmd.Dir = repoPath
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set default branch: %w", err)
	}

	fmt.Printf("Default branch set to '%s' for repository %s/%s\n", branchName, username, repoName)
	return nil
}

// ensureDefaultBranch ensures that a default branch (e.g., 'main') exists and is the HEAD.
// This is necessary because git-http-backend expects a HEAD file pointing to a branch.
func (h *RepositoryHandler) ensureDefaultBranch(repoPath string) error {
	// Change to repository directory
	originalDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}
	defer os.Chdir(originalDir)

	if err := os.Chdir(repoPath); err != nil {
		return fmt.Errorf("failed to change to repository directory: %w", err)
	}

	// Check if 'main' branch exists
	cmd := exec.Command("git", "show-ref", "--verify", "--quiet", "refs/heads/main")
	cmd.Dir = repoPath
	if err := cmd.Run(); err != nil {
		fmt.Printf("Main branch not found, creating...\n")
		// Create the 'main' branch
		cmd := exec.Command("git", "branch", "main")
		cmd.Dir = repoPath
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to create main branch: %w", err)
		}
		fmt.Printf("Main branch created successfully\n")
	}

	// Set HEAD to point to the 'main' branch
	cmd = exec.Command("git", "symbolic-ref", "HEAD", "refs/heads/main")
	cmd.Dir = repoPath
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set HEAD to main branch: %w", err)
	}
	fmt.Printf("HEAD set to main branch successfully\n")

	return nil
}
