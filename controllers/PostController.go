package controllers

import (
	"encoding/json"
	"post-service/controllers/mappers"
	"post-service/controllers/requests"
	"post-service/controllers/responses"
	"post-service/exceptions"
	"post-service/usecases/in"

	"github.com/gofiber/fiber/v2"
)

type PostController struct {
	retrieve in.RetrievePostsInput
	create   in.CreatePostInput
	update   in.UpdatePostInput
	details  in.RetrievePostDetailsInput
}

func NewPostController(
	retrieve *in.RetrievePostsInput,
	create *in.CreatePostInput,
	update *in.UpdatePostInput,
	details *in.RetrievePostDetailsInput,
) PostController {
	return PostController{
		*retrieve,
		*create,
		*update,
		*details,
	}
}

func (controller *PostController) Route(app *fiber.App) {
	app.Post("/posts", controller.Create)
	app.Get("/posts", controller.List)
	app.Get("/posts/:postId", controller.Detail)
	app.Put("/posts/:postId/food", controller.UpdateFood)
}

func (controller *PostController) Create(c *fiber.Ctx) error {
	var request requests.CreatePostRequest
	err := c.BodyParser(&request)
	var storeEntity = mappers.CreatePostRequestToPostMapper(request)

	exceptions.PanicIfNeeded(err)

	controller.create.Create(*storeEntity)
	return c.JSON(responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   storeEntity.Id,
	})
}

func (controller *PostController) List(c *fiber.Ctx) error {
	var filters []string

	if c.Get("filters") != "" {
		err := json.Unmarshal([]byte(c.Get("filters")), &filters)
		exceptions.PanicIfNeeded(err)
	}

	result := controller.retrieve.GetAll(filters)
	return c.JSON(responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	})
}

func (controller *PostController) Detail(c *fiber.Ctx) error {
	result := controller.details.Get(c.Params("postId"))

	return c.JSON(responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	})
}

func (controller *PostController) UpdateFood(c *fiber.Ctx) error {
	var request requests.UpdateFoodRequest
	err := c.BodyParser(&request)

	var foodEntities = mappers.UpdateFoodRequestToFoodMapper(request)

	exceptions.PanicIfNeeded(err)

	controller.update.UpdateFood(c.Params("postId"), foodEntities)
	return c.JSON(responses.WebResponse{
		Code:   200,
		Status: "OK",
	})
}
