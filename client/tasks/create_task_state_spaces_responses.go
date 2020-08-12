// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// CreateTaskStateSpacesReader is a Reader for the CreateTaskStateSpaces structure.
type CreateTaskStateSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTaskStateSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTaskStateSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTaskStateSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTaskStateSpacesOK creates a CreateTaskStateSpacesOK with default headers values
func NewCreateTaskStateSpacesOK() *CreateTaskStateSpacesOK {
	return &CreateTaskStateSpacesOK{}
}

/*CreateTaskStateSpacesOK handles this case with default header values.

TaskResource resource returned
*/
type CreateTaskStateSpacesOK struct {
	Payload *models.TaskResource
}

func (o *CreateTaskStateSpacesOK) Error() string {
	return fmt.Sprintf("[POST /api/{baseSpaceId}/tasks/{id}/state][%d] createTaskStateSpacesOK  %+v", 200, o.Payload)
}

func (o *CreateTaskStateSpacesOK) GetPayload() *models.TaskResource {
	return o.Payload
}

func (o *CreateTaskStateSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTaskStateSpacesBadRequest creates a CreateTaskStateSpacesBadRequest with default headers values
func NewCreateTaskStateSpacesBadRequest() *CreateTaskStateSpacesBadRequest {
	return &CreateTaskStateSpacesBadRequest{}
}

/*CreateTaskStateSpacesBadRequest handles this case with default header values.

No body was provided.
No id parameter was provided.
No reason was provided.
No state was provided.
The task can not be put into the state.  Valid states are: Canceled, Failed, Success.
The task is already in the state.
The task state can not be changed to the selected state.  It can only be changed out of the states: Canceled, Failed, Success.
*/
type CreateTaskStateSpacesBadRequest struct {
}

func (o *CreateTaskStateSpacesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/{baseSpaceId}/tasks/{id}/state][%d] createTaskStateSpacesBadRequest ", 400)
}

func (o *CreateTaskStateSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
