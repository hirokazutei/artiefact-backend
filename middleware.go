package artiefact

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"runtime/debug"
	"strings"

	c "github.com/hirokazu/artiefact-backend/constants"
	"github.com/hirokazu/artiefact-backend/model"
)

var openAPIs = []string{
	"/user/sign-up",
	"/user/sign-in",
	"/user/username-availability",
}

var tokenRegex = regexp.MustCompile(`^(?i)bearer (\w+)$`)

// AuthResponse response object
type AuthResponse struct {
	Token *model.AccessToken
	User  *model.ArtiefactUser
}

// ServeHTTPC
func (h ServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	status, res, err := h.handler(w, r)
	fmt.Println(err)

	w.WriteHeader(status)
	if err := encoder.Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(c.ErrorUnknown)
		return
	}
	return
}

func setJSONHeaderMiddleware(next http.Handler) http.Handler {
	setHeader := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(setHeader)
}

func recoverMiddleware(next http.Handler) http.Handler {
	recovery := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				debug.PrintStack()
				fmt.Printf("system paniced: %+v", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(recovery)
}

func tokenAuthMiddleware(app *App) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var response AuthResponse
			if !isOpenAPI(r.URL.Path) {
				authHeader := r.Header.Get("Authorization")
				matches := tokenRegex.FindStringSubmatch(authHeader)
				if len(matches) != 2 {
					e := c.ErrorInvalidHeader()
					app.sendResponse(w, e.Status, e)
					return
				}
				token := matches[1]
				matchedToken, found, err := model.GetValidToken(app.DB, token)
				if err != nil {
					e := c.ErrorInternalServer()
					app.sendResponse(w, e.GetStatus(), e)
					return
				}
				if !found {
					e := c.ErrorInvalidToken()
					app.sendResponse(w, e.GetStatus(), e)
					return
				}

				user, found, err := model.GetActiveArtiefactUserByID(app.DB, matchedToken.UserID)
				if err != nil {
					e := c.ErrorInternalServer()
					app.sendResponse(w, e.GetStatus(), e)
					return
				}
				if !found {
					e := c.ErrorObjectNotFound("artiefact_user")
					app.sendResponse(w, e.GetStatus(), e)
					return
				}

				response.User = user
				response.Token = matchedToken
				ctx := context.WithValue(r.Context(), contextKeyAuth, response)
				r = r.WithContext(ctx)
			}
			next.ServeHTTP(w, r)
		})
	}
}

func isOpenAPI(path string) bool {
	for _, api := range openAPIs {
		if strings.Index(path, api) == 0 {
			return true
		}
	}
	return false
}
