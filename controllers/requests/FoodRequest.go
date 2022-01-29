package requests

type FoodRequest struct {
	Name   string  `json:"name"`
	Image  string  `json:"image_url"`
	Cost   float32 `json:"cost"`
	Rating int8    `json:"rating"`
}
