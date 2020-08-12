// Code generated by go-swagger; DO NOT EDIT.

package packages

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

// NewGetPackageRepositoryDeltaSignatureSpacesParams creates a new GetPackageRepositoryDeltaSignatureSpacesParams object
// with the default values initialized.
func NewGetPackageRepositoryDeltaSignatureSpacesParams() *GetPackageRepositoryDeltaSignatureSpacesParams {
	var ()
	return &GetPackageRepositoryDeltaSignatureSpacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPackageRepositoryDeltaSignatureSpacesParamsWithTimeout creates a new GetPackageRepositoryDeltaSignatureSpacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPackageRepositoryDeltaSignatureSpacesParamsWithTimeout(timeout time.Duration) *GetPackageRepositoryDeltaSignatureSpacesParams {
	var ()
	return &GetPackageRepositoryDeltaSignatureSpacesParams{

		timeout: timeout,
	}
}

// NewGetPackageRepositoryDeltaSignatureSpacesParamsWithContext creates a new GetPackageRepositoryDeltaSignatureSpacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPackageRepositoryDeltaSignatureSpacesParamsWithContext(ctx context.Context) *GetPackageRepositoryDeltaSignatureSpacesParams {
	var ()
	return &GetPackageRepositoryDeltaSignatureSpacesParams{

		Context: ctx,
	}
}

// NewGetPackageRepositoryDeltaSignatureSpacesParamsWithHTTPClient creates a new GetPackageRepositoryDeltaSignatureSpacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPackageRepositoryDeltaSignatureSpacesParamsWithHTTPClient(client *http.Client) *GetPackageRepositoryDeltaSignatureSpacesParams {
	var ()
	return &GetPackageRepositoryDeltaSignatureSpacesParams{
		HTTPClient: client,
	}
}

/*GetPackageRepositoryDeltaSignatureSpacesParams contains all the parameters to send to the API endpoint
for the get package repository delta signature spaces operation typically these are written to a http.Request
*/
type GetPackageRepositoryDeltaSignatureSpacesParams struct {

	/*BaseSpaceID
	  ID of the space

	*/
	BaseSpaceID string
	/*PackageID
	  Package ID of the package to be uploaded

	*/
	PackageID string
	/*Version
	  The version of the package to be uploaded

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get package repository delta signature spaces params
func (o *GetPackageRepositoryDeltaSignatureSpacesParams) WithTimeout(timeout time.Duration) *GetPackageRepositoryDeltaSignatureSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get package repository delta signature spaces params
func (o *GetPackageRepositoryDeltaSignatureSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get package repository delta signature spaces params
func (o *GetPackageRepositoryDeltaSignatureSpacesParams) WithContext(ctx context.Context) *GetPackageRepositoryDeltaSignatureSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get package repository delta signature spaces params
func (o *GetPackageRepositoryDeltaSignatureSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get package repository delta signature spaces params
func (o *GetPackageRepositoryDeltaSignatureSpacesParams) WithHTTPClient(client *http.Client) *GetPackageRepositoryDeltaSignatureSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get package repository delta signature spaces params
func (o *GetPackageRepositoryDeltaSignatureSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSpaceID adds the baseSpaceID to the get package repository delta signature spaces params
func (o *GetPackageRepositoryDeltaSignatureSpacesParams) WithBaseSpaceID(baseSpaceID string) *GetPackageRepositoryDeltaSignatureSpacesParams {
	o.SetBaseSpaceID(baseSpaceID)
	return o
}

// SetBaseSpaceID adds the baseSpaceId to the get package repository delta signature spaces params
func (o *GetPackageRepositoryDeltaSignatureSpacesParams) SetBaseSpaceID(baseSpaceID string) {
	o.BaseSpaceID = baseSpaceID
}

// WithPackageID adds the packageID to the get package repository delta signature spaces params
func (o *GetPackageRepositoryDeltaSignatureSpacesParams) WithPackageID(packageID string) *GetPackageRepositoryDeltaSignatureSpacesParams {
	o.SetPackageID(packageID)
	return o
}

// SetPackageID adds the packageId to the get package repository delta signature spaces params
func (o *GetPackageRepositoryDeltaSignatureSpacesParams) SetPackageID(packageID string) {
	o.PackageID = packageID
}

// WithVersion adds the version to the get package repository delta signature spaces params
func (o *GetPackageRepositoryDeltaSignatureSpacesParams) WithVersion(version string) *GetPackageRepositoryDeltaSignatureSpacesParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get package repository delta signature spaces params
func (o *GetPackageRepositoryDeltaSignatureSpacesParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetPackageRepositoryDeltaSignatureSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSpaceId
	if err := r.SetPathParam("baseSpaceId", o.BaseSpaceID); err != nil {
		return err
	}

	// path param packageId
	if err := r.SetPathParam("packageId", o.PackageID); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}