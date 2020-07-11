package userresource

import (
	"Food/dto"
	"Food/dto/request"
	RequestUser "Food/dto/request/user"
	"Food/dto/response"
	"Food/dto/response/user"
	"Food/helpers/e"
	"Food/helpers/logging"
	"Food/helpers/util"
	"Food/service"

	"github.com/gin-gonic/gin"
)

// Register register
// @Summary Register
// @Accept  json
// @Produce  json
// @.....Param body body RequestUser.RegisterDTO true "RegisterDTO"
// @.....Success 200 {object} response.APIResponseDTO{data=nil} "desc"
// @Router /api/public/user/register [post]
func Register(c *gin.Context) {
	var registerDTO RequestUser.RegisterDTO
	errCode := request.BindAndValid(c, &registerDTO)
	if errCode != e.SUCCESS {
		response.CreateErrorResponse(c, e.GetMsg(errCode))
		return
	}

	userDTO := dto.UserDTO{Name: registerDTO.Username, Password: registerDTO.Password}

	_, isSuccess := service.CreateUser(userDTO)
	if !isSuccess {
		response.CreateErrorResponse(c, "Register failed!!!")
		return
	}

	response.CreateSuccesResponse(c, nil)
}

func Login(c *gin.Context) {
	var loginDTO RequestUser.LoginDTO
	errCode := request.BindAndValid(c, &loginDTO)
	if errCode != e.SUCCESS {
		response.CreateErrorResponse(c, e.GetMsg(errCode))
		return
	}

	userDTO, isSuccess := service.FindOneUserLogin(loginDTO.Username, loginDTO.Password)
	if !isSuccess {
		response.CreateErrorResponse(c, "UNAUTHORIZED")
		return
	}

	tokenString, error := util.GenerateToken(userDTO.ID, userDTO.Name, userDTO.Password)
	if error != nil {
		logging.Error("Signed token error: ", error)
	}

	response.CreateSuccesResponse(c, user.LoginResponseDTO{
		Token: tokenString,
	})
	return;
}
