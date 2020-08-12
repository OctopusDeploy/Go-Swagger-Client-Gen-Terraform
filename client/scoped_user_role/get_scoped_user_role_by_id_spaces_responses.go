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

// GetScopedUserRoleByIDSpacesReader is a Reader for the GetScopedUserRoleByIDSpaces structure.
type GetScopedUserRoleByIDSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScopedUserRoleByIDSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScopedUserRoleByIDSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetScopedUserRoleByIDSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScopedUserRoleByIDSpacesOK creates a GetScopedUserRoleByIDSpacesOK with default headers values
func NewGetScopedUserRoleByIDSpacesOK() *GetScopedUserRoleByIDSpacesOK {
	return &GetScopedUserRoleByIDSpacesOK{}
}

/*GetScopedUserRoleByIDSpacesOK handles this case with default header values.

Load ScopedUserRoleResource by ID
*/
type GetScopedUserRoleByIDSpacesOK struct {
	Payload *models.ScopedUserRoleResource
}

func (o *GetScopedUserRoleByIDSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/scopeduserroles/{id}][%d] getScopedUserRoleByIdSpacesOK  %+v", 200, o.Payload)
}

func (o *GetScopedUserRoleByIDSpacesOK) GetPayload() *models.ScopedUserRoleResource {
	return o.Payload
}

func (o *GetScopedUserRoleByIDSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScopedUserRoleResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScopedUserRoleByIDSpacesBadRequest creates a GetScopedUserRoleByIDSpacesBadRequest with default headers values
func NewGetScopedUserRoleByIDSpacesBadRequest() *GetScopedUserRoleByIDSpacesBadRequest {
	return &GetScopedUserRoleByIDSpacesBadRequest{}
}

/*GetScopedUserRoleByIDSpacesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetScopedUserRoleByIDSpacesBadRequest struct {
}

func (o *GetScopedUserRoleByIDSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/scopeduserroles/{id}][%d] getScopedUserRoleByIdSpacesBadRequest ", 400)
}

func (o *GetScopedUserRoleByIDSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}