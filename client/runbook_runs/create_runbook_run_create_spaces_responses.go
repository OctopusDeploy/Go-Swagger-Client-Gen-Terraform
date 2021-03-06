// Code generated by go-swagger; DO NOT EDIT.

package runbook_runs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// CreateRunbookRunCreateSpacesReader is a Reader for the CreateRunbookRunCreateSpaces structure.
type CreateRunbookRunCreateSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRunbookRunCreateSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRunbookRunCreateSpacesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRunbookRunCreateSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRunbookRunCreateSpacesCreated creates a CreateRunbookRunCreateSpacesCreated with default headers values
func NewCreateRunbookRunCreateSpacesCreated() *CreateRunbookRunCreateSpacesCreated {
	return &CreateRunbookRunCreateSpacesCreated{}
}

/*CreateRunbookRunCreateSpacesCreated handles this case with default header values.

RunbookRunResource resource returned
*/
type CreateRunbookRunCreateSpacesCreated struct {
	Payload *models.RunbookRunResource
}

func (o *CreateRunbookRunCreateSpacesCreated) Error() string {
	return fmt.Sprintf("[POST /api/{baseSpaceId}/runbookRuns][%d] createRunbookRunCreateSpacesCreated  %+v", 201, o.Payload)
}

func (o *CreateRunbookRunCreateSpacesCreated) GetPayload() *models.RunbookRunResource {
	return o.Payload
}

func (o *CreateRunbookRunCreateSpacesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RunbookRunResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRunbookRunCreateSpacesBadRequest creates a CreateRunbookRunCreateSpacesBadRequest with default headers values
func NewCreateRunbookRunCreateSpacesBadRequest() *CreateRunbookRunCreateSpacesBadRequest {
	return &CreateRunbookRunCreateSpacesBadRequest{}
}

/*CreateRunbookRunCreateSpacesBadRequest handles this case with default header values.

License is not compliant.
No request body was supplied.
RunbookRun creation failed.
RunbookRun Schedule is invalid.
*/
type CreateRunbookRunCreateSpacesBadRequest struct {
}

func (o *CreateRunbookRunCreateSpacesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/{baseSpaceId}/runbookRuns][%d] createRunbookRunCreateSpacesBadRequest ", 400)
}

func (o *CreateRunbookRunCreateSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
