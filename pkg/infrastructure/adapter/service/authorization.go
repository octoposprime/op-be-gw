package infrastructure

import (
	"context"

	me "github.com/octoposprime/op-be-graphql/internal/domain/model/entity"
	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/authorization"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
	tconfig "github.com/octoposprime/op-be-shared/tool/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// CheckAuth returns true if the user is authorized to perform the action.
func (a ServiceAdapter) CheckAuth(ctx context.Context, authRequest *pb.AuthRequest) (*pb.AuthResponse, error) {
	conn, err := grpc.Dial(tconfig.GetInternalConfigInstance().Grpc.AuthorizationHost+":"+tconfig.GetInternalConfigInstance().Grpc.AuthorizationPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Login", userId, err.Error()))
		return &pb.AuthResponse{IsAuthorized: false}, err
	}

	pbResult, err := pb.NewAuthorizationSvcClient(conn).CheckAuth(ctx, authRequest)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Login", userId, err.Error()))
		return &pb.AuthResponse{IsAuthorized: false}, err
	}
	return pbResult, nil
}
