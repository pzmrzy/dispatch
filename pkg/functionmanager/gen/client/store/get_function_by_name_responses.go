///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"gitlab.eng.vmware.com/serverless/serverless/pkg/functionmanager/gen/models"
)

// GetFunctionByNameReader is a Reader for the GetFunctionByName structure.
type GetFunctionByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFunctionByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFunctionByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetFunctionByNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetFunctionByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFunctionByNameOK creates a GetFunctionByNameOK with default headers values
func NewGetFunctionByNameOK() *GetFunctionByNameOK {
	return &GetFunctionByNameOK{}
}

/*GetFunctionByNameOK handles this case with default header values.

Successful operation
*/
type GetFunctionByNameOK struct {
	Payload *models.Function
}

func (o *GetFunctionByNameOK) Error() string {
	return fmt.Sprintf("[GET /{functionName}][%d] getFunctionByNameOK  %+v", 200, o.Payload)
}

func (o *GetFunctionByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Function)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFunctionByNameBadRequest creates a GetFunctionByNameBadRequest with default headers values
func NewGetFunctionByNameBadRequest() *GetFunctionByNameBadRequest {
	return &GetFunctionByNameBadRequest{}
}

/*GetFunctionByNameBadRequest handles this case with default header values.

Invalid Name supplied
*/
type GetFunctionByNameBadRequest struct {
}

func (o *GetFunctionByNameBadRequest) Error() string {
	return fmt.Sprintf("[GET /{functionName}][%d] getFunctionByNameBadRequest ", 400)
}

func (o *GetFunctionByNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFunctionByNameNotFound creates a GetFunctionByNameNotFound with default headers values
func NewGetFunctionByNameNotFound() *GetFunctionByNameNotFound {
	return &GetFunctionByNameNotFound{}
}

/*GetFunctionByNameNotFound handles this case with default header values.

Function not found
*/
type GetFunctionByNameNotFound struct {
}

func (o *GetFunctionByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /{functionName}][%d] getFunctionByNameNotFound ", 404)
}

func (o *GetFunctionByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
