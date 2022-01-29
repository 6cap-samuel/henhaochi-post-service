package in

import "post-service/entities"

type CreatePostInput interface {
	Create(post entities.Post)
}
