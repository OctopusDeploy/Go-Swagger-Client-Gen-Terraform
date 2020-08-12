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

// CreateMigrationPartialExportReader is a Reader for the CreateMigrationPartialExport structure.
type CreateMigrationPartialExportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMigrationPartialExportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateMigrationPartialExportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateMigrationPartialExportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateMigrationPartialExportOK creates a CreateMigrationPartialExportOK with default headers values
func NewCreateMigrationPartialExportOK() *CreateMigrationPartialExportOK {
	return &CreateMigrationPartialExportOK{}
}

/*CreateMigrationPartialExportOK handles this case with default header values.

MigrationPartialExportResource resource returned
*/
type CreateMigrationPartialExportOK struct {
	Payload *models.MigrationPartialExportResource
}

func (o *CreateMigrationPartialExportOK) Error() string {
	return fmt.Sprintf("[POST /api/migrations/partialexport][%d] createMigrationPartialExportOK  %+v", 200, o.Payload)
}

func (o *CreateMigrationPartialExportOK) GetPayload() *models.MigrationPartialExportResource {
	return o.Payload
}

func (o *CreateMigrationPartialExportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MigrationPartialExportResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMigrationPartialExportBadRequest creates a CreateMigrationPartialExportBadRequest with default headers values
func NewCreateMigrationPartialExportBadRequest() *CreateMigrationPartialExportBadRequest {
	return &CreateMigrationPartialExportBadRequest{}
}

/*CreateMigrationPartialExportBadRequest handles this case with default header values.

Missing expected 'Password' parameter.
Missing expected 'Projects' parameter.
No request body was supplied.
*/
type CreateMigrationPartialExportBadRequest struct {
}

func (o *CreateMigrationPartialExportBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/migrations/partialexport][%d] createMigrationPartialExportBadRequest ", 400)
}

func (o *CreateMigrationPartialExportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
