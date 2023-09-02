package middleware

import "net/http"

type UserIdentityMiddleware struct {
}

func NewUserIdentityMiddleware() *UserIdentityMiddleware {
	return &UserIdentityMiddleware{}
}

func (m *UserIdentityMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		next(w, r)
	}
}
