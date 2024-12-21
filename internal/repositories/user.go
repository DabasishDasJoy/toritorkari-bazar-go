package repositories

import (
	"toritorkari-bazar/internal/domain"
	"toritorkari-bazar/internal/models"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func UserDBInstance(db *gorm.DB) domain.IUserRepo {
	return &UserRepo{
		db: db,
	}
}

func (repo UserRepo) SignUp(user models.User) error {
	if err := repo.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (repo UserRepo) GetUserByEmail(email string) (user models.User, err error) {
	var User models.User

	query := `select email, password, role, name from users where email=?`

	err = repo.db.Raw(query, email).First(&User).Error

	if err != nil {
		return models.User{}, err
	}

	return User, nil
}
