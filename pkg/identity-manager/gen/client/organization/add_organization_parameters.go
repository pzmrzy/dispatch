///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// NewAddOrganizationParams creates a new AddOrganizationParams object
// with the default values initialized.
func NewAddOrganizationParams() *AddOrganizationParams {
	var ()
	return &AddOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddOrganizationParamsWithTimeout creates a new AddOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddOrganizationParamsWithTimeout(timeout time.Duration) *AddOrganizationParams {
	var ()
	return &AddOrganizationParams{

		timeout: timeout,
	}
}

// NewAddOrganizationParamsWithContext creates a new AddOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddOrganizationParamsWithContext(ctx context.Context) *AddOrganizationParams {
	var ()
	return &AddOrganizationParams{

		Context: ctx,
	}
}

// NewAddOrganizationParamsWithHTTPClient creates a new AddOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddOrganizationParamsWithHTTPClient(client *http.Client) *AddOrganizationParams {
	var ()
	return &AddOrganizationParams{
		HTTPClient: client,
	}
}

/*AddOrganizationParams contains all the parameters to send to the API endpoint
for the add organization operation typically these are written to a http.Request
*/
type AddOrganizationParams struct {

	/*XDispatchOrg*/
	XDispatchOrg string
	/*Body
	  Organization Object

	*/
	Body *v1.Organization

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add organization params
func (o *AddOrganizationParams) WithTimeout(timeout time.Duration) *AddOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add organization params
func (o *AddOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add organization params
func (o *AddOrganizationParams) WithContext(ctx context.Context) *AddOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add organization params
func (o *AddOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add organization params
func (o *AddOrganizationParams) WithHTTPClient(client *http.Client) *AddOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add organization params
func (o *AddOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDispatchOrg adds the xDispatchOrg to the add organization params
func (o *AddOrganizationParams) WithXDispatchOrg(xDispatchOrg string) *AddOrganizationParams {
	o.SetXDispatchOrg(xDispatchOrg)
	return o
}

// SetXDispatchOrg adds the xDispatchOrg to the add organization params
func (o *AddOrganizationParams) SetXDispatchOrg(xDispatchOrg string) {
	o.XDispatchOrg = xDispatchOrg
}

// WithBody adds the body to the add organization params
func (o *AddOrganizationParams) WithBody(body *v1.Organization) *AddOrganizationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add organization params
func (o *AddOrganizationParams) SetBody(body *v1.Organization) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Dispatch-Org
	if err := r.SetHeaderParam("X-Dispatch-Org", o.XDispatchOrg); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
