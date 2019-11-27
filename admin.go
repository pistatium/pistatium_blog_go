package main

import (
	"context"
	"golang.org/x/crypto/bcrypt"
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


type AdminUserRepoImpl struct {
	projectID string
}

func NewAdminUserRepoImpl() AdminUserRepo {
	return &AdminUserRepoImpl{}
}

func (d AdminUserRepoImpl) GetValidUser(ctx context.Context, username string, password string) (*AdminUser, error) {
	if username != "kimihiro-n" {
		return nil, &LoginError{"Invalid user/password"}
	}
	hash := "$2y$12$IXFmmeezqymra5O00L95lejHgVvMf9n7vDsKrU7f3s3zJTd5aePNS"
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return nil, &LoginError{"Invalid user/password"}
	}
	u := &AdminUser{Username: username, Password:hash}
	return u, nil
}

