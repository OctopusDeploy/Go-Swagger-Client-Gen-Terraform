// Code generated by go-swagger; DO NOT EDIT.

package project_triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteProjectTriggerReader is a Reader for the DeleteProjectTrigger structure.
type DeleteProjectTriggerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectTriggerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProjectTriggerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteProjectTriggerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteProjectTriggerOK creates a DeleteProjectTriggerOK with default headers values
func NewDeleteProjectTriggerOK() *DeleteProjectTriggerOK {
	return &DeleteProjectTriggerOK{}
}

/*DeleteProjectTriggerOK handles this case with default header values.

OK
*/
type DeleteProjectTriggerOK struct {
}

func (o *DeleteProjectTriggerOK) Error() string {
	return fmt.Sprintf("[DELETE /api/projecttriggers/{id}][%d] deleteProjectTriggerOK ", 200)
}

func (o *DeleteProjectTriggerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectTriggerBadRequest creates a DeleteProjectTriggerBadRequest with default headers values
func NewDeleteProjectTriggerBadRequest() *DeleteProjectTriggerBadRequest {
	return &DeleteProjectTriggerBadRequest{}
}

/*DeleteProjectTriggerBadRequest handles this case with default header values.

BadRequest
No id parameter was provided.
This resource is read-only and cannot be deleted.
*/
type DeleteProjectTriggerBadRequest struct {
}

func (o *DeleteProjectTriggerBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/projecttriggers/{id}][%d] deleteProjectTriggerBadRequest ", 400)
}

func (o *DeleteProjectTriggerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
