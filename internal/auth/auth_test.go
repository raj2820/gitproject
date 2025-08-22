package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestHashPassword(t *testing.T) {
	password := "testpassword123"
	hash, err := HashPassword(password)

	if err != nil {
		t.Errorf("HashPassword failed: %v", err)
	}

	if hash == password {
		t.Error("Password was not hashed")
	}

	if len(hash) == 0 {
		t.Error("Hash is empty")
	}
}

func TestCheckPassword(t *testing.T) {
	password := "testpassword123"
	hash, err := HashPassword(password)

	if err != nil {
		t.Errorf("HashPassword failed: %v", err)
	}

	if !CheckPassword(password, hash) {
		t.Error("CheckPassword failed for correct password")
	}

	if CheckPassword("wrongpassword", hash) {
		t.Error("CheckPassword succeeded for wrong password")
	}
}

func TestGenerateToken(t *testing.T) {
	userID := uint(123)
	username := "testuser"
	role := "user"
	secret := "test-secret"

	token, err := GenerateToken(userID, username, role, secret)

	if err != nil {
		t.Errorf("GenerateToken failed: %v", err)
	}

	if len(token) == 0 {
		t.Error("Generated token is empty")
	}
}

func TestValidateToken(t *testing.T) {
	userID := uint(123)
	username := "testuser"
	role := "user"
	secret := "test-secret"

	token, err := GenerateToken(userID, username, role, secret)
	if err != nil {
		t.Errorf("GenerateToken failed: %v", err)
	}

	claims, err := ValidateToken(token, secret)
	if err != nil {
		t.Errorf("ValidateToken failed: %v", err)
	}

	if claims.UserID != userID {
		t.Errorf("Expected UserID %d, got %d", userID, claims.UserID)
	}

	if claims.Username != username {
		t.Errorf("Expected Username %s, got %s", username, claims.Username)
	}

	if claims.Role != role {
		t.Errorf("Expected Role %s, got %s", role, claims.Role)
	}

	// Test with wrong secret
	_, err = ValidateToken(token, "wrong-secret")
	if err == nil {
		t.Error("ValidateToken should fail with wrong secret")
	}
}

func TestTokenExpiration(t *testing.T) {
	userID := uint(123)
	username := "testuser"
	role := "user"
	secret := "test-secret"

	// Generate token with very short expiration
	claims := Claims{
		UserID:   userID,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-1 * time.Second)), // Expired 1 second ago
			IssuedAt:  jwt.NewNumericDate(time.Now().Add(-2 * time.Second)),
			NotBefore: jwt.NewNumericDate(time.Now().Add(-2 * time.Second)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		t.Errorf("Failed to sign token: %v", err)
	}

	// Token should be expired
	_, err = ValidateToken(tokenString, secret)
	if err == nil {
		t.Error("Expired token should not be valid")
	}
}
