package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/Ng1n3/social/internal/store"
	"github.com/go-chi/chi/v5"
)

type userKey string

const UserCtx userKey = "user"

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromCtx(r)
	if user == nil {
		app.notFound(w, r, errors.New("user not found in context"))
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func getUserFromCtx(r *http.Request) *store.User {
	user, _ := r.Context().Value(UserCtx).(*store.User)
	return user
}

func (app *application) usersContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParams := chi.URLParam(r, "userID")
		id, err := strconv.ParseInt(idParams, 10, 64)
		if err != nil {
			app.badRequestResponse(w, r, err)
		}

		ctx := r.Context()

		user, err := app.store.Users.GetByID(ctx, id)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
				app.notFound(w, r, err)
			default:
				app.internalServerError(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, UserCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type FollowUser struct {
	UserID int64 `json:"user_id"`
}


func (app *application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	followedUser := getUserFromCtx(r)

	// TODO:Revert back to auth userID from ctx
	var payload FollowUser
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Followers.Follow(ctx, followedUser.ID, payload.UserID); err != nil {
		switch err {
		case store.ErrConflict:
			app.conflict(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
	}
}


func (app *application) unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	unfollowedUser := getUserFromCtx(r)

	// TODO:Revert back to auth userID from ctx
	var payload FollowUser
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Followers.Unfollow(ctx, unfollowedUser.ID, payload.UserID); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
	}
}
