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

// CreateTenantVariablesReader is a Reader for the CreateTenantVariables structure.
type CreateTenantVariablesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTenantVariablesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTenantVariablesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTenantVariablesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTenantVariablesOK creates a CreateTenantVariablesOK with default headers values
func NewCreateTenantVariablesOK() *CreateTenantVariablesOK {
	return &CreateTenantVariablesOK{}
}

/*CreateTenantVariablesOK handles this case with default header values.

TenantVariableResource resource returned
*/
type CreateTenantVariablesOK struct {
	Payload *models.TenantVariableResource
}

func (o *CreateTenantVariablesOK) Error() string {
	return fmt.Sprintf("[POST /api/tenants/{id}/variables][%d] createTenantVariablesOK  %+v", 200, o.Payload)
}

func (o *CreateTenantVariablesOK) GetPayload() *models.TenantVariableResource {
	return o.Payload
}

func (o *CreateTenantVariablesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantVariableResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTenantVariablesBadRequest creates a CreateTenantVariablesBadRequest with default headers values
func NewCreateTenantVariablesBadRequest() *CreateTenantVariablesBadRequest {
	return &CreateTenantVariablesBadRequest{}
}

/*CreateTenantVariablesBadRequest handles this case with default header values.

No id parameter was provided.
No request body was supplied.
*/
type CreateTenantVariablesBadRequest struct {
}

func (o *CreateTenantVariablesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/tenants/{id}/variables][%d] createTenantVariablesBadRequest ", 400)
}

func (o *CreateTenantVariablesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
