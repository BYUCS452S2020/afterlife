package afterlife

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	LoginRequest
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Timeline []Event
