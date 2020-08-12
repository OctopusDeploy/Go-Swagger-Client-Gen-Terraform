// Code generated by go-swagger; DO NOT EDIT.

package tenant_variables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetTenantVariablesAllSpacesReader is a Reader for the GetTenantVariablesAllSpaces structure.
type GetTenantVariablesAllSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTenantVariablesAllSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTenantVariablesAllSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTenantVariablesAllSpacesOK creates a GetTenantVariablesAllSpacesOK with default headers values
func NewGetTenantVariablesAllSpacesOK() *GetTenantVariablesAllSpacesOK {
	return &GetTenantVariablesAllSpacesOK{}
}

/*GetTenantVariablesAllSpacesOK handles this case with default header values.

OK - Default
*/
type GetTenantVariablesAllSpacesOK struct {
}

func (o *GetTenantVariablesAllSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/tenantvariables/all][%d] getTenantVariablesAllSpacesOK ", 200)
}

func (o *GetTenantVariablesAllSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}