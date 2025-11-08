package usecase

import (
	"focussprint/internal/matter/domain"
	"focussprint/internal/matter/dto"
)

// мапит matter dto matter to domain matter
func (m *MatterUsecase) mapDTOToDomain(matterDTO *dto.MatterDTO) (*domain.Matter, error) {
	matterDomain := &domain.Matter{
		ID:          matterDTO.ID,
		Name:        matterDTO.Name,
		Description: matterDTO.Description,
	}
	return matterDomain, nil
}

// мапит dto matter to domain matter
func (m *MatterUsecase) mapDomainToDTO(matter *domain.Matter) (*dto.MatterDTO, error) {
	matterDTO := &dto.MatterDTO{
		ID:          matter.ID,
		Name:        matter.Name,
		Description: matter.Description,
	}
	return matterDTO, nil
}
