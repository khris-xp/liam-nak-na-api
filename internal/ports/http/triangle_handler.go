package http

import (
	"errors"
	"liam-nak-na-api/internal/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TriangleHandler struct {
	triangleService domain.TriangleService
}

func NewTriangleHandler(service domain.TriangleService) *TriangleHandler {
	return &TriangleHandler{
		triangleService: service,
	}
}

type triangleRequest struct {
	Height string `json:"height"`
	Width  string `json:"width"`
	Base   string `json:"base"`
}

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

func (h *TriangleHandler) ClassifyTriangle(c echo.Context) error {
	var req triangleRequest
	if err := c.Bind(&req); err != nil {
		return newErrorResponse(c, http.StatusBadRequest, "Invalid request format")
	}

	// Validate and convert inputs
	triangle, err := validateAndConvertInput(req)
	if err != nil {
		return newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	// Call service
	classification, err := h.triangleService.ClassifyTriangle(triangle)
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, "Error processing request")
	}

	return newSuccessResponse(c, "Triangle classification successful", classification)
}

func validateAndConvertInput(req triangleRequest) (domain.Triangle, error) {
	height, err := validateAndParseFloat(req.Height)
	if err != nil {
		return domain.Triangle{}, err
	}

	width, err := validateAndParseFloat(req.Width)
	if err != nil {
		return domain.Triangle{}, err
	}

	base, err := validateAndParseFloat(req.Base)
	if err != nil {
		return domain.Triangle{}, err
	}

	return domain.Triangle{
		Height: height,
		Width:  width,
		Base:   base,
	}, nil
}

//TODO: Teetouch Jaknamon ทำต่อจากนี้
//TODO: Check Triangle Validate Input Word format in Thai

func validateAndParseFloat(input string) (float64, error) {
	if input == "" {
		return 0, errors.New("Please fill out this field.")
	}

	parts := strings.Split(input, ".")
	if len(parts) > 1 && len(parts[1]) > 3 {
		return 0, errors.New("Invalid inputs : Enter a number between 0-999999.")
	}
	
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, errors.New("Invalid inputs : Enter a number between 0-999999.")
	}

	if num < 0 || num > 999999 {
		return 0, errors.New("Invalid inputs : Enter a number between 0-999999.")
	}

	return num, nil
}

func newSuccessResponse(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusOK, Response{
		Status:  true,
		Message: message,
		Code:    http.StatusOK,
		Data:    data,
	})
}

func newErrorResponse(c echo.Context, code int, message string) error {
	return c.JSON(code, Response{
		Status:  false,
		Message: message,
		Code:    code,
		Data:    nil,
	})
}
