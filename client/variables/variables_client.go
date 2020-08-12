// Code generated by go-swagger; DO NOT EDIT.

package variables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new variables API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for variables API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetVariableNamesList(params *GetVariableNamesListParams, authInfo runtime.ClientAuthInfoWriter) (*GetVariableNamesListOK, error)

	GetVariableNamesListSpaces(params *GetVariableNamesListSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetVariableNamesListSpacesOK, error)

	GetVariablePreviewList(params *GetVariablePreviewListParams, authInfo runtime.ClientAuthInfoWriter) (*GetVariablePreviewListOK, error)

	GetVariablePreviewListSpaces(params *GetVariablePreviewListSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetVariablePreviewListSpacesOK, error)

	GetVariableSetByID(params *GetVariableSetByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetVariableSetByIDOK, error)

	GetVariableSetByIDSpaces(params *GetVariableSetByIDSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetVariableSetByIDSpacesOK, error)

	ListAllVariableSets(params *ListAllVariableSetsParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllVariableSetsOK, error)

	ListAllVariableSetsSpaces(params *ListAllVariableSetsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllVariableSetsSpacesOK, error)

	UpdateVariableSet(params *UpdateVariableSetParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateVariableSetOK, error)

	UpdateVariableSetSpaces(params *UpdateVariableSetSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateVariableSetSpacesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetVariableNamesList List the names of variables that can be used in deployment actions. If a project is specified, this will include variables in that project.  If a project environments filter is specified, project variables only scoped to the environment will be excluded.
*/
func (a *Client) GetVariableNamesList(params *GetVariableNamesListParams, authInfo runtime.ClientAuthInfoWriter) (*GetVariableNamesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVariableNamesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVariableNamesList",
		Method:             "GET",
		PathPattern:        "/api/variables/names",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVariableNamesListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVariableNamesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVariableNamesList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVariableNamesListSpaces List the names of variables that can be used in deployment actions. If a project is specified, this will include variables in that project.  If a project environments filter is specified, project variables only scoped to the environment will be excluded.
*/
func (a *Client) GetVariableNamesListSpaces(params *GetVariableNamesListSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetVariableNamesListSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVariableNamesListSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVariableNamesList_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/variables/names",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVariableNamesListSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVariableNamesListSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVariableNamesList_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVariablePreviewList List the evaluated variables for a deployment.
*/
func (a *Client) GetVariablePreviewList(params *GetVariablePreviewListParams, authInfo runtime.ClientAuthInfoWriter) (*GetVariablePreviewListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVariablePreviewListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVariablePreviewList",
		Method:             "GET",
		PathPattern:        "/api/variables/preview",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVariablePreviewListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVariablePreviewListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVariablePreviewList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVariablePreviewListSpaces List the evaluated variables for a deployment.
*/
func (a *Client) GetVariablePreviewListSpaces(params *GetVariablePreviewListSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetVariablePreviewListSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVariablePreviewListSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVariablePreviewList_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/variables/preview",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVariablePreviewListSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVariablePreviewListSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVariablePreviewList_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVariableSetByID gets a variable set resource by ID

  Gets a variable set by Id.
*/
func (a *Client) GetVariableSetByID(params *GetVariableSetByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetVariableSetByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVariableSetByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVariableSetById",
		Method:             "GET",
		PathPattern:        "/api/variables/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVariableSetByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVariableSetByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVariableSetById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVariableSetByIDSpaces gets a variable set resource by ID

  Gets a variable set by Id.
*/
func (a *Client) GetVariableSetByIDSpaces(params *GetVariableSetByIDSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetVariableSetByIDSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVariableSetByIDSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVariableSetById_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/variables/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVariableSetByIDSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVariableSetByIDSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVariableSetById_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAllVariableSets gets a list of variable set resources

  Lists all the variable sets in the supplied Octopus Deploy Space.
*/
func (a *Client) ListAllVariableSets(params *ListAllVariableSetsParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllVariableSetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAllVariableSetsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAllVariableSets",
		Method:             "GET",
		PathPattern:        "/api/variables/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListAllVariableSetsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAllVariableSetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAllVariableSets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAllVariableSetsSpaces gets a list of variable set resources

  Lists all the variable sets in the supplied Octopus Deploy Space.
*/
func (a *Client) ListAllVariableSetsSpaces(params *ListAllVariableSetsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllVariableSetsSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAllVariableSetsSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAllVariableSets_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/variables/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListAllVariableSetsSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAllVariableSetsSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAllVariableSets_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateVariableSet Updates a variable set.
*/
func (a *Client) UpdateVariableSet(params *UpdateVariableSetParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateVariableSetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateVariableSetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateVariableSet",
		Method:             "PUT",
		PathPattern:        "/api/variables/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateVariableSetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateVariableSetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateVariableSet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateVariableSetSpaces Updates a variable set.
*/
func (a *Client) UpdateVariableSetSpaces(params *UpdateVariableSetSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateVariableSetSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateVariableSetSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateVariableSet_Spaces",
		Method:             "PUT",
		PathPattern:        "/api/{baseSpaceId}/variables/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateVariableSetSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateVariableSetSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateVariableSet_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}