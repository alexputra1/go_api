package transfer

type LoginTransfer struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}