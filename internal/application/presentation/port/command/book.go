package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/dlr"
)

// DlrCommandPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the application layer.
type DlrCommandPort interface {
	// CreateDlr sends the given dlr to the application layer for creating a new dlr.
	CreateDlr(ctx context.Context, dlr *pb.Dlr) (*pb.Dlr, error)

	// UpdateDlrBase sends the given dlr to the application layer for updating dlr base values.
	UpdateDlrBase(ctx context.Context, dlr *pb.Dlr) (*pb.Dlr, error)

	// UpdateDlrCore sends the given dlr to the application layer for updating dlr core values.
	UpdateDlrCore(ctx context.Context, dlr *pb.Dlr) (*pb.Dlr, error)

	// UpdateDlrStatus sends the given dlr to the application layer for updating dlr status.
	UpdateDlrStatus(ctx context.Context, dlr *pb.Dlr) (*pb.Dlr, error)

	// DeleteDlr sends the given dlr to the application layer for deleting data.
	DeleteDlr(ctx context.Context, dlr *pb.Dlr) (*pb.Dlr, error)
}
