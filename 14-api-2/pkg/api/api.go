package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strings"
	"thinknetica/14-api-2/pkg/services/jwtAuth"
	"thinknetica/14-api-2/pkg/users"
)

type API struct {
	router *mux.Router
	users  *users.Users
}

type LogInfo struct {
	Usr string
	Psw string
}

const AUTH_PATH = "/auth"

func New() *API {
	api := API{}
	api.users = users.New()
	api.router = mux.NewRouter()
	api.endpoints()
	api.router.Use(api.authorized)
	api.router.Use(api.requestID)
	api.router.Use(api.logger)
	return &api
}

func (api *API) Router() *mux.Router {
	return api.router
}

func (api *API) endpoints() {
	api.router.HandleFunc("/", api.mainHandler).Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc(AUTH_PATH, api.authentication).Methods(http.MethodPost, http.MethodOptions)
}

func (a *API) mainHandler(w http.ResponseWriter, r *http.Request) {
	if admin, ok := r.Context().Value("admin").(bool); !ok || !admin {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}
	fmt.Fprint(w, "!!!!!")
}

func (api *API) authentication(w http.ResponseWriter, r *http.Request) {
	var log LogInfo
	err := json.NewDecoder(r.Body).Decode(&log)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usr, err := api.users.Search(log.Usr, log.Psw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	token, err := jwtAuth.GenerateToken(usr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(token)
}

func (api *API) authorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == AUTH_PATH {
			next.ServeHTTP(w, r)
			return
		}
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Empty auth header", http.StatusUnauthorized)
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			http.Error(w, "Invalid auth header", http.StatusUnauthorized)
			return
		}

		claims, err := jwtAuth.ParseToken(headerParts[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(context.WithValue(r.Context(), "admin", claims.Admin), "userID", claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (api *API) requestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "requestID", rand.Intn(1_000_000))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (api *API) logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var userID int
		var requestID int
		if val, ok := r.Context().Value("userID").(int); ok {
			userID = val
		}
		if val, ok := r.Context().Value("requestID").(int); ok {
			requestID = val
		}
		fmt.Printf("Method: %v, Addr: %v, URI: %v, userID: %v, requestID: %v\r\n", r.Method, r.RemoteAddr, r.RequestURI, userID, requestID)
		next.ServeHTTP(w, r)
	})
}
