// Code generated by go-swagger; DO NOT EDIT.

package releases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// CreateReleaseSnapshotVariablesReader is a Reader for the CreateReleaseSnapshotVariables structure.
type CreateReleaseSnapshotVariablesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateReleaseSnapshotVariablesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateReleaseSnapshotVariablesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateReleaseSnapshotVariablesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateReleaseSnapshotVariablesOK creates a CreateReleaseSnapshotVariablesOK with default headers values
func NewCreateReleaseSnapshotVariablesOK() *CreateReleaseSnapshotVariablesOK {
	return &CreateReleaseSnapshotVariablesOK{}
}

/*CreateReleaseSnapshotVariablesOK handles this case with default header values.

ReleaseResource resource returned
*/
type CreateReleaseSnapshotVariablesOK struct {
	Payload *models.ReleaseResource
}

func (o *CreateReleaseSnapshotVariablesOK) Error() string {
	return fmt.Sprintf("[POST /api/releases/{id}/snapshot-variables][%d] createReleaseSnapshotVariablesOK  %+v", 200, o.Payload)
}

func (o *CreateReleaseSnapshotVariablesOK) GetPayload() *models.ReleaseResource {
	return o.Payload
}

func (o *CreateReleaseSnapshotVariablesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReleaseResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReleaseSnapshotVariablesBadRequest creates a CreateReleaseSnapshotVariablesBadRequest with default headers values
func NewCreateReleaseSnapshotVariablesBadRequest() *CreateReleaseSnapshotVariablesBadRequest {
	return &CreateReleaseSnapshotVariablesBadRequest{}
}

/*CreateReleaseSnapshotVariablesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type CreateReleaseSnapshotVariablesBadRequest struct {
}

func (o *CreateReleaseSnapshotVariablesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/releases/{id}/snapshot-variables][%d] createReleaseSnapshotVariablesBadRequest ", 400)
}

func (o *CreateReleaseSnapshotVariablesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
