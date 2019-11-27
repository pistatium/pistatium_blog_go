package main

import (
	"golang.org/x/crypto/bcrypt"
	"cloud.google.com/go/datastore"
	"context"
)

type LoginError struct {
	Msg string
}

func (e *LoginError) Error() string {
	return "ユーザー名またはパスワードが一致しませんでした"
}

type AdminUser struct {
	Username string `datastore:-`
	Password string `datastore:"password"`
}

type AdminUserRepo interface {
	GetValidUser(ctx context.Context, username string, password string) (*AdminUser, error)

}

// ======================================================================

type DatastoreAdminUserRepoImpl struct {
	projectID string
}

func NewDatastoreAdminUserRepoImpl(projectID string) AdminUserRepo {
	return &DatastoreAdminUserRepoImpl{projectID: projectID}
}

func (d DatastoreAdminUserRepoImpl) GetValidUser(ctx context.Context, username string, password string) (*AdminUser, error) {
	client, err := datastore.NewClient(ctx, d.projectID)
	if err != nil {
		return nil, err
	}
	k := datastore.NameKey("AdminUser", username, nil)
	u := new(AdminUser)
	err = client.Get(ctx, k, u)
	if err != nil && err != err.(*datastore.ErrFieldMismatch) {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return nil, &LoginError{err.Error()}
	}
	return u, nil
}
