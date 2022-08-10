package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Get-Service_1 godoc
// @ID write-posts
// @Router /v1/api [GET]
// @Summary get and write posts ti DB
// @Description get and write
// @Tags service 1
// @Accept json
// @Produce json
// @Success 200 {object} models.ResponseModel{data=service_1.Status} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllWrite(c *gin.Context) {

	resp, err := h.services.Service_1().GetAll(c.Request.Context(), &emptypb.Empty{})
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error getting and writing all posts ", err)
		return
	}
	fmt.Println(resp)

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}
