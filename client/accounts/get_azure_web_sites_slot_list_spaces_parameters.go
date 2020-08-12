// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetAzureWebSitesSlotListSpacesParams creates a new GetAzureWebSitesSlotListSpacesParams object
// with the default values initialized.
func NewGetAzureWebSitesSlotListSpacesParams() *GetAzureWebSitesSlotListSpacesParams {
	var ()
	return &GetAzureWebSitesSlotListSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAzureWebSitesSlotListSpacesParamsWithTimeout creates a new GetAzureWebSitesSlotListSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAzureWebSitesSlotListSpacesParamsWithTimeout(timeout time.Duration) *GetAzureWebSitesSlotListSpacesParams {
	var ()
	return &GetAzureWebSitesSlotListSpacesParams{

		timeout: timeout,
	}
}

// NewGetAzureWebSitesSlotListSpacesParamsWithContext creates a new GetAzureWebSitesSlotListSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAzureWebSitesSlotListSpacesParamsWithContext(ctx context.Context) *GetAzureWebSitesSlotListSpacesParams {
	var ()
	return &GetAzureWebSitesSlotListSpacesParams{

		Context: ctx,
	}
}

// NewGetAzureWebSitesSlotListSpacesParamsWithHTTPClient creates a new GetAzureWebSitesSlotListSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAzureWebSitesSlotListSpacesParamsWithHTTPClient(client *http.Client) *GetAzureWebSitesSlotListSpacesParams {
	var ()
	return &GetAzureWebSitesSlotListSpacesParams{
		HTTPClient: client,
	}
}

/*GetAzureWebSitesSlotListSpacesParams contains all the parameters to send to the API endpoint
for the get azure web sites slot list spaces operation typically these are written to a http.Request
*/
type GetAzureWebSitesSlotListSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  Account Id for Azure

	*/
	ID string
	/*ResourceGroupName
	  Resource Group name

	*/
	ResourceGroupName string
	/*WebSiteName
	  Web site name

	*/
	WebSiteName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get azure web sites slot list spaces params
func (o *GetAzureWebSitesSlotListSpacesParams) WithTimeout(timeout time.Duration) *GetAzureWebSitesSlotListSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get azure web sites slot list spaces params
func (o *GetAzureWebSitesSlotListSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get azure web sites slot list spaces params
func (o *GetAzureWebSitesSlotListSpacesParams) WithContext(ctx context.Context) *GetAzureWebSitesSlotListSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get azure web sites slot list spaces params
func (o *GetAzureWebSitesSlotListSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get azure web sites slot list spaces params
func (o *GetAzureWebSitesSlotListSpacesParams) WithHTTPClient(client *http.Client) *GetAzureWebSitesSlotListSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get azure web sites slot list spaces params
func (o *GetAzureWebSitesSlotListSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get azure web sites slot list spaces params
func (o *GetAzureWebSitesSlotListSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetAzureWebSitesSlotListSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get azure web sites slot list spaces params
func (o *GetAzureWebSitesSlotListSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the get azure web sites slot list spaces params
func (o *GetAzureWebSitesSlotListSpacesParams) WithID(id string) *GetAzureWebSitesSlotListSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get azure web sites slot list spaces params
func (o *GetAzureWebSitesSlotListSpacesParams) SetID(id string) {
	o.ID = id
}

// WithResourceGroupName adds the resourceGroupName to the get azure web sites slot list spaces params
func (o *GetAzureWebSitesSlotListSpacesParams) WithResourceGroupName(resourceGroupName string) *GetAzureWebSitesSlotListSpacesParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the get azure web sites slot list spaces params
func (o *GetAzureWebSitesSlotListSpacesParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithWebSiteName adds the webSiteName to the get azure web sites slot list spaces params
func (o *GetAzureWebSitesSlotListSpacesParams) WithWebSiteName(webSiteName string) *GetAzureWebSitesSlotListSpacesParams {
	o.SetWebSiteName(webSiteName)
	return o
}

// SetWebSiteName adds the webSiteName to the get azure web sites slot list spaces params
func (o *GetAzureWebSitesSlotListSpacesParams) SetWebSiteName(webSiteName string) {
	o.WebSiteName = webSiteName
}

// WriteToRequest writes these params to a swagger request
func (o *GetAzureWebSitesSlotListSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSpaceId
	if err := r.SetPathParam("baseSpaceId", o.BaseSpaceID); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param resourceGroupName
	if err := r.SetPathParam("resourceGroupName", o.ResourceGroupName); err != nil {
		return err
	}

	// path param webSiteName
	if err := r.SetPathParam("webSiteName", o.WebSiteName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}