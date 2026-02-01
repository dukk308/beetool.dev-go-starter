package http

import (
	"github.com/dukk308/beetool.dev-go-starter/pkgs/components/gin_comp"
	"github.com/gin-gonic/gin"
)

func (h *Http) HandlerGetNoteByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		ctx := c.Request.Context()
		response, err := h.getNoteQuery.ExecuteByID(ctx, id)
		if err != nil {
			gin_comp.ResponseError(c, err)
			return
		}
		gin_comp.ResponseSuccess(c, response)
	}
}

func (h *Http) HandlerGetNoteBySlug() gin.HandlerFunc {
	return func(c *gin.Context) {
		slug := c.Param("slug")
		ctx := c.Request.Context()
		response, err := h.getNoteQuery.ExecuteBySlug(ctx, slug)
		if err != nil {
			gin_comp.ResponseError(c, err)
			return
		}
		gin_comp.ResponseSuccess(c, response)
	}
}
