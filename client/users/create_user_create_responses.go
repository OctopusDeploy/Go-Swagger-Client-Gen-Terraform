// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// CreateUserCreateReader is a Reader for the CreateUserCreate structure.
type CreateUserCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateUserCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateUserCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateUserCreateCreated creates a CreateUserCreateCreated with default headers values
func NewCreateUserCreateCreated() *CreateUserCreateCreated {
	return &CreateUserCreateCreated{}
}

/*CreateUserCreateCreated handles this case with default header values.

UserResource resource returned
*/
type CreateUserCreateCreated struct {
	Payload *models.UserResource
}

func (o *CreateUserCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/users][%d] createUserCreateCreated  %+v", 201, o.Payload)
}

func (o *CreateUserCreateCreated) GetPayload() *models.UserResource {
	return o.Payload
}

func (o *CreateUserCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserCreateBadRequest creates a CreateUserCreateBadRequest with default headers values
func NewCreateUserCreateBadRequest() *CreateUserCreateBadRequest {
	return &CreateUserCreateBadRequest{}
}

/*CreateUserCreateBadRequest handles this case with default header values.

Model validation failed.
No request body was supplied.
User creation failed.
Username is reserved and can't be used to create new users.
You cannot add multiple logins for the same identity provider.
*/
type CreateUserCreateBadRequest struct {
}

func (o *CreateUserCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/users][%d] createUserCreateBadRequest ", 400)
}

func (o *CreateUserCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
