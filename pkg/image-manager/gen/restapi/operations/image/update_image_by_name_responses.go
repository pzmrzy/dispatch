///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.eng.vmware.com/serverless/serverless/pkg/image-manager/gen/models"
)

// UpdateImageByNameOKCode is the HTTP code returned for type UpdateImageByNameOK
const UpdateImageByNameOKCode int = 200

/*UpdateImageByNameOK updated

swagger:response updateImageByNameOK
*/
type UpdateImageByNameOK struct {

	/*
	  In: Body
	*/
	Payload *models.BaseImage `json:"body,omitempty"`
}

// NewUpdateImageByNameOK creates UpdateImageByNameOK with default headers values
func NewUpdateImageByNameOK() *UpdateImageByNameOK {
	return &UpdateImageByNameOK{}
}

// WithPayload adds the payload to the update image by name o k response
func (o *UpdateImageByNameOK) WithPayload(payload *models.BaseImage) *UpdateImageByNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update image by name o k response
func (o *UpdateImageByNameOK) SetPayload(payload *models.BaseImage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateImageByNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateImageByNameBadRequestCode is the HTTP code returned for type UpdateImageByNameBadRequest
const UpdateImageByNameBadRequestCode int = 400

/*UpdateImageByNameBadRequest Invalid input

swagger:response updateImageByNameBadRequest
*/
type UpdateImageByNameBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateImageByNameBadRequest creates UpdateImageByNameBadRequest with default headers values
func NewUpdateImageByNameBadRequest() *UpdateImageByNameBadRequest {
	return &UpdateImageByNameBadRequest{}
}

// WithPayload adds the payload to the update image by name bad request response
func (o *UpdateImageByNameBadRequest) WithPayload(payload *models.Error) *UpdateImageByNameBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update image by name bad request response
func (o *UpdateImageByNameBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateImageByNameBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateImageByNameNotFoundCode is the HTTP code returned for type UpdateImageByNameNotFound
const UpdateImageByNameNotFoundCode int = 404

/*UpdateImageByNameNotFound Image not found

swagger:response updateImageByNameNotFound
*/
type UpdateImageByNameNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateImageByNameNotFound creates UpdateImageByNameNotFound with default headers values
func NewUpdateImageByNameNotFound() *UpdateImageByNameNotFound {
	return &UpdateImageByNameNotFound{}
}

// WithPayload adds the payload to the update image by name not found response
func (o *UpdateImageByNameNotFound) WithPayload(payload *models.Error) *UpdateImageByNameNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update image by name not found response
func (o *UpdateImageByNameNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateImageByNameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateImageByNameDefault Generic error response

swagger:response updateImageByNameDefault
*/
type UpdateImageByNameDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateImageByNameDefault creates UpdateImageByNameDefault with default headers values
func NewUpdateImageByNameDefault(code int) *UpdateImageByNameDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateImageByNameDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update image by name default response
func (o *UpdateImageByNameDefault) WithStatusCode(code int) *UpdateImageByNameDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update image by name default response
func (o *UpdateImageByNameDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update image by name default response
func (o *UpdateImageByNameDefault) WithPayload(payload *models.Error) *UpdateImageByNameDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update image by name default response
func (o *UpdateImageByNameDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateImageByNameDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
