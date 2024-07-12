package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/book"
)

// GetPagesByFilter returns the pages that match the given filter.
func (a QueryAdapter) GetPagesByFilter(ctx context.Context, pageFilter *pb.PageFilter) (*pb.Pages, error) {
	return a.Service.GetPagesByFilter(ctx, pageFilter)
}
