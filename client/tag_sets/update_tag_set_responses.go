// Code generated by go-swagger; DO NOT EDIT.

package tag_sets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// UpdateTagSetReader is a Reader for the UpdateTagSet structure.
type UpdateTagSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTagSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTagSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateTagSetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateTagSetOK creates a UpdateTagSetOK with default headers values
func NewUpdateTagSetOK() *UpdateTagSetOK {
	return &UpdateTagSetOK{}
}

/*UpdateTagSetOK handles this case with default header values.

TagSetResource Modified
*/
type UpdateTagSetOK struct {
	Payload *models.TagSetResource
}

func (o *UpdateTagSetOK) Error() string {
	return fmt.Sprintf("[PUT /api/tagsets/{id}][%d] updateTagSetOK  %+v", 200, o.Payload)
}

func (o *UpdateTagSetOK) GetPayload() *models.TagSetResource {
	return o.Payload
}

func (o *UpdateTagSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TagSetResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTagSetBadRequest creates a UpdateTagSetBadRequest with default headers values
func NewUpdateTagSetBadRequest() *UpdateTagSetBadRequest {
	return &UpdateTagSetBadRequest{}
}

/*UpdateTagSetBadRequest handles this case with default header values.

Model validation failed.
No id parameter was provided.
No request body was supplied.
This resource is read-only and cannot be modified.
*/
type UpdateTagSetBadRequest struct {
}

func (o *UpdateTagSetBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/tagsets/{id}][%d] updateTagSetBadRequest ", 400)
}

func (o *UpdateTagSetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}