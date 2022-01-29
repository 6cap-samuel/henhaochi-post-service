package usecases

import (
	"post-service/entities"
	"post-service/usecases/in"
	"post-service/usecases/out"
)

type retrievePostDetailsInteractor struct {
	postDataSource out.PostDataSource
}

func NewRetrievePostDetailsInteractor(
	postDataSource *out.PostDataSource,
) in.RetrievePostDetailsInput {
	return &retrievePostDetailsInteractor{
		*postDataSource,
	}
}

func (r retrievePostDetailsInteractor) Get(
	postId string,
) (response entities.Post) {
	return r.postDataSource.Get(postId)
}
