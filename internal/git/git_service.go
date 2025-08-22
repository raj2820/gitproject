package git

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Service struct {
	reposPath string
}

func NewService(reposPath string) *Service {
	return &Service{reposPath: reposPath}
}

func (s *Service) InitBareRepository(username, repoName string) error {
	repoPath := filepath.Join(s.reposPath, username, repoName+".git")

	// Create directory structure
	if err := os.MkdirAll(repoPath, 0755); err != nil {
		return fmt.Errorf("failed to create repo directory: %w", err)
	}

	// Initialize bare repository
	cmd := exec.Command("git", "init", "--bare")
	cmd.Dir = repoPath
	// Remove stdout/stderr redirection to avoid interfering with HTTP responses
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to init bare repository: %w", err)
	}

	return nil
}

func (s *Service) GetRepositoryPath(username, repoName string) string {
	return filepath.Join(s.reposPath, username, repoName+".git")
}

func (s *Service) RepositoryExists(username, repoName string) bool {
	repoPath := s.GetRepositoryPath(username, repoName)
	_, err := os.Stat(filepath.Join(repoPath, "HEAD"))
	return err == nil
}

func (s *Service) ListBranches(username, repoName string) ([]string, error) {
	repoPath := s.GetRepositoryPath(username, repoName)

	cmd := exec.Command("git", "branch", "-r")
	cmd.Dir = repoPath

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list branches: %w", err)
	}

	// Parse output and extract branch names
	// This is a simplified version - you might want to parse more carefully
	branches := []string{}
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.Contains(line, "HEAD") {
			// Remove "origin/" prefix if present
			if strings.HasPrefix(line, "origin/") {
				line = strings.TrimPrefix(line, "origin/")
			}
			branches = append(branches, line)
		}
	}

	return branches, nil
}

func (s *Service) GetLatestCommit(username, repoName, branch string) (string, error) {
	repoPath := s.GetRepositoryPath(username, repoName)

	cmd := exec.Command("git", "rev-parse", branch)
	cmd.Dir = repoPath

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get latest commit: %w", err)
	}

	return strings.TrimSpace(string(output)), nil
}
