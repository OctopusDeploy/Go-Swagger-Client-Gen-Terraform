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

// GetTaskDetailsSpacesReader is a Reader for the GetTaskDetailsSpaces structure.
type GetTaskDetailsSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTaskDetailsSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTaskDetailsSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTaskDetailsSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTaskDetailsSpacesOK creates a GetTaskDetailsSpacesOK with default headers values
func NewGetTaskDetailsSpacesOK() *GetTaskDetailsSpacesOK {
	return &GetTaskDetailsSpacesOK{}
}

/*GetTaskDetailsSpacesOK handles this case with default header values.

TaskDetailsResource resource returned
*/
type GetTaskDetailsSpacesOK struct {
	Payload *models.TaskDetailsResource
}

func (o *GetTaskDetailsSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/tasks/{id}/details][%d] getTaskDetailsSpacesOK  %+v", 200, o.Payload)
}

func (o *GetTaskDetailsSpacesOK) GetPayload() *models.TaskDetailsResource {
	return o.Payload
}

func (o *GetTaskDetailsSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskDetailsResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskDetailsSpacesBadRequest creates a GetTaskDetailsSpacesBadRequest with default headers values
func NewGetTaskDetailsSpacesBadRequest() *GetTaskDetailsSpacesBadRequest {
	return &GetTaskDetailsSpacesBadRequest{}
}

/*GetTaskDetailsSpacesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetTaskDetailsSpacesBadRequest struct {
}

func (o *GetTaskDetailsSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/tasks/{id}/details][%d] getTaskDetailsSpacesBadRequest ", 400)
}

func (o *GetTaskDetailsSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
