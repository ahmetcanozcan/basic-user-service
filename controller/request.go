package controller

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EditUserRequest struct {
	ID       string `param:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserRequest struct {
	ID string `param:"id"`
}

type DeleteUserRequest struct {
	ID string `param:"id"`
}
