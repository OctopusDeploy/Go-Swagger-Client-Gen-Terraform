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

// UpdateRunbookSnapshotSpacesReader is a Reader for the UpdateRunbookSnapshotSpaces structure.
type UpdateRunbookSnapshotSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRunbookSnapshotSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRunbookSnapshotSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRunbookSnapshotSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRunbookSnapshotSpacesOK creates a UpdateRunbookSnapshotSpacesOK with default headers values
func NewUpdateRunbookSnapshotSpacesOK() *UpdateRunbookSnapshotSpacesOK {
	return &UpdateRunbookSnapshotSpacesOK{}
}

/*UpdateRunbookSnapshotSpacesOK handles this case with default header values.

RunbookSnapshotResource Modified
*/
type UpdateRunbookSnapshotSpacesOK struct {
	Payload *models.RunbookSnapshotResource
}

func (o *UpdateRunbookSnapshotSpacesOK) Error() string {
	return fmt.Sprintf("[PUT /api/{baseSpaceId}/runbookSnapshots/{id}][%d] updateRunbookSnapshotSpacesOK  %+v", 200, o.Payload)
}

func (o *UpdateRunbookSnapshotSpacesOK) GetPayload() *models.RunbookSnapshotResource {
	return o.Payload
}

func (o *UpdateRunbookSnapshotSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RunbookSnapshotResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRunbookSnapshotSpacesBadRequest creates a UpdateRunbookSnapshotSpacesBadRequest with default headers values
func NewUpdateRunbookSnapshotSpacesBadRequest() *UpdateRunbookSnapshotSpacesBadRequest {
	return &UpdateRunbookSnapshotSpacesBadRequest{}
}

/*UpdateRunbookSnapshotSpacesBadRequest handles this case with default header values.

Model validation failed.
No id parameter was provided.
No request body was supplied.
This resource is read-only and cannot be modified.
*/
type UpdateRunbookSnapshotSpacesBadRequest struct {
}

func (o *UpdateRunbookSnapshotSpacesBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/{baseSpaceId}/runbookSnapshots/{id}][%d] updateRunbookSnapshotSpacesBadRequest ", 400)
}

func (o *UpdateRunbookSnapshotSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
