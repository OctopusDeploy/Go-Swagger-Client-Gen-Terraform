// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateUserRegisterReader is a Reader for the CreateUserRegister structure.
type CreateUserRegisterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserRegisterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateUserRegisterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateUserRegisterOK creates a CreateUserRegisterOK with default headers values
func NewCreateUserRegisterOK() *CreateUserRegisterOK {
	return &CreateUserRegisterOK{}
}

/*CreateUserRegisterOK handles this case with default header values.

OK - Default
*/
type CreateUserRegisterOK struct {
}

func (o *CreateUserRegisterOK) Error() string {
	return fmt.Sprintf("[POST /api/users/register][%d] createUserRegisterOK ", 200)
}

func (o *CreateUserRegisterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}