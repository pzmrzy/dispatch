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
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"gitlab.eng.vmware.com/serverless/serverless/pkg/image-manager/gen/models"
)

// NewUpdateImageByNameParams creates a new UpdateImageByNameParams object
// with the default values initialized.
func NewUpdateImageByNameParams() UpdateImageByNameParams {
	var ()
	return UpdateImageByNameParams{}
}

// UpdateImageByNameParams contains all the bound params for the update image by name operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateImageByName
type UpdateImageByNameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*
	  In: body
	*/
	Body *models.Image
	/*Name of image to return
	  Required: true
	  Pattern: ^[\w\d\-]+$
	  In: path
	*/
	ImageName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *UpdateImageByNameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Image
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}

	}

	rImageName, rhkImageName, _ := route.Params.GetOK("imageName")
	if err := o.bindImageName(rImageName, rhkImageName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateImageByNameParams) bindImageName(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *UpdateImageByNameParams) validateImageName(formats strfmt.Registry) error {

	if err := validate.Pattern("imageName", "path", o.ImageName, `^[\w\d\-]+$`); err != nil {
		return err
	}

	return nil
}
