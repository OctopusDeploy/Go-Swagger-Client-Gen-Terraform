// Code generated by go-swagger; DO NOT EDIT.

package proxies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteProxySpacesReader is a Reader for the DeleteProxySpaces structure.
type DeleteProxySpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProxySpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProxySpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteProxySpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteProxySpacesOK creates a DeleteProxySpacesOK with default headers values
func NewDeleteProxySpacesOK() *DeleteProxySpacesOK {
	return &DeleteProxySpacesOK{}
}

/*DeleteProxySpacesOK handles this case with default header values.

OK
*/
type DeleteProxySpacesOK struct {
}

func (o *DeleteProxySpacesOK) Error() string {
	return fmt.Sprintf("[DELETE /api/{baseSpaceId}/proxies/{id}][%d] deleteProxySpacesOK ", 200)
}

func (o *DeleteProxySpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProxySpacesBadRequest creates a DeleteProxySpacesBadRequest with default headers values
func NewDeleteProxySpacesBadRequest() *DeleteProxySpacesBadRequest {
	return &DeleteProxySpacesBadRequest{}
}

/*DeleteProxySpacesBadRequest handles this case with default header values.

BadRequest
No id parameter was provided.
This resource is read-only and cannot be deleted.
*/
type DeleteProxySpacesBadRequest struct {
}

func (o *DeleteProxySpacesBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/{baseSpaceId}/proxies/{id}][%d] deleteProxySpacesBadRequest ", 400)
}

func (o *DeleteProxySpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
