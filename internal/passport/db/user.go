package db

import (
	"context"
	"crypto/sha512"
	"database/sql"
	"errors"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	utilspagination "github.com/lvlBA/online_shop/pkg/utils/pagination"

	"github.com/lvlBA/online_shop/internal/passport/models"
)

const tableNameUser = "User"

type UserImpl struct {
	svc service
}

type CreateUserParams struct {
	FirstName    string
	LastName     string
	Age          uint32
	Sex          models.Sex
	Login        string
	HashPassword string
}

func (u *UserImpl) CreateUser(ctx context.Context, params *CreateUserParams) (*models.User, error) {
	model := &models.User{
		Meta:         models.Meta{},
		FirstName:    params.FirstName,
		LastName:     params.LastName,
		Age:          params.Age,
		Sex:          params.Sex,
		Login:        params.Login,
		HashPassword: params.HashPassword,
	}
	model.UpdateMeta()

	id, err := u.svc.create(ctx, tableNameUser, model)
	if err != nil {
		return nil, err
	}
	model.ID = id

	return model, nil
}

func (u *UserImpl) GetUser(ctx context.Context, id string) (*models.User, error) {
	result := &models.User{}

	query, _, err := goqu.From(tableNameUser).Select("*").Where(goqu.Ex{"id": id}).ToSQL()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}
		return nil, fmt.Errorf("failed to create query: %w", err)
	}

	if err = u.svc.GetContext(ctx, result, query); err != nil {
		return nil, err
	}

	return result, nil
}

func (u *UserImpl) DeleteUser(ctx context.Context, id string) error {
	return u.svc.delete(ctx, tableNameUser, id)
}

type ListUserFilter struct {
	Pagination *models.Pagination
}

func (f *ListUserFilter) Filter(ds *goqu.SelectDataset) *goqu.SelectDataset {
	if f.Pagination != nil {
		utilspagination.NewPagination(f.Pagination.Page, f.Pagination.Limit).DataSet(ds)
	}

	return ds
}

func (u *UserImpl) ListUsers(ctx context.Context, filter *ListUserFilter) ([]*models.User, error) {
	ds := goqu.From(tableNameUser).Select("*")
	ds = filter.Filter(ds)
	query, _, err := ds.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to create query: %w", err)
	}
	result := make([]*models.User, 0)
	if err = u.svc.SelectContext(ctx, &result, query); err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserImpl) ChangePass(ctx context.Context, id string, oldPass string, newPass string) error {
	result := &models.User{}

	query, _, err := goqu.From(tableNameUser).Select("*").Where(goqu.Ex{"id": id}).ToSQL()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return fmt.Errorf("failed to create query: %w", err)
	}
	if err = u.svc.GetContext(ctx, result, query); err != nil {
		hashPass := localHash(oldPass)
		if result.ID == id && result.HashPassword == string(hashPass) {
			ds := goqu.From(tableNameUser)

			_, _, _ = ds.Where(goqu.C("id").Eq(result.ID)).Update().Set(
				goqu.Record{"hash_password": string(localHash(newPass))},
			).ToSQL()
		}
	}
	return err
}

func localHash(password string) (hashPass []byte) {

	hash := sha512.Sum512([]byte(password))
	hashP := make([]byte, len(hash))
	copy(hashPass, hash[:len(hash)])
	return hashP

}
