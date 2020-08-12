// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteTenantSpacesReader is a Reader for the DeleteTenantSpaces structure.
type DeleteTenantSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTenantSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTenantSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTenantSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTenantSpacesOK creates a DeleteTenantSpacesOK with default headers values
func NewDeleteTenantSpacesOK() *DeleteTenantSpacesOK {
	return &DeleteTenantSpacesOK{}
}

/*DeleteTenantSpacesOK handles this case with default header values.

OK
*/
type DeleteTenantSpacesOK struct {
}

func (o *DeleteTenantSpacesOK) Error() string {
	return fmt.Sprintf("[DELETE /api/{baseSpaceId}/tenants/{id}][%d] deleteTenantSpacesOK ", 200)
}

func (o *DeleteTenantSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTenantSpacesBadRequest creates a DeleteTenantSpacesBadRequest with default headers values
func NewDeleteTenantSpacesBadRequest() *DeleteTenantSpacesBadRequest {
	return &DeleteTenantSpacesBadRequest{}
}

/*DeleteTenantSpacesBadRequest handles this case with default header values.

BadRequest
No id parameter was provided.
This resource is read-only and cannot be deleted.
*/
type DeleteTenantSpacesBadRequest struct {
}

func (o *DeleteTenantSpacesBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/{baseSpaceId}/tenants/{id}][%d] deleteTenantSpacesBadRequest ", 400)
}

func (o *DeleteTenantSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}