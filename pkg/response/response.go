package response

import (
	"davidasrobot2/go-boilerplate/pkg/constant"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ResponseTemplate struct {
	Code    int
	Message string
}

var (
	ResponseInternalServerError  = ResponseTemplate{500, "Internal Server Error"}
	ResponseServiceUnavailable   = ResponseTemplate{503, "Service Unavailable"}
	ResponseGatewayTimeout       = ResponseTemplate{504, "Gateway Timeout"}
	ResponseBadRequest           = ResponseTemplate{400, "Bad Request"}
	ResponseNotFound             = ResponseTemplate{404, "Not Found"}
	ResponseUnauthorized         = ResponseTemplate{401, "Unauthorized"}
	ResponseForbidden            = ResponseTemplate{403, "Forbidden"}
	ResponseTooManyRequests      = ResponseTemplate{429, "Too Many Requests"}
	ResponseMethodNotAllowed     = ResponseTemplate{405, "Method Not Allowed"}
	ResponsePayloadTooLarge      = ResponseTemplate{413, "Payload Too Large"}
	ResponseUnsupportedMediaType = ResponseTemplate{415, "Unsupported Media Type"}
	ResponseInvalidRequestBody   = ResponseTemplate{422, "Invalid Request Body"}
	ResponseSuccess              = ResponseTemplate{200, "Success"}
	ResponseCreated              = ResponseTemplate{201, "Created"}
	ResponseConflict             = ResponseTemplate{409, "Conflict"}
)

func BuildResponse(template ResponseTemplate, data ...interface{}) fiber.Map {
	resp := fiber.Map{
		"responseCode":    template.Code,
		"responseMessage": template.Message,
	}
	if len(data) > 0 && data[0] != nil {
		resp["responseData"] = data[0]
	}
	return resp
}

func FormatValidationErrors(c *fiber.Ctx, errs validator.FieldError) error {
	translateError := errs.Field() + " " + errs.Tag()
	return c.Status(fiber.StatusBadRequest).JSON(BuildResponse(ResponseBadRequest, translateError))
}

func HandleErrors(c *fiber.Ctx, err error) error {
	switch err {
	case constant.ErrorMessageBadRequest:
		return c.Status(fiber.StatusBadRequest).JSON(BuildResponse(ResponseBadRequest))
	case constant.ErrorMessageUserAlreadyHasMerchant:
		return c.Status(fiber.StatusBadRequest).JSON(BuildResponse(ResponseTemplate{Code: 400, Message: err.Error()}))
	case constant.ErrorMessageNotFound:
		return c.Status(fiber.StatusNotFound).JSON(BuildResponse(ResponseNotFound))

	// Error messages for authentication operations
	case constant.ErrorMessageInvalidToken,
		constant.ErrorMessageMissingOrMalformedJWT,
		constant.ErrorMessageTokenExpired,
		constant.ErrorMessageInvalidCredential:
		return c.Status(fiber.StatusUnauthorized).JSON(BuildResponse(ResponseUnauthorized, err.Error()))
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(BuildResponse(ResponseInternalServerError, err.Error()))
	}
}

func HandleSuccess(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(BuildResponse(ResponseSuccess, data))
}
