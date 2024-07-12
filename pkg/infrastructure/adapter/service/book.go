package infrastructure

import (
	"context"

	me "github.com/octoposprime/op-be-graphql/internal/domain/model/entity"
	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/book"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
	tconfig "github.com/octoposprime/op-be-shared/tool/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// GetPagesByFilter returns the pages that match the given filter.
func (a ServiceAdapter) GetPagesByFilter(ctx context.Context, pageFilter *pb.PageFilter) (*pb.Pages, error) {
	conn, err := grpc.Dial(tconfig.GetInternalConfigInstance().Grpc.BookHost+":"+tconfig.GetInternalConfigInstance().Grpc.BookPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "GetPagesByFilter", userId, err.Error()))
		return &pb.Pages{}, err
	}

	userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	ctx = metadata.AppendToOutgoingContext(ctx, string(smodel.QueryKeyUid), userId)
	pbResults, err := pb.NewBookSvcClient(conn).GetPagesByFilter(ctx, pageFilter)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "GetPagesByFilter", userId, err.Error()))
		return &pb.Pages{}, err
	}
	return pbResults, nil
}

// CreatePage insert a new page in the databse.
func (a ServiceAdapter) CreatePage(ctx context.Context, page *pb.Page) (*pb.Page, error) {
	conn, err := grpc.Dial(tconfig.GetInternalConfigInstance().Grpc.BookHost+":"+tconfig.GetInternalConfigInstance().Grpc.BookPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "CreatePage", userId, err.Error()))
		return &pb.Page{}, err
	}

	userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	ctx = metadata.AppendToOutgoingContext(ctx, string(smodel.QueryKeyUid), userId)
	pbResult, err := pb.NewBookSvcClient(conn).CreatePage(ctx, page)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "CreatePage", userId, err.Error()))
		return &pb.Page{}, err
	}
	return pbResult, nil
}

// UpdatePageBase sends the given page to the application layer for updating page base values.
func (a ServiceAdapter) UpdatePageBase(ctx context.Context, page *pb.Page) (*pb.Page, error) {
	conn, err := grpc.Dial(tconfig.GetInternalConfigInstance().Grpc.BookHost+":"+tconfig.GetInternalConfigInstance().Grpc.BookPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdatePageStatus", userId, err.Error()))
		return &pb.Page{}, err
	}

	userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	ctx = metadata.AppendToOutgoingContext(ctx, string(smodel.QueryKeyUid), userId)
	pbResult, err := pb.NewBookSvcClient(conn).UpdatePageBase(ctx, page)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdatePageStatus", userId, err.Error()))
		return &pb.Page{}, err
	}
	return pbResult, nil
}

// UpdatePageCore sends the given page to the application layer for updating page core values.
func (a ServiceAdapter) UpdatePageCore(ctx context.Context, page *pb.Page) (*pb.Page, error) {
	conn, err := grpc.Dial(tconfig.GetInternalConfigInstance().Grpc.BookHost+":"+tconfig.GetInternalConfigInstance().Grpc.BookPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdatePageStatus", userId, err.Error()))
		return &pb.Page{}, err
	}

	userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	ctx = metadata.AppendToOutgoingContext(ctx, string(smodel.QueryKeyUid), userId)
	pbResult, err := pb.NewBookSvcClient(conn).UpdatePageCore(ctx, page)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdatePageStatus", userId, err.Error()))
		return &pb.Page{}, err
	}
	return pbResult, nil
}

// UpdatePageStatus sends the given page to the application layer for updating page status.
func (a ServiceAdapter) UpdatePageStatus(ctx context.Context, page *pb.Page) (*pb.Page, error) {
	conn, err := grpc.Dial(tconfig.GetInternalConfigInstance().Grpc.BookHost+":"+tconfig.GetInternalConfigInstance().Grpc.BookPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdatePageStatus", userId, err.Error()))
		return &pb.Page{}, err
	}

	userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	ctx = metadata.AppendToOutgoingContext(ctx, string(smodel.QueryKeyUid), userId)
	pbResult, err := pb.NewBookSvcClient(conn).UpdatePageStatus(ctx, page)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdatePageStatus", userId, err.Error()))
		return &pb.Page{}, err
	}
	return pbResult, nil
}

// DeletePage soft-deletes the given page in the database.
func (a ServiceAdapter) DeletePage(ctx context.Context, page *pb.Page) (*pb.Page, error) {
	conn, err := grpc.Dial(tconfig.GetInternalConfigInstance().Grpc.BookHost+":"+tconfig.GetInternalConfigInstance().Grpc.BookPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "DeletePage", userId, err.Error()))
		return &pb.Page{}, err
	}

	userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	ctx = metadata.AppendToOutgoingContext(ctx, string(smodel.QueryKeyUid), userId)
	pbResult, err := pb.NewBookSvcClient(conn).DeletePage(ctx, page)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "DeletePage", userId, err.Error()))
		return &pb.Page{}, err
	}
	return pbResult, nil
}
