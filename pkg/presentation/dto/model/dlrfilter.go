package presentation

import (
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/dlr"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// DlrFilterDto is a struct that provides filtering requested dlr data.
type DlrFilterDto struct {
	ModelInputData *DlrFilterInput
	PbData         *pb.DlrFilter
}

// NewDlrFilterDto creates a new *DlrFilterDto.
func NewDlrFilterDto(modelInputData *DlrFilterInput) *DlrFilterDto {
	return &DlrFilterDto{
		ModelInputData: modelInputData,
		PbData:         new(pb.DlrFilter),
	}
}

// ToPb converts *ModelInputData to *PbData and returns it.
func (u *DlrFilterDto) ToPb() *pb.DlrFilter {
	u.PbData.Id = u.ModelInputData.ID
	if u.ModelInputData.DlrType != nil {
		enumVal := pb.DlrType(DlrType_value[*u.ModelInputData.DlrType])
		u.PbData.DlrType = &enumVal
	}
	if u.ModelInputData.DlrStatus != nil {
		enumVal := pb.DlrStatus(DlrStatus_value[*u.ModelInputData.DlrStatus])
		u.PbData.DlrStatus = &enumVal
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
		enumVal := pb.DlrSortField(DlrSortField_value[*u.ModelInputData.SortField])
		u.PbData.SortField = &enumVal
	}
	u.ModelInputData.Pagination = u.ModelInputData.Pagination.Validate()
	u.PbData.Limit = u.ModelInputData.Pagination.Limit
	u.PbData.Offset = u.ModelInputData.Pagination.Offset
	return u.PbData
}
