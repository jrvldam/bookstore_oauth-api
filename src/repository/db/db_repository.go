package db

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/jrvldam/bookstore_oauth-api/src/clients/cassandra"
	"github.com/jrvldam/bookstore_oauth-api/src/domain/access_token"
	"github.com/jrvldam/bookstore_oauth-api/src/utils/errors"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens (access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	quryUpdateExpires      = "UPDATE access_token SET expires=? WHERE access_token=?;"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type dbRepository struct{}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	var at access_token.AccessToken

	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(&at.AccessToken, &at.UserId, &at.ClientId, &at.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError(fmt.Sprintf("Access token %s not found", id))
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &at, nil

}

func (r *dbRepository) Create(at access_token.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryCreateAccessToken, at.AccessToken, at.UserId, at.ClientId, at.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(quryUpdateExpires, at.Expires, at.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func NewRepository() DbRepository {
	return &dbRepository{}
}
