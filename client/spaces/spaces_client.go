// Code generated by go-swagger; DO NOT EDIT.

package spaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new spaces API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for spaces API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateSpace(params *CreateSpaceParams, authInfo runtime.ClientAuthInfoWriter) (*CreateSpaceCreated, error)

	CreateSpaceLogo(params *CreateSpaceLogoParams, authInfo runtime.ClientAuthInfoWriter) (*CreateSpaceLogoOK, error)

	DeleteSpace(params *DeleteSpaceParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSpaceOK, error)

	GetSpaceByID(params *GetSpaceByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetSpaceByIDOK, error)

	GetSpaceLogo(params *GetSpaceLogoParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer) (*GetSpaceLogoOK, error)

	IndexSpaces(params *IndexSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*IndexSpacesOK, error)

	ListAllSpaces(params *ListAllSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllSpacesOK, error)

	UpdateModifySpace(params *UpdateModifySpaceParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateModifySpaceCreated, error)

	UpdateSpaceLogo(params *UpdateSpaceLogoParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSpaceLogoOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateSpace Creates a new Space.
*/
func (a *Client) CreateSpace(params *CreateSpaceParams, authInfo runtime.ClientAuthInfoWriter) (*CreateSpaceCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSpaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSpace",
		Method:             "POST",
		PathPattern:        "/api/spaces",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateSpaceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateSpaceCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createSpace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateSpaceLogo Updates the logo associated with the space.
*/
func (a *Client) CreateSpaceLogo(params *CreateSpaceLogoParams, authInfo runtime.ClientAuthInfoWriter) (*CreateSpaceLogoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSpaceLogoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSpaceLogo",
		Method:             "POST",
		PathPattern:        "/api/spaces/{id}/logo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateSpaceLogoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateSpaceLogoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createSpaceLogo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteSpace Deletes an existing space.
*/
func (a *Client) DeleteSpace(params *DeleteSpaceParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSpaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSpaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSpace",
		Method:             "DELETE",
		PathPattern:        "/api/spaces/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSpaceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteSpaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteSpace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSpaceByID gets a space resource by ID

  Get a space
*/
func (a *Client) GetSpaceByID(params *GetSpaceByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetSpaceByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSpaceByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSpaceById",
		Method:             "GET",
		PathPattern:        "/api/spaces/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSpaceByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSpaceByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSpaceById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSpaceLogo Gets the logo associated with the space.
*/
func (a *Client) GetSpaceLogo(params *GetSpaceLogoParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer) (*GetSpaceLogoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSpaceLogoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSpaceLogo",
		Method:             "GET",
		PathPattern:        "/api/spaces/{id}/logo",
		ProducesMediaTypes: []string{"image/png"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSpaceLogoReader{formats: a.formats, writer: writer},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSpaceLogoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSpaceLogo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IndexSpaces gets a list of space resources

  Lists all of the spaces in the supplied Octopus Deploy Space. The results will be sorted alphabetically by name.
*/
func (a *Client) IndexSpaces(params *IndexSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*IndexSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndexSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indexSpaces",
		Method:             "GET",
		PathPattern:        "/api/spaces",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IndexSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IndexSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for indexSpaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAllSpaces gets a list of space resources

  Lists all of the spaces in the supplied Octopus Deploy Space. The results will be sorted alphabetically by name.
*/
func (a *Client) ListAllSpaces(params *ListAllSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAllSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAllSpaces",
		Method:             "GET",
		PathPattern:        "/api/spaces/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListAllSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAllSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAllSpaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateModifySpace Modifies an existing space.
*/
func (a *Client) UpdateModifySpace(params *UpdateModifySpaceParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateModifySpaceCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateModifySpaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateModifySpace",
		Method:             "PUT",
		PathPattern:        "/api/spaces/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateModifySpaceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateModifySpaceCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateModifySpace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateSpaceLogo Updates the logo associated with the space.
*/
func (a *Client) UpdateSpaceLogo(params *UpdateSpaceLogoParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSpaceLogoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSpaceLogoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSpaceLogo",
		Method:             "PUT",
		PathPattern:        "/api/spaces/{id}/logo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateSpaceLogoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateSpaceLogoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateSpaceLogo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
