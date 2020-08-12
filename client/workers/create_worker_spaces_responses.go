// Code generated by go-swagger; DO NOT EDIT.

package workers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// CreateWorkerSpacesReader is a Reader for the CreateWorkerSpaces structure.
type CreateWorkerSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateWorkerSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateWorkerSpacesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateWorkerSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateWorkerSpacesCreated creates a CreateWorkerSpacesCreated with default headers values
func NewCreateWorkerSpacesCreated() *CreateWorkerSpacesCreated {
	return &CreateWorkerSpacesCreated{}
}

/*CreateWorkerSpacesCreated handles this case with default header values.

WorkerResource Created
*/
type CreateWorkerSpacesCreated struct {
	Payload *models.WorkerResource
}

func (o *CreateWorkerSpacesCreated) Error() string {
	return fmt.Sprintf("[POST /api/{baseSpaceId}/workers][%d] createWorkerSpacesCreated  %+v", 201, o.Payload)
}

func (o *CreateWorkerSpacesCreated) GetPayload() *models.WorkerResource {
	return o.Payload
}

func (o *CreateWorkerSpacesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkerResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateWorkerSpacesBadRequest creates a CreateWorkerSpacesBadRequest with default headers values
func NewCreateWorkerSpacesBadRequest() *CreateWorkerSpacesBadRequest {
	return &CreateWorkerSpacesBadRequest{}
}

/*CreateWorkerSpacesBadRequest handles this case with default header values.

Model validation failed.
No request body was supplied.
*/
type CreateWorkerSpacesBadRequest struct {
}

func (o *CreateWorkerSpacesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/{baseSpaceId}/workers][%d] createWorkerSpacesBadRequest ", 400)
}

func (o *CreateWorkerSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}