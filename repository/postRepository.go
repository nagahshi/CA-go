package repository

import (
	"../entity"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	Delete(post *entity.Post) error
}

// NewPostRepository
func NewPostRepository() PostRepository {
	return &repo{}
}
