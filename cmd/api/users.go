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

func(app *application) usersContextMiddleware(next http.Handler) http.Handler{
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