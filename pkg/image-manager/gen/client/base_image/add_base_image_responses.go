///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
// Code generated by go-swagger; DO NOT EDIT.

package base_image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"gitlab.eng.vmware.com/serverless/serverless/pkg/image-manager/gen/models"
)

// AddBaseImageReader is a Reader for the AddBaseImage structure.
type AddBaseImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddBaseImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAddBaseImageCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddBaseImageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAddBaseImageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddBaseImageCreated creates a AddBaseImageCreated with default headers values
func NewAddBaseImageCreated() *AddBaseImageCreated {
	return &AddBaseImageCreated{}
}

/*AddBaseImageCreated handles this case with default header values.

created
*/
type AddBaseImageCreated struct {
	Payload *models.BaseImage
}

func (o *AddBaseImageCreated) Error() string {
	return fmt.Sprintf("[POST /base][%d] addBaseImageCreated  %+v", 201, o.Payload)
}

func (o *AddBaseImageCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BaseImage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddBaseImageBadRequest creates a AddBaseImageBadRequest with default headers values
func NewAddBaseImageBadRequest() *AddBaseImageBadRequest {
	return &AddBaseImageBadRequest{}
}

/*AddBaseImageBadRequest handles this case with default header values.

Invalid input
*/
type AddBaseImageBadRequest struct {
	Payload *models.Error
}

func (o *AddBaseImageBadRequest) Error() string {
	return fmt.Sprintf("[POST /base][%d] addBaseImageBadRequest  %+v", 400, o.Payload)
}

func (o *AddBaseImageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddBaseImageDefault creates a AddBaseImageDefault with default headers values
func NewAddBaseImageDefault(code int) *AddBaseImageDefault {
	return &AddBaseImageDefault{
		_statusCode: code,
	}
}

/*AddBaseImageDefault handles this case with default header values.

Generic error response
*/
type AddBaseImageDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the add base image default response
func (o *AddBaseImageDefault) Code() int {
	return o._statusCode
}

func (o *AddBaseImageDefault) Error() string {
	return fmt.Sprintf("[POST /base][%d] addBaseImage default  %+v", o._statusCode, o.Payload)
}

func (o *AddBaseImageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
