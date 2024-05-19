package application

// AuthorizationCommandPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the application layer.
type AuthorizationCommandPort interface {
	// CheckAuth returns true if the user is authorized to perform the action.
	//CheckAuth(ctx context.Context, authRequest *pb.AuthRequest) (*pb.AuthResponse, error)
}
