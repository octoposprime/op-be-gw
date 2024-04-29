package infrastructure

import (
	"context"

	me "github.com/octoposprime/op-be-graphql/internal/domain/model/entity"
	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/dlr"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
	tconfig "github.com/octoposprime/op-be-shared/tool/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// GetDlrsByFilter returns the dlrs that match the given filter.
func (a ServiceAdapter) GetDlrsByFilter(ctx context.Context, dlrFilter *pb.DlrFilter) (*pb.Dlrs, error) {
	conn, err := grpc.Dial(tconfig.GetInternalConfigInstance().Grpc.DlrHost+":"+tconfig.GetInternalConfigInstance().Grpc.DlrPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "GetDlrsByFilter", userId, err.Error()))
		return &pb.Dlrs{}, err
	}

	userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	ctx = metadata.AppendToOutgoingContext(ctx, string(smodel.QueryKeyUid), userId)
	pbResults, err := pb.NewDlrSvcClient(conn).GetDlrsByFilter(ctx, dlrFilter)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "GetDlrsByFilter", userId, err.Error()))
		return &pb.Dlrs{}, err
	}
	return pbResults, nil
}

// CreateDlr insert a new dlr in the databse.
func (a ServiceAdapter) CreateDlr(ctx context.Context, dlr *pb.Dlr) (*pb.Dlr, error) {
	conn, err := grpc.Dial(tconfig.GetInternalConfigInstance().Grpc.DlrHost+":"+tconfig.GetInternalConfigInstance().Grpc.DlrPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "CreateDlr", userId, err.Error()))
		return &pb.Dlr{}, err
	}

	userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	ctx = metadata.AppendToOutgoingContext(ctx, string(smodel.QueryKeyUid), userId)
	pbResult, err := pb.NewDlrSvcClient(conn).CreateDlr(ctx, dlr)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "CreateDlr", userId, err.Error()))
		return &pb.Dlr{}, err
	}
	return pbResult, nil
}

// UpdateDlrBase sends the given dlr to the application layer for updating dlr base values.
func (a ServiceAdapter) UpdateDlrBase(ctx context.Context, dlr *pb.Dlr) (*pb.Dlr, error) {
	conn, err := grpc.Dial(tconfig.GetInternalConfigInstance().Grpc.DlrHost+":"+tconfig.GetInternalConfigInstance().Grpc.DlrPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateDlrStatus", userId, err.Error()))
		return &pb.Dlr{}, err
	}

	userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	ctx = metadata.AppendToOutgoingContext(ctx, string(smodel.QueryKeyUid), userId)
	pbResult, err := pb.NewDlrSvcClient(conn).UpdateDlrBase(ctx, dlr)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateDlrStatus", userId, err.Error()))
		return &pb.Dlr{}, err
	}
	return pbResult, nil
}

// UpdateDlrCore sends the given dlr to the application layer for updating dlr core values.
func (a ServiceAdapter) UpdateDlrCore(ctx context.Context, dlr *pb.Dlr) (*pb.Dlr, error) {
	conn, err := grpc.Dial(tconfig.GetInternalConfigInstance().Grpc.DlrHost+":"+tconfig.GetInternalConfigInstance().Grpc.DlrPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateDlrStatus", userId, err.Error()))
		return &pb.Dlr{}, err
	}

	userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	ctx = metadata.AppendToOutgoingContext(ctx, string(smodel.QueryKeyUid), userId)
	pbResult, err := pb.NewDlrSvcClient(conn).UpdateDlrCore(ctx, dlr)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateDlrStatus", userId, err.Error()))
		return &pb.Dlr{}, err
	}
	return pbResult, nil
}

// UpdateDlrStatus sends the given dlr to the application layer for updating dlr status.
func (a ServiceAdapter) UpdateDlrStatus(ctx context.Context, dlr *pb.Dlr) (*pb.Dlr, error) {
	conn, err := grpc.Dial(tconfig.GetInternalConfigInstance().Grpc.DlrHost+":"+tconfig.GetInternalConfigInstance().Grpc.DlrPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateDlrStatus", userId, err.Error()))
		return &pb.Dlr{}, err
	}

	userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	ctx = metadata.AppendToOutgoingContext(ctx, string(smodel.QueryKeyUid), userId)
	pbResult, err := pb.NewDlrSvcClient(conn).UpdateDlrStatus(ctx, dlr)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateDlrStatus", userId, err.Error()))
		return &pb.Dlr{}, err
	}
	return pbResult, nil
}

// DeleteDlr soft-deletes the given dlr in the database.
func (a ServiceAdapter) DeleteDlr(ctx context.Context, dlr *pb.Dlr) (*pb.Dlr, error) {
	conn, err := grpc.Dial(tconfig.GetInternalConfigInstance().Grpc.DlrHost+":"+tconfig.GetInternalConfigInstance().Grpc.DlrPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "DeleteDlr", userId, err.Error()))
		return &pb.Dlr{}, err
	}

	userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	ctx = metadata.AppendToOutgoingContext(ctx, string(smodel.QueryKeyUid), userId)
	pbResult, err := pb.NewDlrSvcClient(conn).DeleteDlr(ctx, dlr)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "DeleteDlr", userId, err.Error()))
		return &pb.Dlr{}, err
	}
	return pbResult, nil
}
