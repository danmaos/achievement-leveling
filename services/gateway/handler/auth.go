package handler

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"achievement-leveling/gateway/config"
	mw "achievement-leveling/gateway/middleware"
	"achievement-leveling/gateway/oauth"
)

type AuthHandler struct {
	cfg         *config.Config
	googleOAuth *oauth.GoogleOAuth
}

func NewAuthHandler(cfg *config.Config, googleOAuth *oauth.GoogleOAuth) *AuthHandler {
	return &AuthHandler{cfg: cfg, googleOAuth: googleOAuth}
}

func (h *AuthHandler) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	state := generateState()
	// In production, store state in a cookie for CSRF validation
	http.SetCookie(w, &http.Cookie{
		Name:     "oauth_state",
		Value:    state,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   300,
	})
	url := h.googleOAuth.GetAuthURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *AuthHandler) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	// Validate state
	stateCookie, err := r.Cookie("oauth_state")
	if err != nil || stateCookie.Value != r.URL.Query().Get("state") {
		http.Error(w, `{"error":"invalid state parameter"}`, http.StatusBadRequest)
		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, `{"error":"missing authorization code"}`, http.StatusBadRequest)
		return
	}

	token, err := h.googleOAuth.Exchange(r.Context(), code)
	if err != nil {
		log.Printf("token exchange error: %v", err)
		http.Error(w, `{"error":"failed to exchange token"}`, http.StatusInternalServerError)
		return
	}

	googleUser, err := h.googleOAuth.GetUserInfo(r.Context(), token)
	if err != nil {
		log.Printf("get user info error: %v", err)
		http.Error(w, `{"error":"failed to get user info"}`, http.StatusInternalServerError)
		return
	}

	// Upsert user via achievement service
	upsertBody, _ := json.Marshal(map[string]string{
		"google_id": googleUser.ID,
		"email":     googleUser.Email,
		"name":      googleUser.Name,
		"picture":   googleUser.Picture,
	})

	resp, err := http.Post(h.cfg.AchievementSvcURL+"/api/users", "application/json", bytes.NewReader(upsertBody))
	if err != nil {
		log.Printf("upsert user error: %v", err)
		http.Error(w, `{"error":"failed to create user"}`, http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var user struct {
		ID      string `json:"id"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
	}
	json.NewDecoder(resp.Body).Decode(&user)

	// Generate JWT
	jwtToken, err := mw.GenerateToken(h.cfg.JWTSecret, user.ID, user.Email, user.Name, user.Picture)
	if err != nil {
		http.Error(w, `{"error":"failed to generate token"}`, http.StatusInternalServerError)
		return
	}

	// Redirect to frontend with token
	redirectURL := fmt.Sprintf("%s/auth/callback?token=%s", h.cfg.FrontendURL, jwtToken)
	http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
}

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	claims := mw.GetClaims(r)
	if claims == nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	// Fetch fresh user data from achievement service
	resp, err := http.Get(h.cfg.AchievementSvcURL + "/api/users/" + claims.UserID)
	if err != nil {
		http.Error(w, `{"error":"failed to fetch user"}`, http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	var user json.RawMessage
	json.NewDecoder(resp.Body).Decode(&user)
	json.NewEncoder(w).Encode(user)
}

func generateState() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}
