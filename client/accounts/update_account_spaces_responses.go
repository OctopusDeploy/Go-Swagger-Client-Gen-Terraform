// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// UpdateAccountSpacesReader is a Reader for the UpdateAccountSpaces structure.
type UpdateAccountSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAccountSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAccountSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateAccountSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateAccountSpacesOK creates a UpdateAccountSpacesOK with default headers values
func NewUpdateAccountSpacesOK() *UpdateAccountSpacesOK {
	return &UpdateAccountSpacesOK{}
}

/*UpdateAccountSpacesOK handles this case with default header values.

AccountResource Modified
*/
type UpdateAccountSpacesOK struct {
	Payload *models.AccountResource
}

func (o *UpdateAccountSpacesOK) Error() string {
	return fmt.Sprintf("[PUT /api/{baseSpaceId}/accounts/{id}][%d] updateAccountSpacesOK  %+v", 200, o.Payload)
}

func (o *UpdateAccountSpacesOK) GetPayload() *models.AccountResource {
	return o.Payload
}

func (o *UpdateAccountSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAccountSpacesBadRequest creates a UpdateAccountSpacesBadRequest with default headers values
func NewUpdateAccountSpacesBadRequest() *UpdateAccountSpacesBadRequest {
	return &UpdateAccountSpacesBadRequest{}
}

/*UpdateAccountSpacesBadRequest handles this case with default header values.

Model validation failed.
No id parameter was provided.
No request body was supplied.
This resource is read-only and cannot be modified.
*/
type UpdateAccountSpacesBadRequest struct {
}

func (o *UpdateAccountSpacesBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/{baseSpaceId}/accounts/{id}][%d] updateAccountSpacesBadRequest ", 400)
}

func (o *UpdateAccountSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
