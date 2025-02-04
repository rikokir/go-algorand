// Package private provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package private

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Aborts a catchpoint catchup.
	// (DELETE /v2/catchup/{catchpoint})
	AbortCatchup(ctx echo.Context, catchpoint string) error
	// Starts a catchpoint catchup.
	// (POST /v2/catchup/{catchpoint})
	StartCatchup(ctx echo.Context, catchpoint string) error
	// Return a list of participation keys
	// (GET /v2/participation)
	GetParticipationKeys(ctx echo.Context) error
	// Add a participation key to the node
	// (POST /v2/participation)
	AddParticipationKey(ctx echo.Context) error
	// Delete a given participation key by ID
	// (DELETE /v2/participation/{participation-id})
	DeleteParticipationKeyByID(ctx echo.Context, participationId string) error
	// Get participation key info given a participation ID
	// (GET /v2/participation/{participation-id})
	GetParticipationKeyByID(ctx echo.Context, participationId string) error
	// Append state proof keys to a participation key
	// (POST /v2/participation/{participation-id})
	AppendKeys(ctx echo.Context, participationId string) error

	// (POST /v2/shutdown)
	ShutdownNode(ctx echo.Context, params ShutdownNodeParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AbortCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) AbortCatchup(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameter("simple", false, "catchpoint", ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AbortCatchup(ctx, catchpoint)
	return err
}

// StartCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) StartCatchup(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameter("simple", false, "catchpoint", ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.StartCatchup(ctx, catchpoint)
	return err
}

// GetParticipationKeys converts echo context to params.
func (w *ServerInterfaceWrapper) GetParticipationKeys(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetParticipationKeys(ctx)
	return err
}

// AddParticipationKey converts echo context to params.
func (w *ServerInterfaceWrapper) AddParticipationKey(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddParticipationKey(ctx)
	return err
}

// DeleteParticipationKeyByID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteParticipationKeyByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteParticipationKeyByID(ctx, participationId)
	return err
}

// GetParticipationKeyByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetParticipationKeyByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetParticipationKeyByID(ctx, participationId)
	return err
}

// AppendKeys converts echo context to params.
func (w *ServerInterfaceWrapper) AppendKeys(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AppendKeys(ctx, participationId)
	return err
}

// ShutdownNode converts echo context to params.
func (w *ServerInterfaceWrapper) ShutdownNode(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty":  true,
		"timeout": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ShutdownNodeParams
	// ------------- Optional query parameter "timeout" -------------
	if paramValue := ctx.QueryParam("timeout"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "timeout", ctx.QueryParams(), &params.Timeout)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timeout: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ShutdownNode(ctx, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.DELETE("/v2/catchup/:catchpoint", wrapper.AbortCatchup, m...)
	router.POST("/v2/catchup/:catchpoint", wrapper.StartCatchup, m...)
	router.GET("/v2/participation", wrapper.GetParticipationKeys, m...)
	router.POST("/v2/participation", wrapper.AddParticipationKey, m...)
	router.DELETE("/v2/participation/:participation-id", wrapper.DeleteParticipationKeyByID, m...)
	router.GET("/v2/participation/:participation-id", wrapper.GetParticipationKeyByID, m...)
	router.POST("/v2/participation/:participation-id", wrapper.AppendKeys, m...)
	router.POST("/v2/shutdown", wrapper.ShutdownNode, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
