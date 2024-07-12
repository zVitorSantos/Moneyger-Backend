package auth

import (
    "crypto/rand"
    "encoding/base64"
    "errors"
    "time"

    "github.com/dgrijalva/jwt-go"
    "golang.org/x/crypto/argon2"
    "gorm.io/gorm"
)

var jwtKey = []byte("your_secret_key")

// JWT Structure
type Claims struct {
    UserID uint `json:"user_id"`
    jwt.StandardClaims
}

// Register new user
func RegisterUser(db *gorm.DB, user *User) error {
    salt := make([]byte, 16)
    _, err := rand.Read(salt)
    if err != nil {
        return err
    }

    hashedPassword := hashPassword(user.PasswordHash, salt)
    user.PasswordHash = base64.RawStdEncoding.EncodeToString(hashedPassword)
    user.PasswordHash = ""

    return db.Create(user).Error
}

// Authenticate a new user and returns a JWT
func AuthenticateUser(db *gorm.DB, username, password string) (*Token, error) {
    var user User
    if err := db.Where("username = ?", username).First(&user).Error; err != nil {
        return nil, err
    }

    salt, err := base64.RawStdEncoding.DecodeString(user.PasswordHash)
    if err != nil {
        return nil, err
    }

    hashedPassword := hashPassword(password, salt)
    if base64.RawStdEncoding.EncodeToString(hashedPassword) != user.PasswordHash {
        return nil, errors.New("invalid username or password")
    }

    // Generate a new JWT
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        UserID: user.ID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return nil, err
    }

    return &Token{AccessToken: tokenString}, nil
}

// Hashing the password with argon2id
func hashPassword(password string, salt []byte) []byte {
    return argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
}
