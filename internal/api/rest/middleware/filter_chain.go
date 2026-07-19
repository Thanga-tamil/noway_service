package middleware

import (
	"net/http"
	"context"

	"gateway/internal/config"
)

func MyMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	tenantId := r.Header.Get("tenant-x")

	db := config.TenantDBs[tenantId]

    ctx := context.WithValue(r.Context(), tenantId, db)

    // call the next handler in the chain, passing the response writer and
    // the updated request object with the new context value.
    //
    // note: context.Context values are nested, so any previously set
    // values will be accessible as well, and the new `"user"` key
    // will be accessible from this point forward.
    next.ServeHTTP(w, r.WithContext(ctx))
  })
}

