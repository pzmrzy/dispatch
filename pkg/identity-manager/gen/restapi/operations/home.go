///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// HomeHandlerFunc turns a function with the right signature into a home handler
type HomeHandlerFunc func(HomeParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn HomeHandlerFunc) Handle(params HomeParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// HomeHandler interface for that can handle valid home params
type HomeHandler interface {
	Handle(HomeParams, interface{}) middleware.Responder
}

// NewHome creates a new http.Handler for the home operation
func NewHome(ctx *middleware.Context, handler HomeHandler) *Home {
	return &Home{Context: ctx, Handler: handler}
}

/*Home swagger:route GET /v1/iam/home home

an placehold home page, will be redirected to if successfully logged in

*/
type Home struct {
	Context *middleware.Context
	Handler HomeHandler
}

func (o *Home) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewHomeParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
