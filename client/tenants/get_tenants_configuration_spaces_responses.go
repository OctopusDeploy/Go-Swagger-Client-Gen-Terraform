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

// GetTenantsConfigurationSpacesReader is a Reader for the GetTenantsConfigurationSpaces structure.
type GetTenantsConfigurationSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTenantsConfigurationSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTenantsConfigurationSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTenantsConfigurationSpacesOK creates a GetTenantsConfigurationSpacesOK with default headers values
func NewGetTenantsConfigurationSpacesOK() *GetTenantsConfigurationSpacesOK {
	return &GetTenantsConfigurationSpacesOK{}
}

/*GetTenantsConfigurationSpacesOK handles this case with default header values.

MultiTenancyStatusResource resource returned
*/
type GetTenantsConfigurationSpacesOK struct {
	Payload *models.MultiTenancyStatusResource
}

func (o *GetTenantsConfigurationSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/tenants/status][%d] getTenantsConfigurationSpacesOK  %+v", 200, o.Payload)
}

func (o *GetTenantsConfigurationSpacesOK) GetPayload() *models.MultiTenancyStatusResource {
	return o.Payload
}

func (o *GetTenantsConfigurationSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MultiTenancyStatusResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
