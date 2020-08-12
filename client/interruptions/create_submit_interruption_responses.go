// Code generated by go-swagger; DO NOT EDIT.

package interruptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// CreateSubmitInterruptionReader is a Reader for the CreateSubmitInterruption structure.
type CreateSubmitInterruptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSubmitInterruptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSubmitInterruptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateSubmitInterruptionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateSubmitInterruptionOK creates a CreateSubmitInterruptionOK with default headers values
func NewCreateSubmitInterruptionOK() *CreateSubmitInterruptionOK {
	return &CreateSubmitInterruptionOK{}
}

/*CreateSubmitInterruptionOK handles this case with default header values.

InterruptionResource resource returned
*/
type CreateSubmitInterruptionOK struct {
	Payload *models.InterruptionResource
}

func (o *CreateSubmitInterruptionOK) Error() string {
	return fmt.Sprintf("[POST /api/interruptions/{id}/submit][%d] createSubmitInterruptionOK  %+v", 200, o.Payload)
}

func (o *CreateSubmitInterruptionOK) GetPayload() *models.InterruptionResource {
	return o.Payload
}

func (o *CreateSubmitInterruptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InterruptionResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSubmitInterruptionBadRequest creates a CreateSubmitInterruptionBadRequest with default headers values
func NewCreateSubmitInterruptionBadRequest() *CreateSubmitInterruptionBadRequest {
	return &CreateSubmitInterruptionBadRequest{}
}

/*CreateSubmitInterruptionBadRequest handles this case with default header values.

No id parameter was provided.
No request body was supplied.
Someone else has taken responsibility for this interruption, and only they are allowed to submit it.
The interruption has already been submitted.
There were errors with the form.
You must take responsibility for this interruption before attempting to submit it.
*/
type CreateSubmitInterruptionBadRequest struct {
}

func (o *CreateSubmitInterruptionBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/interruptions/{id}/submit][%d] createSubmitInterruptionBadRequest ", 400)
}

func (o *CreateSubmitInterruptionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}