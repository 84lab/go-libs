package http

import "net/http"

type Error struct {
	Code int    `json:"code"`
	Slug string `json:"slug"`
}

func NewError(code int, slug string) *Error {
	return &Error{
		Code: code,
		Slug: slug,
	}
}

func (e Error) GetSlug() string {
	return e.Slug
}

func (e Error) GetStatusCode() int {
	return e.Code
}

func NewBadRequestError(slug string) *Error {
	return NewError(http.StatusBadRequest, slug)
}

func NewUnauthorizedError(slug string) *Error {
	return NewError(http.StatusUnauthorized, slug)
}

func NewConflictError(slug string) *Error {
	return NewError(http.StatusConflict, slug)
}

func NewInternalServerError(slug string) *Error {
	return NewError(http.StatusInternalServerError, slug)
}

func NewBadGatewayError(slug string) *Error {
	return NewError(http.StatusBadGateway, slug)
}
