// Code generated by go-swagger; DO NOT EDIT.

package channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteChannelSpacesReader is a Reader for the DeleteChannelSpaces structure.
type DeleteChannelSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteChannelSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteChannelSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteChannelSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteChannelSpacesOK creates a DeleteChannelSpacesOK with default headers values
func NewDeleteChannelSpacesOK() *DeleteChannelSpacesOK {
	return &DeleteChannelSpacesOK{}
}

/*DeleteChannelSpacesOK handles this case with default header values.

OK
*/
type DeleteChannelSpacesOK struct {
}

func (o *DeleteChannelSpacesOK) Error() string {
	return fmt.Sprintf("[DELETE /api/{baseSpaceId}/channels/{id}][%d] deleteChannelSpacesOK ", 200)
}

func (o *DeleteChannelSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteChannelSpacesBadRequest creates a DeleteChannelSpacesBadRequest with default headers values
func NewDeleteChannelSpacesBadRequest() *DeleteChannelSpacesBadRequest {
	return &DeleteChannelSpacesBadRequest{}
}

/*DeleteChannelSpacesBadRequest handles this case with default header values.

BadRequest
No id parameter was provided.
The channel you are about to remove is used by '{action.Name}' step.
This resource is read-only and cannot be deleted.
You cannot delete the default channel.
You cannot delete the only remaining channel
*/
type DeleteChannelSpacesBadRequest struct {
}

func (o *DeleteChannelSpacesBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/{baseSpaceId}/channels/{id}][%d] deleteChannelSpacesBadRequest ", 400)
}

func (o *DeleteChannelSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
