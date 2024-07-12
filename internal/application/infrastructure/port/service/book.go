package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/book"
)

// PageServicePort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the other servies.
type BookServicePort interface {
	// GetPagesByFilter returns the pages that match the given filter.
	GetPagesByFilter(ctx context.Context, pageFilter *pb.PageFilter) (*pb.Pages, error)

	// CreatePage sends the given page to the service of the infrastructure layer for creatimng page.
	CreatePage(ctx context.Context, page *pb.Page) (*pb.Page, error)

	// UpdatePageBase sends the given page to the service of the infrastructure layer for updating page base values.
	UpdatePageBase(ctx context.Context, page *pb.Page) (*pb.Page, error)

	// UpdatePageCore sends the given page to the service of the infrastructure layer for updating page core values.
	UpdatePageCore(ctx context.Context, page *pb.Page) (*pb.Page, error)

	// UpdatePageStatus sends the given page to the service of the infrastructure layer for updating page status.
	UpdatePageStatus(ctx context.Context, page *pb.Page) (*pb.Page, error)

	// DeletePage sends the given page to the service of the infrastructure layer for deleting page.
	DeletePage(ctx context.Context, page *pb.Page) (*pb.Page, error)
}
