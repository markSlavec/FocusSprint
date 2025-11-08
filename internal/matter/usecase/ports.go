package usecase

import "focussprint/internal/matter/domain"

type MatterInfra interface {
	GetMatter(id int) (*domain.Matter, error)
	UpdateMatter(idMatter int, newMatter *domain.Matter) (id int, err error)
	DeleteMatter(id int) error
	CreateMatter(newMatter *domain.Matter) (id int, err error)
}
