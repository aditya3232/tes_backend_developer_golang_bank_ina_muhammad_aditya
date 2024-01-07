package users

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]User, error)
	GetOne(id int) (User, error)
	Create(User) (User, error)
	Update(User) (User, error)
	Delete(id int) error
	GetByEmail(email string) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]User, error) {
	var users []User

	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *repository) GetOne(id int) (User, error) {
	var user User

	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Create(user User) (User, error) {
	err := r.db.Model(&user).Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Model(&user).Where("id = ?", user.ID).Updates(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Delete(id int) error {
	var user User

	err := r.db.Delete(&user, "id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetByEmail(email string) (bool, error) {
	var user User

	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
