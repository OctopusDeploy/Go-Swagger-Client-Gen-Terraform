// Code generated by go-swagger; DO NOT EDIT.

package releases

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
	"github.com/go-openapi/swag"

	"models"
)

// NewCreateReleaseParams creates a new CreateReleaseParams object
// with the default values initialized.
func NewCreateReleaseParams() *CreateReleaseParams {
	var ()
	return &CreateReleaseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateReleaseParamsWithTimeout creates a new CreateReleaseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateReleaseParamsWithTimeout(timeout time.Duration) *CreateReleaseParams {
	var ()
	return &CreateReleaseParams{

		timeout: timeout,
	}
}

// NewCreateReleaseParamsWithContext creates a new CreateReleaseParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateReleaseParamsWithContext(ctx context.Context) *CreateReleaseParams {
	var ()
	return &CreateReleaseParams{

		Context: ctx,
	}
}

// NewCreateReleaseParamsWithHTTPClient creates a new CreateReleaseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateReleaseParamsWithHTTPClient(client *http.Client) *CreateReleaseParams {
	var ()
	return &CreateReleaseParams{
		HTTPClient: client,
	}
}

/*CreateReleaseParams contains all the parameters to send to the API endpoint
for the create release operation typically these are written to a http.Request
*/
type CreateReleaseParams struct {

	/*ReleaseResource
	  The ReleaseResource resource to create

	*/
	ReleaseResource *models.ReleaseResource
	/*IgnoreChannelRules
	  Ignore channel rules.

	*/
	IgnoreChannelRules *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create release params
func (o *CreateReleaseParams) WithTimeout(timeout time.Duration) *CreateReleaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create release params
func (o *CreateReleaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create release params
func (o *CreateReleaseParams) WithContext(ctx context.Context) *CreateReleaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create release params
func (o *CreateReleaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create release params
func (o *CreateReleaseParams) WithHTTPClient(client *http.Client) *CreateReleaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create release params
func (o *CreateReleaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReleaseResource adds the releaseResource to the create release params
func (o *CreateReleaseParams) WithReleaseResource(releaseResource *models.ReleaseResource) *CreateReleaseParams {
	o.SetReleaseResource(releaseResource)
	return o
}

// SetReleaseResource adds the releaseResource to the create release params
func (o *CreateReleaseParams) SetReleaseResource(releaseResource *models.ReleaseResource) {
	o.ReleaseResource = releaseResource
}

// WithIgnoreChannelRules adds the ignoreChannelRules to the create release params
func (o *CreateReleaseParams) WithIgnoreChannelRules(ignoreChannelRules *bool) *CreateReleaseParams {
	o.SetIgnoreChannelRules(ignoreChannelRules)
	return o
}

// SetIgnoreChannelRules adds the ignoreChannelRules to the create release params
func (o *CreateReleaseParams) SetIgnoreChannelRules(ignoreChannelRules *bool) {
	o.IgnoreChannelRules = ignoreChannelRules
}

// WriteToRequest writes these params to a swagger request
func (o *CreateReleaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ReleaseResource != nil {
		if err := r.SetBodyParam(o.ReleaseResource); err != nil {
			return err
		}
	}

	if o.IgnoreChannelRules != nil {

		// query param ignoreChannelRules
		var qrIgnoreChannelRules bool
		if o.IgnoreChannelRules != nil {
			qrIgnoreChannelRules = *o.IgnoreChannelRules
		}
		qIgnoreChannelRules := swag.FormatBool(qrIgnoreChannelRules)
		if qIgnoreChannelRules != "" {
			if err := r.SetQueryParam("ignoreChannelRules", qIgnoreChannelRules); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
