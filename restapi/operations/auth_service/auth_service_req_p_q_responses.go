// Code generated by go-swagger; DO NOT EDIT.

package auth_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/habib-web-go/diffie-hellman/models"
)

// AuthServiceReqPQOKCode is the HTTP code returned for type AuthServiceReqPQOK
const AuthServiceReqPQOKCode int = 200

/*
AuthServiceReqPQOK A successful response.

swagger:response authServiceReqPQOK
*/
type AuthServiceReqPQOK struct {

	/*
	  In: Body
	*/
	Payload *models.AuthDHReqPQResponse `json:"body,omitempty"`
}

// NewAuthServiceReqPQOK creates AuthServiceReqPQOK with default headers values
func NewAuthServiceReqPQOK() *AuthServiceReqPQOK {

	return &AuthServiceReqPQOK{}
}

// WithPayload adds the payload to the auth service req p q o k response
func (o *AuthServiceReqPQOK) WithPayload(payload *models.AuthDHReqPQResponse) *AuthServiceReqPQOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the auth service req p q o k response
func (o *AuthServiceReqPQOK) SetPayload(payload *models.AuthDHReqPQResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AuthServiceReqPQOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
AuthServiceReqPQDefault An unexpected error response.

swagger:response authServiceReqPQDefault
*/
type AuthServiceReqPQDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.RuntimeError `json:"body,omitempty"`
}

// NewAuthServiceReqPQDefault creates AuthServiceReqPQDefault with default headers values
func NewAuthServiceReqPQDefault(code int) *AuthServiceReqPQDefault {
	if code <= 0 {
		code = 500
	}

	return &AuthServiceReqPQDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the auth service req p q default response
func (o *AuthServiceReqPQDefault) WithStatusCode(code int) *AuthServiceReqPQDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the auth service req p q default response
func (o *AuthServiceReqPQDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the auth service req p q default response
func (o *AuthServiceReqPQDefault) WithPayload(payload *models.RuntimeError) *AuthServiceReqPQDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the auth service req p q default response
func (o *AuthServiceReqPQDefault) SetPayload(payload *models.RuntimeError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AuthServiceReqPQDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
