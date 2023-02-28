package site

import (
	"context"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api "github.com/lvlBA/online_shop/pkg/management/v1"
)

func (s ServiceImpl) GetSite(ctx context.Context, req *api.GetSiteRequest) (*api.GetSiteResponse, error) {
	if err := validateGetSideReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	site, err := s.ctrlSite.Get(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, "error get site")
	}

	return &api.GetSiteResponse{
		Site: adaptSiteToApi(site),
	}, nil
}

func validateGetSideReq(req *api.GetSiteRequest) error {
	return validation.Errors{
		"name": validation.Validate(req.Id, validation.Required, is.UUIDv4),
	}.Filter()
}