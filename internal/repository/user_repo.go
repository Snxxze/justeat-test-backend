package repository

import (
	"justeat/internal/models"
	"gorm.io/gorm"
)

// ตัว manager ของข้อมูล user
type UserRepository struct {
	DB *gorm.DB		// เชื่อม DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// create ผู้ใช้ใหม่
func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error	// ถ้า error จะ return โอยไม่สร้าง
}

// FindByEmail ค้นหา user โดยใช้ email
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.DB.Where("email = ?", email).First(&user)

	return &user, result.Error
}

// FindByID ค้นหา user ตาม ID
func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, id).Error; err != nil {		// WHERE id = ?
		return nil, err
	}
	return &user, nil
}