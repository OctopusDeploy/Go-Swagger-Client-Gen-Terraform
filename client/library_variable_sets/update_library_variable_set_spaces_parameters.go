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

// NewUpdateLibraryVariableSetSpacesParams creates a new UpdateLibraryVariableSetSpacesParams object
// with the default values initialized.
func NewUpdateLibraryVariableSetSpacesParams() *UpdateLibraryVariableSetSpacesParams {
	var ()
	return &UpdateLibraryVariableSetSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateLibraryVariableSetSpacesParamsWithTimeout creates a new UpdateLibraryVariableSetSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateLibraryVariableSetSpacesParamsWithTimeout(timeout time.Duration) *UpdateLibraryVariableSetSpacesParams {
	var ()
	return &UpdateLibraryVariableSetSpacesParams{

		timeout: timeout,
	}
}

// NewUpdateLibraryVariableSetSpacesParamsWithContext creates a new UpdateLibraryVariableSetSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateLibraryVariableSetSpacesParamsWithContext(ctx context.Context) *UpdateLibraryVariableSetSpacesParams {
	var ()
	return &UpdateLibraryVariableSetSpacesParams{

		Context: ctx,
	}
}

// NewUpdateLibraryVariableSetSpacesParamsWithHTTPClient creates a new UpdateLibraryVariableSetSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateLibraryVariableSetSpacesParamsWithHTTPClient(client *http.Client) *UpdateLibraryVariableSetSpacesParams {
	var ()
	return &UpdateLibraryVariableSetSpacesParams{
		HTTPClient: client,
	}
}

/*UpdateLibraryVariableSetSpacesParams contains all the parameters to send to the API endpoint
for the update library variable set spaces operation typically these are written to a http.Request
*/
type UpdateLibraryVariableSetSpacesParams struct {

	/*LibraryVariableSetResource
	  The LibraryVariableSetResource resource to create

	*/
	LibraryVariableSetResource *models.LibraryVariableSetResource
	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the LibraryVariableSetResource to modify

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update library variable set spaces params
func (o *UpdateLibraryVariableSetSpacesParams) WithTimeout(timeout time.Duration) *UpdateLibraryVariableSetSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update library variable set spaces params
func (o *UpdateLibraryVariableSetSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update library variable set spaces params
func (o *UpdateLibraryVariableSetSpacesParams) WithContext(ctx context.Context) *UpdateLibraryVariableSetSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update library variable set spaces params
func (o *UpdateLibraryVariableSetSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update library variable set spaces params
func (o *UpdateLibraryVariableSetSpacesParams) WithHTTPClient(client *http.Client) *UpdateLibraryVariableSetSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update library variable set spaces params
func (o *UpdateLibraryVariableSetSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLibraryVariableSetResource adds the libraryVariableSetResource to the update library variable set spaces params
func (o *UpdateLibraryVariableSetSpacesParams) WithLibraryVariableSetResource(libraryVariableSetResource *models.LibraryVariableSetResource) *UpdateLibraryVariableSetSpacesParams {
	o.SetLibraryVariableSetResource(libraryVariableSetResource)
	return o
}

// SetLibraryVariableSetResource adds the libraryVariableSetResource to the update library variable set spaces params
func (o *UpdateLibraryVariableSetSpacesParams) SetLibraryVariableSetResource(libraryVariableSetResource *models.LibraryVariableSetResource) {
	o.LibraryVariableSetResource = libraryVariableSetResource
}

// WithBaseSpaceID adds the baseSpaceID to the update library variable set spaces params
func (o *UpdateLibraryVariableSetSpacesParams) WithBaseSpaceID(baseSpaceID string) *UpdateLibraryVariableSetSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the update library variable set spaces params
func (o *UpdateLibraryVariableSetSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the update library variable set spaces params
func (o *UpdateLibraryVariableSetSpacesParams) WithID(id string) *UpdateLibraryVariableSetSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update library variable set spaces params
func (o *UpdateLibraryVariableSetSpacesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateLibraryVariableSetSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
