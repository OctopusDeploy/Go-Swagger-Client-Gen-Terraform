// Code generated by go-swagger; DO NOT EDIT.

package tag_sets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteTagSetReader is a Reader for the DeleteTagSet structure.
type DeleteTagSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTagSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTagSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTagSetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTagSetOK creates a DeleteTagSetOK with default headers values
func NewDeleteTagSetOK() *DeleteTagSetOK {
	return &DeleteTagSetOK{}
}

/*DeleteTagSetOK handles this case with default header values.

OK
*/
type DeleteTagSetOK struct {
}

func (o *DeleteTagSetOK) Error() string {
	return fmt.Sprintf("[DELETE /api/tagsets/{id}][%d] deleteTagSetOK ", 200)
}

func (o *DeleteTagSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTagSetBadRequest creates a DeleteTagSetBadRequest with default headers values
func NewDeleteTagSetBadRequest() *DeleteTagSetBadRequest {
	return &DeleteTagSetBadRequest{}
}

/*DeleteTagSetBadRequest handles this case with default header values.

BadRequest
No id parameter was provided.
This resource is read-only and cannot be deleted.
*/
type DeleteTagSetBadRequest struct {
}

func (o *DeleteTagSetBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/tagsets/{id}][%d] deleteTagSetBadRequest ", 400)
}

func (o *DeleteTagSetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
