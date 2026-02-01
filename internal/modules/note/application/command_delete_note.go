package application

import (
	"context"

	"github.com/dukk308/beetool.dev-go-starter/internal/modules/note/domain"
	"github.com/dukk308/beetool.dev-go-starter/pkgs/ddd"
)

type DeleteNoteCommand struct {
	repository domain.INoteRepository
}

func NewDeleteNoteCommand(repository domain.INoteRepository) *DeleteNoteCommand {
	return &DeleteNoteCommand{
		repository: repository,
	}
}

func (c *DeleteNoteCommand) Execute(ctx context.Context, id string) error {
	return ddd.ToDomainError(c.repository.Delete(ctx, id))
}
