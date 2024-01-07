package users

import "time"

type UserGetFormatter struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UserCreateFormatter struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UserUpdateFormatter struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func UserGetFormat(user User) UserGetFormatter {
	var userFormatter UserGetFormatter

	userFormatter.ID = user.ID
	userFormatter.Name = user.Name
	userFormatter.Email = user.Email
	userFormatter.Password = user.Password
	userFormatter.CreatedAt = user.CreatedAt
	userFormatter.UpdatedAt = user.UpdatedAt

	return userFormatter
}

func UserGetAllFormat(users []User) []UserGetFormatter {
	formatter := []UserGetFormatter{}

	for _, user := range users {
		userGetFormatter := UserGetFormat(user)         // format data satu persatu
		formatter = append(formatter, userGetFormatter) // append data formatter ke slice formatter
	}

	return formatter
}

func UserCreateFormat(user User) UserCreateFormatter {
	var userFormatter UserCreateFormatter

	userFormatter.ID = user.ID
	userFormatter.Name = user.Name
	userFormatter.Email = user.Email
	userFormatter.Password = user.Password
	userFormatter.CreatedAt = user.CreatedAt
	userFormatter.UpdatedAt = user.UpdatedAt

	return userFormatter
}

func UserUpdateFormat(user User) UserUpdateFormatter {
	var userFormatter UserUpdateFormatter

	userFormatter.ID = user.ID
	userFormatter.Name = user.Name
	userFormatter.Email = user.Email
	userFormatter.Password = user.Password
	userFormatter.CreatedAt = user.CreatedAt
	userFormatter.UpdatedAt = user.UpdatedAt

	return userFormatter
}
