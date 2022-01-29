package requests

type CreatePostRequest struct {
	StoreName     string   `json:"store_name"`
	StoreImage    string   `json:"store_image"`
	StoreLocation string   `json:"store_location"`
	StoreLat      float32  `json:"store_lat"`
	StoreLong     float32  `json:"store_long"`
	Description   string   `json:"description"`
	Rating        int8     `json:"rating"`
	Neutrals      []string `json:"neutrals"`
	Positives     []string `json:"positives"`
	Negatives     []string `json:"negatives"`
	HashTags      []string `json:"hash_tags"`
}
