package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/doug-martin/goqu/v9"

	"github.com/lvlBA/online_shop/internal/passport/models"
	utilspagination "github.com/lvlBA/online_shop/pkg/utils/pagination"
)

const tableNameService = "resources"

type CreateServiceParams struct {
	Urn    string
	Access bool
}

type ResourceImpl struct {
	svc service
}

type ListServiceFilter struct {
	Pagination *models.Pagination
}

func (f *ListServiceFilter) Filter(ds *goqu.SelectDataset) *goqu.SelectDataset {
	if f.Pagination != nil {
		utilspagination.NewPagination(f.Pagination.Page, f.Pagination.Limit).DataSet(ds)
	}

	return ds
}

func (r *ResourceImpl) CreateResource(ctx context.Context, params *CreateServiceParams) (*models.Resource, error) {
	model := &models.Resource{
		Meta:   models.Meta{},
		Urn:    params.Urn,
		Access: params.Access,
	}
	model.UpdateMeta()

	id, err := r.svc.create(ctx, tableNameService, model)
	if err != nil {
		return nil, err
	}
	model.ID = id

	return model, nil
}

type GetResourceParams struct {
	ID       *string
	Resource *string
}

func (p *GetResourceParams) filter(sd *goqu.SelectDataset) (*goqu.SelectDataset, error) {
	switch {
	case p.ID != nil:
		return sd.Where(goqu.Ex{"id": *p.ID}), nil
	case p.Resource != nil:
		return sd.Where(goqu.Ex{"resource": *p.Resource}), nil
	default:
		return nil, errors.New("undefined behavior: id is not set and resource is not set")
	}
}

func (r *ResourceImpl) GetResource(ctx context.Context, params *GetResourceParams) (*models.Resource, error) {
	sd, err := params.filter(goqu.From(tableNameService).Select("*"))
	if err != nil {
		return nil, fmt.Errorf("failed to create filter: %w", err)
	}

	query, _, err := sd.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to create query: %w", err)
	}

	result := &models.Resource{}
	if err = r.svc.GetContext(ctx, result, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}

		return nil, err
	}

	return result, nil
}

func (r *ResourceImpl) DeleteResource(ctx context.Context, id string) error {
	return r.svc.delete(ctx, tableNameService, id)
}

func (r *ResourceImpl) ListResource(ctx context.Context, filter *ListServiceFilter) ([]*models.Resource, error) {
	ds := goqu.From(tableNameService).Select("*")
	ds = filter.Filter(ds)
	query, _, err := ds.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to create query: %w", err)
	}
	result := make([]*models.Resource, 0)
	if err = r.svc.SelectContext(ctx, &result, query); err != nil {
		return nil, err
	}
	return result, nil
}
