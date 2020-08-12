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

// CreateReleaseSnapshotVariablesSpacesReader is a Reader for the CreateReleaseSnapshotVariablesSpaces structure.
type CreateReleaseSnapshotVariablesSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateReleaseSnapshotVariablesSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateReleaseSnapshotVariablesSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateReleaseSnapshotVariablesSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateReleaseSnapshotVariablesSpacesOK creates a CreateReleaseSnapshotVariablesSpacesOK with default headers values
func NewCreateReleaseSnapshotVariablesSpacesOK() *CreateReleaseSnapshotVariablesSpacesOK {
	return &CreateReleaseSnapshotVariablesSpacesOK{}
}

/*CreateReleaseSnapshotVariablesSpacesOK handles this case with default header values.

ReleaseResource resource returned
*/
type CreateReleaseSnapshotVariablesSpacesOK struct {
	Payload *models.ReleaseResource
}

func (o *CreateReleaseSnapshotVariablesSpacesOK) Error() string {
	return fmt.Sprintf("[POST /api/{baseSpaceId}/releases/{id}/snapshot-variables][%d] createReleaseSnapshotVariablesSpacesOK  %+v", 200, o.Payload)
}

func (o *CreateReleaseSnapshotVariablesSpacesOK) GetPayload() *models.ReleaseResource {
	return o.Payload
}

func (o *CreateReleaseSnapshotVariablesSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReleaseResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReleaseSnapshotVariablesSpacesBadRequest creates a CreateReleaseSnapshotVariablesSpacesBadRequest with default headers values
func NewCreateReleaseSnapshotVariablesSpacesBadRequest() *CreateReleaseSnapshotVariablesSpacesBadRequest {
	return &CreateReleaseSnapshotVariablesSpacesBadRequest{}
}

/*CreateReleaseSnapshotVariablesSpacesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type CreateReleaseSnapshotVariablesSpacesBadRequest struct {
}

func (o *CreateReleaseSnapshotVariablesSpacesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/{baseSpaceId}/releases/{id}/snapshot-variables][%d] createReleaseSnapshotVariablesSpacesBadRequest ", 400)
}

func (o *CreateReleaseSnapshotVariablesSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}