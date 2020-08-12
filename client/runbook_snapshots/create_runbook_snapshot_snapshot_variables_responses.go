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

// CreateRunbookSnapshotSnapshotVariablesReader is a Reader for the CreateRunbookSnapshotSnapshotVariables structure.
type CreateRunbookSnapshotSnapshotVariablesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRunbookSnapshotSnapshotVariablesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRunbookSnapshotSnapshotVariablesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRunbookSnapshotSnapshotVariablesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRunbookSnapshotSnapshotVariablesOK creates a CreateRunbookSnapshotSnapshotVariablesOK with default headers values
func NewCreateRunbookSnapshotSnapshotVariablesOK() *CreateRunbookSnapshotSnapshotVariablesOK {
	return &CreateRunbookSnapshotSnapshotVariablesOK{}
}

/*CreateRunbookSnapshotSnapshotVariablesOK handles this case with default header values.

RunbookSnapshotResource resource returned
*/
type CreateRunbookSnapshotSnapshotVariablesOK struct {
	Payload *models.RunbookSnapshotResource
}

func (o *CreateRunbookSnapshotSnapshotVariablesOK) Error() string {
	return fmt.Sprintf("[POST /api/runbookSnapshots/{id}/snapshot-variables][%d] createRunbookSnapshotSnapshotVariablesOK  %+v", 200, o.Payload)
}

func (o *CreateRunbookSnapshotSnapshotVariablesOK) GetPayload() *models.RunbookSnapshotResource {
	return o.Payload
}

func (o *CreateRunbookSnapshotSnapshotVariablesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RunbookSnapshotResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRunbookSnapshotSnapshotVariablesBadRequest creates a CreateRunbookSnapshotSnapshotVariablesBadRequest with default headers values
func NewCreateRunbookSnapshotSnapshotVariablesBadRequest() *CreateRunbookSnapshotSnapshotVariablesBadRequest {
	return &CreateRunbookSnapshotSnapshotVariablesBadRequest{}
}

/*CreateRunbookSnapshotSnapshotVariablesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type CreateRunbookSnapshotSnapshotVariablesBadRequest struct {
}

func (o *CreateRunbookSnapshotSnapshotVariablesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/runbookSnapshots/{id}/snapshot-variables][%d] createRunbookSnapshotSnapshotVariablesBadRequest ", 400)
}

func (o *CreateRunbookSnapshotSnapshotVariablesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
