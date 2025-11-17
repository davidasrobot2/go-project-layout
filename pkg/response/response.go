package response

import (
	"davidasrobot/project-layout/pkg/constant"

	"github.com/gofiber/fiber/v2"
)

type ResponseStruct struct {
	ResponseCode    string      `json:"responseCode"`
	ResponseMessage string      `json:"responseMessage"`
	ResponseData    interface{} `json:"responseData"`
}

// NewSuccessResponse creates a new success response.
func NewSuccessResponse(code string, message string, data interface{}) ResponseStruct {
	return ResponseStruct{
		ResponseCode:    code,
		ResponseMessage: message,
		ResponseData:    data,
	}
}

// NewErrorResponse creates a new error response.
func NewErrorResponse(code string, message string, data ...interface{}) ResponseStruct {
	return ResponseStruct{
		ResponseCode:    code,
		ResponseMessage: message,
		ResponseData:    nil, // Error responses typically don't have data
	}
}

func Success(c *fiber.Ctx, data interface{}) error {
	response := NewSuccessResponse(constant.SuccessCodeOK, "Success", data)
	return c.Status(fiber.StatusOK).JSON(response)
}

func Error(c *fiber.Ctx, errCode int, errMessage interface{}) error {
	response := ErrorClassifier(errCode, errMessage.(string))
	return c.Status(errCode).JSON(response)
}

func ErrorClassifier(errCode int, errMessage string) ResponseStruct {
	switch errCode {
	case fiber.ErrBadRequest.Code:
		return NewErrorResponse(constant.ErrCodeBadRequest, constant.ErrMessageBadRequest)
	case fiber.ErrUnauthorized.Code:
		return NewErrorResponse(constant.ErrCodeUnauthorized, constant.ErrMessageUnauthorized)
	case fiber.ErrForbidden.Code:
		return NewErrorResponse(constant.ErrCodeForbidden, constant.ErrMessageForbidden)
	case fiber.ErrNotFound.Code:
		return NewErrorResponse(constant.ErrCodeNotFound, constant.ErrMessageNotFound)
	case 0:
		return NewErrorResponse(constant.ErrCodeNotFound, constant.ErrMessageInvalidRequestBody, errMessage)
	default:
		return NewErrorResponse(constant.ErrCodeInternal, constant.ErrMessageInternalServerError)
	}
}
