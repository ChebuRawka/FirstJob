// Package users provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package users

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// Task defines model for Task.
type Task struct {
	// CreatedAt Timestamp when the task was created
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// DeletedAt Timestamp when the task was soft-deleted
	DeletedAt *time.Time `json:"deleted_at"`

	// Id Unique identifier for the task
	Id *int `json:"id,omitempty"`

	// IsDone Indicates whether the task is completed
	IsDone *bool `json:"is_done,omitempty"`

	// Task Description of the task
	Task *string `json:"task,omitempty"`

	// UpdatedAt Timestamp when the task was last updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	// UserId ID of the user associated with the task
	UserId *int `json:"user_id,omitempty"`
}

// User defines model for User.
type User struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt"`
	Email     *string    `json:"email,omitempty"`
	Id        *int64     `json:"id,omitempty"`
	Password  *string    `json:"password,omitempty"`

	// Tasks List of tasks assigned to the user
	Tasks     *[]Task    `json:"tasks,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// PostUserJSONRequestBody defines body for PostUser for application/json ContentType.
type PostUserJSONRequestBody = User

// PatchUserByIDJSONRequestBody defines body for PatchUserByID for application/json ContentType.
type PatchUserByIDJSONRequestBody = User

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all users
	// (GET /users)
	GetUsers(ctx echo.Context) error
	// Create new user
	// (POST /users)
	PostUser(ctx echo.Context) error
	// Delete user by ID
	// (DELETE /users/{id})
	DeleteUserByID(ctx echo.Context, id int64) error
	// Get user by ID
	// (GET /users/{id})
	GetUserByID(ctx echo.Context, id int64) error
	// Update user by ID
	// (PATCH /users/{id})
	PatchUserByID(ctx echo.Context, id int64) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetUsers converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUsers(ctx)
	return err
}

// PostUser converts echo context to params.
func (w *ServerInterfaceWrapper) PostUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostUser(ctx)
	return err
}

// DeleteUserByID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteUserByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteUserByID(ctx, id)
	return err
}

// GetUserByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetUserByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUserByID(ctx, id)
	return err
}

// PatchUserByID converts echo context to params.
func (w *ServerInterfaceWrapper) PatchUserByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchUserByID(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/users", wrapper.GetUsers)
	router.POST(baseURL+"/users", wrapper.PostUser)
	router.DELETE(baseURL+"/users/:id", wrapper.DeleteUserByID)
	router.GET(baseURL+"/users/:id", wrapper.GetUserByID)
	router.PATCH(baseURL+"/users/:id", wrapper.PatchUserByID)

}

type GetUsersRequestObject struct {
}

type GetUsersResponseObject interface {
	VisitGetUsersResponse(w http.ResponseWriter) error
}

type GetUsers200JSONResponse []User

func (response GetUsers200JSONResponse) VisitGetUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostUserRequestObject struct {
	Body *PostUserJSONRequestBody
}

type PostUserResponseObject interface {
	VisitPostUserResponse(w http.ResponseWriter) error
}

type PostUser201Response struct {
}

func (response PostUser201Response) VisitPostUserResponse(w http.ResponseWriter) error {
	w.WriteHeader(201)
	return nil
}

type DeleteUserByIDRequestObject struct {
	Id int64 `json:"id"`
}

type DeleteUserByIDResponseObject interface {
	VisitDeleteUserByIDResponse(w http.ResponseWriter) error
}

type DeleteUserByID204Response struct {
}

func (response DeleteUserByID204Response) VisitDeleteUserByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeleteUserByID404Response struct {
}

func (response DeleteUserByID404Response) VisitDeleteUserByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetUserByIDRequestObject struct {
	Id int64 `json:"id"`
}

type GetUserByIDResponseObject interface {
	VisitGetUserByIDResponse(w http.ResponseWriter) error
}

type GetUserByID200JSONResponse User

func (response GetUserByID200JSONResponse) VisitGetUserByIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetUserByID404Response struct {
}

func (response GetUserByID404Response) VisitGetUserByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type PatchUserByIDRequestObject struct {
	Id   int64 `json:"id"`
	Body *PatchUserByIDJSONRequestBody
}

type PatchUserByIDResponseObject interface {
	VisitPatchUserByIDResponse(w http.ResponseWriter) error
}

type PatchUserByID200Response struct {
}

func (response PatchUserByID200Response) VisitPatchUserByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PatchUserByID404Response struct {
}

func (response PatchUserByID404Response) VisitPatchUserByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get all users
	// (GET /users)
	GetUsers(ctx context.Context, request GetUsersRequestObject) (GetUsersResponseObject, error)
	// Create new user
	// (POST /users)
	PostUser(ctx context.Context, request PostUserRequestObject) (PostUserResponseObject, error)
	// Delete user by ID
	// (DELETE /users/{id})
	DeleteUserByID(ctx context.Context, request DeleteUserByIDRequestObject) (DeleteUserByIDResponseObject, error)
	// Get user by ID
	// (GET /users/{id})
	GetUserByID(ctx context.Context, request GetUserByIDRequestObject) (GetUserByIDResponseObject, error)
	// Update user by ID
	// (PATCH /users/{id})
	PatchUserByID(ctx context.Context, request PatchUserByIDRequestObject) (PatchUserByIDResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetUsers operation middleware
func (sh *strictHandler) GetUsers(ctx echo.Context) error {
	var request GetUsersRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetUsers(ctx.Request().Context(), request.(GetUsersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUsers")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetUsersResponseObject); ok {
		return validResponse.VisitGetUsersResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostUser operation middleware
func (sh *strictHandler) PostUser(ctx echo.Context) error {
	var request PostUserRequestObject

	var body PostUserJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostUser(ctx.Request().Context(), request.(PostUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostUser")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostUserResponseObject); ok {
		return validResponse.VisitPostUserResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteUserByID operation middleware
func (sh *strictHandler) DeleteUserByID(ctx echo.Context, id int64) error {
	var request DeleteUserByIDRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteUserByID(ctx.Request().Context(), request.(DeleteUserByIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteUserByID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteUserByIDResponseObject); ok {
		return validResponse.VisitDeleteUserByIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetUserByID operation middleware
func (sh *strictHandler) GetUserByID(ctx echo.Context, id int64) error {
	var request GetUserByIDRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetUserByID(ctx.Request().Context(), request.(GetUserByIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUserByID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetUserByIDResponseObject); ok {
		return validResponse.VisitGetUserByIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PatchUserByID operation middleware
func (sh *strictHandler) PatchUserByID(ctx echo.Context, id int64) error {
	var request PatchUserByIDRequestObject

	request.Id = id

	var body PatchUserByIDJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchUserByID(ctx.Request().Context(), request.(PatchUserByIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchUserByID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchUserByIDResponseObject); ok {
		return validResponse.VisitPatchUserByIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
