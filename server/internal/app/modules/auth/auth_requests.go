package auth

type RegisterPayload struct {
	Name     string `json:"name,omitempty" validate:"required,min=2,max=50"`
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,min=8,max=50"`
}
