// Code generated by go-swagger; DO NOT EDIT.

package performance_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new performance configuration API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for performance configuration API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetPerformanceConfiguration(params *GetPerformanceConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*GetPerformanceConfigurationOK, error)

	UpdatePerformanceConfiguration(params *UpdatePerformanceConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*UpdatePerformanceConfigurationOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetPerformanceConfiguration Get performance settings for the Octopus Server
*/
func (a *Client) GetPerformanceConfiguration(params *GetPerformanceConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*GetPerformanceConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPerformanceConfigurationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPerformanceConfiguration",
		Method:             "GET",
		PathPattern:        "/api/performanceconfiguration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPerformanceConfigurationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPerformanceConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPerformanceConfiguration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdatePerformanceConfiguration Updates the performance settings for the Octopus Server
*/
func (a *Client) UpdatePerformanceConfiguration(params *UpdatePerformanceConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*UpdatePerformanceConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePerformanceConfigurationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updatePerformanceConfiguration",
		Method:             "PUT",
		PathPattern:        "/api/performanceconfiguration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdatePerformanceConfigurationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdatePerformanceConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updatePerformanceConfiguration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
