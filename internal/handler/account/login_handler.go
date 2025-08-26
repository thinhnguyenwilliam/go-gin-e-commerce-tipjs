// internal/handler/account/login_handler.go
package account

import (
	"github.com/gin-gonic/gin"
	"github.com/thinhcompany/ecommerce-ver-2/internal/service"
	"github.com/thinhcompany/ecommerce-ver-2/pkg/response"
)

// managemment handler login user
var Login = new(cUserLogin)

type cUserLogin struct{}

func (c *cUserLogin) Login(ctx *gin.Context) {
	err := service.UserLogin().Login(ctx)
	if err != nil {
		res := response.ErrorResponse(response.ErrorCodeParamInvalid, nil)
		ctx.JSON(400, res) // <-- write JSON response
		return
	}

	res := response.SuccessResponse("Login success")
	ctx.JSON(200, res) // <-- write JSON response
}
