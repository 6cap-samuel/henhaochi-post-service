package usecases

import (
	"post-service/entities"
	"post-service/usecases/in"
	"post-service/usecases/out"
)

type retrievePostInput struct {
	postDataSource out.PostDataSource
}

func NewRetrievePostInteractor(
	postDataSoruce *out.PostDataSource,
) in.RetrievePostsInput {
	return &retrievePostInput{
		*postDataSoruce,
	}
}

func (r *retrievePostInput) GetAll(filters []string) (response []entities.Post) {
	if len(filters) == 0 {
		return r.postDataSource.GetAll()
	} else {
		return r.postDataSource.GetAllWith(filters)
	}
}
