// Code generated by go-swagger; DO NOT EDIT.

package lifecycles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// UpdateLifecycleReader is a Reader for the UpdateLifecycle structure.
type UpdateLifecycleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateLifecycleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateLifecycleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateLifecycleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateLifecycleOK creates a UpdateLifecycleOK with default headers values
func NewUpdateLifecycleOK() *UpdateLifecycleOK {
	return &UpdateLifecycleOK{}
}

/*UpdateLifecycleOK handles this case with default header values.

LifecycleResource Modified
*/
type UpdateLifecycleOK struct {
	Payload *models.LifecycleResource
}

func (o *UpdateLifecycleOK) Error() string {
	return fmt.Sprintf("[PUT /api/lifecycles/{id}][%d] updateLifecycleOK  %+v", 200, o.Payload)
}

func (o *UpdateLifecycleOK) GetPayload() *models.LifecycleResource {
	return o.Payload
}

func (o *UpdateLifecycleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LifecycleResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLifecycleBadRequest creates a UpdateLifecycleBadRequest with default headers values
func NewUpdateLifecycleBadRequest() *UpdateLifecycleBadRequest {
	return &UpdateLifecycleBadRequest{}
}

/*UpdateLifecycleBadRequest handles this case with default header values.

Model validation failed.
No id parameter was provided.
No request body was supplied.
This resource is read-only and cannot be modified.
*/
type UpdateLifecycleBadRequest struct {
}

func (o *UpdateLifecycleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/lifecycles/{id}][%d] updateLifecycleBadRequest ", 400)
}

func (o *UpdateLifecycleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
