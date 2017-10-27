// Copyright 2017, Dell EMC, Inc.

// Code generated by go-swagger; DO NOT EDIT.

package switch_firmware

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// SwitchFirmwareHandlerFunc turns a function with the right signature into a switch firmware handler
type SwitchFirmwareHandlerFunc func(SwitchFirmwareParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn SwitchFirmwareHandlerFunc) Handle(params SwitchFirmwareParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// SwitchFirmwareHandler interface for that can handle valid switch firmware params
type SwitchFirmwareHandler interface {
	Handle(SwitchFirmwareParams, interface{}) middleware.Responder
}

// NewSwitchFirmware creates a new http.Handler for the switch firmware operation
func NewSwitchFirmware(ctx *middleware.Context, handler SwitchFirmwareHandler) *SwitchFirmware {
	return &SwitchFirmware{Context: ctx, Handler: handler}
}

/*SwitchFirmware swagger:route POST /switchFirmware /switchFirmware switchFirmware

Get switch Firmware Version

Get switch Firmware Version

*/
type SwitchFirmware struct {
	Context *middleware.Context
	Handler SwitchFirmwareHandler
}

func (o *SwitchFirmware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSwitchFirmwareParams()

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
