package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/book"
)

// GetPagesByFilter returns the pages that match the given filter.
func (a *Service) GetPagesByFilter(ctx context.Context, pageFilter *pb.PageFilter) (*pb.Pages, error) {
	return a.ServicePort.GetPagesByFilter(ctx, pageFilter)
}

// CreatePage sends the given page to the service of the infrastructure layer for creating a new page.
func (a *Service) CreatePage(ctx context.Context, page *pb.Page) (*pb.Page, error) {
	return a.ServicePort.CreatePage(ctx, page)
}

// UpdatePageBase sends the given page to the service of the infrastructure layer for updating page base values.
func (a *Service) UpdatePageBase(ctx context.Context, page *pb.Page) (*pb.Page, error) {
	return a.ServicePort.UpdatePageBase(ctx, page)
}

// UpdatePageCore sends the given page to the service of the infrastructure layer for updating page core values.
func (a *Service) UpdatePageCore(ctx context.Context, page *pb.Page) (*pb.Page, error) {
	return a.ServicePort.UpdatePageCore(ctx, page)
}

// UpdatePageStatus sends the given page to the service of the infrastructure layer for updating page status.
func (a *Service) UpdatePageStatus(ctx context.Context, page *pb.Page) (*pb.Page, error) {
	return a.ServicePort.UpdatePageStatus(ctx, page)
}

// DeletePage sends the given page to the service of the infrastructure layer for deleting data.
func (a *Service) DeletePage(ctx context.Context, page *pb.Page) (*pb.Page, error) {
	return a.ServicePort.DeletePage(ctx, page)
}
