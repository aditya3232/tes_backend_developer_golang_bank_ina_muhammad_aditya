package tasks

type TaskCreateInput struct {
	UserID      int    `form:"user_id" binding:"required"`
	Title       string `form:"title" binding:"required"`
	Description string `form:"description" binding:"required"`
	Status      string `form:"status"`
}

type TaskUpdateInput struct {
	ID          int    `form:"id"` // buat update
	UserID      int    `form:"user_id"`
	Title       string `form:"title"`
	Description string `form:"description"`
	Status      string `form:"status"`
}

type TaskGetOneByIdInput struct {
	ID int `uri:"id" binding:"required"`
}
