package usecase

import "awesomeProjectADV/internal/domain/model"

type SessionUsecase interface {
	CreateSession(*model.Session) error
	GetAllSessions() ([]*model.Session, error)
	GetSessionByID(string) (*model.Session, error)
}

type sessionUsecase struct {
	repo interface {
		Create(*model.Session) error
		GetAll() ([]*model.Session, error)
		GetByID(string) (*model.Session, error)
	}
}

func NewSessionUsecase(repo interface {
	Create(*model.Session) error
	GetAll() ([]*model.Session, error)
	GetByID(string) (*model.Session, error)
}) SessionUsecase {
	return &sessionUsecase{repo: repo}
}

func (u *sessionUsecase) CreateSession(s *model.Session) error {
	return u.repo.Create(s)
}

func (u *sessionUsecase) GetAllSessions() ([]*model.Session, error) {
	return u.repo.GetAll()
}

func (u *sessionUsecase) GetSessionByID(id string) (*model.Session, error) {
	return u.repo.GetByID(id)
}
