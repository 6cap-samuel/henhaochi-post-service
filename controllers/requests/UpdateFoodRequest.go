package requests

type UpdateFoodRequest struct {
	Foods []FoodRequest `json:"foods"`
}
