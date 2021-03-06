// Code generated by go-swagger; DO NOT EDIT.

package library_variable_sets

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

	"models"
)

// NewCreateLibraryVariableSetSpacesParams creates a new CreateLibraryVariableSetSpacesParams object
// with the default values initialized.
func NewCreateLibraryVariableSetSpacesParams() *CreateLibraryVariableSetSpacesParams {
	var ()
	return &CreateLibraryVariableSetSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateLibraryVariableSetSpacesParamsWithTimeout creates a new CreateLibraryVariableSetSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateLibraryVariableSetSpacesParamsWithTimeout(timeout time.Duration) *CreateLibraryVariableSetSpacesParams {
	var ()
	return &CreateLibraryVariableSetSpacesParams{

		timeout: timeout,
	}
}

// NewCreateLibraryVariableSetSpacesParamsWithContext creates a new CreateLibraryVariableSetSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateLibraryVariableSetSpacesParamsWithContext(ctx context.Context) *CreateLibraryVariableSetSpacesParams {
	var ()
	return &CreateLibraryVariableSetSpacesParams{

		Context: ctx,
	}
}

// NewCreateLibraryVariableSetSpacesParamsWithHTTPClient creates a new CreateLibraryVariableSetSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateLibraryVariableSetSpacesParamsWithHTTPClient(client *http.Client) *CreateLibraryVariableSetSpacesParams {
	var ()
	return &CreateLibraryVariableSetSpacesParams{
		HTTPClient: client,
	}
}

/*CreateLibraryVariableSetSpacesParams contains all the parameters to send to the API endpoint
for the create library variable set spaces operation typically these are written to a http.Request
*/
type CreateLibraryVariableSetSpacesParams struct {

	/*LibraryVariableSetResource
	  The LibraryVariableSetResource resource to create

	*/
	LibraryVariableSetResource *models.LibraryVariableSetResource
	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create library variable set spaces params
func (o *CreateLibraryVariableSetSpacesParams) WithTimeout(timeout time.Duration) *CreateLibraryVariableSetSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create library variable set spaces params
func (o *CreateLibraryVariableSetSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create library variable set spaces params
func (o *CreateLibraryVariableSetSpacesParams) WithContext(ctx context.Context) *CreateLibraryVariableSetSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create library variable set spaces params
func (o *CreateLibraryVariableSetSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create library variable set spaces params
func (o *CreateLibraryVariableSetSpacesParams) WithHTTPClient(client *http.Client) *CreateLibraryVariableSetSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create library variable set spaces params
func (o *CreateLibraryVariableSetSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLibraryVariableSetResource adds the libraryVariableSetResource to the create library variable set spaces params
func (o *CreateLibraryVariableSetSpacesParams) WithLibraryVariableSetResource(libraryVariableSetResource *models.LibraryVariableSetResource) *CreateLibraryVariableSetSpacesParams {
	o.SetLibraryVariableSetResource(libraryVariableSetResource)
	return o
}

// SetLibraryVariableSetResource adds the libraryVariableSetResource to the create library variable set spaces params
func (o *CreateLibraryVariableSetSpacesParams) SetLibraryVariableSetResource(libraryVariableSetResource *models.LibraryVariableSetResource) {
	o.LibraryVariableSetResource = libraryVariableSetResource
}

// WithBaseSpaceID adds the baseSpaceID to the create library variable set spaces params
func (o *CreateLibraryVariableSetSpacesParams) WithBaseSpaceID(baseSpaceID string) *CreateLibraryVariableSetSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the create library variable set spaces params
func (o *CreateLibraryVariableSetSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateLibraryVariableSetSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.LibraryVariableSetResource != nil {
		if err := r.SetBodyParam(o.LibraryVariableSetResource); err != nil {
			return err
		}
	}

	// path param baseSpaceId
	if err := r.SetPathParam("baseSpaceId", o.BaseSpaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
