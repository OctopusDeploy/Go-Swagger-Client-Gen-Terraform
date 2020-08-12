// Code generated by go-swagger; DO NOT EDIT.

package machines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new machines API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for machines API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateDeploymentTarget(params *CreateDeploymentTargetParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDeploymentTargetCreated, error)

	CreateDeploymentTargetSpaces(params *CreateDeploymentTargetSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDeploymentTargetSpacesCreated, error)

	DeleteDeploymentTarget(params *DeleteDeploymentTargetParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteDeploymentTargetOK, error)

	DeleteDeploymentTargetSpaces(params *DeleteDeploymentTargetSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteDeploymentTargetSpacesOK, error)

	GetDeploymentTargetByID(params *GetDeploymentTargetByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeploymentTargetByIDOK, error)

	GetDeploymentTargetByIDSpaces(params *GetDeploymentTargetByIDSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeploymentTargetByIDSpacesOK, error)

	GetDeploymentTargetConnectionStatus(params *GetDeploymentTargetConnectionStatusParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeploymentTargetConnectionStatusOK, error)

	GetDeploymentTargetConnectionStatusSpaces(params *GetDeploymentTargetConnectionStatusSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeploymentTargetConnectionStatusSpacesOK, error)

	GetDeploymentTargetOperatingSystemNamesListAll(params *GetDeploymentTargetOperatingSystemNamesListAllParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeploymentTargetOperatingSystemNamesListAllOK, error)

	GetDeploymentTargetOperatingSystemNamesListAllSpaces(params *GetDeploymentTargetOperatingSystemNamesListAllSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeploymentTargetOperatingSystemNamesListAllSpacesOK, error)

	GetDeploymentTargetOperatingSystemShellNameListAll(params *GetDeploymentTargetOperatingSystemShellNameListAllParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeploymentTargetOperatingSystemShellNameListAllOK, error)

	GetDeploymentTargetOperatingSystemShellNameListAllSpaces(params *GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeploymentTargetOperatingSystemShellNameListAllSpacesOK, error)

	GetDiscoverDeploymentTarget(params *GetDiscoverDeploymentTargetParams, authInfo runtime.ClientAuthInfoWriter) (*GetDiscoverDeploymentTargetOK, error)

	GetDiscoverDeploymentTargetSpaces(params *GetDiscoverDeploymentTargetSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetDiscoverDeploymentTargetSpacesOK, error)

	IndexDeploymentTargetTasks(params *IndexDeploymentTargetTasksParams, authInfo runtime.ClientAuthInfoWriter) (*IndexDeploymentTargetTasksOK, error)

	IndexDeploymentTargetTasksSpaces(params *IndexDeploymentTargetTasksSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*IndexDeploymentTargetTasksSpacesOK, error)

	IndexDeploymentTargets(params *IndexDeploymentTargetsParams, authInfo runtime.ClientAuthInfoWriter) (*IndexDeploymentTargetsOK, error)

	IndexDeploymentTargetsSpaces(params *IndexDeploymentTargetsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*IndexDeploymentTargetsSpacesOK, error)

	ListAllDeploymentTargets(params *ListAllDeploymentTargetsParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllDeploymentTargetsOK, error)

	ListAllDeploymentTargetsSpaces(params *ListAllDeploymentTargetsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllDeploymentTargetsSpacesOK, error)

	UpdateDeploymentTarget(params *UpdateDeploymentTargetParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateDeploymentTargetOK, error)

	UpdateDeploymentTargetSpaces(params *UpdateDeploymentTargetSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateDeploymentTargetSpacesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateDeploymentTarget creates a deployment target resource

  Creates a new machine.
*/
func (a *Client) CreateDeploymentTarget(params *CreateDeploymentTargetParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDeploymentTargetCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDeploymentTargetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createDeploymentTarget",
		Method:             "POST",
		PathPattern:        "/api/machines",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateDeploymentTargetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateDeploymentTargetCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createDeploymentTarget: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateDeploymentTargetSpaces creates a deployment target resource

  Creates a new machine.
*/
func (a *Client) CreateDeploymentTargetSpaces(params *CreateDeploymentTargetSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDeploymentTargetSpacesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDeploymentTargetSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createDeploymentTarget_Spaces",
		Method:             "POST",
		PathPattern:        "/api/{baseSpaceId}/machines",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateDeploymentTargetSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateDeploymentTargetSpacesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createDeploymentTarget_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteDeploymentTarget deletes a deployment target resource by ID

  Deletes an existing machine.
*/
func (a *Client) DeleteDeploymentTarget(params *DeleteDeploymentTargetParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteDeploymentTargetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDeploymentTargetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteDeploymentTarget",
		Method:             "DELETE",
		PathPattern:        "/api/machines/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteDeploymentTargetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteDeploymentTargetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteDeploymentTarget: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteDeploymentTargetSpaces deletes a deployment target resource by ID

  Deletes an existing machine.
*/
func (a *Client) DeleteDeploymentTargetSpaces(params *DeleteDeploymentTargetSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteDeploymentTargetSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDeploymentTargetSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteDeploymentTarget_Spaces",
		Method:             "DELETE",
		PathPattern:        "/api/{baseSpaceId}/machines/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteDeploymentTargetSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteDeploymentTargetSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteDeploymentTarget_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDeploymentTargetByID gets a deployment target resource by ID

  Gets a single machine by ID.
*/
func (a *Client) GetDeploymentTargetByID(params *GetDeploymentTargetByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeploymentTargetByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentTargetByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDeploymentTargetById",
		Method:             "GET",
		PathPattern:        "/api/machines/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDeploymentTargetByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeploymentTargetByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeploymentTargetById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDeploymentTargetByIDSpaces gets a deployment target resource by ID

  Gets a single machine by ID.
*/
func (a *Client) GetDeploymentTargetByIDSpaces(params *GetDeploymentTargetByIDSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeploymentTargetByIDSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentTargetByIDSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDeploymentTargetById_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/machines/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDeploymentTargetByIDSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeploymentTargetByIDSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeploymentTargetById_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDeploymentTargetConnectionStatus Get the status of the network connection between the Octopus server and a machine.
*/
func (a *Client) GetDeploymentTargetConnectionStatus(params *GetDeploymentTargetConnectionStatusParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeploymentTargetConnectionStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentTargetConnectionStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDeploymentTargetConnectionStatus",
		Method:             "GET",
		PathPattern:        "/api/machines/{id}/connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDeploymentTargetConnectionStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeploymentTargetConnectionStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeploymentTargetConnectionStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDeploymentTargetConnectionStatusSpaces Get the status of the network connection between the Octopus server and a machine.
*/
func (a *Client) GetDeploymentTargetConnectionStatusSpaces(params *GetDeploymentTargetConnectionStatusSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeploymentTargetConnectionStatusSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentTargetConnectionStatusSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDeploymentTargetConnectionStatus_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/machines/{id}/connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDeploymentTargetConnectionStatusSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeploymentTargetConnectionStatusSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeploymentTargetConnectionStatus_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDeploymentTargetOperatingSystemNamesListAll Gets all operating system names for deployment targets. The result will be a string array.
*/
func (a *Client) GetDeploymentTargetOperatingSystemNamesListAll(params *GetDeploymentTargetOperatingSystemNamesListAllParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeploymentTargetOperatingSystemNamesListAllOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentTargetOperatingSystemNamesListAllParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDeploymentTargetOperatingSystemNamesListAll",
		Method:             "GET",
		PathPattern:        "/api/machines/operatingsystem/names/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDeploymentTargetOperatingSystemNamesListAllReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeploymentTargetOperatingSystemNamesListAllOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeploymentTargetOperatingSystemNamesListAll: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDeploymentTargetOperatingSystemNamesListAllSpaces Gets all operating system names for deployment targets. The result will be a string array.
*/
func (a *Client) GetDeploymentTargetOperatingSystemNamesListAllSpaces(params *GetDeploymentTargetOperatingSystemNamesListAllSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeploymentTargetOperatingSystemNamesListAllSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentTargetOperatingSystemNamesListAllSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDeploymentTargetOperatingSystemNamesListAll_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/machines/operatingsystem/names/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDeploymentTargetOperatingSystemNamesListAllSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeploymentTargetOperatingSystemNamesListAllSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeploymentTargetOperatingSystemNamesListAll_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDeploymentTargetOperatingSystemShellNameListAll Gets all operating system shell names for deployment targets. The result will be a string array.
*/
func (a *Client) GetDeploymentTargetOperatingSystemShellNameListAll(params *GetDeploymentTargetOperatingSystemShellNameListAllParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeploymentTargetOperatingSystemShellNameListAllOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentTargetOperatingSystemShellNameListAllParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDeploymentTargetOperatingSystemShellNameListAll",
		Method:             "GET",
		PathPattern:        "/api/machines/operatingsystem/shells/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDeploymentTargetOperatingSystemShellNameListAllReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeploymentTargetOperatingSystemShellNameListAllOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeploymentTargetOperatingSystemShellNameListAll: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDeploymentTargetOperatingSystemShellNameListAllSpaces Gets all operating system shell names for deployment targets. The result will be a string array.
*/
func (a *Client) GetDeploymentTargetOperatingSystemShellNameListAllSpaces(params *GetDeploymentTargetOperatingSystemShellNameListAllSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeploymentTargetOperatingSystemShellNameListAllSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentTargetOperatingSystemShellNameListAllSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDeploymentTargetOperatingSystemShellNameListAll_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/machines/operatingsystem/shells/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDeploymentTargetOperatingSystemShellNameListAllSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeploymentTargetOperatingSystemShellNameListAllSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeploymentTargetOperatingSystemShellNameListAll_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDiscoverDeploymentTarget Interrogate a machine for communication details so that it may be added to the installation.
*/
func (a *Client) GetDiscoverDeploymentTarget(params *GetDiscoverDeploymentTargetParams, authInfo runtime.ClientAuthInfoWriter) (*GetDiscoverDeploymentTargetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDiscoverDeploymentTargetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDiscoverDeploymentTarget",
		Method:             "GET",
		PathPattern:        "/api/machines/discover",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDiscoverDeploymentTargetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDiscoverDeploymentTargetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDiscoverDeploymentTarget: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDiscoverDeploymentTargetSpaces Interrogate a machine for communication details so that it may be added to the installation.
*/
func (a *Client) GetDiscoverDeploymentTargetSpaces(params *GetDiscoverDeploymentTargetSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetDiscoverDeploymentTargetSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDiscoverDeploymentTargetSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDiscoverDeploymentTarget_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/machines/discover",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDiscoverDeploymentTargetSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDiscoverDeploymentTargetSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDiscoverDeploymentTarget_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IndexDeploymentTargetTasks gets a list of task resources for the given deployment target resource

  Get the history of related tasks for a machine.
*/
func (a *Client) IndexDeploymentTargetTasks(params *IndexDeploymentTargetTasksParams, authInfo runtime.ClientAuthInfoWriter) (*IndexDeploymentTargetTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndexDeploymentTargetTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indexDeploymentTargetTasks",
		Method:             "GET",
		PathPattern:        "/api/machines/{id}/tasks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IndexDeploymentTargetTasksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IndexDeploymentTargetTasksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for indexDeploymentTargetTasks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IndexDeploymentTargetTasksSpaces gets a list of task resources for the given deployment target resource

  Get the history of related tasks for a machine.
*/
func (a *Client) IndexDeploymentTargetTasksSpaces(params *IndexDeploymentTargetTasksSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*IndexDeploymentTargetTasksSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndexDeploymentTargetTasksSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indexDeploymentTargetTasks_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/machines/{id}/tasks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IndexDeploymentTargetTasksSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IndexDeploymentTargetTasksSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for indexDeploymentTargetTasks_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IndexDeploymentTargets gets a list of deployment target resources

  Lists all of the registered machines in the supplied Octopus Deploy Space, from all environments. The results will be sorted alphabetically by name.
*/
func (a *Client) IndexDeploymentTargets(params *IndexDeploymentTargetsParams, authInfo runtime.ClientAuthInfoWriter) (*IndexDeploymentTargetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndexDeploymentTargetsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indexDeploymentTargets",
		Method:             "GET",
		PathPattern:        "/api/machines",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IndexDeploymentTargetsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IndexDeploymentTargetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for indexDeploymentTargets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IndexDeploymentTargetsSpaces gets a list of deployment target resources

  Lists all of the registered machines in the supplied Octopus Deploy Space, from all environments. The results will be sorted alphabetically by name.
*/
func (a *Client) IndexDeploymentTargetsSpaces(params *IndexDeploymentTargetsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*IndexDeploymentTargetsSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndexDeploymentTargetsSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indexDeploymentTargets_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/machines",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IndexDeploymentTargetsSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IndexDeploymentTargetsSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for indexDeploymentTargets_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAllDeploymentTargets gets a list of deployment target resources

  Lists all of the deployment targets in the supplied Octopus Deploy Space. The results will be sorted alphabetically by name.
*/
func (a *Client) ListAllDeploymentTargets(params *ListAllDeploymentTargetsParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllDeploymentTargetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAllDeploymentTargetsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAllDeploymentTargets",
		Method:             "GET",
		PathPattern:        "/api/machines/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListAllDeploymentTargetsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAllDeploymentTargetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAllDeploymentTargets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAllDeploymentTargetsSpaces gets a list of deployment target resources

  Lists all of the deployment targets in the supplied Octopus Deploy Space. The results will be sorted alphabetically by name.
*/
func (a *Client) ListAllDeploymentTargetsSpaces(params *ListAllDeploymentTargetsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllDeploymentTargetsSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAllDeploymentTargetsSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAllDeploymentTargets_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/machines/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListAllDeploymentTargetsSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAllDeploymentTargetsSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAllDeploymentTargets_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateDeploymentTarget modifies a deployment target resource by ID

  Modifies an existing machine.
*/
func (a *Client) UpdateDeploymentTarget(params *UpdateDeploymentTargetParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateDeploymentTargetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDeploymentTargetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateDeploymentTarget",
		Method:             "PUT",
		PathPattern:        "/api/machines/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateDeploymentTargetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateDeploymentTargetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateDeploymentTarget: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateDeploymentTargetSpaces modifies a deployment target resource by ID

  Modifies an existing machine.
*/
func (a *Client) UpdateDeploymentTargetSpaces(params *UpdateDeploymentTargetSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateDeploymentTargetSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDeploymentTargetSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateDeploymentTarget_Spaces",
		Method:             "PUT",
		PathPattern:        "/api/{baseSpaceId}/machines/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateDeploymentTargetSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateDeploymentTargetSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateDeploymentTarget_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}