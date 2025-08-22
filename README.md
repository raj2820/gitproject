# GitLab-like Tool

A minimal GitHub/GitLab-like tool built in Go, implementing core features like user authentication, repository management, and Git operations over HTTP.

## Features

### Phase 1 - Foundations ✅
- **User Authentication**: JWT-based signup/login system
- **Repository Management**: Create, list, and delete repositories
- **Git over HTTP**: Support for `git clone`, `git push`, and `git pull`
- **Database Storage**: PostgreSQL with GORM for data persistence

### Phase 2 - Collaboration Features (Planned)
- Commits & Branches management
- Branch and commit history UI
- Repository collaboration

### Phase 3 - Advanced Features (Planned)
- Merge Requests (Pull Requests)
- Issues & Discussions
- CI/CD integration

## Prerequisites

- Go 1.21 or later
- PostgreSQL database
- Git installed on the system

## Installation

1. **Clone the repository**
   ```bash
   git clone <your-repo-url>
   cd gitlab-tool
   ```

2. **Install Go dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up PostgreSQL database**
   ```bash
   # Option A: Using Docker Compose (Recommended)
   docker-compose up -d postgres
   
   # Option B: Manual PostgreSQL setup
   # Create database
   createdb gitlab_tool
   
   # Or using psql
   psql -U postgres
   CREATE DATABASE gitlab_tool;
   ```

4. **Set environment variables (optional - defaults are provided)**
   ```bash
   export DATABASE_URL="postgres://postgres:password@localhost:5432/gitlab_tool?sslmode=disable"
   export JWT_SECRET="your-secret-key-change-in-production"
   export REPOS_PATH="/tmp/repos"
   export PORT="8080"
   ```

5. **Run the application**
   ```bash
   # Option A: Using the startup script
   ./start.sh
   
   # Option B: Direct execution
   go run main.go
   ```

The server will start on `http://localhost:8080`

## Usage

### Web Interface

1. Open your browser and navigate to `http://localhost:8080`
2. Sign up with a new account
3. Create repositories
4. Use the provided clone URLs for Git operations

### API Endpoints

#### Authentication
- `POST /auth/signup` - User registration
- `POST /auth/login` - User login

#### Repository Management
- `POST /api/repos` - Create repository
- `GET /api/repos` - List user repositories
- `GET /api/repos/:id` - Get repository details
- `DELETE /api/repos/:id` - Delete repository

#### Git Operations
- `GET /git/:username/:repo/info/refs` - List references
- `POST /git/:username/:repo/git-upload-pack` - Clone/fetch
- `POST /git/:username/:repo/git-receive-pack` - Push

### Git Commands

```bash
# Clone a repository
git clone http://localhost:8080/git/username/repo-name.git

# Add remote to existing repository
git remote add origin http://localhost:8080/git/username/repo-name.git

# Push to repository
git push origin main

# Pull from repository
git pull origin main
```

## Project Structure

```
gitlab-tool/
├── main.go                 # Application entry point
├── go.mod                  # Go module file
├── internal/               # Internal packages
│   ├── auth/              # Authentication utilities
│   ├── config/            # Configuration management
│   ├── database/          # Database connection and migrations
│   ├── git/               # Git operations service
│   ├── handlers/          # HTTP request handlers
│   ├── middleware/        # HTTP middleware
│   ├── models/            # Database models
│   └── repository/        # Data access layer
├── static/                # Static web assets
│   └── index.html         # Web interface
└── README.md              # This file
```

## Configuration

The application can be configured using environment variables:

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Server port |
| `DATABASE_URL` | `postgres://postgres:your_new_password@localhost:5432/gitlab_tool?sslmode=disable` | PostgreSQL connection string |
| `JWT_SECRET` | `your-secret-key-change-in-production` | Secret key for JWT tokens |
| `REPOS_PATH` | `/tmp/repos` | Directory to store Git repositories |

## Development

### Adding New Features

1. **Models**: Add new structs in `internal/models/`
2. **Database**: Update migrations in `internal/database/`
3. **Repository**: Add data access methods in `internal/repository/`
4. **Handlers**: Implement HTTP endpoints in `internal/handlers/`
5. **Routes**: Register new endpoints in `main.go`

### Testing

```bash
# Run tests
go test ./...

# Run with coverage
go test -cover ./...
```

## Security Considerations

- **JWT Secret**: Use a strong, unique secret in production
- **Database**: Use strong passwords and consider SSL connections
- **File Permissions**: Ensure repository directories have appropriate permissions
- **Authentication**: Implement rate limiting for login attempts

## Production Deployment

1. **Environment**: Set production environment variables
2. **Database**: Use production-grade PostgreSQL instance
3. **Storage**: Consider using object storage for large repositories
4. **Reverse Proxy**: Use Nginx or similar for SSL termination
5. **Monitoring**: Add logging and health check endpoints

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is open source and available under the [MIT License](LICENSE).

## Roadmap

- [ ] Phase 2: Branches and commits management
- [ ] Phase 3: Merge requests and issues
- [ ] Phase 4: CI/CD and scaling features
- [ ] SSH support for Git operations
- [ ] Webhook system
- [ ] Advanced permissions and collaboration
- [ ] Performance optimizations with libgit2 # gitproject
