// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetTenantTagTestSpacesReader is a Reader for the GetTenantTagTestSpaces structure.
type GetTenantTagTestSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTenantTagTestSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTenantTagTestSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTenantTagTestSpacesOK creates a GetTenantTagTestSpacesOK with default headers values
func NewGetTenantTagTestSpacesOK() *GetTenantTagTestSpacesOK {
	return &GetTenantTagTestSpacesOK{}
}

/*GetTenantTagTestSpacesOK handles this case with default header values.

OK - Default
*/
type GetTenantTagTestSpacesOK struct {
}

func (o *GetTenantTagTestSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/tenants/tag-test][%d] getTenantTagTestSpacesOK ", 200)
}

func (o *GetTenantTagTestSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
