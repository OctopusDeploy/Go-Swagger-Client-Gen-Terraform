// Code generated by go-swagger; DO NOT EDIT.

package user_onboarding

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user onboarding API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user onboarding API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetOnboarding(params *GetOnboardingParams, authInfo runtime.ClientAuthInfoWriter) (*GetOnboardingOK, error)

	GetOnboardingSpaces(params *GetOnboardingSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetOnboardingSpacesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetOnboarding Gets information about how far the Octopus Server is towards having its first release deployed.
*/
func (a *Client) GetOnboarding(params *GetOnboardingParams, authInfo runtime.ClientAuthInfoWriter) (*GetOnboardingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOnboardingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOnboarding",
		Method:             "GET",
		PathPattern:        "/api/useronboarding",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetOnboardingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOnboardingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOnboarding: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOnboardingSpaces Gets information about how far the Octopus Server is towards having its first release deployed.
*/
func (a *Client) GetOnboardingSpaces(params *GetOnboardingSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetOnboardingSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOnboardingSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOnboarding_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/useronboarding",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetOnboardingSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOnboardingSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOnboarding_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
