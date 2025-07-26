package presentation

import (
	"net/http"

	"github.com/tesso57/error-handling-sample/internal/application"
	"github.com/tesso57/error-handling-sample/internal/infrastructure"
)

// NewRouter sets up HTTP routes for the API.
func NewRouter() http.Handler {
	mux := http.NewServeMux()
	repo := infrastructure.NewInMemoryUserRepository()
	svc := application.NewUserService(repo)
	mux.Handle("/user.v1.UserService/RegisterUser", RegisterUserHandler(svc))
	mux.Handle("/user.v1.UserService/GetUser", GetUserHandler(svc))
	return mux
}
