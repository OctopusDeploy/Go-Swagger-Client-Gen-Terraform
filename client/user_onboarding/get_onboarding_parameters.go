// Code generated by go-swagger; DO NOT EDIT.

package user_onboarding

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

// NewGetOnboardingParams creates a new GetOnboardingParams object
// with the default values initialized.
func NewGetOnboardingParams() *GetOnboardingParams {

	return &GetOnboardingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOnboardingParamsWithTimeout creates a new GetOnboardingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOnboardingParamsWithTimeout(timeout time.Duration) *GetOnboardingParams {

	return &GetOnboardingParams{

		timeout: timeout,
	}
}

// NewGetOnboardingParamsWithContext creates a new GetOnboardingParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOnboardingParamsWithContext(ctx context.Context) *GetOnboardingParams {

	return &GetOnboardingParams{

		Context: ctx,
	}
}

// NewGetOnboardingParamsWithHTTPClient creates a new GetOnboardingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOnboardingParamsWithHTTPClient(client *http.Client) *GetOnboardingParams {

	return &GetOnboardingParams{
		HTTPClient: client,
	}
}

/*GetOnboardingParams contains all the parameters to send to the API endpoint
for the get onboarding operation typically these are written to a http.Request
*/
type GetOnboardingParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get onboarding params
func (o *GetOnboardingParams) WithTimeout(timeout time.Duration) *GetOnboardingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get onboarding params
func (o *GetOnboardingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get onboarding params
func (o *GetOnboardingParams) WithContext(ctx context.Context) *GetOnboardingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get onboarding params
func (o *GetOnboardingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get onboarding params
func (o *GetOnboardingParams) WithHTTPClient(client *http.Client) *GetOnboardingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get onboarding params
func (o *GetOnboardingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetOnboardingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
