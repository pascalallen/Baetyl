package Auth

import (
	"github.com/gin-gonic/gin"
	"github.com/pascalallen/Baetyl/src/Adapter/Http/Responder/JSend"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/User"
	"net/http"
)

func InternalServerErrorResponse(c *gin.Context, error error) {
	c.JSON(
		http.StatusInternalServerError,
		JSend.ErrorResponse[string]{
			Status:  "error",
			Message: error.Error(),
		},
	)

	return
}

func UnprocessableEntityResponse(c *gin.Context, error error) {
	c.JSON(
		http.StatusUnprocessableEntity,
		JSend.ErrorResponse[string]{
			Status:  "error",
			Message: error.Error(),
		},
	)

	return
}

func BadRequestResponse(c *gin.Context, error error) {
	c.JSON(
		http.StatusBadRequest,
		JSend.FailResponse[string]{
			Status: "fail",
			Data:   error.Error(),
		},
	)

	return
}

func CreatedResponse(c *gin.Context, user *User.User) {
	c.JSON(
		http.StatusCreated,
		JSend.SuccessResponse[User.User]{
			Status: "success",
			Data:   *user,
		},
	)

	return
}
