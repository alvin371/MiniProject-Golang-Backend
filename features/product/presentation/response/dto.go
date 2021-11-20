package response

import "time"

type Product struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Price     float32   `json:"price"`
	Satuan    int       `json:"satuan"`
}
