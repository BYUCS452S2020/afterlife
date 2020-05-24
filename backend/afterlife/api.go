package afterlife

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	LoginRequest
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Timeline []Event
