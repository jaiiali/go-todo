package todo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaiiali/go-todo/internal/core/ports"
)

type Handler struct {
	uc ports.TodoUseCase
}

func NewHandler(uc ports.TodoUseCase, app fiber.Router) *Handler {
	handler := &Handler{
		uc: uc,
	}

	app.Post("/todos", handler.Create)
	app.Get("/todos", handler.FindAll)
	app.Get("/todos/:id<len(26)>", handler.FindByID)

	return handler
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var todo = &todoReq{}

	if err := c.BodyParser(todo); err != nil {
		return h.customError(c, err)
	}

	result, err := h.uc.Create(todo.Title, todo.Description)
	if err != nil {
		return h.customError(c, err)
	}

	var resp = &todoResp{}
	resp.fromDomain(result)

	err = c.JSON(resp)
	if err != nil {
		return h.customError(c, err)
	}

	return nil
}

func (h *Handler) FindAll(c *fiber.Ctx) error {
	result, err := h.uc.FindAll()
	if err != nil {
		return h.customError(c, err)
	}

	var resp = todoListResp{}
	resp.fromDomain(result)

	err = c.JSON(resp)
	if err != nil {
		return h.customError(c, err)
	}

	return nil
}

func (h *Handler) FindByID(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := h.uc.FindByID(id)
	if err != nil {
		return h.customError(c, err)
	}

	var resp = &todoResp{}
	resp.fromDomain(result)

	err = c.JSON(resp)
	if err != nil {
		return h.customError(c, err)
	}

	return nil
}

func (h *Handler) customError(c *fiber.Ctx, err error) error {
	//nolint: errcheck
	c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": err.Error(),
	})

	return nil
}
