package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/authorization"
)

// AuthorizationServicePort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the other servies.
type AuthorizationServicePort interface {
	// CheckAuth returns true if the user is authorized to perform the action.
	CheckAuth(ctx context.Context, authRequest *pb.AuthRequest) (*pb.AuthResponse, error)
}
