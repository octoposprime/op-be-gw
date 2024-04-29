package presentation

import (
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/dlr"
)

// DlrDto is a struct that represents the dto of a dlr basic values.
type DlrDto struct {
	ModelData      *Dlr
	ModelInputData *DlrInput
	PbData         *pb.Dlr
}

// NewDlrDto creates a new *DlrDto.
func NewDlrDto(modelInputData *DlrInput) *DlrDto {
	pbData := new(pb.Dlr)
	return &DlrDto{
		ModelData:      new(Dlr),
		ModelInputData: modelInputData,
		PbData:         pbData,
	}
}

// ToPb converts *ModelInputData to *PbData and returns it.
func (u *DlrDto) ToPb() *pb.Dlr {
	if u.ModelInputData.ID != nil {
		u.PbData.Id = *u.ModelInputData.ID
	}
	if u.ModelInputData.DlrCore != nil {
		if u.ModelInputData.DlrCore.DlrData != nil {
			u.PbData.DlrData = *u.ModelInputData.DlrCore.DlrData
		}
	}
	if u.ModelInputData.DlrType != nil {
		u.PbData.DlrType = pb.DlrType(DlrType_value[*u.ModelInputData.DlrType])
	}
	if u.ModelInputData.DlrStatus != nil {
		u.PbData.DlrStatus = pb.DlrStatus(DlrStatus_value[*u.ModelInputData.DlrStatus])
	}
	if u.ModelInputData.DlrBase != nil {
		if u.ModelInputData.DlrBase.Tags != nil {
			u.PbData.Tags = u.ModelInputData.DlrBase.Tags
		}
	}
	return u.PbData
}

// ToModel converts *PbData to *ModelData and returns it.
func (u *DlrDto) ToModel() *Dlr {
	u.ModelData.DlrBase = new(DlrBase)
	u.ModelData.DlrCore = new(DlrCore)
	u.ModelData.ID = u.PbData.Id
	u.ModelData.DlrCore.DlrData = u.PbData.DlrData
	u.ModelData.DlrType = DlrType(DlrType_name[int32(u.PbData.DlrType)])
	u.ModelData.DlrStatus = DlrStatus(DlrStatus_name[int32(u.PbData.DlrStatus)])
	u.ModelData.DlrBase.Tags = u.PbData.Tags

	// Only for view
	u.ModelData.CreatedAt = u.PbData.CreatedAt.AsTime()
	u.ModelData.UpdatedAt = u.PbData.UpdatedAt.AsTime()
	return u.ModelData
}
