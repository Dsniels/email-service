package router

import (
	"net/http"

	"github.com/dsniels/email-service/pkg"
)

func RecoverMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				pkg.HandleError(w, err)
			}
		}()

		next.ServeHTTP(w, r)
	})

}
