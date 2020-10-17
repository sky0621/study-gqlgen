package graph

import (
	"context"
	"net/http"

	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/99designs/gqlgen/graphql"
)

type AppError struct {
	httpStatusCode int          // http.StatusCodeXXXXXXX を入れる
	appErrorCode   AppErrorCode // サービス固有に定義したエラーコード

	/*
	 * 以下、全てのエラー表現に必須ではない要素（オプションとして設定可能）
	 */
	field string
	value string
}

func (e *AppError) AddGraphQLError(ctx context.Context) {
	extensions := map[string]interface{}{
		"status_code": e.httpStatusCode,
		"error_code":  e.appErrorCode,
	}
	if e.field != "" {
		extensions["field"] = e.field
	}
	if e.value != "" {
		extensions["value"] = e.value
	}
	graphql.AddError(ctx, &gqlerror.Error{
		Message:    "",
		Extensions: extensions,
	})
}

func NewAppError(httpStatusCode int, appErrorCode AppErrorCode, opts ...AppErrorOption) *AppError {
	a := &AppError{
		httpStatusCode: httpStatusCode,
		appErrorCode:   appErrorCode,
	}

	for _, o := range opts {
		o(a)
	}

	return a
}

// 認証エラー用
func NewAuthenticationError(opts ...AppErrorOption) *AppError {
	return NewAppError(http.StatusUnauthorized, AppErrorCodeAuthenticationFailure, opts...)
}

// 認可エラー用
func NewAuthorizationError(opts ...AppErrorOption) *AppError {
	return NewAppError(http.StatusForbidden, AppErrorCodeAuthorizationFailure, opts...)
}

// バリデーションエラー用
func NewValidationError(field, value string, opts ...AppErrorOption) *AppError {
	options := []AppErrorOption{WithField(field), WithValue(value)}
	for _, opt := range opts {
		options = append(options, opt)
	}
	return NewAppError(http.StatusBadRequest, AppErrorCodeValidationFailure, options...)
}

// その他エラー用
func NewInternalServerError(opts ...AppErrorOption) *AppError {
	return NewAppError(http.StatusInternalServerError, AppErrorCodeUnexpectedFailure, opts...)
}

type AppErrorCode string

// MEMO: サービスの定義によっては意味のある文字列よりもコード体系を決めるのもあり。
const (
	// 認証エラー
	AppErrorCodeAuthenticationFailure AppErrorCode = "AUTHENTICATION_FAILURE"
	// 認可エラー
	AppErrorCodeAuthorizationFailure AppErrorCode = "AUTHORIZATION_FAILURE"
	// バリデーションエラー
	AppErrorCodeValidationFailure AppErrorCode = "VALIDATION_FAILURE"

	// その他の予期せぬエラー
	AppErrorCodeUnexpectedFailure AppErrorCode = "UNEXPECTED_FAILURE"
)

type AppErrorOption func(*AppError)

func WithField(v string) AppErrorOption {
	return func(a *AppError) {
		a.field = v
	}
}

func WithValue(v string) AppErrorOption {
	return func(a *AppError) {
		a.value = v
	}
}
