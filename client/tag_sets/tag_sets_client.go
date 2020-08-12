// Code generated by go-swagger; DO NOT EDIT.

package tag_sets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new tag sets API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tag sets API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateTagSet(params *CreateTagSetParams, authInfo runtime.ClientAuthInfoWriter) (*CreateTagSetCreated, error)

	CreateTagSetSpaces(params *CreateTagSetSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*CreateTagSetSpacesCreated, error)

	DeleteTagSet(params *DeleteTagSetParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteTagSetOK, error)

	DeleteTagSetSpaces(params *DeleteTagSetSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteTagSetSpacesOK, error)

	GetTagSetByID(params *GetTagSetByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetTagSetByIDOK, error)

	GetTagSetByIDSpaces(params *GetTagSetByIDSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetTagSetByIDSpacesOK, error)

	IndexTagSets(params *IndexTagSetsParams, authInfo runtime.ClientAuthInfoWriter) (*IndexTagSetsOK, error)

	IndexTagSetsSpaces(params *IndexTagSetsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*IndexTagSetsSpacesOK, error)

	ListAllTagSets(params *ListAllTagSetsParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllTagSetsOK, error)

	ListAllTagSetsSpaces(params *ListAllTagSetsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllTagSetsSpacesOK, error)

	UpdateSortTagSets(params *UpdateSortTagSetsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSortTagSetsOK, error)

	UpdateSortTagSetsSpaces(params *UpdateSortTagSetsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSortTagSetsSpacesOK, error)

	UpdateTagSet(params *UpdateTagSetParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateTagSetOK, error)

	UpdateTagSetSpaces(params *UpdateTagSetSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateTagSetSpacesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateTagSet creates a tag set resource

  Creates a new tag set.
*/
func (a *Client) CreateTagSet(params *CreateTagSetParams, authInfo runtime.ClientAuthInfoWriter) (*CreateTagSetCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTagSetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createTagSet",
		Method:             "POST",
		PathPattern:        "/api/tagsets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateTagSetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateTagSetCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createTagSet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateTagSetSpaces creates a tag set resource

  Creates a new tag set.
*/
func (a *Client) CreateTagSetSpaces(params *CreateTagSetSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*CreateTagSetSpacesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTagSetSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createTagSet_Spaces",
		Method:             "POST",
		PathPattern:        "/api/{baseSpaceId}/tagsets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateTagSetSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateTagSetSpacesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createTagSet_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteTagSet deletes a tag set resource by ID

  Deletes an existing tag set.
*/
func (a *Client) DeleteTagSet(params *DeleteTagSetParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteTagSetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTagSetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteTagSet",
		Method:             "DELETE",
		PathPattern:        "/api/tagsets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteTagSetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteTagSetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteTagSet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteTagSetSpaces deletes a tag set resource by ID

  Deletes an existing tag set.
*/
func (a *Client) DeleteTagSetSpaces(params *DeleteTagSetSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteTagSetSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTagSetSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteTagSet_Spaces",
		Method:             "DELETE",
		PathPattern:        "/api/{baseSpaceId}/tagsets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteTagSetSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteTagSetSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteTagSet_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTagSetByID gets a tag set resource by ID

  Gets a tag set by ID.
*/
func (a *Client) GetTagSetByID(params *GetTagSetByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetTagSetByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTagSetByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTagSetById",
		Method:             "GET",
		PathPattern:        "/api/tagsets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTagSetByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTagSetByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTagSetById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTagSetByIDSpaces gets a tag set resource by ID

  Gets a tag set by ID.
*/
func (a *Client) GetTagSetByIDSpaces(params *GetTagSetByIDSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetTagSetByIDSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTagSetByIDSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTagSetById_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/tagsets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTagSetByIDSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTagSetByIDSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTagSetById_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IndexTagSets gets a list of tag set resources

  Lists all of the tag sets in the supplied Octopus Deploy Space. The results will be sorted alphabetically by the `SortOrder` field on each tag set.
*/
func (a *Client) IndexTagSets(params *IndexTagSetsParams, authInfo runtime.ClientAuthInfoWriter) (*IndexTagSetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndexTagSetsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indexTagSets",
		Method:             "GET",
		PathPattern:        "/api/tagsets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IndexTagSetsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IndexTagSetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for indexTagSets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IndexTagSetsSpaces gets a list of tag set resources

  Lists all of the tag sets in the supplied Octopus Deploy Space. The results will be sorted alphabetically by the `SortOrder` field on each tag set.
*/
func (a *Client) IndexTagSetsSpaces(params *IndexTagSetsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*IndexTagSetsSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndexTagSetsSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indexTagSets_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/tagsets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IndexTagSetsSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IndexTagSetsSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for indexTagSets_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAllTagSets gets a list of tag set resources

  Lists the details of all of the tag sets in the supplied Octopus Deploy Space. The results will be sorted by the `SortOrder` field on each tag set.
*/
func (a *Client) ListAllTagSets(params *ListAllTagSetsParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllTagSetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAllTagSetsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAllTagSets",
		Method:             "GET",
		PathPattern:        "/api/tagsets/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListAllTagSetsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAllTagSetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAllTagSets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAllTagSetsSpaces gets a list of tag set resources

  Lists the details of all of the tag sets in the supplied Octopus Deploy Space. The results will be sorted by the `SortOrder` field on each tag set.
*/
func (a *Client) ListAllTagSetsSpaces(params *ListAllTagSetsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllTagSetsSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAllTagSetsSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAllTagSets_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/tagsets/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListAllTagSetsSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAllTagSetsSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAllTagSets_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateSortTagSets Takes an array of tag set IDs as the request body, uses the order of items in the array to sort the tag sets on the server. The ID of every tag set must be specified.
*/
func (a *Client) UpdateSortTagSets(params *UpdateSortTagSetsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSortTagSetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSortTagSetsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSortTagSets",
		Method:             "PUT",
		PathPattern:        "/api/tagsets/sortorder",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateSortTagSetsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateSortTagSetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateSortTagSets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateSortTagSetsSpaces Takes an array of tag set IDs as the request body, uses the order of items in the array to sort the tag sets on the server. The ID of every tag set must be specified.
*/
func (a *Client) UpdateSortTagSetsSpaces(params *UpdateSortTagSetsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSortTagSetsSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSortTagSetsSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSortTagSets_Spaces",
		Method:             "PUT",
		PathPattern:        "/api/{baseSpaceId}/tagsets/sortorder",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateSortTagSetsSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateSortTagSetsSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateSortTagSets_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateTagSet modifies a tag set resource by ID

  Modifies an existing tag set.
*/
func (a *Client) UpdateTagSet(params *UpdateTagSetParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateTagSetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTagSetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateTagSet",
		Method:             "PUT",
		PathPattern:        "/api/tagsets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateTagSetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateTagSetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateTagSet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateTagSetSpaces modifies a tag set resource by ID

  Modifies an existing tag set.
*/
func (a *Client) UpdateTagSetSpaces(params *UpdateTagSetSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateTagSetSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTagSetSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateTagSet_Spaces",
		Method:             "PUT",
		PathPattern:        "/api/{baseSpaceId}/tagsets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateTagSetSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateTagSetSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateTagSet_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}