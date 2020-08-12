// Code generated by go-swagger; DO NOT EDIT.

package machine_policies

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

// NewIndexMachinePolicyDeploymentTargetsSpacesParams creates a new IndexMachinePolicyDeploymentTargetsSpacesParams object
// with the default values initialized.
func NewIndexMachinePolicyDeploymentTargetsSpacesParams() *IndexMachinePolicyDeploymentTargetsSpacesParams {
	var ()
	return &IndexMachinePolicyDeploymentTargetsSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIndexMachinePolicyDeploymentTargetsSpacesParamsWithTimeout creates a new IndexMachinePolicyDeploymentTargetsSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIndexMachinePolicyDeploymentTargetsSpacesParamsWithTimeout(timeout time.Duration) *IndexMachinePolicyDeploymentTargetsSpacesParams {
	var ()
	return &IndexMachinePolicyDeploymentTargetsSpacesParams{

		timeout: timeout,
	}
}

// NewIndexMachinePolicyDeploymentTargetsSpacesParamsWithContext creates a new IndexMachinePolicyDeploymentTargetsSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewIndexMachinePolicyDeploymentTargetsSpacesParamsWithContext(ctx context.Context) *IndexMachinePolicyDeploymentTargetsSpacesParams {
	var ()
	return &IndexMachinePolicyDeploymentTargetsSpacesParams{

		Context: ctx,
	}
}

// NewIndexMachinePolicyDeploymentTargetsSpacesParamsWithHTTPClient creates a new IndexMachinePolicyDeploymentTargetsSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIndexMachinePolicyDeploymentTargetsSpacesParamsWithHTTPClient(client *http.Client) *IndexMachinePolicyDeploymentTargetsSpacesParams {
	var ()
	return &IndexMachinePolicyDeploymentTargetsSpacesParams{
		HTTPClient: client,
	}
}

/*IndexMachinePolicyDeploymentTargetsSpacesParams contains all the parameters to send to the API endpoint
for the index machine policy deployment targets spaces operation typically these are written to a http.Request
*/
type IndexMachinePolicyDeploymentTargetsSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*ID
	  ID of the MachinePolicy

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

// WithTimeout adds the timeout to the index machine policy deployment targets spaces params
func (o *IndexMachinePolicyDeploymentTargetsSpacesParams) WithTimeout(timeout time.Duration) *IndexMachinePolicyDeploymentTargetsSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the index machine policy deployment targets spaces params
func (o *IndexMachinePolicyDeploymentTargetsSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the index machine policy deployment targets spaces params
func (o *IndexMachinePolicyDeploymentTargetsSpacesParams) WithContext(ctx context.Context) *IndexMachinePolicyDeploymentTargetsSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the index machine policy deployment targets spaces params
func (o *IndexMachinePolicyDeploymentTargetsSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the index machine policy deployment targets spaces params
func (o *IndexMachinePolicyDeploymentTargetsSpacesParams) WithHTTPClient(client *http.Client) *IndexMachinePolicyDeploymentTargetsSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the index machine policy deployment targets spaces params
func (o *IndexMachinePolicyDeploymentTargetsSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the index machine policy deployment targets spaces params
func (o *IndexMachinePolicyDeploymentTargetsSpacesParams) WithBaseSpaceID(baseSpaceID string) *IndexMachinePolicyDeploymentTargetsSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the index machine policy deployment targets spaces params
func (o *IndexMachinePolicyDeploymentTargetsSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithID adds the id to the index machine policy deployment targets spaces params
func (o *IndexMachinePolicyDeploymentTargetsSpacesParams) WithID(id string) *IndexMachinePolicyDeploymentTargetsSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the index machine policy deployment targets spaces params
func (o *IndexMachinePolicyDeploymentTargetsSpacesParams) SetID(id string) {
	o.ID = id
}

// WithSkip adds the skip to the index machine policy deployment targets spaces params
func (o *IndexMachinePolicyDeploymentTargetsSpacesParams) WithSkip(skip *int32) *IndexMachinePolicyDeploymentTargetsSpacesParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the index machine policy deployment targets spaces params
func (o *IndexMachinePolicyDeploymentTargetsSpacesParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithTake adds the take to the index machine policy deployment targets spaces params
func (o *IndexMachinePolicyDeploymentTargetsSpacesParams) WithTake(take *int32) *IndexMachinePolicyDeploymentTargetsSpacesParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the index machine policy deployment targets spaces params
func (o *IndexMachinePolicyDeploymentTargetsSpacesParams) SetTake(take *int32) {
	o.Take = take
}

// WriteToRequest writes these params to a swagger request
func (o *IndexMachinePolicyDeploymentTargetsSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
