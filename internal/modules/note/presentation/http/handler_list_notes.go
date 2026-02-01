package http

import (
	"github.com/dukk308/beetool.dev-go-starter/pkgs/components/gin_comp"
	"github.com/gin-gonic/gin"
)

func (h *Http) HandlerListNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		response, err := h.listNotesQuery.Execute(ctx)
		if err != nil {
			gin_comp.ResponseError(c, err)
			return
		}
		gin_comp.ResponseSuccess(c, response)
	}
}
