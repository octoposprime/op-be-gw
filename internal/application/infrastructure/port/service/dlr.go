package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/dlr"
)

// DlrServicePort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the other servies.
type DlrServicePort interface {
	// GetDlrsByFilter returns the dlrs that match the given filter.
	GetDlrsByFilter(ctx context.Context, dlrFilter *pb.DlrFilter) (*pb.Dlrs, error)

	// CreateDlr sends the given dlr to the service of the infrastructure layer for creatimng dlr.
	CreateDlr(ctx context.Context, dlr *pb.Dlr) (*pb.Dlr, error)

	// UpdateDlrBase sends the given dlr to the service of the infrastructure layer for updating dlr base values.
	UpdateDlrBase(ctx context.Context, dlr *pb.Dlr) (*pb.Dlr, error)

	// UpdateDlrCore sends the given dlr to the service of the infrastructure layer for updating dlr core values.
	UpdateDlrCore(ctx context.Context, dlr *pb.Dlr) (*pb.Dlr, error)

	// UpdateDlrStatus sends the given dlr to the service of the infrastructure layer for updating dlr status.
	UpdateDlrStatus(ctx context.Context, dlr *pb.Dlr) (*pb.Dlr, error)

	// DeleteDlr sends the given dlr to the service of the infrastructure layer for deleting dlr.
	DeleteDlr(ctx context.Context, dlr *pb.Dlr) (*pb.Dlr, error)
}
