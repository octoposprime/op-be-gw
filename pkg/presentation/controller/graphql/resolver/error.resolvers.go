package presentation

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"

	mo "github.com/octoposprime/op-be-graphql/internal/domain/model/object"
	presentation "github.com/octoposprime/op-be-graphql/pkg/presentation/dto/model"
	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/error"
)

// Errors is the resolver for the errors field.
func (r *queryResolver) Errors(ctx context.Context) ([]*presentation.Error, error) {
	var results []*presentation.Error
	resultDatas, err := r.QueryHandler.GetErrors(ctx)
	if err != nil {
		return nil, err
	}

	//Shared Errors
	for _, resultData := range smodel.ERRORS {
		if resultData != nil {
			dtoData := presentation.NewErrorDto()
			dtoData.PbData = &pb.Error{
				Error: resultData.Error(),
			}
			results = append(results, dtoData.ToModel())
		}
	}

	//Service Errors
	for _, resultData := range resultDatas.Errors {
		if resultData != nil {
			dtoData := presentation.NewErrorDto()
			dtoData.PbData = resultData
			results = append(results, dtoData.ToModel())
		}
	}

	//Graphql Errors
	for _, resultData := range mo.ERRORS {
		if resultData != nil {
			dtoData := presentation.NewErrorDto()
			dtoData.PbData = &pb.Error{
				Error: resultData.Error(),
			}
			results = append(results, dtoData.ToModel())
		}
	}

	return results, nil
}
