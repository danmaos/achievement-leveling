package handler

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	mw "achievement-leveling/gateway/middleware"
)

func NewProxy(achievementSvcURL string) http.Handler {
	target, _ := url.Parse(achievementSvcURL)

	proxy := httputil.NewSingleHostReverseProxy(target)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Inject user ID header from JWT claims
		claims := mw.GetClaims(r)
		if claims != nil {
			r.Header.Set("X-User-ID", claims.UserID)
			r.Header.Set("X-User-Email", claims.Email)
		}
		proxy.ServeHTTP(w, r)
	})
}
