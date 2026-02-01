package http

import (
	"github.com/dukk308/beetool.dev-go-starter/internal/modules/note/application"
	"github.com/gin-gonic/gin"
)

type Http struct {
	createNoteCommand *application.CreateNoteCommand
	getNoteQuery      *application.GetNoteQuery
	listNotesQuery    *application.ListNotesQuery
	updateNoteCommand *application.UpdateNoteCommand
	deleteNoteCommand *application.DeleteNoteCommand
}

func NewHttp(
	createNoteCommand *application.CreateNoteCommand,
	getNoteQuery *application.GetNoteQuery,
	listNotesQuery *application.ListNotesQuery,
	updateNoteCommand *application.UpdateNoteCommand,
	deleteNoteCommand *application.DeleteNoteCommand,
) *Http {
	return &Http{
		createNoteCommand: createNoteCommand,
		getNoteQuery:      getNoteQuery,
		listNotesQuery:    listNotesQuery,
		updateNoteCommand: updateNoteCommand,
		deleteNoteCommand: deleteNoteCommand,
	}
}

func (h *Http) RegisterRoutes(router *gin.RouterGroup) {
	notesGroup := router.Group("/v1/notes")
	{
		notesGroup.POST("", h.HandlerCreateNote())
		notesGroup.GET("", h.HandlerListNotes())
		notesGroup.GET("/:id", h.HandlerGetNoteByID())
		notesGroup.GET("/slug/:slug", h.HandlerGetNoteBySlug())
		notesGroup.PUT("/:id", h.HandlerUpdateNote())
		notesGroup.DELETE("/:id", h.HandlerDeleteNote())
	}
}
