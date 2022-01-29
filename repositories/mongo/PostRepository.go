package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"post-service/configurations"
	"post-service/entities"
	"post-service/exceptions"
	"post-service/usecases/out"
)

type postRepository struct {
	Collection *mongo.Collection
}

func NewPostRepository(database *mongo.Database) out.PostDataSource {
	return &postRepository{
		Collection: database.Collection("posts"),
	}
}

func (p postRepository) GetAll() (
	response []entities.Post,
) {
	ctx, cancel := configurations.NewMongoContext()
	defer cancel()

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"date_created", -1}})

	cursor, err := p.Collection.Find(ctx, bson.M{}, findOptions)
	exceptions.PanicIfNeeded(err)

	for cursor.Next(ctx) {
		var post entities.Post

		err := cursor.Decode(&post)
		exceptions.PanicIfNeeded(err)

		response = append(response, post)
	}
	return response
}

func (p postRepository) GetAllWith(
	filters []string,
) (response []entities.Post) {
	ctx, cancel := configurations.NewMongoContext()
	defer cancel()

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"date_created", -1}})

	var bsonMap bson.A
	for _, item := range filters {
		bsonMap = append(bsonMap, item)
	}

	cursor, err := p.Collection.Find(
		ctx,
		bson.D{
			{"hash_tags",
				bson.D{
					{"$all",
						bsonMap,
					},
				},
			},
		},
		findOptions,
	)

	exceptions.PanicIfNeeded(err)

	for cursor.Next(ctx) {
		var post entities.Post

		err := cursor.Decode(&post)
		exceptions.PanicIfNeeded(err)

		response = append(response, post)
	}
	return response
}

func (p postRepository) Create(post entities.Post) {
	ctx, cancel := configurations.NewMongoContext()
	defer cancel()

	_, err := p.Collection.InsertOne(ctx, post)

	exceptions.PanicIfNeeded(err)
}

func (p postRepository) UpdateFood(
	postId string,
	food []entities.Food,
) {
	ctx, cancel := configurations.NewMongoContext()
	defer cancel()

	_, err := p.Collection.UpdateOne(
		ctx,
		bson.D{
			{"_id", postId},
		},
		bson.D{
			{"$set",
				bson.D{
					{"foods", food},
				},
			},
		},
	)

	exceptions.PanicIfNeeded(err)
}

func (p postRepository) Get(
	postId string,
) (response entities.Post) {
	ctx, cancel := configurations.NewMongoContext()
	defer cancel()

	err := p.Collection.FindOne(
		ctx,
		bson.D{
			{"_id",
				postId,
			},
		},
	).Decode(&response)

	exceptions.PanicIfNeeded(err)

	return response
}
