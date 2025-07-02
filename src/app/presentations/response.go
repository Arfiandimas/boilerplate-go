package presentations

//go:generate easytags $GOFILE json,form

// ExampleResponse Presentation Response Model Example
type ExampleResponse struct {
	ID      uint64 `json:"id,omitempty"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}
