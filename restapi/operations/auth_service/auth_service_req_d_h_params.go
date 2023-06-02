// Code generated by go-swagger; DO NOT EDIT.

package auth_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AuthServiceReqDHParamsHandlerFunc turns a function with the right signature into a auth service req d h params handler
type AuthServiceReqDHParamsHandlerFunc func(AuthServiceReqDHParamsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AuthServiceReqDHParamsHandlerFunc) Handle(params AuthServiceReqDHParamsParams) middleware.Responder {
	return fn(params)
}

// AuthServiceReqDHParamsHandler interface for that can handle valid auth service req d h params params
type AuthServiceReqDHParamsHandler interface {
	Handle(AuthServiceReqDHParamsParams) middleware.Responder
}

// NewAuthServiceReqDHParams creates a new http.Handler for the auth service req d h params operation
func NewAuthServiceReqDHParams(ctx *middleware.Context, handler AuthServiceReqDHParamsHandler) *AuthServiceReqDHParams {
	return &AuthServiceReqDHParams{Context: ctx, Handler: handler}
}

/*
	AuthServiceReqDHParams swagger:route POST /auth/dh AuthService authServiceReqDHParams

AuthServiceReqDHParams auth service req d h params API
*/
type AuthServiceReqDHParams struct {
	Context *middleware.Context
	Handler AuthServiceReqDHParamsHandler
}

func (o *AuthServiceReqDHParams) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAuthServiceReqDHParamsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
