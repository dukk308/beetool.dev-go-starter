package persistence

import (
	"context"

	"gorm.io/gorm"

	"github.com/dukk308/beetool.dev-go-starter/internal/modules/note/domain"
)

type NoteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) domain.INoteRepository {
	return &NoteRepository{
		db: db,
	}
}

func (r *NoteRepository) Create(ctx context.Context, note *domain.Note) error {
	sqlNote := &SQLNote{}
	sqlNote.FromDomain(note)
	return r.db.WithContext(ctx).Create(sqlNote).Error
}

func (r *NoteRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&SQLNote{}).Error
}

func (r *NoteRepository) GetAll(ctx context.Context) ([]*domain.Note, error) {
	var sqlNotes []SQLNote
	if err := r.db.WithContext(ctx).Find(&sqlNotes).Error; err != nil {
		return nil, err
	}
	notes := make([]*domain.Note, len(sqlNotes))
	for i, n := range sqlNotes {
		notes[i] = n.ToDomain()
	}
	return notes, nil
}

func (r *NoteRepository) GetBySlug(ctx context.Context, slug string) (*domain.Note, error) {
	var sqlNote SQLNote
	if err := r.db.WithContext(ctx).Where("slug = ?", slug).First(&sqlNote).Error; err != nil {
		return nil, err
	}
	return sqlNote.ToDomain(), nil
}

func (r *NoteRepository) GetByID(ctx context.Context, id string) (*domain.Note, error) {
	var sqlNote SQLNote
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&sqlNote).Error; err != nil {
		return nil, err
	}
	return sqlNote.ToDomain(), nil
}

func (r *NoteRepository) Update(ctx context.Context, note *domain.Note) error {
	sqlNote := &SQLNote{}
	sqlNote.FromDomain(note)
	return r.db.WithContext(ctx).Save(sqlNote).Error
}
