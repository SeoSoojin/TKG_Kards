package models

import "errors"

var (
	ErrEmptyUserId       = errors.New("userId is empty")
	ErrEmptyCollectionId = errors.New("collectionId is empty")
)
