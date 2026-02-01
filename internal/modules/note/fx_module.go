package note

import (
	"github.com/dukk308/beetool.dev-go-starter/internal/modules/note/application"
	"github.com/dukk308/beetool.dev-go-starter/internal/modules/note/domain"
	"github.com/dukk308/beetool.dev-go-starter/internal/modules/note/infrastructure/persistence"
	note_http "github.com/dukk308/beetool.dev-go-starter/internal/modules/note/presentation/http"
	"go.uber.org/fx"
)

var Module = fx.Module("note",
	fx.Provide(
		fx.Annotate(
			persistence.NewNoteRepository,
			fx.As(new(domain.INoteRepository)),
		),
	),
	fx.Provide(application.NewCreateNoteCommand),
	fx.Provide(application.NewGetNoteQuery),
	fx.Provide(application.NewListNotesQuery),
	fx.Provide(application.NewUpdateNoteCommand),
	fx.Provide(application.NewDeleteNoteCommand),
	fx.Provide(
		func(
			createNoteCommand *application.CreateNoteCommand,
			getNoteQuery *application.GetNoteQuery,
			listNotesQuery *application.ListNotesQuery,
			updateNoteCommand *application.UpdateNoteCommand,
			deleteNoteCommand *application.DeleteNoteCommand,
		) *note_http.Http {
			return note_http.NewHttp(
				createNoteCommand,
				getNoteQuery,
				listNotesQuery,
				updateNoteCommand,
				deleteNoteCommand,
			)
		},
	),
)
