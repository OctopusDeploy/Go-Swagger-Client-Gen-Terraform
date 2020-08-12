// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetUserGetSpacesReader is a Reader for the GetUserGetSpaces structure.
type GetUserGetSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserGetSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewGetUserGetSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserGetSpacesBadRequest creates a GetUserGetSpacesBadRequest with default headers values
func NewGetUserGetSpacesBadRequest() *GetUserGetSpacesBadRequest {
	return &GetUserGetSpacesBadRequest{}
}

/*GetUserGetSpacesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetUserGetSpacesBadRequest struct {
}

func (o *GetUserGetSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/users/{id}/spaces][%d] getUserGetSpacesBadRequest ", 400)
}

func (o *GetUserGetSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}