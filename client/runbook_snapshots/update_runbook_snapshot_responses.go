// Code generated by go-swagger; DO NOT EDIT.

package runbook_snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// UpdateRunbookSnapshotReader is a Reader for the UpdateRunbookSnapshot structure.
type UpdateRunbookSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRunbookSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRunbookSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRunbookSnapshotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRunbookSnapshotOK creates a UpdateRunbookSnapshotOK with default headers values
func NewUpdateRunbookSnapshotOK() *UpdateRunbookSnapshotOK {
	return &UpdateRunbookSnapshotOK{}
}

/*UpdateRunbookSnapshotOK handles this case with default header values.

RunbookSnapshotResource Modified
*/
type UpdateRunbookSnapshotOK struct {
	Payload *models.RunbookSnapshotResource
}

func (o *UpdateRunbookSnapshotOK) Error() string {
	return fmt.Sprintf("[PUT /api/runbookSnapshots/{id}][%d] updateRunbookSnapshotOK  %+v", 200, o.Payload)
}

func (o *UpdateRunbookSnapshotOK) GetPayload() *models.RunbookSnapshotResource {
	return o.Payload
}

func (o *UpdateRunbookSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RunbookSnapshotResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRunbookSnapshotBadRequest creates a UpdateRunbookSnapshotBadRequest with default headers values
func NewUpdateRunbookSnapshotBadRequest() *UpdateRunbookSnapshotBadRequest {
	return &UpdateRunbookSnapshotBadRequest{}
}

/*UpdateRunbookSnapshotBadRequest handles this case with default header values.

Model validation failed.
No id parameter was provided.
No request body was supplied.
This resource is read-only and cannot be modified.
*/
type UpdateRunbookSnapshotBadRequest struct {
}

func (o *UpdateRunbookSnapshotBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/runbookSnapshots/{id}][%d] updateRunbookSnapshotBadRequest ", 400)
}

func (o *UpdateRunbookSnapshotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
