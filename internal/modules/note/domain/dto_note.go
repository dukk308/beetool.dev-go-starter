package domain

import "time"

type DTOCreateNote struct {
	Title   string `json:"title" binding:"required"`
	Slug    string `json:"slug" binding:"required"`
	Content string `json:"content"`
}

type DTONoteResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	Content   string    `json:"content"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func NewDTONoteResponse(note *Note) *DTONoteResponse {
	return &DTONoteResponse{
		ID:        note.ID.String(),
		Title:     note.Title,
		Slug:      note.Slug,
		Content:   note.Content,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}
}
