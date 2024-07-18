package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/dlr"
)

// GetDlrsByFilter returns the dlrs that match the given filter.
func (a *Service) GetDlrsByFilter(ctx context.Context, dlrFilter *pb.DlrFilter) (*pb.Dlrs, error) {
	return a.ServicePort.GetDlrsByFilter(ctx, dlrFilter)
}

// CreateDlr sends the given dlr to the service of the infrastructure layer for creating a new dlr.
func (a *Service) CreateDlr(ctx context.Context, dlr *pb.Dlr) (*pb.Dlr, error) {
	return a.ServicePort.CreateDlr(ctx, dlr)
}

// UpdateDlrBase sends the given dlr to the service of the infrastructure layer for updating dlr base values.
func (a *Service) UpdateDlrBase(ctx context.Context, dlr *pb.Dlr) (*pb.Dlr, error) {
	return a.ServicePort.UpdateDlrBase(ctx, dlr)
}

// UpdateDlrCore sends the given dlr to the service of the infrastructure layer for updating dlr core values.
func (a *Service) UpdateDlrCore(ctx context.Context, dlr *pb.Dlr) (*pb.Dlr, error) {
	return a.ServicePort.UpdateDlrCore(ctx, dlr)
}


// UpdateDlrStatus sends the given dlr to the service of the infrastructure layer for updating dlr status.
func (a *Service) UpdateDlrStatus(ctx context.Context, dlr *pb.Dlr) (*pb.Dlr, error) {
	return a.ServicePort.UpdateDlrStatus(ctx, dlr)
}

// DeleteDlr sends the given dlr to the service of the infrastructure layer for deleting data.
func (a *Service) DeleteDlr(ctx context.Context, dlr *pb.Dlr) (*pb.Dlr, error) {
	return a.ServicePort.DeleteDlr(ctx, dlr)
}