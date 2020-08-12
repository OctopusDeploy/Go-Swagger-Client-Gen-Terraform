// Code generated by go-swagger; DO NOT EDIT.

package configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateConfigurationItemReader is a Reader for the UpdateConfigurationItem structure.
type UpdateConfigurationItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateConfigurationItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateConfigurationItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateConfigurationItemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateConfigurationItemOK creates a UpdateConfigurationItemOK with default headers values
func NewUpdateConfigurationItemOK() *UpdateConfigurationItemOK {
	return &UpdateConfigurationItemOK{}
}

/*UpdateConfigurationItemOK handles this case with default header values.

Object resource returned
*/
type UpdateConfigurationItemOK struct {
	Payload interface{}
}

func (o *UpdateConfigurationItemOK) Error() string {
	return fmt.Sprintf("[PUT /api/configuration/{id}/values][%d] updateConfigurationItemOK  %+v", 200, o.Payload)
}

func (o *UpdateConfigurationItemOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateConfigurationItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConfigurationItemBadRequest creates a UpdateConfigurationItemBadRequest with default headers values
func NewUpdateConfigurationItemBadRequest() *UpdateConfigurationItemBadRequest {
	return &UpdateConfigurationItemBadRequest{}
}

/*UpdateConfigurationItemBadRequest handles this case with default header values.

No id parameter was provided.
No request body was supplied.
*/
type UpdateConfigurationItemBadRequest struct {
}

func (o *UpdateConfigurationItemBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/configuration/{id}/values][%d] updateConfigurationItemBadRequest ", 400)
}

func (o *UpdateConfigurationItemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}