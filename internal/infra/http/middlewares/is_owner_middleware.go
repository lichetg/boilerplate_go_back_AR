package middlewares

/*
import (
	"errors"
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/controllers"
)

type Userable interface {
	GetUserId() uint64
}

func IsOwnerMiddleware[domainType Userable]() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			userVal := ctx.Value(controllers.GetUserKey())
			user, ok := userVal.(domain.User)
			if !ok {
				log.Println("IsOwnerMiddleware: no user in context or wrong type")
				controllers.Unauthorized(w, errors.New("unauthorized"))
				return
			}
			obj := controllers.GetPathValFromCtx[domainType](ctx)

			if obj.GetUserId() != user.Id {
				err := errors.New("you have no access to this object")
				controllers.Forbidden(w, err)
				return
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(hfn)
	}
}

func (s Subscription) GetUserId() uint64 {
	return s.UserId
}
*/
