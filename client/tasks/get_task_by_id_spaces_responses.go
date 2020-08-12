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

// GetTaskByIDSpacesReader is a Reader for the GetTaskByIDSpaces structure.
type GetTaskByIDSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTaskByIDSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTaskByIDSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTaskByIDSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTaskByIDSpacesOK creates a GetTaskByIDSpacesOK with default headers values
func NewGetTaskByIDSpacesOK() *GetTaskByIDSpacesOK {
	return &GetTaskByIDSpacesOK{}
}

/*GetTaskByIDSpacesOK handles this case with default header values.

Load TaskResource by ID
*/
type GetTaskByIDSpacesOK struct {
	Payload *models.TaskResource
}

func (o *GetTaskByIDSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/tasks/{id}][%d] getTaskByIdSpacesOK  %+v", 200, o.Payload)
}

func (o *GetTaskByIDSpacesOK) GetPayload() *models.TaskResource {
	return o.Payload
}

func (o *GetTaskByIDSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskByIDSpacesBadRequest creates a GetTaskByIDSpacesBadRequest with default headers values
func NewGetTaskByIDSpacesBadRequest() *GetTaskByIDSpacesBadRequest {
	return &GetTaskByIDSpacesBadRequest{}
}

/*GetTaskByIDSpacesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetTaskByIDSpacesBadRequest struct {
}

func (o *GetTaskByIDSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/tasks/{id}][%d] getTaskByIdSpacesBadRequest ", 400)
}

func (o *GetTaskByIDSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
