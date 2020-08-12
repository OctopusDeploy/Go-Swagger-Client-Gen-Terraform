// Code generated by go-swagger; DO NOT EDIT.

package api_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAPIKeyReader is a Reader for the DeleteAPIKey structure.
type DeleteAPIKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAPIKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAPIKeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAPIKeyOK creates a DeleteAPIKeyOK with default headers values
func NewDeleteAPIKeyOK() *DeleteAPIKeyOK {
	return &DeleteAPIKeyOK{}
}

/*DeleteAPIKeyOK handles this case with default header values.

OK
*/
type DeleteAPIKeyOK struct {
}

func (o *DeleteAPIKeyOK) Error() string {
	return fmt.Sprintf("[DELETE /api/users/{userId}/apikeys/{id}][%d] deleteApiKeyOK ", 200)
}

func (o *DeleteAPIKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAPIKeyBadRequest creates a DeleteAPIKeyBadRequest with default headers values
func NewDeleteAPIKeyBadRequest() *DeleteAPIKeyBadRequest {
	return &DeleteAPIKeyBadRequest{}
}

/*DeleteAPIKeyBadRequest handles this case with default header values.

BadRequest
No id parameter was provided.
No userId parameter was provided.
This ApiKey does not belong to the specified user.
This resource is read-only and cannot be deleted.
*/
type DeleteAPIKeyBadRequest struct {
}

func (o *DeleteAPIKeyBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/users/{userId}/apikeys/{id}][%d] deleteApiKeyBadRequest ", 400)
}

func (o *DeleteAPIKeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}