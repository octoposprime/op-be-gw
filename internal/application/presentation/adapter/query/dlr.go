package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/dlr"
)

// GetDlrsByFilter returns the dlrs that match the given filter.
func (a QueryAdapter) GetDlrsByFilter(ctx context.Context, dlrFilter *pb.DlrFilter) (*pb.Dlrs, error) {
	return a.Service.GetDlrsByFilter(ctx, dlrFilter)
}
