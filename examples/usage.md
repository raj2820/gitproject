# Usage Examples

This document provides practical examples of how to use the GitLab-like tool.

## Quick Start

### 1. Start the Application

```bash
# Start PostgreSQL (if using Docker)
docker-compose up -d postgres

# Start the application
./start.sh
```

The web interface will be available at `http://localhost:8080`

### 2. Create Your First Repository

1. Open `http://localhost:8080` in your browser
2. Sign up with a new account
3. Create a repository named "my-first-repo"
4. Note the clone URL: `http://localhost:8080/git/username/my-first-repo.git`

### 3. Clone and Use the Repository

```bash
# Clone the repository
git clone http://localhost:8080/git/username/my-first-repo.git
cd my-first-repo

# Make some changes
echo "# My First Repository" > README.md
echo "This is a test repository" >> README.md

# Commit and push
git add .
git commit -m "Initial commit"
git push origin main
```

## API Examples

### Authentication

```bash
# Sign up
curl -X POST http://localhost:8080/auth/signup \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123"
  }'

# Login
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123"
  }'
```

### Repository Management

```bash
# Get the token from login response
TOKEN="your-jwt-token-here"

# Create repository
curl -X POST http://localhost:8080/api/repos \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "api-test-repo",
    "description": "Repository created via API",
    "visibility": "public"
  }'

# List repositories
curl -X GET http://localhost:8080/api/repos \
  -H "Authorization: Bearer $TOKEN"

# Get repository details
curl -X GET http://localhost:8080/api/repos/1 \
  -H "Authorization: Bearer $TOKEN"
```

## Git Operations

### Clone Repository

```bash
git clone http://localhost:8080/git/username/repo-name.git
```

### Push to Repository

```bash
cd repo-name
echo "Hello World" > hello.txt
git add hello.txt
git commit -m "Add hello.txt"
git push origin main
```

### Pull from Repository

```bash
git pull origin main
```

## Advanced Usage

### Multiple Users

1. Create multiple user accounts
2. Each user can create their own repositories
3. Users can only access their own repositories (unless public)

### Repository Visibility

- **Private**: Only the owner can access
- **Public**: Anyone can clone (but only owner can push)

### Working with Branches

```bash
# Create a new branch
git checkout -b feature-branch

# Make changes and commit
echo "New feature" > feature.txt
git add feature.txt
git commit -m "Add new feature"

# Push the new branch
git push origin feature-branch
```

## Troubleshooting

### Common Issues

1. **Database Connection Error**
   - Ensure PostgreSQL is running
   - Check DATABASE_URL environment variable
   - Verify database exists

2. **Git Operations Fail**
   - Check if repository exists
   - Verify user permissions
   - Ensure git is installed on the system

3. **Authentication Errors**
   - Check JWT token expiration
   - Verify token format (Bearer token)
   - Ensure user exists

### Debug Mode

To see detailed logs, set the log level:

```bash
export GIN_MODE=debug
./start.sh
```

## Next Steps

After mastering the basics, you can:

1. **Implement Phase 2**: Add branch and commit management
2. **Add SSH Support**: Implement SSH-based Git operations
3. **Build CI/CD**: Add continuous integration features
4. **Scale Up**: Use object storage for large repositories
5. **Add Webhooks**: Notify external services on repository events

## Contributing

Feel free to contribute by:

1. Reporting bugs
2. Suggesting new features
3. Submitting pull requests
4. Improving documentation 