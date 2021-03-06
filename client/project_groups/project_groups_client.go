// Code generated by go-swagger; DO NOT EDIT.

package project_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new project groups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for project groups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateProjectGroup(params *CreateProjectGroupParams, authInfo runtime.ClientAuthInfoWriter) (*CreateProjectGroupCreated, error)

	CreateProjectGroupSpaces(params *CreateProjectGroupSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*CreateProjectGroupSpacesCreated, error)

	DeleteProjectGroup(params *DeleteProjectGroupParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteProjectGroupOK, error)

	DeleteProjectGroupSpaces(params *DeleteProjectGroupSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteProjectGroupSpacesOK, error)

	GetProjectGroupByID(params *GetProjectGroupByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetProjectGroupByIDOK, error)

	GetProjectGroupByIDSpaces(params *GetProjectGroupByIDSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetProjectGroupByIDSpacesOK, error)

	IndexProjectGroupProjects(params *IndexProjectGroupProjectsParams, authInfo runtime.ClientAuthInfoWriter) (*IndexProjectGroupProjectsOK, error)

	IndexProjectGroupProjectsSpaces(params *IndexProjectGroupProjectsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*IndexProjectGroupProjectsSpacesOK, error)

	IndexProjectGroups(params *IndexProjectGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*IndexProjectGroupsOK, error)

	IndexProjectGroupsSpaces(params *IndexProjectGroupsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*IndexProjectGroupsSpacesOK, error)

	ListAllProjectGroups(params *ListAllProjectGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllProjectGroupsOK, error)

	ListAllProjectGroupsSpaces(params *ListAllProjectGroupsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllProjectGroupsSpacesOK, error)

	UpdateProjectGroup(params *UpdateProjectGroupParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateProjectGroupOK, error)

	UpdateProjectGroupSpaces(params *UpdateProjectGroupSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateProjectGroupSpacesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateProjectGroup creates a project group resource

  Creates a new project group.
*/
func (a *Client) CreateProjectGroup(params *CreateProjectGroupParams, authInfo runtime.ClientAuthInfoWriter) (*CreateProjectGroupCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateProjectGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createProjectGroup",
		Method:             "POST",
		PathPattern:        "/api/projectgroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateProjectGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateProjectGroupCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createProjectGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateProjectGroupSpaces creates a project group resource

  Creates a new project group.
*/
func (a *Client) CreateProjectGroupSpaces(params *CreateProjectGroupSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*CreateProjectGroupSpacesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateProjectGroupSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createProjectGroup_Spaces",
		Method:             "POST",
		PathPattern:        "/api/{baseSpaceId}/projectgroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateProjectGroupSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateProjectGroupSpacesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createProjectGroup_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteProjectGroup deletes a project group resource by ID

  Deletes an existing project group.
*/
func (a *Client) DeleteProjectGroup(params *DeleteProjectGroupParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteProjectGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProjectGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteProjectGroup",
		Method:             "DELETE",
		PathPattern:        "/api/projectgroups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteProjectGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteProjectGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteProjectGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteProjectGroupSpaces deletes a project group resource by ID

  Deletes an existing project group.
*/
func (a *Client) DeleteProjectGroupSpaces(params *DeleteProjectGroupSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteProjectGroupSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProjectGroupSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteProjectGroup_Spaces",
		Method:             "DELETE",
		PathPattern:        "/api/{baseSpaceId}/projectgroups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteProjectGroupSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteProjectGroupSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteProjectGroup_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetProjectGroupByID gets a project group resource by ID

  Gets a single project group by ID.
*/
func (a *Client) GetProjectGroupByID(params *GetProjectGroupByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetProjectGroupByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProjectGroupByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProjectGroupById",
		Method:             "GET",
		PathPattern:        "/api/projectgroups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetProjectGroupByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetProjectGroupByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getProjectGroupById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetProjectGroupByIDSpaces gets a project group resource by ID

  Gets a single project group by ID.
*/
func (a *Client) GetProjectGroupByIDSpaces(params *GetProjectGroupByIDSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetProjectGroupByIDSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProjectGroupByIDSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProjectGroupById_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/projectgroups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetProjectGroupByIDSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetProjectGroupByIDSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getProjectGroupById_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IndexProjectGroupProjects gets a list of project resources for the given project group resource

  Lists all of the projects that belong to the given project group.
*/
func (a *Client) IndexProjectGroupProjects(params *IndexProjectGroupProjectsParams, authInfo runtime.ClientAuthInfoWriter) (*IndexProjectGroupProjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndexProjectGroupProjectsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indexProjectGroupProjects",
		Method:             "GET",
		PathPattern:        "/api/projectgroups/{id}/projects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IndexProjectGroupProjectsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IndexProjectGroupProjectsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for indexProjectGroupProjects: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IndexProjectGroupProjectsSpaces gets a list of project resources for the given project group resource

  Lists all of the projects that belong to the given project group.
*/
func (a *Client) IndexProjectGroupProjectsSpaces(params *IndexProjectGroupProjectsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*IndexProjectGroupProjectsSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndexProjectGroupProjectsSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indexProjectGroupProjects_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/projectgroups/{id}/projects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IndexProjectGroupProjectsSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IndexProjectGroupProjectsSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for indexProjectGroupProjects_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IndexProjectGroups gets a list of project group resources

  Lists all of the project groups in the supplied Octopus Deploy Space. The results will be sorted alphabetically by name.
*/
func (a *Client) IndexProjectGroups(params *IndexProjectGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*IndexProjectGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndexProjectGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indexProjectGroups",
		Method:             "GET",
		PathPattern:        "/api/projectgroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IndexProjectGroupsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IndexProjectGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for indexProjectGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IndexProjectGroupsSpaces gets a list of project group resources

  Lists all of the project groups in the supplied Octopus Deploy Space. The results will be sorted alphabetically by name.
*/
func (a *Client) IndexProjectGroupsSpaces(params *IndexProjectGroupsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*IndexProjectGroupsSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndexProjectGroupsSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indexProjectGroups_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/projectgroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IndexProjectGroupsSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IndexProjectGroupsSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for indexProjectGroups_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAllProjectGroups gets a list of project group resources

  Lists the name and ID of all of the project groups in the supplied Octopus Deploy Space. The results will be sorted alphabetically by name.
*/
func (a *Client) ListAllProjectGroups(params *ListAllProjectGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllProjectGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAllProjectGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAllProjectGroups",
		Method:             "GET",
		PathPattern:        "/api/projectgroups/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListAllProjectGroupsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAllProjectGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAllProjectGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAllProjectGroupsSpaces gets a list of project group resources

  Lists the name and ID of all of the project groups in the supplied Octopus Deploy Space. The results will be sorted alphabetically by name.
*/
func (a *Client) ListAllProjectGroupsSpaces(params *ListAllProjectGroupsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllProjectGroupsSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAllProjectGroupsSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAllProjectGroups_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/projectgroups/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListAllProjectGroupsSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAllProjectGroupsSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAllProjectGroups_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateProjectGroup modifies a project group resource by ID

  Modifies an existing project group.
*/
func (a *Client) UpdateProjectGroup(params *UpdateProjectGroupParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateProjectGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateProjectGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateProjectGroup",
		Method:             "PUT",
		PathPattern:        "/api/projectgroups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateProjectGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateProjectGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateProjectGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateProjectGroupSpaces modifies a project group resource by ID

  Modifies an existing project group.
*/
func (a *Client) UpdateProjectGroupSpaces(params *UpdateProjectGroupSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateProjectGroupSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateProjectGroupSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateProjectGroup_Spaces",
		Method:             "PUT",
		PathPattern:        "/api/{baseSpaceId}/projectgroups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateProjectGroupSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateProjectGroupSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateProjectGroup_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
