// Code generated by go-swagger; DO NOT EDIT.

package spaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteSpaceReader is a Reader for the DeleteSpace structure.
type DeleteSpaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSpaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSpaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteSpaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteSpaceOK creates a DeleteSpaceOK with default headers values
func NewDeleteSpaceOK() *DeleteSpaceOK {
	return &DeleteSpaceOK{}
}

/*DeleteSpaceOK handles this case with default header values.

OK
*/
type DeleteSpaceOK struct {
}

func (o *DeleteSpaceOK) Error() string {
	return fmt.Sprintf("[DELETE /api/spaces/{id}][%d] deleteSpaceOK ", 200)
}

func (o *DeleteSpaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSpaceBadRequest creates a DeleteSpaceBadRequest with default headers values
func NewDeleteSpaceBadRequest() *DeleteSpaceBadRequest {
	return &DeleteSpaceBadRequest{}
}

/*DeleteSpaceBadRequest handles this case with default header values.

No id parameter was provided.
*/
type DeleteSpaceBadRequest struct {
}

func (o *DeleteSpaceBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/spaces/{id}][%d] deleteSpaceBadRequest ", 400)
}

func (o *DeleteSpaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
