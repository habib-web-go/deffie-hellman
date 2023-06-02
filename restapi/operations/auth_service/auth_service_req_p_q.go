// Code generated by go-swagger; DO NOT EDIT.

package auth_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AuthServiceReqPQHandlerFunc turns a function with the right signature into a auth service req p q handler
type AuthServiceReqPQHandlerFunc func(AuthServiceReqPQParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AuthServiceReqPQHandlerFunc) Handle(params AuthServiceReqPQParams) middleware.Responder {
	return fn(params)
}

// AuthServiceReqPQHandler interface for that can handle valid auth service req p q params
type AuthServiceReqPQHandler interface {
	Handle(AuthServiceReqPQParams) middleware.Responder
}

// NewAuthServiceReqPQ creates a new http.Handler for the auth service req p q operation
func NewAuthServiceReqPQ(ctx *middleware.Context, handler AuthServiceReqPQHandler) *AuthServiceReqPQ {
	return &AuthServiceReqPQ{Context: ctx, Handler: handler}
}

/*
	AuthServiceReqPQ swagger:route POST /auth/pq AuthService authServiceReqPQ

AuthServiceReqPQ auth service req p q API
*/
type AuthServiceReqPQ struct {
	Context *middleware.Context
	Handler AuthServiceReqPQHandler
}

func (o *AuthServiceReqPQ) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAuthServiceReqPQParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}