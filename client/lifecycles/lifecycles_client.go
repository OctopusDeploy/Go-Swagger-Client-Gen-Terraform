// Code generated by go-swagger; DO NOT EDIT.

package lifecycles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new lifecycles API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for lifecycles API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateLifecycle(params *CreateLifecycleParams, authInfo runtime.ClientAuthInfoWriter) (*CreateLifecycleCreated, error)

	CreateLifecycleSpaces(params *CreateLifecycleSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*CreateLifecycleSpacesCreated, error)

	DeleteLifecycle(params *DeleteLifecycleParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLifecycleOK, error)

	DeleteLifecycleSpaces(params *DeleteLifecycleSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLifecycleSpacesOK, error)

	GetLifecycleByID(params *GetLifecycleByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetLifecycleByIDOK, error)

	GetLifecycleByIDSpaces(params *GetLifecycleByIDSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetLifecycleByIDSpacesOK, error)

	GetLifecyclePreview(params *GetLifecyclePreviewParams, authInfo runtime.ClientAuthInfoWriter) (*GetLifecyclePreviewOK, error)

	GetLifecyclePreviewSpaces(params *GetLifecyclePreviewSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetLifecyclePreviewSpacesOK, error)

	GetLifecycleProjects(params *GetLifecycleProjectsParams, authInfo runtime.ClientAuthInfoWriter) (*GetLifecycleProjectsOK, error)

	GetLifecycleProjectsSpaces(params *GetLifecycleProjectsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetLifecycleProjectsSpacesOK, error)

	IndexLifecycles(params *IndexLifecyclesParams, authInfo runtime.ClientAuthInfoWriter) (*IndexLifecyclesOK, error)

	IndexLifecyclesSpaces(params *IndexLifecyclesSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*IndexLifecyclesSpacesOK, error)

	ListAllLifecycles(params *ListAllLifecyclesParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllLifecyclesOK, error)

	ListAllLifecyclesSpaces(params *ListAllLifecyclesSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllLifecyclesSpacesOK, error)

	UpdateLifecycle(params *UpdateLifecycleParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateLifecycleOK, error)

	UpdateLifecycleSpaces(params *UpdateLifecycleSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateLifecycleSpacesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateLifecycle creates a lifecycle resource

  Creates a new lifecycle.
*/
func (a *Client) CreateLifecycle(params *CreateLifecycleParams, authInfo runtime.ClientAuthInfoWriter) (*CreateLifecycleCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateLifecycleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createLifecycle",
		Method:             "POST",
		PathPattern:        "/api/lifecycles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateLifecycleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateLifecycleCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createLifecycle: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateLifecycleSpaces creates a lifecycle resource

  Creates a new lifecycle.
*/
func (a *Client) CreateLifecycleSpaces(params *CreateLifecycleSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*CreateLifecycleSpacesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateLifecycleSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createLifecycle_Spaces",
		Method:             "POST",
		PathPattern:        "/api/{baseSpaceId}/lifecycles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateLifecycleSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateLifecycleSpacesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createLifecycle_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteLifecycle deletes a lifecycle resource by ID

  Deletes an existing lifecycle.
*/
func (a *Client) DeleteLifecycle(params *DeleteLifecycleParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLifecycleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLifecycleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteLifecycle",
		Method:             "DELETE",
		PathPattern:        "/api/lifecycles/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteLifecycleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteLifecycleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteLifecycle: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteLifecycleSpaces deletes a lifecycle resource by ID

  Deletes an existing lifecycle.
*/
func (a *Client) DeleteLifecycleSpaces(params *DeleteLifecycleSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLifecycleSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLifecycleSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteLifecycle_Spaces",
		Method:             "DELETE",
		PathPattern:        "/api/{baseSpaceId}/lifecycles/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteLifecycleSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteLifecycleSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteLifecycle_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetLifecycleByID gets a lifecycle resource by ID

  Gets a single lifecycle by ID.
*/
func (a *Client) GetLifecycleByID(params *GetLifecycleByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetLifecycleByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLifecycleByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLifecycleById",
		Method:             "GET",
		PathPattern:        "/api/lifecycles/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLifecycleByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLifecycleByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLifecycleById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetLifecycleByIDSpaces gets a lifecycle resource by ID

  Gets a single lifecycle by ID.
*/
func (a *Client) GetLifecycleByIDSpaces(params *GetLifecycleByIDSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetLifecycleByIDSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLifecycleByIDSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLifecycleById_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/lifecycles/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLifecycleByIDSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLifecycleByIDSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLifecycleById_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetLifecyclePreview Gets a single lifecycle by ID, as a preview.
*/
func (a *Client) GetLifecyclePreview(params *GetLifecyclePreviewParams, authInfo runtime.ClientAuthInfoWriter) (*GetLifecyclePreviewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLifecyclePreviewParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLifecyclePreview",
		Method:             "GET",
		PathPattern:        "/api/lifecycles/{id}/preview",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLifecyclePreviewReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLifecyclePreviewOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLifecyclePreview: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetLifecyclePreviewSpaces Gets a single lifecycle by ID, as a preview.
*/
func (a *Client) GetLifecyclePreviewSpaces(params *GetLifecyclePreviewSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetLifecyclePreviewSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLifecyclePreviewSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLifecyclePreview_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/lifecycles/{id}/preview",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLifecyclePreviewSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLifecyclePreviewSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLifecyclePreview_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetLifecycleProjects Gets a all projects that use this lifecycle.
*/
func (a *Client) GetLifecycleProjects(params *GetLifecycleProjectsParams, authInfo runtime.ClientAuthInfoWriter) (*GetLifecycleProjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLifecycleProjectsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLifecycleProjects",
		Method:             "GET",
		PathPattern:        "/api/lifecycles/{id}/projects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLifecycleProjectsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLifecycleProjectsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLifecycleProjects: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetLifecycleProjectsSpaces Gets a all projects that use this lifecycle.
*/
func (a *Client) GetLifecycleProjectsSpaces(params *GetLifecycleProjectsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetLifecycleProjectsSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLifecycleProjectsSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLifecycleProjects_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/lifecycles/{id}/projects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLifecycleProjectsSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLifecycleProjectsSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLifecycleProjects_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IndexLifecycles gets a list of lifecycle resources

  Lists all of the lifecycles in the supplied Octopus Deploy Space. The results will be sorted alphabetically by name.
*/
func (a *Client) IndexLifecycles(params *IndexLifecyclesParams, authInfo runtime.ClientAuthInfoWriter) (*IndexLifecyclesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndexLifecyclesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indexLifecycles",
		Method:             "GET",
		PathPattern:        "/api/lifecycles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IndexLifecyclesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IndexLifecyclesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for indexLifecycles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IndexLifecyclesSpaces gets a list of lifecycle resources

  Lists all of the lifecycles in the supplied Octopus Deploy Space. The results will be sorted alphabetically by name.
*/
func (a *Client) IndexLifecyclesSpaces(params *IndexLifecyclesSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*IndexLifecyclesSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndexLifecyclesSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indexLifecycles_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/lifecycles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IndexLifecyclesSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IndexLifecyclesSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for indexLifecycles_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAllLifecycles gets a list of lifecycle resources

  Lists all the lifecycles in the supplied Octopus Deploy Space.
*/
func (a *Client) ListAllLifecycles(params *ListAllLifecyclesParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllLifecyclesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAllLifecyclesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAllLifecycles",
		Method:             "GET",
		PathPattern:        "/api/lifecycles/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListAllLifecyclesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAllLifecyclesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAllLifecycles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAllLifecyclesSpaces gets a list of lifecycle resources

  Lists all the lifecycles in the supplied Octopus Deploy Space.
*/
func (a *Client) ListAllLifecyclesSpaces(params *ListAllLifecyclesSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllLifecyclesSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAllLifecyclesSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAllLifecycles_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/lifecycles/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListAllLifecyclesSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAllLifecyclesSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAllLifecycles_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateLifecycle modifies a lifecycle resource by ID

  Modifies an existing lifecycle.
*/
func (a *Client) UpdateLifecycle(params *UpdateLifecycleParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateLifecycleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateLifecycleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateLifecycle",
		Method:             "PUT",
		PathPattern:        "/api/lifecycles/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateLifecycleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateLifecycleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateLifecycle: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateLifecycleSpaces modifies a lifecycle resource by ID

  Modifies an existing lifecycle.
*/
func (a *Client) UpdateLifecycleSpaces(params *UpdateLifecycleSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateLifecycleSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateLifecycleSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateLifecycle_Spaces",
		Method:             "PUT",
		PathPattern:        "/api/{baseSpaceId}/lifecycles/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateLifecycleSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateLifecycleSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateLifecycle_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}