package handler

//
//import "focussprint/internal/matter/dto"
//
//type SubMatterHandler struct {
//	matterUsecase MatterUsecase
//}
//
//type SubMatterUsecase interface {
//	GetSubMatter(id int) (*dto.MatterDTO, error)
//	UpdateMatter(newMatter *dto.MatterDTO) (id int, err error)
//}
//
//func NewSubMatterHandler(m MatterUsecase) *SubMatterHandler {
//	return &SubMatterHandler{}
//}
//
//func (m *SubMatterHandler) GetSubMatter(id int) (*dto.MatterDTO, error) {
//	matter, err := m.GetSubMatter(id)
//	if err != nil {
//		return nil, err
//	}
//	return matter, nil
//}
//
//func (m *SubMatterHandler) UpdateSubMatter(newMatter *dto.MatterDTO) (id int, err error) {
//	id, err = m.UpdateMatter(newMatter)
//	if err != nil {
//		return 0, err
//	}
//	return id, nil
//}
//
//func (m *SubMatterHandler) CreateSubMatter(newMatter *dto.MatterDTO) (id int, err error) {
//	id, err = m.CreateMatter(newMatter)
//	if err != nil {
//		return 0, err
//	}
//	return id, nil
//}
//
//func (m *SubMatterHandler) DeleteSubMatter(id int) error {
//	return m.DeleteMatter(id)
//
//}
