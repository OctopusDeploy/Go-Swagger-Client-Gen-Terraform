// Code generated by go-swagger; DO NOT EDIT.

package migration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// CreateMigrationImportReader is a Reader for the CreateMigrationImport structure.
type CreateMigrationImportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMigrationImportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateMigrationImportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateMigrationImportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateMigrationImportOK creates a CreateMigrationImportOK with default headers values
func NewCreateMigrationImportOK() *CreateMigrationImportOK {
	return &CreateMigrationImportOK{}
}

/*CreateMigrationImportOK handles this case with default header values.

MigrationImportResource resource returned
*/
type CreateMigrationImportOK struct {
	Payload *models.MigrationImportResource
}

func (o *CreateMigrationImportOK) Error() string {
	return fmt.Sprintf("[POST /api/migrations/import][%d] createMigrationImportOK  %+v", 200, o.Payload)
}

func (o *CreateMigrationImportOK) GetPayload() *models.MigrationImportResource {
	return o.Payload
}

func (o *CreateMigrationImportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MigrationImportResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMigrationImportBadRequest creates a CreateMigrationImportBadRequest with default headers values
func NewCreateMigrationImportBadRequest() *CreateMigrationImportBadRequest {
	return &CreateMigrationImportBadRequest{}
}

/*CreateMigrationImportBadRequest handles this case with default header values.

Missing expected 'PackageId' parameter.
Missing expected 'PackageVersion' parameter.
Missing expected 'Password' parameter.
No request body was supplied.
*/
type CreateMigrationImportBadRequest struct {
}

func (o *CreateMigrationImportBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/migrations/import][%d] createMigrationImportBadRequest ", 400)
}

func (o *CreateMigrationImportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
