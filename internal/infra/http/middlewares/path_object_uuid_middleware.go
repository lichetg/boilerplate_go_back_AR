package middlewares

import (
	"context"
	"errors"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
)

type FindableUUID interface {
	FindGUID(id uuid.UUID) (any, error)
}

func PathObjectUUID(
	pathKey string,
	ctxKey controllers.CtxKey,
	service FindableUUID,
) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			raw := chi.URLParam(r, pathKey)

			id, err := uuid.Parse(raw)
			if err != nil {
				controllers.BadRequest(w, errors.New("invalid uuid"))
				return
			}

			obj, err := service.FindGUID(id)
			if err != nil {
				controllers.NotFound(w, errors.New("record not found"))
				return
			}

			ctx := context.WithValue(r.Context(), ctxKey, obj)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
