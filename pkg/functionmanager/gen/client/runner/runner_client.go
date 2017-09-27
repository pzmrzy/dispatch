///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
// Code generated by go-swagger; DO NOT EDIT.

package runner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new runner API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for runner API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetRunByName gets function run by its name
*/
func (a *Client) GetRunByName(params *GetRunByNameParams) (*GetRunByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRunByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRunByName",
		Method:             "GET",
		PathPattern:        "/{functionName}/runs/{runName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRunByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRunByNameOK), nil

}

/*
GetRuns gets function runs that are being executed
*/
func (a *Client) GetRuns(params *GetRunsParams) (*GetRunsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRunsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRuns",
		Method:             "GET",
		PathPattern:        "/{functionName}/runs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRunsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRunsOK), nil

}

/*
RunFunction runs a function
*/
func (a *Client) RunFunction(params *RunFunctionParams) (*RunFunctionOK, *RunFunctionAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunFunctionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "runFunction",
		Method:             "POST",
		PathPattern:        "/{functionName}/runs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RunFunctionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *RunFunctionOK:
		return value, nil, nil
	case *RunFunctionAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
