package internalErrors

import "errors"

var ErrInternal error = errors.New("internal server error")
var ErrEmptyName error = errors.New("name should not be empty")
var ErrEmptyContent error = errors.New("content should not be empty")
var ErrEmptyContacts error = errors.New("contacts should not be empty")
