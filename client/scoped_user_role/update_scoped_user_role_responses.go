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

// UpdateScopedUserRoleReader is a Reader for the UpdateScopedUserRole structure.
type UpdateScopedUserRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateScopedUserRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateScopedUserRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateScopedUserRoleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateScopedUserRoleOK creates a UpdateScopedUserRoleOK with default headers values
func NewUpdateScopedUserRoleOK() *UpdateScopedUserRoleOK {
	return &UpdateScopedUserRoleOK{}
}

/*UpdateScopedUserRoleOK handles this case with default header values.

ScopedUserRoleResource Modified
*/
type UpdateScopedUserRoleOK struct {
	Payload *models.ScopedUserRoleResource
}

func (o *UpdateScopedUserRoleOK) Error() string {
	return fmt.Sprintf("[PUT /api/scopeduserroles/{id}][%d] updateScopedUserRoleOK  %+v", 200, o.Payload)
}

func (o *UpdateScopedUserRoleOK) GetPayload() *models.ScopedUserRoleResource {
	return o.Payload
}

func (o *UpdateScopedUserRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScopedUserRoleResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateScopedUserRoleBadRequest creates a UpdateScopedUserRoleBadRequest with default headers values
func NewUpdateScopedUserRoleBadRequest() *UpdateScopedUserRoleBadRequest {
	return &UpdateScopedUserRoleBadRequest{}
}

/*UpdateScopedUserRoleBadRequest handles this case with default header values.

Model validation failed.
No id parameter was provided.
No request body was supplied.
This resource is read-only and cannot be modified.
*/
type UpdateScopedUserRoleBadRequest struct {
}

func (o *UpdateScopedUserRoleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/scopeduserroles/{id}][%d] updateScopedUserRoleBadRequest ", 400)
}

func (o *UpdateScopedUserRoleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}