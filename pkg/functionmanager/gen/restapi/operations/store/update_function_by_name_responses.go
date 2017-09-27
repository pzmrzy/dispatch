///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.eng.vmware.com/serverless/serverless/pkg/functionmanager/gen/models"
)

// UpdateFunctionByNameOKCode is the HTTP code returned for type UpdateFunctionByNameOK
const UpdateFunctionByNameOKCode int = 200

/*UpdateFunctionByNameOK Successful update

swagger:response updateFunctionByNameOK
*/
type UpdateFunctionByNameOK struct {

	/*
	  In: Body
	*/
	Payload *models.Function `json:"body,omitempty"`
}

// NewUpdateFunctionByNameOK creates UpdateFunctionByNameOK with default headers values
func NewUpdateFunctionByNameOK() *UpdateFunctionByNameOK {
	return &UpdateFunctionByNameOK{}
}

// WithPayload adds the payload to the update function by name o k response
func (o *UpdateFunctionByNameOK) WithPayload(payload *models.Function) *UpdateFunctionByNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update function by name o k response
func (o *UpdateFunctionByNameOK) SetPayload(payload *models.Function) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateFunctionByNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateFunctionByNameBadRequestCode is the HTTP code returned for type UpdateFunctionByNameBadRequest
const UpdateFunctionByNameBadRequestCode int = 400

/*UpdateFunctionByNameBadRequest Invalid input

swagger:response updateFunctionByNameBadRequest
*/
type UpdateFunctionByNameBadRequest struct {
}

// NewUpdateFunctionByNameBadRequest creates UpdateFunctionByNameBadRequest with default headers values
func NewUpdateFunctionByNameBadRequest() *UpdateFunctionByNameBadRequest {
	return &UpdateFunctionByNameBadRequest{}
}

// WriteResponse to the client
func (o *UpdateFunctionByNameBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}

// UpdateFunctionByNameNotFoundCode is the HTTP code returned for type UpdateFunctionByNameNotFound
const UpdateFunctionByNameNotFoundCode int = 404

/*UpdateFunctionByNameNotFound Function not found

swagger:response updateFunctionByNameNotFound
*/
type UpdateFunctionByNameNotFound struct {
}

// NewUpdateFunctionByNameNotFound creates UpdateFunctionByNameNotFound with default headers values
func NewUpdateFunctionByNameNotFound() *UpdateFunctionByNameNotFound {
	return &UpdateFunctionByNameNotFound{}
}

// WriteResponse to the client
func (o *UpdateFunctionByNameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}
