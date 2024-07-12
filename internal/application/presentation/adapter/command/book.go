package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/book"
)

// CreatePage sends the given page to the application layer for creating a new page.
func (a CommandAdapter) CreatePage(ctx context.Context, page *pb.Page) (*pb.Page, error) {
	return a.Service.CreatePage(ctx, page)
}

// UpdatePageBase sends the given page to the application layer for updating page values.
func (a CommandAdapter) UpdatePageBase(ctx context.Context, page *pb.Page) (*pb.Page, error) {
	return a.Service.UpdatePageBase(ctx, page)
}

// UpdatePageCore sends the given page to the application layer for updating page values.
func (a CommandAdapter) UpdatePageCore(ctx context.Context, page *pb.Page) (*pb.Page, error) {
	return a.Service.UpdatePageCore(ctx, page)
}

// UpdatePageStatus sends the given page to the application layer for updating page status.
func (a CommandAdapter) UpdatePageStatus(ctx context.Context, page *pb.Page) (*pb.Page, error) {
	return a.Service.UpdatePageStatus(ctx, page)
}

// DeletePage sends the given page to the application layer for deleting data.
func (a CommandAdapter) DeletePage(ctx context.Context, page *pb.Page) (*pb.Page, error) {
	return a.Service.DeletePage(ctx, page)
}
