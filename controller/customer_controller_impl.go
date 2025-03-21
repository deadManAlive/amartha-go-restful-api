package controller

import (
	"strconv"

	"github.com/aronipurwanto/go-restful-api/exception"
	"github.com/aronipurwanto/go-restful-api/model/web"
	"github.com/aronipurwanto/go-restful-api/service"
	"github.com/gofiber/fiber/v2"
)

type CustomerControllerImpl struct {
	CustomerService service.CustomerService
}

func NewCustomerController(customerService service.CustomerService) CustomerController {
	return &CustomerControllerImpl{
		CustomerService: customerService,
	}
}

func (controller *CustomerControllerImpl) Create(c *fiber.Ctx) error {
	customerCreateRequest := new(web.CustomerCreateRequest)
	if err := c.BodyParser(customerCreateRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	}

	customerResponse, err := controller.CustomerService.Create(c.Context(), *customerCreateRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   customerResponse,
	})
}

func (controller *CustomerControllerImpl) Update(c *fiber.Ctx) error {
	customerUpdateRequest := new(web.CustomerUpdateRequest)
	if err := c.BodyParser(customerUpdateRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	}

	id, err := strconv.Atoi(c.Params("customerId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Invalid Category ID",
			Data:   err.Error(),
		})
	}
	customerUpdateRequest.Id = uint64(id)

	customerResponse, err := controller.CustomerService.Update(c.Context(), *customerUpdateRequest)
	if err != nil {
		if _, ok := err.(exception.NotFoundError); ok {
			return c.Status(fiber.StatusNotFound).JSON(web.WebResponse{
				Code:   fiber.StatusNotFound,
				Status: "Not Found",
				Data:   err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(web.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   customerResponse,
	})
}

func (controller *CustomerControllerImpl) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("customerId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Invalid Customer ID",
			Data:   err.Error(),
		})
	}

	err = controller.CustomerService.Delete(c.Context(), uint64(id))
	if err != nil {
		if _, ok := err.(exception.NotFoundError); ok {
			return c.Status(fiber.StatusNotFound).JSON(web.WebResponse{
				Code:   fiber.StatusNotFound,
				Status: "Not Found",
				Data:   err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(web.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "Deleted Successfully",
	})
}

func (controller *CustomerControllerImpl) FindById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("customerId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Invalid Customer ID",
			Data:   err.Error(),
		})
	}

	customerResponse, err := controller.CustomerService.FindById(c.Context(), uint64(id))
	if err != nil {
		if _, ok := err.(exception.NotFoundError); ok {
			return c.Status(fiber.StatusNotFound).JSON(web.WebResponse{
				Code:   fiber.StatusNotFound,
				Status: "Not Found",
				Data:   err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(web.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   customerResponse,
	})
}

func (controller *CustomerControllerImpl) FindAll(c *fiber.Ctx) error {
	customerResponse, err := controller.CustomerService.FindAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   customerResponse,
	})
}
