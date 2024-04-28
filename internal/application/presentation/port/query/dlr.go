package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/dlr"
)

// QueryPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the application layer.
type DlrQueryPort interface {
	// GetDlrsByFilter returns the dlrs that match the given filter.
	GetDlrsByFilter(ctx context.Context, dlrFilter *pb.DlrFilter) (*pb.Dlrs, error)
}
