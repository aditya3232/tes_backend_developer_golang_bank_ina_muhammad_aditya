package tasks

import "time"

type TaskGetFormatter struct {
	ID          int        `json:"id"`
	UserID      int        `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type TaskCreateFormatter struct {
	ID          int        `json:"id"`
	UserID      int        `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type TaskUpdateFormatter struct {
	ID          int        `json:"id"`
	UserID      int        `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func TaskGetFormat(task Task) TaskGetFormatter {
	var taskFormatter TaskGetFormatter

	taskFormatter.ID = task.ID
	taskFormatter.UserID = task.UserID
	taskFormatter.Title = task.Title
	taskFormatter.Description = task.Description
	taskFormatter.Status = task.Status
	taskFormatter.CreatedAt = task.CreatedAt
	taskFormatter.UpdatedAt = task.UpdatedAt

	return taskFormatter
}

func TaskGetAllFormat(tasks []Task) []TaskGetFormatter {
	formatter := []TaskGetFormatter{}

	for _, task := range tasks {
		taskGetFormatter := TaskGetFormat(task)         // format data satu persatu
		formatter = append(formatter, taskGetFormatter) // append data formatter ke slice formatter
	}

	return formatter
}

func TaskCreateFormat(task Task) TaskCreateFormatter {
	var taskFormatter TaskCreateFormatter

	taskFormatter.ID = task.ID
	taskFormatter.UserID = task.UserID
	taskFormatter.Title = task.Title
	taskFormatter.Description = task.Description
	taskFormatter.Status = task.Status
	taskFormatter.CreatedAt = task.CreatedAt
	taskFormatter.UpdatedAt = task.UpdatedAt

	return taskFormatter
}

func TaskUpdateFormat(task Task) TaskUpdateFormatter {
	var taskFormatter TaskUpdateFormatter

	taskFormatter.ID = task.ID
	taskFormatter.UserID = task.UserID
	taskFormatter.Title = task.Title
	taskFormatter.Description = task.Description
	taskFormatter.Status = task.Status
	taskFormatter.CreatedAt = task.CreatedAt
	taskFormatter.UpdatedAt = task.UpdatedAt

	return taskFormatter
}
