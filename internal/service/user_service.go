package service

import (
	"justeat/internal/models"     // ดึง models ที่สร้างไว้
	"justeat/internal/repository" // ใช้ในการจัดการฐานข้อมูล
	"justeat/internal/utils"

	"errors"

	"golang.org/x/crypto/bcrypt" // hashpassword
)

// จัดการ logic ของ user
type UserService struct {
	Repo *repository.UserRepository
}

// เหมือนสร้าง constructor
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

// Register
func (s *UserService) Register(name, email, password string) (*models.User, error) {
	// ตรวจ email ซ้ำ
	existingUser, _ := s.Repo.FindByEmail(email)
	if existingUser.ID != 0 {
		return nil, errors.New("email นี้ถูกใช้งานแล้ว")
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("ไม่สามารถเข้ารหัสผ่านได้")
	}

	// สร้าง user ใหม่
	user := &models.User{
		Name:	name,
		Email: email,
		Password: string(hashedPassword),
	}

	// บันทึก user ลงฐานข้อมูล
	err = s.Repo.Create(user)
	if err != nil {
		return nil, errors.New("เกิดข้อผิดพลาดขณะบันทึกข้อมูล")
	}

	return user, nil
}

// Login
func (s *UserService) Login(email, password, secret string) (string, error) {
	user, err := s.Repo.FindByEmail(email)
	if err != nil || user.ID == 0 {
		return "",errors.New("email หรือ password ไม่ถูกต้อง")
	}

	// ตรวจสอบรหัส
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("email หรือ password ไม่ถูกต้อง")
	}

	// สร้าง JWT Token
	token, err := utils.GenerateJWT(user.ID, secret)
	if err != nil {
		return "", errors.New("ไม่สามารถสร้าง token ได้")
	}

	return token, nil
}

// ค้นหา user ด้วย id
func (s *UserService) GetByID(id uint) (*models.User, error) {
	return s.Repo.FindByID(id)
}