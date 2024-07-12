package presentation

import (
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/book"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// PageFilterDto is a struct that provides filtering requested page data.
type PageFilterDto struct {
	ModelInputData *PageFilterInput
	PbData         *pb.PageFilter
}

// NewPageFilterDto creates a new *PageFilterDto.
func NewPageFilterDto(modelInputData *PageFilterInput) *PageFilterDto {
	return &PageFilterDto{
		ModelInputData: modelInputData,
		PbData:         new(pb.PageFilter),
	}
}

// ToPb converts *ModelInputData to *PbData and returns it.
func (u *PageFilterDto) ToPb() *pb.PageFilter {
	u.PbData.Id = u.ModelInputData.ID
	if u.ModelInputData.PageType != nil {
		enumVal := pb.PageType(PageType_value[*u.ModelInputData.PageType])
		u.PbData.PageType = &enumVal
	}
	if u.ModelInputData.PageStatus != nil {
		enumVal := pb.PageStatus(PageStatus_value[*u.ModelInputData.PageStatus])
		u.PbData.PageStatus = &enumVal
	}
	u.PbData.Tags = u.ModelInputData.Tags
	if u.ModelInputData.CreatedAtFrom != nil && u.ModelInputData.CreatedAtTo != nil {
		u.PbData.CreatedAtFrom = timestamppb.New(*u.ModelInputData.CreatedAtFrom)
		u.PbData.CreatedAtTo = timestamppb.New(*u.ModelInputData.CreatedAtTo)
	}
	if u.ModelInputData.UpdatedAtFrom != nil && u.ModelInputData.UpdatedAtTo != nil {
		u.PbData.UpdatedAtFrom = timestamppb.New(*u.ModelInputData.UpdatedAtFrom)
		u.PbData.UpdatedAtTo = timestamppb.New(*u.ModelInputData.UpdatedAtTo)
	}
	u.PbData.SearchText = u.ModelInputData.SearchText
	u.PbData.SortType = u.ModelInputData.SortType
	if u.ModelInputData.SortField != nil {
		enumVal := pb.PageSortField(PageSortField_value[*u.ModelInputData.SortField])
		u.PbData.SortField = &enumVal
	}
	u.ModelInputData.Pagination = u.ModelInputData.Pagination.Validate()
	u.PbData.Limit = u.ModelInputData.Pagination.Limit
	u.PbData.Offset = u.ModelInputData.Pagination.Offset
	return u.PbData
}
