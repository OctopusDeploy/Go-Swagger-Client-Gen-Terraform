// Code generated by go-swagger; DO NOT EDIT.

package tag_sets

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

// NewDeleteTagSetParams creates a new DeleteTagSetParams object
// with the default values initialized.
func NewDeleteTagSetParams() *DeleteTagSetParams {
	var ()
	return &DeleteTagSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTagSetParamsWithTimeout creates a new DeleteTagSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteTagSetParamsWithTimeout(timeout time.Duration) *DeleteTagSetParams {
	var ()
	return &DeleteTagSetParams{

		timeout: timeout,
	}
}

// NewDeleteTagSetParamsWithContext creates a new DeleteTagSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteTagSetParamsWithContext(ctx context.Context) *DeleteTagSetParams {
	var ()
	return &DeleteTagSetParams{

		Context: ctx,
	}
}

// NewDeleteTagSetParamsWithHTTPClient creates a new DeleteTagSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteTagSetParamsWithHTTPClient(client *http.Client) *DeleteTagSetParams {
	var ()
	return &DeleteTagSetParams{
		HTTPClient: client,
	}
}

/*DeleteTagSetParams contains all the parameters to send to the API endpoint
for the delete tag set operation typically these are written to a http.Request
*/
type DeleteTagSetParams struct {

	/*ID
	  ID of the TagSetResource to delete

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete tag set params
func (o *DeleteTagSetParams) WithTimeout(timeout time.Duration) *DeleteTagSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete tag set params
func (o *DeleteTagSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete tag set params
func (o *DeleteTagSetParams) WithContext(ctx context.Context) *DeleteTagSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete tag set params
func (o *DeleteTagSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete tag set params
func (o *DeleteTagSetParams) WithHTTPClient(client *http.Client) *DeleteTagSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete tag set params
func (o *DeleteTagSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete tag set params
func (o *DeleteTagSetParams) WithID(id string) *DeleteTagSetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete tag set params
func (o *DeleteTagSetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTagSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
