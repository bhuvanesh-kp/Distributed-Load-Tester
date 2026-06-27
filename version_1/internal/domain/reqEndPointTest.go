package domain

type Endpoint struct {
	Url string `json:"endpoint" binding:"required"`
}
