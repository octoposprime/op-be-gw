package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/book"
)

// QueryPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the application layer.
type BookQueryPort interface {
	// GetPagesByFilter returns the pages that match the given filter.
	GetPagesByFilter(ctx context.Context, pageFilter *pb.PageFilter) (*pb.Pages, error)
}
