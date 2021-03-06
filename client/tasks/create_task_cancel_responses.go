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

// CreateTaskCancelReader is a Reader for the CreateTaskCancel structure.
type CreateTaskCancelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTaskCancelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTaskCancelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTaskCancelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTaskCancelOK creates a CreateTaskCancelOK with default headers values
func NewCreateTaskCancelOK() *CreateTaskCancelOK {
	return &CreateTaskCancelOK{}
}

/*CreateTaskCancelOK handles this case with default header values.

TaskResource resource returned
*/
type CreateTaskCancelOK struct {
	Payload *models.TaskResource
}

func (o *CreateTaskCancelOK) Error() string {
	return fmt.Sprintf("[POST /api/tasks/{id}/cancel][%d] createTaskCancelOK  %+v", 200, o.Payload)
}

func (o *CreateTaskCancelOK) GetPayload() *models.TaskResource {
	return o.Payload
}

func (o *CreateTaskCancelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTaskCancelBadRequest creates a CreateTaskCancelBadRequest with default headers values
func NewCreateTaskCancelBadRequest() *CreateTaskCancelBadRequest {
	return &CreateTaskCancelBadRequest{}
}

/*CreateTaskCancelBadRequest handles this case with default header values.

No id parameter was provided.
*/
type CreateTaskCancelBadRequest struct {
}

func (o *CreateTaskCancelBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/tasks/{id}/cancel][%d] createTaskCancelBadRequest ", 400)
}

func (o *CreateTaskCancelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
