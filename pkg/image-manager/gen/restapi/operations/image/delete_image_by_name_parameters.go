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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteImageByNameParams creates a new DeleteImageByNameParams object
// with the default values initialized.
func NewDeleteImageByNameParams() DeleteImageByNameParams {
	var ()
	return DeleteImageByNameParams{}
}

// DeleteImageByNameParams contains all the bound params for the delete image by name operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteImageByName
type DeleteImageByNameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Name of image to return
	  Required: true
	  Pattern: ^[\w\d\-]+$
	  In: path
	*/
	ImageName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DeleteImageByNameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rImageName, rhkImageName, _ := route.Params.GetOK("imageName")
	if err := o.bindImageName(rImageName, rhkImageName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteImageByNameParams) bindImageName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ImageName = raw

	if err := o.validateImageName(formats); err != nil {
		return err
	}

	return nil
}

func (o *DeleteImageByNameParams) validateImageName(formats strfmt.Registry) error {

	if err := validate.Pattern("imageName", "path", o.ImageName, `^[\w\d\-]+$`); err != nil {
		return err
	}

	return nil
}
