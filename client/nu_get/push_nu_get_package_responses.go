// Code generated by go-swagger; DO NOT EDIT.

package nu_get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PushNuGetPackageReader is a Reader for the PushNuGetPackage structure.
type PushNuGetPackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PushNuGetPackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPushNuGetPackageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPushNuGetPackageCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPushNuGetPackageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPushNuGetPackageConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPushNuGetPackageOK creates a PushNuGetPackageOK with default headers values
func NewPushNuGetPackageOK() *PushNuGetPackageOK {
	return &PushNuGetPackageOK{}
}

/*PushNuGetPackageOK handles this case with default header values.

OK
*/
type PushNuGetPackageOK struct {
}

func (o *PushNuGetPackageOK) Error() string {
	return fmt.Sprintf("[PUT /nuget/packages][%d] pushNuGetPackageOK ", 200)
}

func (o *PushNuGetPackageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPushNuGetPackageCreated creates a PushNuGetPackageCreated with default headers values
func NewPushNuGetPackageCreated() *PushNuGetPackageCreated {
	return &PushNuGetPackageCreated{}
}

/*PushNuGetPackageCreated handles this case with default header values.

Created
*/
type PushNuGetPackageCreated struct {
}

func (o *PushNuGetPackageCreated) Error() string {
	return fmt.Sprintf("[PUT /nuget/packages][%d] pushNuGetPackageCreated ", 201)
}

func (o *PushNuGetPackageCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPushNuGetPackageBadRequest creates a PushNuGetPackageBadRequest with default headers values
func NewPushNuGetPackageBadRequest() *PushNuGetPackageBadRequest {
	return &PushNuGetPackageBadRequest{}
}

/*PushNuGetPackageBadRequest handles this case with default header values.

A package file must be provided
Package Name is too long.
The uploaded package file had length equal to 0. Please upload a non-empty file.
*/
type PushNuGetPackageBadRequest struct {
}

func (o *PushNuGetPackageBadRequest) Error() string {
	return fmt.Sprintf("[PUT /nuget/packages][%d] pushNuGetPackageBadRequest ", 400)
}

func (o *PushNuGetPackageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPushNuGetPackageConflict creates a PushNuGetPackageConflict with default headers values
func NewPushNuGetPackageConflict() *PushNuGetPackageConflict {
	return &PushNuGetPackageConflict{}
}

/*PushNuGetPackageConflict handles this case with default header values.

A package with the same ID and version already exists. To proceed anyway, specify an overwriteMode of OverwriteExisting or IgnoreIfExists.
*/
type PushNuGetPackageConflict struct {
}

func (o *PushNuGetPackageConflict) Error() string {
	return fmt.Sprintf("[PUT /nuget/packages][%d] pushNuGetPackageConflict ", 409)
}

func (o *PushNuGetPackageConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
