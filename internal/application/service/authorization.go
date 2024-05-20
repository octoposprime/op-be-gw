package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/authorization"
)

// CheckAuth returns true if the user is authorized to perform the action.
func (a *Service) CheckAuth(ctx context.Context, authRequest *pb.AuthRequest) (*pb.AuthResponse, error) {
	return a.ServicePort.CheckAuth(ctx, authRequest)
}
