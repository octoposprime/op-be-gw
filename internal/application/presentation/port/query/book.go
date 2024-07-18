package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/book"
)

// QueryPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the application layer.
type BookQueryPort interface {
	// GetBooksByFilter returns the dlrs that match the given filter.
	GetBooksByFilter(ctx context.Context, bookFilter *pb.PageFilter) (*pb.Pages, error)
}
