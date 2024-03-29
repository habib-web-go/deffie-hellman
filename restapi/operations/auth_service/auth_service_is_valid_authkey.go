// Code generated by go-swagger; DO NOT EDIT.

package auth_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AuthServiceIsValidAuthkeyHandlerFunc turns a function with the right signature into a auth service is valid authkey handler
type AuthServiceIsValidAuthkeyHandlerFunc func(AuthServiceIsValidAuthkeyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AuthServiceIsValidAuthkeyHandlerFunc) Handle(params AuthServiceIsValidAuthkeyParams) middleware.Responder {
	return fn(params)
}

// AuthServiceIsValidAuthkeyHandler interface for that can handle valid auth service is valid authkey params
type AuthServiceIsValidAuthkeyHandler interface {
	Handle(AuthServiceIsValidAuthkeyParams) middleware.Responder
}

// NewAuthServiceIsValidAuthkey creates a new http.Handler for the auth service is valid authkey operation
func NewAuthServiceIsValidAuthkey(ctx *middleware.Context, handler AuthServiceIsValidAuthkeyHandler) *AuthServiceIsValidAuthkey {
	return &AuthServiceIsValidAuthkey{Context: ctx, Handler: handler}
}

/*
	AuthServiceIsValidAuthkey swagger:route GET /auth/is_valid_authkey/{authkey} AuthService authServiceIsValidAuthkey

AuthServiceIsValidAuthkey auth service is valid authkey API
*/
type AuthServiceIsValidAuthkey struct {
	Context *middleware.Context
	Handler AuthServiceIsValidAuthkeyHandler
}

func (o *AuthServiceIsValidAuthkey) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAuthServiceIsValidAuthkeyParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
