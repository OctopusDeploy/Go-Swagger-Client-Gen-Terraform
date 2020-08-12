// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteTenantReader is a Reader for the DeleteTenant structure.
type DeleteTenantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTenantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTenantOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTenantBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTenantOK creates a DeleteTenantOK with default headers values
func NewDeleteTenantOK() *DeleteTenantOK {
	return &DeleteTenantOK{}
}

/*DeleteTenantOK handles this case with default header values.

OK
*/
type DeleteTenantOK struct {
}

func (o *DeleteTenantOK) Error() string {
	return fmt.Sprintf("[DELETE /api/tenants/{id}][%d] deleteTenantOK ", 200)
}

func (o *DeleteTenantOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTenantBadRequest creates a DeleteTenantBadRequest with default headers values
func NewDeleteTenantBadRequest() *DeleteTenantBadRequest {
	return &DeleteTenantBadRequest{}
}

/*DeleteTenantBadRequest handles this case with default header values.

BadRequest
No id parameter was provided.
This resource is read-only and cannot be deleted.
*/
type DeleteTenantBadRequest struct {
}

func (o *DeleteTenantBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/tenants/{id}][%d] deleteTenantBadRequest ", 400)
}

func (o *DeleteTenantBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}