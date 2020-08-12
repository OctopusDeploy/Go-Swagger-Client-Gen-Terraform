// Code generated by go-swagger; DO NOT EDIT.

package runbook_processes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// UpdateRunbookProcessSpacesReader is a Reader for the UpdateRunbookProcessSpaces structure.
type UpdateRunbookProcessSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRunbookProcessSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRunbookProcessSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRunbookProcessSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRunbookProcessSpacesOK creates a UpdateRunbookProcessSpacesOK with default headers values
func NewUpdateRunbookProcessSpacesOK() *UpdateRunbookProcessSpacesOK {
	return &UpdateRunbookProcessSpacesOK{}
}

/*UpdateRunbookProcessSpacesOK handles this case with default header values.

RunbookProcessResource resource returned
*/
type UpdateRunbookProcessSpacesOK struct {
	Payload *models.RunbookProcessResource
}

func (o *UpdateRunbookProcessSpacesOK) Error() string {
	return fmt.Sprintf("[PUT /api/{baseSpaceId}/runbookProcesses/{id}][%d] updateRunbookProcessSpacesOK  %+v", 200, o.Payload)
}

func (o *UpdateRunbookProcessSpacesOK) GetPayload() *models.RunbookProcessResource {
	return o.Payload
}

func (o *UpdateRunbookProcessSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RunbookProcessResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRunbookProcessSpacesBadRequest creates a UpdateRunbookProcessSpacesBadRequest with default headers values
func NewUpdateRunbookProcessSpacesBadRequest() *UpdateRunbookProcessSpacesBadRequest {
	return &UpdateRunbookProcessSpacesBadRequest{}
}

/*UpdateRunbookProcessSpacesBadRequest handles this case with default header values.

Changes to this process could not be saved, because another user has made changes to the process between when you started editing and when you saved your changes. Please reload or open a new tab to make your changes.
No id parameter was provided.
No request body was supplied.
Validation failed.
You cannot modify this process because it is frozen. You'll need to create a new process instead.
*/
type UpdateRunbookProcessSpacesBadRequest struct {
}

func (o *UpdateRunbookProcessSpacesBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/{baseSpaceId}/runbookProcesses/{id}][%d] updateRunbookProcessSpacesBadRequest ", 400)
}

func (o *UpdateRunbookProcessSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
