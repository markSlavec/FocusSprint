package usecase

import (
	"focussprint/internal/matter/dto"
)

type MatterUsecase struct {
	matterInfra MatterInfra
}

func NewMatterUsecase(m MatterInfra) *MatterUsecase {
	return &MatterUsecase{matterInfra: m}
}

func (m *MatterUsecase) GetMatter(id int) (*dto.MatterDTO, error) {
	// получаем нашу задачу
	matter, err := m.matterInfra.GetMatter(id)
	if err != nil {
		return nil, err
	}

	// мапим в DTO
	matterDTO, err := m.mapDomainToDTO(matter)
	if err != nil {
		return nil, err
	}
	return matterDTO, nil
}

func (m *MatterUsecase) UpdateMatter(idMatter int, newMatter *dto.MatterDTO) (id int, err error) {
	matter, err := m.mapDTOToDomain(newMatter)
	if err != nil {
		return 0, err
	}
	id, err = m.matterInfra.UpdateMatter(idMatter, matter)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *MatterUsecase) DeleteMatter(id int) error {
	return m.matterInfra.DeleteMatter(id)
}

func (m *MatterUsecase) CreateMatter(newMatter *dto.MatterDTO) (id int, err error) {
	matter, err := m.mapDTOToDomain(newMatter)
	if err != nil {
		return 0, err
	}

	id, err = m.matterInfra.CreateMatter(matter)
	if err != nil {
		return 0, err
	}

	return id, nil
}
