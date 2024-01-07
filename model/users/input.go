package users

type UserCreateInput struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type UserUpdateInput struct {
	ID       int    `form:"id"` // buat update
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

type UserGetOneByIdInput struct {
	ID int `uri:"id" binding:"required"`
}
