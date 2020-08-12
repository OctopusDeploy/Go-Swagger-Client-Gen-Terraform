// Code generated by go-swagger; DO NOT EDIT.

package invitations

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

// NewCreateInvitationSpacesParams creates a new CreateInvitationSpacesParams object
// with the default values initialized.
func NewCreateInvitationSpacesParams() *CreateInvitationSpacesParams {
	var ()
	return &CreateInvitationSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateInvitationSpacesParamsWithTimeout creates a new CreateInvitationSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateInvitationSpacesParamsWithTimeout(timeout time.Duration) *CreateInvitationSpacesParams {
	var ()
	return &CreateInvitationSpacesParams{

		timeout: timeout,
	}
}

// NewCreateInvitationSpacesParamsWithContext creates a new CreateInvitationSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateInvitationSpacesParamsWithContext(ctx context.Context) *CreateInvitationSpacesParams {
	var ()
	return &CreateInvitationSpacesParams{

		Context: ctx,
	}
}

// NewCreateInvitationSpacesParamsWithHTTPClient creates a new CreateInvitationSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateInvitationSpacesParamsWithHTTPClient(client *http.Client) *CreateInvitationSpacesParams {
	var ()
	return &CreateInvitationSpacesParams{
		HTTPClient: client,
	}
}

/*CreateInvitationSpacesParams contains all the parameters to send to the API endpoint
for the create invitation spaces operation typically these are written to a http.Request
*/
type CreateInvitationSpacesParams struct {

	/*InvitationResource
	  The InvitationResource resource to create

	*/
	InvitationResource *models.InvitationResource
	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create invitation spaces params
func (o *CreateInvitationSpacesParams) WithTimeout(timeout time.Duration) *CreateInvitationSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create invitation spaces params
func (o *CreateInvitationSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create invitation spaces params
func (o *CreateInvitationSpacesParams) WithContext(ctx context.Context) *CreateInvitationSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create invitation spaces params
func (o *CreateInvitationSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create invitation spaces params
func (o *CreateInvitationSpacesParams) WithHTTPClient(client *http.Client) *CreateInvitationSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create invitation spaces params
func (o *CreateInvitationSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInvitationResource adds the invitationResource to the create invitation spaces params
func (o *CreateInvitationSpacesParams) WithInvitationResource(invitationResource *models.InvitationResource) *CreateInvitationSpacesParams {
	o.SetInvitationResource(invitationResource)
	return o
}

// SetInvitationResource adds the invitationResource to the create invitation spaces params
func (o *CreateInvitationSpacesParams) SetInvitationResource(invitationResource *models.InvitationResource) {
	o.InvitationResource = invitationResource
}

// WithBaseSpaceID adds the baseSpaceID to the create invitation spaces params
func (o *CreateInvitationSpacesParams) WithBaseSpaceID(baseSpaceID string) *CreateInvitationSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the create invitation spaces params
func (o *CreateInvitationSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateInvitationSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.InvitationResource != nil {
		if err := r.SetBodyParam(o.InvitationResource); err != nil {
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
