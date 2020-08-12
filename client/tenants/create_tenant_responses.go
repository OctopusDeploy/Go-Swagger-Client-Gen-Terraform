// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// CreateTenantReader is a Reader for the CreateTenant structure.
type CreateTenantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTenantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateTenantCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTenantBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTenantCreated creates a CreateTenantCreated with default headers values
func NewCreateTenantCreated() *CreateTenantCreated {
	return &CreateTenantCreated{}
}

/*CreateTenantCreated handles this case with default header values.

TenantResource Created
*/
type CreateTenantCreated struct {
	Payload *models.TenantResource
}

func (o *CreateTenantCreated) Error() string {
	return fmt.Sprintf("[POST /api/tenants][%d] createTenantCreated  %+v", 201, o.Payload)
}

func (o *CreateTenantCreated) GetPayload() *models.TenantResource {
	return o.Payload
}

func (o *CreateTenantCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTenantBadRequest creates a CreateTenantBadRequest with default headers values
func NewCreateTenantBadRequest() *CreateTenantBadRequest {
	return &CreateTenantBadRequest{}
}

/*CreateTenantBadRequest handles this case with default header values.

Model validation failed.
No request body was supplied.
*/
type CreateTenantBadRequest struct {
}

func (o *CreateTenantBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/tenants][%d] createTenantBadRequest ", 400)
}

func (o *CreateTenantBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}