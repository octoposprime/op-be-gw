package presentation

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"

	mo "github.com/octoposprime/op-be-graphql/internal/domain/model/object"
	presentation "github.com/octoposprime/op-be-graphql/pkg/presentation/dto/model"
	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	pb_dlr "github.com/octoposprime/op-be-shared/pkg/proto/pb/dlr"
	pb_user "github.com/octoposprime/op-be-shared/pkg/proto/pb/user"

)

// CreateDlr is the resolver for the createDlr field.
func (r *mutationResolver) CreateDlr(ctx context.Context, dlr presentation.DlrInput) (*presentation.Dlr, error) {
	userId := ctx.Value(smodel.QueryKeyUid).(string)
	userType := ctx.Value(smodel.QueryKeyUType).(pb_user.UserType)
	if userType != pb_user.UserType_UserTypeADMIN {
		user.ID = &userId
	}
	dtoData := presentation.NewDlrDto(&dlr)
	resultData, err := r.CommandHandler.CreateDlr(ctx, dtoData.ToPb())
	if err != nil {
		return nil, err
	}
	dtoData.PbData = resultData
	return dtoData.ToModel(), nil
}

// UpdateDlrBase is the resolver for the updateDlrBase field.
func (r *mutationResolver) UpdateDlrBase(ctx context.Context, dlr presentation.DlrInput) (*presentation.Dlr, error) {
	userId := ctx.Value(smodel.QueryKeyUid).(string)
	userType := ctx.Value(smodel.QueryKeyUType).(pb_user.UserType)
	if userType != pb_user.UserType_UserTypeADMIN {
		user.ID = &userId
	}
	dtoData := presentation.NewDlrDto(&dlr)
	resultData, err := r.CommandHandler.UpdateDlrBase(ctx, dtoData.ToPb())
	if err != nil {
		return nil, err
	}
	dtoData.PbData = resultData
	return dtoData.ToModel(), nil
}

// UpdateDlrCore is the resolver for the updateDlrCore field.
func (r *mutationResolver) UpdateDlrCore(ctx context.Context, dlr presentation.DlrInput) (*presentation.Dlr, error) {
	userId := ctx.Value(smodel.QueryKeyUid).(string)
	userType := ctx.Value(smodel.QueryKeyUType).(pb_user.UserType)
	if userType != pb_user.UserType_UserTypeADMIN {
		user.ID = &userId
	}
	dtoData := presentation.NewDlrDto(&dlr)
	resultData, err := r.CommandHandler.UpdateDlrCore(ctx, dtoData.ToPb())
	if err != nil {
		return nil, err
	}
	dtoData.PbData = resultData
	return dtoData.ToModel(), nil
}

// UpdateDlrStatus is the resolver for the updateDlrStatus field.
func (r *mutationResolver) UpdateDlrStatus(ctx context.Context, dlr presentation.DlrInput) (*presentation.Dlr, error) {
	userId := ctx.Value(smodel.QueryKeyUid).(string)
	userType := ctx.Value(smodel.QueryKeyUType).(pb_user.UserType)
	if userType != pb_user.UserType_UserTypeADMIN {
		user.ID = &userId
	}
	dtoData := presentation.NewDlrDto(&dlr)
	resultData, err := r.CommandHandler.UpdateDlrStatus(ctx, dtoData.ToPb())
	if err != nil {
		return nil, err
	}
	dtoData.PbData = resultData
	return dtoData.ToModel(), nil
}

// DeleteDlr is the resolver for the deleteDlr field.
func (r *mutationResolver) DeleteDlr(ctx context.Context, id string) (*presentation.Dlr, error) {
	userId := ctx.Value(smodel.QueryKeyUid).(string)
	userType := ctx.Value(smodel.QueryKeyUType).(pb_user.UserType)
	if userType != pb_user.UserType_UserTypeADMIN {
		id = userId
	}
	inData := new(presentation.DlrInput)
	inData.ID = &id
	dtoData := presentation.NewDlrDto(inData)
	resultData, err := r.CommandHandler.DeleteDlr(ctx, dtoData.ToPb())
	if err != nil {
		return nil, err
	}
	dtoData.PbData = resultData
	return dtoData.ToModel(), nil
}

// Dlr is the resolver for the dlr field.
func (r *queryResolver) Dlr(ctx context.Context, id string) (*presentation.Dlr, error) {
	userId := ctx.Value(smodel.QueryKeyUid).(string)
	userType := ctx.Value(smodel.QueryKeyUType).(pb_user.UserType)
	if userType != pb_user.UserType_UserTypeADMIN {
		id = userId
	}
	var filter presentation.DlrFilterInput
	filter.ID = &id
	dtoFilter := presentation.NewDlrFilterDto(&filter)
	dtoData := presentation.NewDlrDto(new(presentation.DlrInput))
	resultDatas, err := r.QueryHandler.GetDlrsByFilter(ctx, dtoFilter.ToPb())
	if err != nil {
		return nil, err
	}
	if len(resultDatas.Dlrs) > 0 {
		dtoData.PbData = resultDatas.Dlrs[0]
		return dtoData.ToModel(), nil
	} else {
		return nil, mo.ErrorDlrNotFound
	}
}

// Dlrs is the resolver for the dlrs field.
func (r *queryResolver) Dlrs(ctx context.Context, filter *presentation.DlrFilterInput) (*presentation.Dlrs, error) {
	userId := ctx.Value(smodel.QueryKeyUid).(string)
	userType := ctx.Value(smodel.QueryKeyUType).(pb_user.UserType)
	if userType != pb_user.UserType_UserTypeADMIN {
		filter.ID = &userId
	}
	var results presentation.Dlrs
	if filter == nil {
		filter = &presentation.DlrFilterInput{}
	}
	dtoFilter := presentation.NewDlrFilterDto(filter)

	resultDatas, err := r.QueryHandler.GetDlrsByFilter(ctx, dtoFilter.ToPb())
	if err != nil {
		return nil, err
	}

	for _, resultData := range resultDatas.Dlrs {
		dtoData := presentation.NewDlrDto(new(presentation.DlrInput))
		dtoData.PbData = resultData
		results.Dlrs = append(results.Dlrs, dtoData.ToModel())
	}
	results.Total = int32(resultDatas.TotalRows)

	return &results, nil
}
