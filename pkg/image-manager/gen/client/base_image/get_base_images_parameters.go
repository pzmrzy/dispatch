///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
// Code generated by go-swagger; DO NOT EDIT.

package base_image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetBaseImagesParams creates a new GetBaseImagesParams object
// with the default values initialized.
func NewGetBaseImagesParams() *GetBaseImagesParams {
	var ()
	return &GetBaseImagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBaseImagesParamsWithTimeout creates a new GetBaseImagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBaseImagesParamsWithTimeout(timeout time.Duration) *GetBaseImagesParams {
	var ()
	return &GetBaseImagesParams{

		timeout: timeout,
	}
}

// NewGetBaseImagesParamsWithContext creates a new GetBaseImagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBaseImagesParamsWithContext(ctx context.Context) *GetBaseImagesParams {
	var ()
	return &GetBaseImagesParams{

		Context: ctx,
	}
}

// NewGetBaseImagesParamsWithHTTPClient creates a new GetBaseImagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBaseImagesParamsWithHTTPClient(client *http.Client) *GetBaseImagesParams {
	var ()
	return &GetBaseImagesParams{
		HTTPClient: client,
	}
}

/*GetBaseImagesParams contains all the parameters to send to the API endpoint
for the get base images operation typically these are written to a http.Request
*/
type GetBaseImagesParams struct {

	/*Runtime
	  Base image runtime/language

	*/
	Runtime *string
	/*Tags
	  Filter on base image tags

	*/
	Tags []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get base images params
func (o *GetBaseImagesParams) WithTimeout(timeout time.Duration) *GetBaseImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get base images params
func (o *GetBaseImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get base images params
func (o *GetBaseImagesParams) WithContext(ctx context.Context) *GetBaseImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get base images params
func (o *GetBaseImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get base images params
func (o *GetBaseImagesParams) WithHTTPClient(client *http.Client) *GetBaseImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get base images params
func (o *GetBaseImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRuntime adds the runtime to the get base images params
func (o *GetBaseImagesParams) WithRuntime(runtime *string) *GetBaseImagesParams {
	o.SetRuntime(runtime)
	return o
}

// SetRuntime adds the runtime to the get base images params
func (o *GetBaseImagesParams) SetRuntime(runtime *string) {
	o.Runtime = runtime
}

// WithTags adds the tags to the get base images params
func (o *GetBaseImagesParams) WithTags(tags []string) *GetBaseImagesParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the get base images params
func (o *GetBaseImagesParams) SetTags(tags []string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *GetBaseImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Runtime != nil {

		// query param runtime
		var qrRuntime string
		if o.Runtime != nil {
			qrRuntime = *o.Runtime
		}
		qRuntime := qrRuntime
		if qRuntime != "" {
			if err := r.SetQueryParam("runtime", qRuntime); err != nil {
				return err
			}
		}

	}

	valuesTags := o.Tags

	joinedTags := swag.JoinByFormat(valuesTags, "multi")
	// query array param tags
	if err := r.SetQueryParam("tags", joinedTags...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
