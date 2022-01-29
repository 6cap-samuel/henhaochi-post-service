package mappers

import (
	"post-service/configurations"
	"post-service/controllers/requests"
	"post-service/entities"
	"time"
)

func CreatePostRequestToPostMapper(request requests.CreatePostRequest) *entities.Post {
	return &entities.Post{
		Id:          configurations.NewIdentity(),
		Description: request.Description,
		Positives:   request.Positives,
		Neutrals:    request.Neutrals,
		Negatives:   request.Negatives,
		Store: entities.Store{
			Id:       configurations.NewIdentity(),
			Name:     request.StoreName,
			Image:    request.StoreImage,
			Location: request.StoreLocation,
			Lat:      request.StoreLat,
			Long:     request.StoreLong,
		},
		Rating:      request.Rating,
		HashTags:    request.HashTags,
		DateCreated: time.Now(),
	}
}
