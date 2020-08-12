// Code generated by go-swagger; DO NOT EDIT.

package subscription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new subscription API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for subscription API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateSubscription(params *CreateSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*CreateSubscriptionCreated, error)

	CreateSubscriptionSpaces(params *CreateSubscriptionSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*CreateSubscriptionSpacesCreated, error)

	DeleteSubscription(params *DeleteSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSubscriptionOK, error)

	DeleteSubscriptionSpaces(params *DeleteSubscriptionSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSubscriptionSpacesOK, error)

	GetSubscriptionByID(params *GetSubscriptionByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetSubscriptionByIDOK, error)

	GetSubscriptionByIDSpaces(params *GetSubscriptionByIDSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetSubscriptionByIDSpacesOK, error)

	IndexSubscriptions(params *IndexSubscriptionsParams, authInfo runtime.ClientAuthInfoWriter) (*IndexSubscriptionsOK, error)

	IndexSubscriptionsSpaces(params *IndexSubscriptionsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*IndexSubscriptionsSpacesOK, error)

	ListAllSubscriptions(params *ListAllSubscriptionsParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllSubscriptionsOK, error)

	ListAllSubscriptionsSpaces(params *ListAllSubscriptionsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllSubscriptionsSpacesOK, error)

	UpdateSubscription(params *UpdateSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSubscriptionOK, error)

	UpdateSubscriptionSpaces(params *UpdateSubscriptionSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSubscriptionSpacesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateSubscription creates a subscription resource

  Creates a new subscription
*/
func (a *Client) CreateSubscription(params *CreateSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*CreateSubscriptionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSubscriptionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSubscription",
		Method:             "POST",
		PathPattern:        "/api/subscriptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateSubscriptionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateSubscriptionCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createSubscription: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateSubscriptionSpaces creates a subscription resource

  Creates a new subscription
*/
func (a *Client) CreateSubscriptionSpaces(params *CreateSubscriptionSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*CreateSubscriptionSpacesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSubscriptionSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSubscription_Spaces",
		Method:             "POST",
		PathPattern:        "/api/{baseSpaceId}/subscriptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateSubscriptionSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateSubscriptionSpacesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createSubscription_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteSubscription deletes a subscription resource by ID

  Deletes an existing subscription.
*/
func (a *Client) DeleteSubscription(params *DeleteSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSubscriptionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSubscription",
		Method:             "DELETE",
		PathPattern:        "/api/subscriptions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSubscriptionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteSubscriptionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteSubscription: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteSubscriptionSpaces deletes a subscription resource by ID

  Deletes an existing subscription.
*/
func (a *Client) DeleteSubscriptionSpaces(params *DeleteSubscriptionSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSubscriptionSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSubscriptionSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSubscription_Spaces",
		Method:             "DELETE",
		PathPattern:        "/api/{baseSpaceId}/subscriptions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSubscriptionSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteSubscriptionSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteSubscription_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSubscriptionByID gets a subscription resource by ID

  Get a subscription
*/
func (a *Client) GetSubscriptionByID(params *GetSubscriptionByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetSubscriptionByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSubscriptionByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSubscriptionById",
		Method:             "GET",
		PathPattern:        "/api/subscriptions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSubscriptionByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSubscriptionByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSubscriptionById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSubscriptionByIDSpaces gets a subscription resource by ID

  Get a subscription
*/
func (a *Client) GetSubscriptionByIDSpaces(params *GetSubscriptionByIDSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetSubscriptionByIDSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSubscriptionByIDSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSubscriptionById_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/subscriptions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSubscriptionByIDSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSubscriptionByIDSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSubscriptionById_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IndexSubscriptions gets a list of subscription resources

  Lists all of the subscriptions in the supplied Octopus Deploy Space. The results will be sorted alphabetically by name.
*/
func (a *Client) IndexSubscriptions(params *IndexSubscriptionsParams, authInfo runtime.ClientAuthInfoWriter) (*IndexSubscriptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndexSubscriptionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indexSubscriptions",
		Method:             "GET",
		PathPattern:        "/api/subscriptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IndexSubscriptionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IndexSubscriptionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for indexSubscriptions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IndexSubscriptionsSpaces gets a list of subscription resources

  Lists all of the subscriptions in the supplied Octopus Deploy Space. The results will be sorted alphabetically by name.
*/
func (a *Client) IndexSubscriptionsSpaces(params *IndexSubscriptionsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*IndexSubscriptionsSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndexSubscriptionsSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indexSubscriptions_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/subscriptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IndexSubscriptionsSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IndexSubscriptionsSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for indexSubscriptions_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAllSubscriptions gets a list of subscription resources

  Lists all the subscriptions in the supplied Octopus Deploy Space.
*/
func (a *Client) ListAllSubscriptions(params *ListAllSubscriptionsParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllSubscriptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAllSubscriptionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAllSubscriptions",
		Method:             "GET",
		PathPattern:        "/api/subscriptions/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListAllSubscriptionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAllSubscriptionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAllSubscriptions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAllSubscriptionsSpaces gets a list of subscription resources

  Lists all the subscriptions in the supplied Octopus Deploy Space.
*/
func (a *Client) ListAllSubscriptionsSpaces(params *ListAllSubscriptionsSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllSubscriptionsSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAllSubscriptionsSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAllSubscriptions_Spaces",
		Method:             "GET",
		PathPattern:        "/api/{baseSpaceId}/subscriptions/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListAllSubscriptionsSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAllSubscriptionsSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAllSubscriptions_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateSubscription modifies a subscription resource by ID

  Updates an existing subscription
*/
func (a *Client) UpdateSubscription(params *UpdateSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSubscriptionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSubscription",
		Method:             "PUT",
		PathPattern:        "/api/subscriptions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateSubscriptionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateSubscriptionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateSubscription: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateSubscriptionSpaces modifies a subscription resource by ID

  Updates an existing subscription
*/
func (a *Client) UpdateSubscriptionSpaces(params *UpdateSubscriptionSpacesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSubscriptionSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSubscriptionSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSubscription_Spaces",
		Method:             "PUT",
		PathPattern:        "/api/{baseSpaceId}/subscriptions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateSubscriptionSpacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateSubscriptionSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateSubscription_Spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}