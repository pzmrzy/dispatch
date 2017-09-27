///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"gitlab.eng.vmware.com/serverless/serverless/pkg/functionmanager/gen/restapi/operations/runner"
	"gitlab.eng.vmware.com/serverless/serverless/pkg/functionmanager/gen/restapi/operations/store"
)

// NewFunctionManagerAPI creates a new FunctionManager instance
func NewFunctionManagerAPI(spec *loads.Document) *FunctionManagerAPI {
	return &FunctionManagerAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		StoreAddFunctionHandler: store.AddFunctionHandlerFunc(func(params store.AddFunctionParams) middleware.Responder {
			return middleware.NotImplemented("operation StoreAddFunction has not yet been implemented")
		}),
		StoreDeleteFunctionByNameHandler: store.DeleteFunctionByNameHandlerFunc(func(params store.DeleteFunctionByNameParams) middleware.Responder {
			return middleware.NotImplemented("operation StoreDeleteFunctionByName has not yet been implemented")
		}),
		StoreGetFunctionByNameHandler: store.GetFunctionByNameHandlerFunc(func(params store.GetFunctionByNameParams) middleware.Responder {
			return middleware.NotImplemented("operation StoreGetFunctionByName has not yet been implemented")
		}),
		StoreGetFunctionsHandler: store.GetFunctionsHandlerFunc(func(params store.GetFunctionsParams) middleware.Responder {
			return middleware.NotImplemented("operation StoreGetFunctions has not yet been implemented")
		}),
		RunnerGetRunByNameHandler: runner.GetRunByNameHandlerFunc(func(params runner.GetRunByNameParams) middleware.Responder {
			return middleware.NotImplemented("operation RunnerGetRunByName has not yet been implemented")
		}),
		RunnerGetRunsHandler: runner.GetRunsHandlerFunc(func(params runner.GetRunsParams) middleware.Responder {
			return middleware.NotImplemented("operation RunnerGetRuns has not yet been implemented")
		}),
		RunnerRunFunctionHandler: runner.RunFunctionHandlerFunc(func(params runner.RunFunctionParams) middleware.Responder {
			return middleware.NotImplemented("operation RunnerRunFunction has not yet been implemented")
		}),
		StoreUpdateFunctionByNameHandler: store.UpdateFunctionByNameHandlerFunc(func(params store.UpdateFunctionByNameParams) middleware.Responder {
			return middleware.NotImplemented("operation StoreUpdateFunctionByName has not yet been implemented")
		}),
	}
}

/*FunctionManagerAPI This is the API server for the serverless function manager service.
 */
type FunctionManagerAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// StoreAddFunctionHandler sets the operation handler for the add function operation
	StoreAddFunctionHandler store.AddFunctionHandler
	// StoreDeleteFunctionByNameHandler sets the operation handler for the delete function by name operation
	StoreDeleteFunctionByNameHandler store.DeleteFunctionByNameHandler
	// StoreGetFunctionByNameHandler sets the operation handler for the get function by name operation
	StoreGetFunctionByNameHandler store.GetFunctionByNameHandler
	// StoreGetFunctionsHandler sets the operation handler for the get functions operation
	StoreGetFunctionsHandler store.GetFunctionsHandler
	// RunnerGetRunByNameHandler sets the operation handler for the get run by name operation
	RunnerGetRunByNameHandler runner.GetRunByNameHandler
	// RunnerGetRunsHandler sets the operation handler for the get runs operation
	RunnerGetRunsHandler runner.GetRunsHandler
	// RunnerRunFunctionHandler sets the operation handler for the run function operation
	RunnerRunFunctionHandler runner.RunFunctionHandler
	// StoreUpdateFunctionByNameHandler sets the operation handler for the update function by name operation
	StoreUpdateFunctionByNameHandler store.UpdateFunctionByNameHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *FunctionManagerAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *FunctionManagerAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *FunctionManagerAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *FunctionManagerAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *FunctionManagerAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *FunctionManagerAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *FunctionManagerAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the FunctionManagerAPI
func (o *FunctionManagerAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.StoreAddFunctionHandler == nil {
		unregistered = append(unregistered, "store.AddFunctionHandler")
	}

	if o.StoreDeleteFunctionByNameHandler == nil {
		unregistered = append(unregistered, "store.DeleteFunctionByNameHandler")
	}

	if o.StoreGetFunctionByNameHandler == nil {
		unregistered = append(unregistered, "store.GetFunctionByNameHandler")
	}

	if o.StoreGetFunctionsHandler == nil {
		unregistered = append(unregistered, "store.GetFunctionsHandler")
	}

	if o.RunnerGetRunByNameHandler == nil {
		unregistered = append(unregistered, "runner.GetRunByNameHandler")
	}

	if o.RunnerGetRunsHandler == nil {
		unregistered = append(unregistered, "runner.GetRunsHandler")
	}

	if o.RunnerRunFunctionHandler == nil {
		unregistered = append(unregistered, "runner.RunFunctionHandler")
	}

	if o.StoreUpdateFunctionByNameHandler == nil {
		unregistered = append(unregistered, "store.UpdateFunctionByNameHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *FunctionManagerAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *FunctionManagerAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *FunctionManagerAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *FunctionManagerAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *FunctionManagerAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *FunctionManagerAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the function manager API
func (o *FunctionManagerAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *FunctionManagerAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"][""] = store.NewAddFunction(o.context, o.StoreAddFunctionHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/{functionName}"] = store.NewDeleteFunctionByName(o.context, o.StoreDeleteFunctionByNameHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/{functionName}"] = store.NewGetFunctionByName(o.context, o.StoreGetFunctionByNameHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"][""] = store.NewGetFunctions(o.context, o.StoreGetFunctionsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/{functionName}/runs/{runName}"] = runner.NewGetRunByName(o.context, o.RunnerGetRunByNameHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/{functionName}/runs"] = runner.NewGetRuns(o.context, o.RunnerGetRunsHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/{functionName}/runs"] = runner.NewRunFunction(o.context, o.RunnerRunFunctionHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/{functionName}"] = store.NewUpdateFunctionByName(o.context, o.StoreUpdateFunctionByNameHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *FunctionManagerAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *FunctionManagerAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}
