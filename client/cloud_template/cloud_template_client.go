// Code generated by go-swagger; DO NOT EDIT.

package cloud_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new cloud template API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cloud template API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateCloudTemplateMetadata(params *CreateCloudTemplateMetadataParams, authInfo runtime.ClientAuthInfoWriter) (*CreateCloudTemplateMetadataOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateCloudTemplateMetadata Provides parameter metadata for a cloud (AWS etc) resource template
*/
func (a *Client) CreateCloudTemplateMetadata(params *CreateCloudTemplateMetadataParams, authInfo runtime.ClientAuthInfoWriter) (*CreateCloudTemplateMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCloudTemplateMetadataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createCloudTemplateMetadata",
		Method:             "POST",
		PathPattern:        "/api/cloudtemplate/{id}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateCloudTemplateMetadataReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateCloudTemplateMetadataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createCloudTemplateMetadata: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
