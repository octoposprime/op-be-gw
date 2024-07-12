package presentation

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"

	presentation "github.com/octoposprime/op-be-graphql/pkg/presentation/dto/model"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, loginRequest presentation.LoginRequestInput) (*presentation.Token, error) {
	dtoData := presentation.NewLoginRequestDto(&loginRequest)
	resultData, err := r.CommandHandler.Login(ctx, dtoData.ToPb())
	if err != nil {
		return nil, err
	}
	dtoResultData := presentation.NewTokenDto(new(presentation.TokenInput))
	dtoResultData.PbData = resultData
	return dtoResultData.ToModel(), nil
}

// Refresh is the resolver for the refresh field.
func (r *mutationResolver) Refresh(ctx context.Context, token presentation.TokenInput) (*presentation.Token, error) {
	dtoData := presentation.NewTokenDto(&token)
	resultData, err := r.CommandHandler.Refresh(ctx, dtoData.ToPb())
	if err != nil {
		return nil, err
	}
	dtoData.PbData = resultData
	return dtoData.ToModel(), nil
}

// Logout is the resolver for the logout field.
func (r *mutationResolver) Logout(ctx context.Context, token presentation.TokenInput) (*bool, error) {
	dtoData := presentation.NewTokenDto(&token)
	err := r.CommandHandler.Logout(ctx, dtoData.ToPb())
	if err != nil {
		result := bool(false)
		return &result, err
	}
	result := bool(true)
	return &result, nil
}
