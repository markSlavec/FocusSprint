package handler

import "focussprint/internal/matter/dto"

type MatterHandler struct {
	matterUsecase MatterUsecase
}

type MatterUsecase interface {
	GetMatter(id int) (*dto.MatterDTO, error)
	UpdateMatter(idMatter int, newMatter *dto.MatterDTO) (id int, err error)
	DeleteMatter(id int) error
	CreateMatter(newMatter *dto.MatterDTO) (id int, err error)
}

func NewMatterHandler(m MatterUsecase) *MatterHandler {
	return &MatterHandler{matterUsecase: m}
}

func (m *MatterHandler) GetMatter(id int) (*dto.MatterDTO, error) {
	matter, err := m.matterUsecase.GetMatter(id)
	if err != nil {
		return nil, err
	}
	return matter, nil
}

func (m *MatterHandler) UpdateMatter(idMatter int, newMatter *dto.MatterDTO) (id int, err error) {
	id, err = m.matterUsecase.UpdateMatter(idMatter, newMatter)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *MatterHandler) CreateMatter(newMatter *dto.MatterDTO) (id int, err error) {
	id, err = m.matterUsecase.CreateMatter(newMatter)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *MatterHandler) DeleteMatter(id int) error {
	return m.matterUsecase.DeleteMatter(id)
}
