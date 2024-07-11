package handler

import (
	"net/http"

	"github.com/PParist/go-otp-service/entities"
	"github.com/PParist/go-otp-service/service"
	"github.com/labstack/echo/v4"
)

type mailOtpHandler struct {
	service service.OTPEmailService
}

func NewMailOtpService(service service.OTPEmailService) mailOtpHandler {
	return mailOtpHandler{service: service}
}

func (h *mailOtpHandler) CreateMailOTP(c echo.Context) error {
	var email entities.Email
	err := c.Bind(&email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	otpuid, err := h.service.CreateMailOTP(email.Receiver)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "bad bad request")
	}
	response := map[string]interface{}{
		"message": otpuid,
	}
	return c.JSON(http.StatusOK, response)
}
func (h *mailOtpHandler) ValidateOTP(c echo.Context) error {
	var req entities.OtpRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	err = h.service.ValidateOTP(req.Otp_Uid, req.Otp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	response := map[string]interface{}{
		"message": "otp is valid",
	}
	return c.JSON(http.StatusOK, response)
}
