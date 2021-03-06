// Code generated by go-swagger; DO NOT EDIT.

package runbooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// UpdateRunbookSpacesReader is a Reader for the UpdateRunbookSpaces structure.
type UpdateRunbookSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRunbookSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRunbookSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRunbookSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRunbookSpacesOK creates a UpdateRunbookSpacesOK with default headers values
func NewUpdateRunbookSpacesOK() *UpdateRunbookSpacesOK {
	return &UpdateRunbookSpacesOK{}
}

/*UpdateRunbookSpacesOK handles this case with default header values.

RunbookResource Modified
*/
type UpdateRunbookSpacesOK struct {
	Payload *models.RunbookResource
}

func (o *UpdateRunbookSpacesOK) Error() string {
	return fmt.Sprintf("[PUT /api/{baseSpaceId}/runbooks/{id}][%d] updateRunbookSpacesOK  %+v", 200, o.Payload)
}

func (o *UpdateRunbookSpacesOK) GetPayload() *models.RunbookResource {
	return o.Payload
}

func (o *UpdateRunbookSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RunbookResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRunbookSpacesBadRequest creates a UpdateRunbookSpacesBadRequest with default headers values
func NewUpdateRunbookSpacesBadRequest() *UpdateRunbookSpacesBadRequest {
	return &UpdateRunbookSpacesBadRequest{}
}

/*UpdateRunbookSpacesBadRequest handles this case with default header values.

Model validation failed.
No id parameter was provided.
No request body was supplied.
This resource is read-only and cannot be modified.
*/
type UpdateRunbookSpacesBadRequest struct {
}

func (o *UpdateRunbookSpacesBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/{baseSpaceId}/runbooks/{id}][%d] updateRunbookSpacesBadRequest ", 400)
}

func (o *UpdateRunbookSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
