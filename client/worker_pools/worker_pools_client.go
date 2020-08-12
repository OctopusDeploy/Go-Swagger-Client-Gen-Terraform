// Code generated by go-swagger; DO NOT EDIT.

package worker_pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new worker pools API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for worker pools API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateWorkerPool(params *CreateWorkerPoolParams, authInfo runtime.ClientAuthInfoWriter) (*CreateWorkerPoolCreated, error)

	CreateWorkerPoolSpaces(params *CreateWorkerPoolSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*CreateWorkerPoolSpacesCreated, error)

	DeleteWorkerPool(params *DeleteWorkerPoolParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteWorkerPoolOK, error)

	DeleteWorkerPoolSpaces(params *DeleteWorkerPoolSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteWorkerPoolSpacesOK, error)

	GetWorkerPoolByID(params *GetWorkerPoolByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkerPoolByIDOK, error)

	GetWorkerPoolByIDSpaces(params *GetWorkerPoolByIDSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkerPoolByIDSpacesOK, error)

	GetWorkerPoolDynamicWorkerTypes(params *GetWorkerPoolDynamicWorkerTypesParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkerPoolDynamicWorkerTypesOK, error)

	GetWorkerPoolDynamicWorkerTypesSpaces(params *GetWorkerPoolDynamicWorkerTypesSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkerPoolDynamicWorkerTypesSpacesOK, error)

	GetWorkerPoolSupportedTypes(params *GetWorkerPoolSupportedTypesParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkerPoolSupportedTypesOK, error)

	GetWorkerPoolSupportedTypesSpaces(params *GetWorkerPoolSupportedTypesSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkerPoolSupportedTypesSpacesOK, error)

	GetWorkerPoolsSummary(params *GetWorkerPoolsSummaryParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkerPoolsSummaryOK, error)

	GetWorkerPoolsSummarySpaces(params *GetWorkerPoolsSummarySpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkerPoolsSummarySpacesOK, error)

	IndexWorkerPoolWorkers(params *IndexWorkerPoolWorkersParams, authInfo runtime.ClientAuthInfoWriter) (*IndexWorkerPoolWorkersOK, error)

	IndexWorkerPoolWorkersSpaces(params *IndexWorkerPoolWorkersSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*IndexWorkerPoolWorkersSpacesOK, error)

	IndexWorkerPools(params *IndexWorkerPoolsParams, authInfo runtime.ClientAuthInfoWriter) (*IndexWorkerPoolsOK, error)

	IndexWorkerPoolsSpaces(params *IndexWorkerPoolsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*IndexWorkerPoolsSpacesOK, error)

	ListAllWorkerPools(params *ListAllWorkerPoolsParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllWorkerPoolsOK, error)

	ListAllWorkerPoolsSpaces(params *ListAllWorkerPoolsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllWorkerPoolsSpacesOK, error)

	UpdateSortWorkerPools(params *UpdateSortWorkerPoolsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSortWorkerPoolsOK, error)

	UpdateSortWorkerPoolsSpaces(params *UpdateSortWorkerPoolsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSortWorkerPoolsSpacesOK, error)

	UpdateWorkerPool(params *UpdateWorkerPoolParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateWorkerPoolOK, error)

	UpdateWorkerPoolSpaces(params *UpdateWorkerPoolSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateWorkerPoolSpacesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateWorkerPool creates a worker pool resource

  Creates a new worker pool.
*/
func (a *Client) CreateWorkerPool(params *CreateWorkerPoolParams, authInfo runtime.ClientAuthInfoWriter) (*CreateWorkerPoolCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateWorkerPoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createWorkerPool",
		Method:             "POST",
		PathPattern:        "/api/workerpools",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateWorkerPoolReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateWorkerPoolCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createWorkerPool: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateWorkerPoolSpaces creates a worker pool resource

  Creates a new worker pool.
*/
func (a *Client) CreateWorkerPoolSpaces(params *CreateWorkerPoolSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*CreateWorkerPoolSpacesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateWorkerPoolSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createWorkerPool_Spaces",
		Method:             "POST",
		PathPattern:        "/api/{baseSpaceId}/workerpools",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateWorkerPoolSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateWorkerPoolSpacesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createWorkerPool_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteWorkerPool deletes a worker pool resource by ID

  Deletes an existing pool.
*/
func (a *Client) DeleteWorkerPool(params *DeleteWorkerPoolParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteWorkerPoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteWorkerPoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteWorkerPool",
		Method:             "DELETE",
		PathPattern:        "/api/workerpools/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteWorkerPoolReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteWorkerPoolOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteWorkerPool: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteWorkerPoolSpaces deletes a worker pool resource by ID

  Deletes an existing pool.
*/
func (a *Client) DeleteWorkerPoolSpaces(params *DeleteWorkerPoolSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteWorkerPoolSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteWorkerPoolSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteWorkerPool_Spaces",
		Method:             "DELETE",
		PathPattern:        "/api/{baseSpaceId}/workerpools/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteWorkerPoolSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteWorkerPoolSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteWorkerPool_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetWorkerPoolByID gets a worker pool resource by ID

  Gets a single worker pool by ID.
*/
func (a *Client) GetWorkerPoolByID(params *GetWorkerPoolByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkerPoolByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkerPoolByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWorkerPoolById",
		Method:             "GET",
		PathPattern:        "/api/workerpools/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWorkerPoolByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetWorkerPoolByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getWorkerPoolById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetWorkerPoolByIDSpaces gets a worker pool resource by ID

  Gets a single worker pool by ID.
*/
func (a *Client) GetWorkerPoolByIDSpaces(params *GetWorkerPoolByIDSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkerPoolByIDSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkerPoolByIDSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWorkerPoolById_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/workerpools/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWorkerPoolByIDSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetWorkerPoolByIDSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getWorkerPoolById_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetWorkerPoolDynamicWorkerTypes Lists the available Worker Types for the Dynamic Worker Pool
*/
func (a *Client) GetWorkerPoolDynamicWorkerTypes(params *GetWorkerPoolDynamicWorkerTypesParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkerPoolDynamicWorkerTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkerPoolDynamicWorkerTypesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWorkerPoolDynamicWorkerTypes",
		Method:             "GET",
		PathPattern:        "/api/workerpools/dynamicworkertypes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWorkerPoolDynamicWorkerTypesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetWorkerPoolDynamicWorkerTypesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getWorkerPoolDynamicWorkerTypes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetWorkerPoolDynamicWorkerTypesSpaces Lists the available Worker Types for the Dynamic Worker Pool
*/
func (a *Client) GetWorkerPoolDynamicWorkerTypesSpaces(params *GetWorkerPoolDynamicWorkerTypesSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkerPoolDynamicWorkerTypesSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkerPoolDynamicWorkerTypesSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWorkerPoolDynamicWorkerTypes_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/workerpools/dynamicworkertypes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWorkerPoolDynamicWorkerTypesSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetWorkerPoolDynamicWorkerTypesSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getWorkerPoolDynamicWorkerTypes_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetWorkerPoolSupportedTypes Lists the available Worker Pool types.
*/
func (a *Client) GetWorkerPoolSupportedTypes(params *GetWorkerPoolSupportedTypesParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkerPoolSupportedTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkerPoolSupportedTypesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWorkerPoolSupportedTypes",
		Method:             "GET",
		PathPattern:        "/api/workerpools/supportedtypes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWorkerPoolSupportedTypesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetWorkerPoolSupportedTypesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getWorkerPoolSupportedTypes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetWorkerPoolSupportedTypesSpaces Lists the available Worker Pool types.
*/
func (a *Client) GetWorkerPoolSupportedTypesSpaces(params *GetWorkerPoolSupportedTypesSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkerPoolSupportedTypesSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkerPoolSupportedTypesSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWorkerPoolSupportedTypes_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/workerpools/supportedtypes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWorkerPoolSupportedTypesSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetWorkerPoolSupportedTypesSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getWorkerPoolSupportedTypes_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetWorkerPoolsSummary Lists all worker pools, including a summary of machine information
*/
func (a *Client) GetWorkerPoolsSummary(params *GetWorkerPoolsSummaryParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkerPoolsSummaryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkerPoolsSummaryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWorkerPoolsSummary",
		Method:             "GET",
		PathPattern:        "/api/workerpools/summary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWorkerPoolsSummaryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetWorkerPoolsSummaryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getWorkerPoolsSummary: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetWorkerPoolsSummarySpaces Lists all worker pools, including a summary of machine information
*/
func (a *Client) GetWorkerPoolsSummarySpaces(params *GetWorkerPoolsSummarySpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkerPoolsSummarySpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkerPoolsSummarySpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWorkerPoolsSummary_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/workerpools/summary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWorkerPoolsSummarySpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetWorkerPoolsSummarySpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getWorkerPoolsSummary_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IndexWorkerPoolWorkers gets a list of worker resources for the given worker pool resource

  Lists all of the machines that belong to the given worker pool.
*/
func (a *Client) IndexWorkerPoolWorkers(params *IndexWorkerPoolWorkersParams, authInfo runtime.ClientAuthInfoWriter) (*IndexWorkerPoolWorkersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndexWorkerPoolWorkersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indexWorkerPoolWorkers",
		Method:             "GET",
		PathPattern:        "/api/workerpools/{id}/workers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IndexWorkerPoolWorkersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IndexWorkerPoolWorkersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for indexWorkerPoolWorkers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IndexWorkerPoolWorkersSpaces gets a list of worker resources for the given worker pool resource

  Lists all of the machines that belong to the given worker pool.
*/
func (a *Client) IndexWorkerPoolWorkersSpaces(params *IndexWorkerPoolWorkersSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*IndexWorkerPoolWorkersSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndexWorkerPoolWorkersSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indexWorkerPoolWorkers_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/workerpools/{id}/workers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IndexWorkerPoolWorkersSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IndexWorkerPoolWorkersSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for indexWorkerPoolWorkers_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IndexWorkerPools gets a list of worker pool resources

  Lists all of the worker pools in the supplied Octopus Deploy Space. The results will be sorted by the `SortOrder` field on each environment.
*/
func (a *Client) IndexWorkerPools(params *IndexWorkerPoolsParams, authInfo runtime.ClientAuthInfoWriter) (*IndexWorkerPoolsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndexWorkerPoolsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indexWorkerPools",
		Method:             "GET",
		PathPattern:        "/api/workerpools",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IndexWorkerPoolsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IndexWorkerPoolsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for indexWorkerPools: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IndexWorkerPoolsSpaces gets a list of worker pool resources

  Lists all of the worker pools in the supplied Octopus Deploy Space. The results will be sorted by the `SortOrder` field on each environment.
*/
func (a *Client) IndexWorkerPoolsSpaces(params *IndexWorkerPoolsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*IndexWorkerPoolsSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndexWorkerPoolsSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indexWorkerPools_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/workerpools",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IndexWorkerPoolsSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IndexWorkerPoolsSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for indexWorkerPools_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAllWorkerPools gets a list of worker pool resources

  Lists the name and ID of all of the worker pools in the supplied Octopus Deploy Space. The results will be sorted by the `SortOrder` field on each pool.
*/
func (a *Client) ListAllWorkerPools(params *ListAllWorkerPoolsParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllWorkerPoolsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAllWorkerPoolsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAllWorkerPools",
		Method:             "GET",
		PathPattern:        "/api/workerpools/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListAllWorkerPoolsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAllWorkerPoolsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAllWorkerPools: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAllWorkerPoolsSpaces gets a list of worker pool resources

  Lists the name and ID of all of the worker pools in the supplied Octopus Deploy Space. The results will be sorted by the `SortOrder` field on each pool.
*/
func (a *Client) ListAllWorkerPoolsSpaces(params *ListAllWorkerPoolsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllWorkerPoolsSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAllWorkerPoolsSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAllWorkerPools_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/workerpools/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListAllWorkerPoolsSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAllWorkerPoolsSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAllWorkerPools_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateSortWorkerPools Takes an array of work pool IDs as the request body, uses the order of items in the array to sort the worker pools on the server. The ID of every worker pool must be specified.
*/
func (a *Client) UpdateSortWorkerPools(params *UpdateSortWorkerPoolsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSortWorkerPoolsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSortWorkerPoolsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSortWorkerPools",
		Method:             "PUT",
		PathPattern:        "/api/workerpools/sortorder",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateSortWorkerPoolsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateSortWorkerPoolsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateSortWorkerPools: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateSortWorkerPoolsSpaces Takes an array of work pool IDs as the request body, uses the order of items in the array to sort the worker pools on the server. The ID of every worker pool must be specified.
*/
func (a *Client) UpdateSortWorkerPoolsSpaces(params *UpdateSortWorkerPoolsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSortWorkerPoolsSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSortWorkerPoolsSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSortWorkerPools_Spaces",
		Method:             "PUT",
		PathPattern:        "/api/{baseSpaceId}/workerpools/sortorder",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateSortWorkerPoolsSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateSortWorkerPoolsSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateSortWorkerPools_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateWorkerPool modifies a worker pool resource by ID

  Modifies an existing worker pool.
*/
func (a *Client) UpdateWorkerPool(params *UpdateWorkerPoolParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateWorkerPoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateWorkerPoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateWorkerPool",
		Method:             "PUT",
		PathPattern:        "/api/workerpools/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateWorkerPoolReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateWorkerPoolOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateWorkerPool: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateWorkerPoolSpaces modifies a worker pool resource by ID

  Modifies an existing worker pool.
*/
func (a *Client) UpdateWorkerPoolSpaces(params *UpdateWorkerPoolSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateWorkerPoolSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateWorkerPoolSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateWorkerPool_Spaces",
		Method:             "PUT",
		PathPattern:        "/api/{baseSpaceId}/workerpools/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateWorkerPoolSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateWorkerPoolSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateWorkerPool_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
