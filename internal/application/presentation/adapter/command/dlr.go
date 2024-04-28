package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/dlr"
)

// CreateDlr sends the given dlr to the application layer for creating a new dlr.
func (a CommandAdapter) CreateDlr(ctx context.Context, dlr *pb.Dlr) (*pb.Dlr, error) {
	return a.Service.CreateDlr(ctx, dlr)
}

// UpdateDlrBase sends the given dlr to the application layer for updating dlr values.
func (a CommandAdapter) UpdateDlrBase(ctx context.Context, dlr *pb.Dlr) (*pb.Dlr, error) {
	return a.Service.UpdateDlrBase(ctx, dlr)
}

// UpdateDlrCore sends the given dlr to the application layer for updating dlr values.
func (a CommandAdapter) UpdateDlrCore(ctx context.Context, dlr *pb.Dlr) (*pb.Dlr, error) {
	return a.Service.UpdateDlrCore(ctx, dlr)
}

// UpdateDlrStatus sends the given dlr to the application layer for updating dlr status.
func (a CommandAdapter) UpdateDlrStatus(ctx context.Context, dlr *pb.Dlr) (*pb.Dlr, error) {
	return a.Service.UpdateDlrStatus(ctx, dlr)
}

// DeleteDlr sends the given dlr to the application layer for deleting data.
func (a CommandAdapter) DeleteDlr(ctx context.Context, dlr *pb.Dlr) (*pb.Dlr, error) {
	return a.Service.DeleteDlr(ctx, dlr)
}
