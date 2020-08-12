// Code generated by go-swagger; DO NOT EDIT.

package scoped_user_role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// UpdateScopedUserRoleSpacesReader is a Reader for the UpdateScopedUserRoleSpaces structure.
type UpdateScopedUserRoleSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateScopedUserRoleSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateScopedUserRoleSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateScopedUserRoleSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateScopedUserRoleSpacesOK creates a UpdateScopedUserRoleSpacesOK with default headers values
func NewUpdateScopedUserRoleSpacesOK() *UpdateScopedUserRoleSpacesOK {
	return &UpdateScopedUserRoleSpacesOK{}
}

/*UpdateScopedUserRoleSpacesOK handles this case with default header values.

ScopedUserRoleResource Modified
*/
type UpdateScopedUserRoleSpacesOK struct {
	Payload *models.ScopedUserRoleResource
}

func (o *UpdateScopedUserRoleSpacesOK) Error() string {
	return fmt.Sprintf("[PUT /api/{baseSpaceId}/scopeduserroles/{id}][%d] updateScopedUserRoleSpacesOK  %+v", 200, o.Payload)
}

func (o *UpdateScopedUserRoleSpacesOK) GetPayload() *models.ScopedUserRoleResource {
	return o.Payload
}

func (o *UpdateScopedUserRoleSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScopedUserRoleResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateScopedUserRoleSpacesBadRequest creates a UpdateScopedUserRoleSpacesBadRequest with default headers values
func NewUpdateScopedUserRoleSpacesBadRequest() *UpdateScopedUserRoleSpacesBadRequest {
	return &UpdateScopedUserRoleSpacesBadRequest{}
}

/*UpdateScopedUserRoleSpacesBadRequest handles this case with default header values.

Model validation failed.
No id parameter was provided.
No request body was supplied.
This resource is read-only and cannot be modified.
*/
type UpdateScopedUserRoleSpacesBadRequest struct {
}

func (o *UpdateScopedUserRoleSpacesBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/{baseSpaceId}/scopeduserroles/{id}][%d] updateScopedUserRoleSpacesBadRequest ", 400)
}

func (o *UpdateScopedUserRoleSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}