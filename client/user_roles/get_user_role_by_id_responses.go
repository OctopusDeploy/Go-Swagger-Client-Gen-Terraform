// Code generated by go-swagger; DO NOT EDIT.

package user_roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetUserRoleByIDReader is a Reader for the GetUserRoleByID structure.
type GetUserRoleByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserRoleByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserRoleByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserRoleByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserRoleByIDOK creates a GetUserRoleByIDOK with default headers values
func NewGetUserRoleByIDOK() *GetUserRoleByIDOK {
	return &GetUserRoleByIDOK{}
}

/*GetUserRoleByIDOK handles this case with default header values.

Load UserRoleResource by ID
*/
type GetUserRoleByIDOK struct {
	Payload *models.UserRoleResource
}

func (o *GetUserRoleByIDOK) Error() string {
	return fmt.Sprintf("[GET /api/userroles/{id}][%d] getUserRoleByIdOK  %+v", 200, o.Payload)
}

func (o *GetUserRoleByIDOK) GetPayload() *models.UserRoleResource {
	return o.Payload
}

func (o *GetUserRoleByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserRoleResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRoleByIDBadRequest creates a GetUserRoleByIDBadRequest with default headers values
func NewGetUserRoleByIDBadRequest() *GetUserRoleByIDBadRequest {
	return &GetUserRoleByIDBadRequest{}
}

/*GetUserRoleByIDBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetUserRoleByIDBadRequest struct {
}

func (o *GetUserRoleByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/userroles/{id}][%d] getUserRoleByIdBadRequest ", 400)
}

func (o *GetUserRoleByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}