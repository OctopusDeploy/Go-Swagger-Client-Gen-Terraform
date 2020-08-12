// Code generated by go-swagger; DO NOT EDIT.

package lifecycles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteLifecycleReader is a Reader for the DeleteLifecycle structure.
type DeleteLifecycleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLifecycleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteLifecycleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLifecycleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteLifecycleOK creates a DeleteLifecycleOK with default headers values
func NewDeleteLifecycleOK() *DeleteLifecycleOK {
	return &DeleteLifecycleOK{}
}

/*DeleteLifecycleOK handles this case with default header values.

OK
*/
type DeleteLifecycleOK struct {
}

func (o *DeleteLifecycleOK) Error() string {
	return fmt.Sprintf("[DELETE /api/lifecycles/{id}][%d] deleteLifecycleOK ", 200)
}

func (o *DeleteLifecycleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLifecycleBadRequest creates a DeleteLifecycleBadRequest with default headers values
func NewDeleteLifecycleBadRequest() *DeleteLifecycleBadRequest {
	return &DeleteLifecycleBadRequest{}
}

/*DeleteLifecycleBadRequest handles this case with default header values.

BadRequest
No id parameter was provided.
This resource is read-only and cannot be deleted.
*/
type DeleteLifecycleBadRequest struct {
}

func (o *DeleteLifecycleBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/lifecycles/{id}][%d] deleteLifecycleBadRequest ", 400)
}

func (o *DeleteLifecycleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
