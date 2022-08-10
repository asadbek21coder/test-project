package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/asadbek21coder/test-project/gateway/genproto/service_2"
	"github.com/gin-gonic/gin"
)

// Get-Service_2 godoc
// @ID get-posts
// @Router /v1/posts [GET]
// @Summary get posts
// @Description get posts
// @Tags service 2
// @Accept json
// @Produce json
// @Param search query string false "search"
// @Param limit query integer false "limit"
// @Param offset query integer false "offset"
// @Success 200 {object} models.ResponseModel{data=service_2.GetAllResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAll(c *gin.Context) {
	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}
	fmt.Println(c.Query("search"))
	resp, err := h.services.Service_2().GetAll(
		c.Request.Context(),
		&service_2.GetAllRequest{
			Search: c.Query("search"),
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error getting all posts", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// Service_2 godoc
// @ID update-post
// @Router /v1/posts [PUT]
// @Summary update post
// @Description update post
// @Tags service 2
// @Accept json
// @Produce json
// @Param post body service_2.Post true "post"
// @Success 200 {object} models.ResponseModel{data=service_2.Post} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) Update(c *gin.Context) {
	var posts service_2.Post

	if err := c.BindJSON(&posts); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.Service_2().Update(c.Request.Context(), &posts)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while updating post", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Post godoc
// @ID delete-post
// @Router /v1/posts [DELETE]
// @Summary delete post
// @Description delete post
// @Tags service 2
// @Accept json
// @Produce json
// @Param post body service_2.Id true "post"
// @Success 200 {object} models.ResponseModel{data=service_2.Id} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) Delete(c *gin.Context) {
	var id service_2.Id

	if err := c.BindJSON(&id); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.Service_2().Delete(c.Request.Context(), &id)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while deleting post", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Post godoc
// @ID getbyid-post
// @Router /v1/posts/{id} [GET]
// @Summary getById post
// @Description GetById post
// @Tags service 2
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.ResponseModel{data=service_2.Post} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetById(c *gin.Context) {
	id, err := (strconv.Atoi(c.Param("id")))
	id1 := int32(id)

	resp, err := h.services.Service_2().GetById(
		c.Request.Context(),

		&service_2.Id{
			Id: id1,
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error getting all posts", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)

}
