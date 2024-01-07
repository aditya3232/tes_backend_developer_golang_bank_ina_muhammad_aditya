package tasks

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Task, error)
	GetOne(id int) (Task, error)
	Create(Task) (Task, error)
	Update(Task) (Task, error)
	Delete(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Task, error) {
	var tasks []Task

	err := r.db.Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *repository) GetOne(id int) (Task, error) {
	var task Task

	err := r.db.First(&task, "id = ?", id).Error
	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *repository) Create(task Task) (Task, error) {
	err := r.db.Model(&task).Create(&task).Error
	if err != nil {
		return task, err
	}
	return task, nil
}

func (r *repository) Update(task Task) (Task, error) {
	err := r.db.Model(&task).Where("id = ?", task.ID).Updates(&task).Error
	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *repository) Delete(id int) error {
	var task Task

	err := r.db.Delete(&task, "id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}
