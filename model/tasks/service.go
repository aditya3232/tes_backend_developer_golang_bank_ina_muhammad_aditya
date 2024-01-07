package tasks

type Service interface {
	GetAll() ([]Task, error)
	GetOne(input TaskGetOneByIdInput) (Task, error)
	Create(input TaskCreateInput) (Task, error)
	Update(input TaskUpdateInput) (Task, error)
	Delete(input TaskGetOneByIdInput) error
}

type service struct {
	taskRepository Repository
}

func NewService(taskRepository Repository) *service {
	return &service{taskRepository}
}

func (s *service) GetAll() ([]Task, error) {
	tasks, err := s.taskRepository.GetAll()
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (s *service) GetOne(input TaskGetOneByIdInput) (Task, error) {
	task, err := s.taskRepository.GetOne(input.ID)
	if err != nil {
		return task, err
	}

	return task, nil
}

func (s *service) Create(input TaskCreateInput) (Task, error) {
	task := Task{
		UserID:      input.UserID,
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
	}

	newTask, err := s.taskRepository.Create(task)
	if err != nil {
		return newTask, err
	}

	return newTask, nil
}

func (s *service) Update(input TaskUpdateInput) (Task, error) {
	_, err := s.taskRepository.GetOne(input.ID)
	if err != nil {
		return Task{}, err
	}

	task := Task{
		ID:          input.ID,
		UserID:      input.UserID,
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
	}

	newTask, err := s.taskRepository.Update(task)
	if err != nil {
		return newTask, err
	}

	return newTask, nil
}

func (s *service) Delete(input TaskGetOneByIdInput) error {
	_, err := s.taskRepository.GetOne(input.ID)
	if err != nil {
		return err
	}

	err = s.taskRepository.Delete(input.ID)
	if err != nil {
		return err
	}

	return nil
}
