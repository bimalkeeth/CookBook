package middleware

import (
	"context"
	"net/http"
)

func newMiddlewareContext(ctx context.Context, r *http.Request) context.Context {
	FoodId := r.Header.Get("X-Food-ID")
	if FoodId == "" {
		FoodId = "Bar"
	}
	return context.WithValue(ctx, "FoodId", FoodId)
}

func ContextHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			ctx := newMiddlewareContext(r.Context(), r)
			next.ServeHTTP(w, r.WithContext(ctx))
		}()
	})
}
