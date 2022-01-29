package in

import "post-service/entities"

type UpdatePostInput interface {
	UpdateFood(postId string, food []entities.Food)
}
