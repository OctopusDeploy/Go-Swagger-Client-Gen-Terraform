// Code generated by go-swagger; DO NOT EDIT.

package machines

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
)

// NewIndexDeploymentTargetTasksParams creates a new IndexDeploymentTargetTasksParams object
// with the default values initialized.
func NewIndexDeploymentTargetTasksParams() *IndexDeploymentTargetTasksParams {
	var ()
	return &IndexDeploymentTargetTasksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIndexDeploymentTargetTasksParamsWithTimeout creates a new IndexDeploymentTargetTasksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIndexDeploymentTargetTasksParamsWithTimeout(timeout time.Duration) *IndexDeploymentTargetTasksParams {
	var ()
	return &IndexDeploymentTargetTasksParams{

		timeout: timeout,
	}
}

// NewIndexDeploymentTargetTasksParamsWithContext creates a new IndexDeploymentTargetTasksParams object
// with the default values initialized, and the ability to set a context for a request
func NewIndexDeploymentTargetTasksParamsWithContext(ctx context.Context) *IndexDeploymentTargetTasksParams {
	var ()
	return &IndexDeploymentTargetTasksParams{

		Context: ctx,
	}
}

// NewIndexDeploymentTargetTasksParamsWithHTTPClient creates a new IndexDeploymentTargetTasksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIndexDeploymentTargetTasksParamsWithHTTPClient(client *http.Client) *IndexDeploymentTargetTasksParams {
	var ()
	return &IndexDeploymentTargetTasksParams{
		HTTPClient: client,
	}
}

/*IndexDeploymentTargetTasksParams contains all the parameters to send to the API endpoint
for the index deployment target tasks operation typically these are written to a http.Request
*/
type IndexDeploymentTargetTasksParams struct {

	/*ID
	  ID of the DeploymentTarget

	*/
	ID string
	/*Skip
	  Number of items to skip

	*/
	Skip *int32
	/*Take
	  Number of items per page

	*/
	Take *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the index deployment target tasks params
func (o *IndexDeploymentTargetTasksParams) WithTimeout(timeout time.Duration) *IndexDeploymentTargetTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the index deployment target tasks params
func (o *IndexDeploymentTargetTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the index deployment target tasks params
func (o *IndexDeploymentTargetTasksParams) WithContext(ctx context.Context) *IndexDeploymentTargetTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the index deployment target tasks params
func (o *IndexDeploymentTargetTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the index deployment target tasks params
func (o *IndexDeploymentTargetTasksParams) WithHTTPClient(client *http.Client) *IndexDeploymentTargetTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the index deployment target tasks params
func (o *IndexDeploymentTargetTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the index deployment target tasks params
func (o *IndexDeploymentTargetTasksParams) WithID(id string) *IndexDeploymentTargetTasksParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the index deployment target tasks params
func (o *IndexDeploymentTargetTasksParams) SetID(id string) {
	o.ID = id
}

// WithSkip adds the skip to the index deployment target tasks params
func (o *IndexDeploymentTargetTasksParams) WithSkip(skip *int32) *IndexDeploymentTargetTasksParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the index deployment target tasks params
func (o *IndexDeploymentTargetTasksParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithTake adds the take to the index deployment target tasks params
func (o *IndexDeploymentTargetTasksParams) WithTake(take *int32) *IndexDeploymentTargetTasksParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the index deployment target tasks params
func (o *IndexDeploymentTargetTasksParams) SetTake(take *int32) {
	o.Take = take
}

// WriteToRequest writes these params to a swagger request
func (o *IndexDeploymentTargetTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Skip != nil {

		// query param skip
		var qrSkip int32
		if o.Skip != nil {
			qrSkip = *o.Skip
		}
		qSkip := swag.FormatInt32(qrSkip)
		if qSkip != "" {
			if err := r.SetQueryParam("skip", qSkip); err != nil {
				return err
			}
		}

	}

	if o.Take != nil {

		// query param take
		var qrTake int32
		if o.Take != nil {
			qrTake = *o.Take
		}
		qTake := swag.FormatInt32(qrTake)
		if qTake != "" {
			if err := r.SetQueryParam("take", qTake); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
