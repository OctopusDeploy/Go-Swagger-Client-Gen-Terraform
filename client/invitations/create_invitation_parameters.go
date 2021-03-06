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

// NewCreateInvitationParams creates a new CreateInvitationParams object
// with the default values initialized.
func NewCreateInvitationParams() *CreateInvitationParams {
	var ()
	return &CreateInvitationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateInvitationParamsWithTimeout creates a new CreateInvitationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateInvitationParamsWithTimeout(timeout time.Duration) *CreateInvitationParams {
	var ()
	return &CreateInvitationParams{

		timeout: timeout,
	}
}

// NewCreateInvitationParamsWithContext creates a new CreateInvitationParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateInvitationParamsWithContext(ctx context.Context) *CreateInvitationParams {
	var ()
	return &CreateInvitationParams{

		Context: ctx,
	}
}

// NewCreateInvitationParamsWithHTTPClient creates a new CreateInvitationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateInvitationParamsWithHTTPClient(client *http.Client) *CreateInvitationParams {
	var ()
	return &CreateInvitationParams{
		HTTPClient: client,
	}
}

/*CreateInvitationParams contains all the parameters to send to the API endpoint
for the create invitation operation typically these are written to a http.Request
*/
type CreateInvitationParams struct {

	/*InvitationResource
	  The InvitationResource resource to create

	*/
	InvitationResource *models.InvitationResource

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create invitation params
func (o *CreateInvitationParams) WithTimeout(timeout time.Duration) *CreateInvitationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create invitation params
func (o *CreateInvitationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create invitation params
func (o *CreateInvitationParams) WithContext(ctx context.Context) *CreateInvitationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create invitation params
func (o *CreateInvitationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create invitation params
func (o *CreateInvitationParams) WithHTTPClient(client *http.Client) *CreateInvitationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create invitation params
func (o *CreateInvitationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInvitationResource adds the invitationResource to the create invitation params
func (o *CreateInvitationParams) WithInvitationResource(invitationResource *models.InvitationResource) *CreateInvitationParams {
	o.SetInvitationResource(invitationResource)
	return o
}

// SetInvitationResource adds the invitationResource to the create invitation params
func (o *CreateInvitationParams) SetInvitationResource(invitationResource *models.InvitationResource) {
	o.InvitationResource = invitationResource
}

// WriteToRequest writes these params to a swagger request
func (o *CreateInvitationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.InvitationResource != nil {
		if err := r.SetBodyParam(o.InvitationResource); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
