// Code generated by go-swagger; DO NOT EDIT.

package project_triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// UpdateProjectTriggerSpacesReader is a Reader for the UpdateProjectTriggerSpaces structure.
type UpdateProjectTriggerSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProjectTriggerSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProjectTriggerSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateProjectTriggerSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateProjectTriggerSpacesOK creates a UpdateProjectTriggerSpacesOK with default headers values
func NewUpdateProjectTriggerSpacesOK() *UpdateProjectTriggerSpacesOK {
	return &UpdateProjectTriggerSpacesOK{}
}

/*UpdateProjectTriggerSpacesOK handles this case with default header values.

ProjectTriggerResource Modified
*/
type UpdateProjectTriggerSpacesOK struct {
	Payload *models.ProjectTriggerResource
}

func (o *UpdateProjectTriggerSpacesOK) Error() string {
	return fmt.Sprintf("[PUT /api/{baseSpaceId}/projecttriggers/{id}][%d] updateProjectTriggerSpacesOK  %+v", 200, o.Payload)
}

func (o *UpdateProjectTriggerSpacesOK) GetPayload() *models.ProjectTriggerResource {
	return o.Payload
}

func (o *UpdateProjectTriggerSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectTriggerResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectTriggerSpacesBadRequest creates a UpdateProjectTriggerSpacesBadRequest with default headers values
func NewUpdateProjectTriggerSpacesBadRequest() *UpdateProjectTriggerSpacesBadRequest {
	return &UpdateProjectTriggerSpacesBadRequest{}
}

/*UpdateProjectTriggerSpacesBadRequest handles this case with default header values.

Model validation failed.
No id parameter was provided.
No request body was supplied.
This resource is read-only and cannot be modified.
*/
type UpdateProjectTriggerSpacesBadRequest struct {
}

func (o *UpdateProjectTriggerSpacesBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/{baseSpaceId}/projecttriggers/{id}][%d] updateProjectTriggerSpacesBadRequest ", 400)
}

func (o *UpdateProjectTriggerSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}