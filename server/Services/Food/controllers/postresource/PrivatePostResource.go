package PostResource

import (
	"Food/dto/response"
	"Food/dto/response/post"
	"Food/helpers/pagination"
	"Food/service"

	"github.com/gin-gonic/gin"
)

// Post create
// @Summary CreatePost
// @Tags PrivatePost
// @Accept json
// @Security ApiKeyAuth
// @Success 200 {object} response.APIResponseDTO{data=post.PostListResponseDTO} "desc"
// @Router /api/private/post [post]
func CreatePost(c *gin.Context) {
	pageable := pagination.GetPage(c)

	page := service.FindPagePost(pageable)

	response.CreateSuccesResponse(c, post.CreatePostListResponseDTOFromPage(page))
}