// Code generated by go-swagger; DO NOT EDIT.

package packages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// CreatePackageRepositoryUploadReader is a Reader for the CreatePackageRepositoryUpload structure.
type CreatePackageRepositoryUploadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePackageRepositoryUploadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreatePackageRepositoryUploadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreatePackageRepositoryUploadCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreatePackageRepositoryUploadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreatePackageRepositoryUploadConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreatePackageRepositoryUploadOK creates a CreatePackageRepositoryUploadOK with default headers values
func NewCreatePackageRepositoryUploadOK() *CreatePackageRepositoryUploadOK {
	return &CreatePackageRepositoryUploadOK{}
}

/*CreatePackageRepositoryUploadOK handles this case with default header values.

OK
*/
type CreatePackageRepositoryUploadOK struct {
}

func (o *CreatePackageRepositoryUploadOK) Error() string {
	return fmt.Sprintf("[POST /api/packages/raw][%d] createPackageRepositoryUploadOK ", 200)
}

func (o *CreatePackageRepositoryUploadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreatePackageRepositoryUploadCreated creates a CreatePackageRepositoryUploadCreated with default headers values
func NewCreatePackageRepositoryUploadCreated() *CreatePackageRepositoryUploadCreated {
	return &CreatePackageRepositoryUploadCreated{}
}

/*CreatePackageRepositoryUploadCreated handles this case with default header values.

PackageFromBuiltInFeedResource resource returned
*/
type CreatePackageRepositoryUploadCreated struct {
	Payload *models.PackageFromBuiltInFeedResource
}

func (o *CreatePackageRepositoryUploadCreated) Error() string {
	return fmt.Sprintf("[POST /api/packages/raw][%d] createPackageRepositoryUploadCreated  %+v", 201, o.Payload)
}

func (o *CreatePackageRepositoryUploadCreated) GetPayload() *models.PackageFromBuiltInFeedResource {
	return o.Payload
}

func (o *CreatePackageRepositoryUploadCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PackageFromBuiltInFeedResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePackageRepositoryUploadBadRequest creates a CreatePackageRepositoryUploadBadRequest with default headers values
func NewCreatePackageRepositoryUploadBadRequest() *CreatePackageRepositoryUploadBadRequest {
	return &CreatePackageRepositoryUploadBadRequest{}
}

/*CreatePackageRepositoryUploadBadRequest handles this case with default header values.

A package file must be provided
Package Name is too long.
The uploaded package file had length equal to 0. Please upload a non-empty file.
*/
type CreatePackageRepositoryUploadBadRequest struct {
}

func (o *CreatePackageRepositoryUploadBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/packages/raw][%d] createPackageRepositoryUploadBadRequest ", 400)
}

func (o *CreatePackageRepositoryUploadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreatePackageRepositoryUploadConflict creates a CreatePackageRepositoryUploadConflict with default headers values
func NewCreatePackageRepositoryUploadConflict() *CreatePackageRepositoryUploadConflict {
	return &CreatePackageRepositoryUploadConflict{}
}

/*CreatePackageRepositoryUploadConflict handles this case with default header values.

A package with the same ID and version already exists. To proceed anyway, specify an overwriteMode of OverwriteExisting or IgnoreIfExists.
*/
type CreatePackageRepositoryUploadConflict struct {
}

func (o *CreatePackageRepositoryUploadConflict) Error() string {
	return fmt.Sprintf("[POST /api/packages/raw][%d] createPackageRepositoryUploadConflict ", 409)
}

func (o *CreatePackageRepositoryUploadConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
