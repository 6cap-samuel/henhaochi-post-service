package usecases

import (
	"post-service/entities"
	"post-service/usecases/in"
	"post-service/usecases/out"
)

type createPostInput struct {
	postDataSource out.PostDataSource
}

func NewCreatePostInput(
	post *out.PostDataSource,
) in.CreatePostInput {
	return &createPostInput{
		*post,
	}
}

func (c *createPostInput) Create(
	post entities.Post,
) {
	go c.postDataSource.Create(post)
}
