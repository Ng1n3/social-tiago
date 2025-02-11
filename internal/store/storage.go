package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("resource not found")
	ErrConflict          = errors.New("conflict: resource has been modified")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, int64) (*Post, error)
		Update(context.Context, *Post) error
		Delete(context.Context, int64) error
		DeleteSeedAll(context.Context) error
	}
	Users interface {
		Create(context.Context, *User) error
		DeleteSeedAll(context.Context) error
		GetByID(context.Context, int64) (*User, error)
	}
	Followers interface {
		Follow(ctx context.Context, followerId, userId int64) error
		Unfollow(ctx context.Context, followerId, userId int64) error
	}
	Comments interface {
		GetByPostID(context.Context, int64) (*[]Comment, error)
		Create(context.Context, *Comment) error
		DeleteSeedAll(context.Context) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:    &PostStore{db},
		Users:    &UserStore{db},
		Comments: &CommentStore{db},
		Followers: &FollowerStore{db},
	}
}
