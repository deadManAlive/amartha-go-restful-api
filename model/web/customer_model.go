package web

type CustomerCreateRequest struct {
	Name string `validate:"required,min=1,max=100" json:"name"`
}

type CustomerUpdateRequest struct {
	Id      uint64 `validate:"required"`
	Name    string `validate:"required,max=200,min=1" json:"name"`
	Email   string `validate:"required,max=200,min=1" json:"email"`
	Phone   string `validate:"required,max=200,min=1" json:"phone"`
	Address string `validate:"required,max=200,min=1" json:"address"`
}

type CustomerResponse struct {
	Id         uint64 `json:"customer_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	LoyaltyPts int    `json:"loyalty_points"`
}
