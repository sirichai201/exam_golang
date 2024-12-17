package services

import (
	"exam_go/config"
	"exam_go/models"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var secretKey string

func SetSecretKey() {
	secretKey = os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		log.Println("JWT_SECRET_KEY is not set") // Log สำหรับ debug
	}
}

// HashPassword: ฟังก์ชันเข้ารหัสรหัสผ่าน
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// CheckPassword: ฟังก์ชันตรวจสอบรหัสผ่าน
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// GenerateToken: สร้าง JWT token
func GenerateToken(user models.RegisterUsers) (string, error) {
	if secretKey == "" {
		log.Println("JWT_SECRET_KEY is not set")
		return "", fmt.Errorf("JWT_SECRET_KEY is not set")
	}

	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Printf("Error signing token: %v\n", err)
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return signedToken, nil
}

// ฟังก์ชันตรวจสอบ token
func ValidateToken(tokenString string) bool {
	if secretKey == "" {
		log.Println("JWT_SECRET_KEY is not set")
		return false
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// ตรวจสอบ signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method: %v", token.Method)
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		log.Println("Invalid token:", err)
		return false
	}

	return true
}

// RegisterUser: ลงทะเบียนผู้ใช้
func RegisterUser(user *models.RegisterUsers) error {
	// ตรวจสอบว่า email ซ้ำหรือไม่
	if CheckDuplicateEmail(user.Email) {
		return fmt.Errorf("email %s is already registered", user.Email)
	}

	// เข้ารหัสรหัสผ่าน
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}
	user.Password = hashedPassword

	// สร้าง account_id ใหม่
	user.Acount_ID = uint(uuid.New().ID())

	// บันทึกข้อมูลลงฐานข้อมูล
	if err := config.DB.Create(user).Error; err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	return nil
}

func AuthenticateUser(credentials models.RegisterUsers) (models.RegisterUsers, error) {
	var user models.RegisterUsers

	// ค้นหาผู้ใช้จาก username
	if err := config.DB.Where("username = ?", credentials.Username).First(&user).Error; err != nil {
		log.Printf("User not found for username: %s\n", credentials.Username) // Log สำหรับ debug
		return user, fmt.Errorf("Invalid username or password")
	}

	// ตรวจสอบรหัสผ่าน
	if err := CheckPassword(user.Password, credentials.Password); err != nil {
		log.Printf("Password mismatch for username: %s\n", credentials.Username) // Log สำหรับ debug
		return user, fmt.Errorf("Invalid username or password")
	}

	return user, nil
}

func CheckDuplicateEmail(email string) bool {
	var user models.RegisterUsers
	if err := config.DB.Where("email = ?", email).First(&user).Error; err == nil {
		// ถ้ามี email อยู่แล้วในฐานข้อมูล
		return true
	}
	return false
}
