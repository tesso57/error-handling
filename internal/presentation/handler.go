package presentation

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	"github.com/tesso57/error-handling-sample/internal/application"
	"github.com/tesso57/error-handling-sample/internal/domain"
	"github.com/tesso57/error-handling-sample/internal/errors"
)

// Request/Response types for Connect handlers.
type RegisterUserRequest struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GetUserRequest struct {
	ID string `json:"id"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// RegisterUserHandler returns a Connect handler for registering users.
func RegisterUserHandler(svc *application.UserService) http.Handler {

	return connect.NewUnaryHandler(
		"/user.v1.UserService/RegisterUser",
		func(ctx context.Context, req *connect.Request[RegisterUserRequest]) (*connect.Response[UserResponse], error) {
			r := req.Msg
			user, err := svc.RegisterUser(domain.UserID(r.ID), r.Name, domain.Email(r.Email))
			if err != nil {
				// Check for application errors
				if errors.Is(err, application.ErrEmailAlreadyRegistered) {
					return nil, errors.NewErrorResponse(connect.CodeAlreadyExists, "email already registered", err).ToConnectError()
				}
				if errors.Is(err, application.ErrUserNotFound) {
					return nil, errors.NewErrorResponse(connect.CodeNotFound, "user not found", err).ToConnectError()
				}
				// For any other errors, use default ToConnectError
				return nil, errors.ToConnectError(err)
			}
			resp := &UserResponse{ID: string(user.ID), Name: user.Name, Email: string(user.Email)}
			return connect.NewResponse(resp), nil
		},
	)
}

// GetUserHandler returns a Connect handler for retrieving users.
func GetUserHandler(svc *application.UserService) http.Handler {

	return connect.NewUnaryHandler(
		"/user.v1.UserService/GetUser",
		func(ctx context.Context, req *connect.Request[GetUserRequest]) (*connect.Response[UserResponse], error) {
			r := req.Msg
			user, err := svc.GetUser(domain.UserID(r.ID))
			if err != nil {
				// Check for application errors
				if errors.Is(err, application.ErrEmailAlreadyRegistered) {
					return nil, errors.NewErrorResponse(connect.CodeAlreadyExists, "email already registered", err).ToConnectError()
				}
				if errors.Is(err, application.ErrUserNotFound) {
					return nil, errors.NewErrorResponse(connect.CodeNotFound, "user not found", err).ToConnectError()
				}
				// For any other errors, use default ToConnectError
				return nil, errors.ToConnectError(err)
			}
			resp := &UserResponse{ID: string(user.ID), Name: user.Name, Email: string(user.Email)}
			return connect.NewResponse(resp), nil
		},
	)
}
