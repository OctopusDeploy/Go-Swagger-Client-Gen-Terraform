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

// CreateTenantSpacesReader is a Reader for the CreateTenantSpaces structure.
type CreateTenantSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTenantSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateTenantSpacesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTenantSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTenantSpacesCreated creates a CreateTenantSpacesCreated with default headers values
func NewCreateTenantSpacesCreated() *CreateTenantSpacesCreated {
	return &CreateTenantSpacesCreated{}
}

/*CreateTenantSpacesCreated handles this case with default header values.

TenantResource Created
*/
type CreateTenantSpacesCreated struct {
	Payload *models.TenantResource
}

func (o *CreateTenantSpacesCreated) Error() string {
	return fmt.Sprintf("[POST /api/{baseSpaceId}/tenants][%d] createTenantSpacesCreated  %+v", 201, o.Payload)
}

func (o *CreateTenantSpacesCreated) GetPayload() *models.TenantResource {
	return o.Payload
}

func (o *CreateTenantSpacesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTenantSpacesBadRequest creates a CreateTenantSpacesBadRequest with default headers values
func NewCreateTenantSpacesBadRequest() *CreateTenantSpacesBadRequest {
	return &CreateTenantSpacesBadRequest{}
}

/*CreateTenantSpacesBadRequest handles this case with default header values.

Model validation failed.
No request body was supplied.
*/
type CreateTenantSpacesBadRequest struct {
}

func (o *CreateTenantSpacesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/{baseSpaceId}/tenants][%d] createTenantSpacesBadRequest ", 400)
}

func (o *CreateTenantSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}