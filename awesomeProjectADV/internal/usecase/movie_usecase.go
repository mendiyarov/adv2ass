package usecase

import (
	"awesomeProjectADV/internal/domain/model"
)

type MovieUsecase interface {
	CreateMovie(movie *model.Movie) error
	GetAllMovies() ([]*model.Movie, error)
	GetMovieByID(id string) (*model.Movie, error)
	Delete(string) error // ✅ добавить сюда!

}

type movieUsecase struct {
	repo interface {
		Create(*model.Movie) error
		GetAll() ([]*model.Movie, error)
		GetByID(string) (*model.Movie, error)
		Delete(string) error
	}
}

func NewMovieUsecase(repo interface {
	Create(*model.Movie) error
	GetAll() ([]*model.Movie, error)
	GetByID(string) (*model.Movie, error)
	Delete(string) error // ✅ добавляем сюда
}) MovieUsecase {
	return &movieUsecase{repo: repo}
}

func (u *movieUsecase) CreateMovie(movie *model.Movie) error {
	return u.repo.Create(movie)
}

func (u *movieUsecase) GetAllMovies() ([]*model.Movie, error) {
	return u.repo.GetAll()
}

func (u *movieUsecase) GetMovieByID(id string) (*model.Movie, error) {
	return u.repo.GetByID(id)
}
func (u *movieUsecase) Delete(id string) error {
	return u.repo.Delete(id)
}
