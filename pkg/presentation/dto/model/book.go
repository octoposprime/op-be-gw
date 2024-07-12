package presentation

import (
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/book"
)

// PageDto is a struct that represents the dto of a page basic values.
type PageDto struct {
	ModelData      *Page
	ModelInputData *PageInput
	PbData         *pb.Page
}

// NewPageDto creates a new *PageDto.
func NewPageDto(modelInputData *PageInput) *PageDto {
	pbData := new(pb.Page)
	return &PageDto{
		ModelData:      new(Page),
		ModelInputData: modelInputData,
		PbData:         pbData,
	}
}

// ToPb converts *ModelInputData to *PbData and returns it.
func (u *PageDto) ToPb() *pb.Page {
	if u.ModelInputData.ID != nil {
		u.PbData.Id = *u.ModelInputData.ID
	}
	if u.ModelInputData.PageCore != nil {
		if u.ModelInputData.PageCore.PageData != nil {
			u.PbData.PageData = *u.ModelInputData.PageCore.PageData
		}
	}
	if u.ModelInputData.PageType != nil {
		u.PbData.PageType = pb.PageType(PageType_value[*u.ModelInputData.PageType])
	}
	if u.ModelInputData.PageStatus != nil {
		u.PbData.PageStatus = pb.PageStatus(PageStatus_value[*u.ModelInputData.PageStatus])
	}
	if u.ModelInputData.PageBase != nil {
		if u.ModelInputData.PageBase.Tags != nil {
			u.PbData.Tags = u.ModelInputData.PageBase.Tags
		}
	}
	return u.PbData
}

// ToModel converts *PbData to *ModelData and returns it.
func (u *PageDto) ToModel() *Page {
	u.ModelData.PageBase = new(PageBase)
	u.ModelData.PageCore = new(PageCore)
	u.ModelData.ID = u.PbData.Id
	u.ModelData.PageCore.PageData = u.PbData.PageData
	u.ModelData.PageType = PageType(PageType_name[int32(u.PbData.PageType)])
	u.ModelData.PageStatus = PageStatus(PageStatus_name[int32(u.PbData.PageStatus)])
	u.ModelData.PageBase.Tags = u.PbData.Tags

	// Only for view
	u.ModelData.CreatedAt = u.PbData.CreatedAt.AsTime()
	u.ModelData.UpdatedAt = u.PbData.UpdatedAt.AsTime()
	return u.ModelData
}
