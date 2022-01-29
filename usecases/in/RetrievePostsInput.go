package in

import "post-service/entities"

type RetrievePostsInput interface {
	GetAll(filters []string) (response []entities.Post)
}
