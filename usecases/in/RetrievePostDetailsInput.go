package in

import "post-service/entities"

type RetrievePostDetailsInput interface {
	Get(postId string) (response entities.Post)
}
