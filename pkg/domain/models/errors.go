package models

import "errors"

var (
	ErrEmptyUserId       = errors.New("userId is empty")
	ErrEmptyCollectionId = errors.New("collectionId is empty")
	ErrEmptyIdolId       = errors.New("idolId is empty")
	ErrEmptyAlbumId      = errors.New("albumId is empty")
	ErrEmptyGroupId      = errors.New("groupId is empty")
)
