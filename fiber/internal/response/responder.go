package response

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
)

type (
	Meta struct {
		Page  int `json:"page"`
		Size  int `json:"size"`
		Total int `json:"total"`
		Start int `json:"-"`
	}
	Response struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Meta    *Meta       `json:"meta,omitempty"`
		Data    interface{} `json:"data"`
	}
	Option func(*Response)
)

func New(options ...Option) *Response {
	response := Response{
		Code:    fiber.StatusOK,
		Message: "Success",
	}
	for _, option := range options {
		option(&response)
	}
	return &response
}

func WithCode(code int) Option {
	return func(r *Response) {
		r.Code = code
	}
}

func WithMessage(message string) Option {
	return func(r *Response) {
		r.Message = message
	}
}

func WithMeta(meta *Meta) Option {
	return func(r *Response) {
		r.Meta = meta
	}
}

func WithData(data interface{}) Option {
	return func(r *Response) {
		if data == nil {
			r.Data = struct{}{}
		}
		r.Data = data
	}
}

func (r *Response) WithMeta(meta *Meta) *Response {
	r.Meta = meta
	return r
}

func (r *Response) WithData(data interface{}) *Response {
	r.Data = data
	return r
}

// Success Sending a successful response
func Success(ctx fiber.Ctx, options ...Option) error {
	response := Response{
		Code:    fiber.StatusOK,
		Message: "Success",
	}
	for _, option := range options {
		option(&response)
	}

	if response.Data == nil {
		response.Data = struct{}{}
	}

	return ctx.JSON(response)
}

// UnAuthenticated Authentication Failure
func UnAuthenticated(ctx fiber.Ctx, err error) error {
	response := Response{
		Code:    http.StatusUnauthorized,
		Message: err.Error(),
		Data:    struct{}{},
	}

	return ctx.JSON(response)
}

// NotFound Sending a not found response
func NotFound(ctx fiber.Ctx, err error) error {
	response := Response{
		Code:    http.StatusNotFound,
		Message: err.Error(),
		Data:    struct{}{},
	}

	return ctx.JSON(response)
}

// ValidationError Sending a validation error response
func ValidationError(ctx fiber.Ctx, err error) error {
	response := Response{
		Code:    http.StatusUnprocessableEntity,
		Message: err.Error(),
		Data:    struct{}{},
	}

	return ctx.JSON(response)
}

// InternalServerError Sending an internal server error response
func InternalServerError(ctx fiber.Ctx, err error) error {
	response := Response{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
		Data:    struct{}{},
	}

	return ctx.JSON(response)
}

// Send Sending a basic response
func Send(ctx fiber.Ctx, options ...Option) error {
	response := Response{
		Code:    fiber.StatusOK,
		Message: "Success",
	}
	for _, option := range options {
		option(&response)
	}

	if response.Data == nil {
		response.Data = struct{}{}
	}

	return ctx.JSON(response)
}
