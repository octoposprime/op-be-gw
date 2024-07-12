package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/book"
)

// BookCommandPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the application layer.
type BookCommandPort interface {
	// CreatePage sends the given page to the application layer for creating a new page.
	CreatePage(ctx context.Context, page *pb.Page) (*pb.Page, error)

	// UpdatePageBase sends the given page to the application layer for updating page base values.
	UpdatePageBase(ctx context.Context, page *pb.Page) (*pb.Page, error)

	// UpdatePageCore sends the given page to the application layer for updating page core values.
	UpdatePageCore(ctx context.Context, page *pb.Page) (*pb.Page, error)

	// UpdatePageStatus sends the given page to the application layer for updating page status.
	UpdatePageStatus(ctx context.Context, page *pb.Page) (*pb.Page, error)

	// DeletePage sends the given page to the application layer for deleting data.
	DeletePage(ctx context.Context, page *pb.Page) (*pb.Page, error)
}
