// Code generated by go-swagger; DO NOT EDIT.

package user_teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user teams API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user teams API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetUserGetTeams(params *GetUserGetTeamsParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserGetTeamsOK, error)

	GetUserGetTeamsSpaces(params *GetUserGetTeamsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserGetTeamsSpacesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetUserGetTeams Gets a list of teams (id and name only) for the specified user, including any from external auth-provider security groups.
*/
func (a *Client) GetUserGetTeams(params *GetUserGetTeamsParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserGetTeamsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserGetTeamsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserGetTeams",
		Method:             "GET",
		PathPattern:        "/api/users/{id}/teams",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUserGetTeamsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUserGetTeamsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUserGetTeams: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUserGetTeamsSpaces Gets a list of teams (id and name only) for the specified user, including any from external auth-provider security groups.
*/
func (a *Client) GetUserGetTeamsSpaces(params *GetUserGetTeamsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserGetTeamsSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserGetTeamsSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserGetTeams_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/users/{id}/teams",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUserGetTeamsSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUserGetTeamsSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUserGetTeams_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
