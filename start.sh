#!/bin/bash

# GitLab-like Tool Startup Script

echo "Starting GitLab-like Tool..."

# Set default environment variables
export PORT=${PORT:-8080}
export DATABASE_URL=${DATABASE_URL:-"postgres://postgres:your_new_password@localhost:5432/gitlab_tool?sslmode=disable"}
export JWT_SECRET=${JWT_SECRET:-"your-secret-key-change-in-production"}
export REPOS_PATH=${REPOS_PATH:-"/tmp/repos"}

echo "Configuration:"
echo "  Port: $PORT"
echo "  Database: $DATABASE_URL"
echo "  Repos Path: $REPOS_PATH"
echo "  JWT Secret: ${JWT_SECRET:0:10}..."

# Create repos directory if it doesn't exist
mkdir -p "$REPOS_PATH"

# Check if PostgreSQL is running
if ! pg_isready -h localhost -p 5432 > /dev/null 2>&1; then
    echo "Warning: PostgreSQL is not running on localhost:5432"
    echo "Please start PostgreSQL or update DATABASE_URL"
    echo "You can use: sudo systemctl start postgresql"
fi

# Start the application
echo "Starting server on port $PORT..."
go run main.go 