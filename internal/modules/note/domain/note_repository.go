package domain

import "context"

type INoteRepository interface {
	GetByID(ctx context.Context, id string) (*Note, error)
	GetBySlug(ctx context.Context, slug string) (*Note, error)
	GetAll(ctx context.Context) ([]*Note, error)
	Create(ctx context.Context, note *Note) error
	Update(ctx context.Context, note *Note) error
	Delete(ctx context.Context, id string) error
}
