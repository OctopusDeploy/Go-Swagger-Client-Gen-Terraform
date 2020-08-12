// Code generated by go-swagger; DO NOT EDIT.

package lifecycles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteLifecycleSpacesReader is a Reader for the DeleteLifecycleSpaces structure.
type DeleteLifecycleSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLifecycleSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteLifecycleSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLifecycleSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteLifecycleSpacesOK creates a DeleteLifecycleSpacesOK with default headers values
func NewDeleteLifecycleSpacesOK() *DeleteLifecycleSpacesOK {
	return &DeleteLifecycleSpacesOK{}
}

/*DeleteLifecycleSpacesOK handles this case with default header values.

OK
*/
type DeleteLifecycleSpacesOK struct {
}

func (o *DeleteLifecycleSpacesOK) Error() string {
	return fmt.Sprintf("[DELETE /api/{baseSpaceId}/lifecycles/{id}][%d] deleteLifecycleSpacesOK ", 200)
}

func (o *DeleteLifecycleSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLifecycleSpacesBadRequest creates a DeleteLifecycleSpacesBadRequest with default headers values
func NewDeleteLifecycleSpacesBadRequest() *DeleteLifecycleSpacesBadRequest {
	return &DeleteLifecycleSpacesBadRequest{}
}

/*DeleteLifecycleSpacesBadRequest handles this case with default header values.

BadRequest
No id parameter was provided.
This resource is read-only and cannot be deleted.
*/
type DeleteLifecycleSpacesBadRequest struct {
}

func (o *DeleteLifecycleSpacesBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/{baseSpaceId}/lifecycles/{id}][%d] deleteLifecycleSpacesBadRequest ", 400)
}

func (o *DeleteLifecycleSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}