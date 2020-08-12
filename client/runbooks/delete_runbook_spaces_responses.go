// Code generated by go-swagger; DO NOT EDIT.

package runbooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteRunbookSpacesReader is a Reader for the DeleteRunbookSpaces structure.
type DeleteRunbookSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRunbookSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRunbookSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRunbookSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRunbookSpacesOK creates a DeleteRunbookSpacesOK with default headers values
func NewDeleteRunbookSpacesOK() *DeleteRunbookSpacesOK {
	return &DeleteRunbookSpacesOK{}
}

/*DeleteRunbookSpacesOK handles this case with default header values.

OK
*/
type DeleteRunbookSpacesOK struct {
}

func (o *DeleteRunbookSpacesOK) Error() string {
	return fmt.Sprintf("[DELETE /api/{baseSpaceId}/runbooks/{id}][%d] deleteRunbookSpacesOK ", 200)
}

func (o *DeleteRunbookSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRunbookSpacesBadRequest creates a DeleteRunbookSpacesBadRequest with default headers values
func NewDeleteRunbookSpacesBadRequest() *DeleteRunbookSpacesBadRequest {
	return &DeleteRunbookSpacesBadRequest{}
}

/*DeleteRunbookSpacesBadRequest handles this case with default header values.

BadRequest
No id parameter was provided.
This resource is read-only and cannot be deleted.
*/
type DeleteRunbookSpacesBadRequest struct {
}

func (o *DeleteRunbookSpacesBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/{baseSpaceId}/runbooks/{id}][%d] deleteRunbookSpacesBadRequest ", 400)
}

func (o *DeleteRunbookSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
