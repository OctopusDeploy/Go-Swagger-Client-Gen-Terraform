// Code generated by go-swagger; DO NOT EDIT.

package environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// CreateEnvironmentReader is a Reader for the CreateEnvironment structure.
type CreateEnvironmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateEnvironmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateEnvironmentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateEnvironmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateEnvironmentCreated creates a CreateEnvironmentCreated with default headers values
func NewCreateEnvironmentCreated() *CreateEnvironmentCreated {
	return &CreateEnvironmentCreated{}
}

/*CreateEnvironmentCreated handles this case with default header values.

EnvironmentResource Created
*/
type CreateEnvironmentCreated struct {
	Payload *models.EnvironmentResource
}

func (o *CreateEnvironmentCreated) Error() string {
	return fmt.Sprintf("[POST /api/environments][%d] createEnvironmentCreated  %+v", 201, o.Payload)
}

func (o *CreateEnvironmentCreated) GetPayload() *models.EnvironmentResource {
	return o.Payload
}

func (o *CreateEnvironmentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnvironmentResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEnvironmentBadRequest creates a CreateEnvironmentBadRequest with default headers values
func NewCreateEnvironmentBadRequest() *CreateEnvironmentBadRequest {
	return &CreateEnvironmentBadRequest{}
}

/*CreateEnvironmentBadRequest handles this case with default header values.

Model validation failed.
No request body was supplied.
*/
type CreateEnvironmentBadRequest struct {
}

func (o *CreateEnvironmentBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/environments][%d] createEnvironmentBadRequest ", 400)
}

func (o *CreateEnvironmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
