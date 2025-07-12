package handler

import (
	"strings"

	"gapi/auth/service"
	"gapi/pkg/helper"

	"github.com/gofiber/fiber/v2"
)

type subHandler struct {
	si service.ServiceInterface
}

func NewHandler(si service.ServiceInterface) *subHandler {
	return &subHandler{
		si: si,
	}
}

func (h *subHandler) LoginRedirect(f *fiber.Ctx) error {
	url, err := h.si.RedirectUrl()
	if err != nil {
		if strings.Contains(err.Error(), "error") {
			return f.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(err.Error()))
		}

		return f.Status(fiber.StatusInternalServerError).JSON(helper.ErrorResponse(err.Error()))

	}
	return f.Redirect(url)
}

func (h *subHandler) GetAccessToken(f *fiber.Ctx) error {
	code := f.Query("code")

	data, err := h.si.AccessToken(code,f.Context())
	if err != nil {
		if strings.Contains(err.Error(), "error") {
			return  f.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(err.Error()))
		}
		return f.Status(fiber.StatusInternalServerError).JSON(helper.ErrorResponse(err.Error()))
	}

	return f.Status(fiber.StatusOK).JSON(helper.SuccessWithDataResponse("success get data",data))
}


