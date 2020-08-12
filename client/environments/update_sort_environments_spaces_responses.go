// Code generated by go-swagger; DO NOT EDIT.

package environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateSortEnvironmentsSpacesReader is a Reader for the UpdateSortEnvironmentsSpaces structure.
type UpdateSortEnvironmentsSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSortEnvironmentsSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSortEnvironmentsSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateSortEnvironmentsSpacesOK creates a UpdateSortEnvironmentsSpacesOK with default headers values
func NewUpdateSortEnvironmentsSpacesOK() *UpdateSortEnvironmentsSpacesOK {
	return &UpdateSortEnvironmentsSpacesOK{}
}

/*UpdateSortEnvironmentsSpacesOK handles this case with default header values.

OK - Default
*/
type UpdateSortEnvironmentsSpacesOK struct {
}

func (o *UpdateSortEnvironmentsSpacesOK) Error() string {
	return fmt.Sprintf("[PUT /api/{baseSpaceId}/environments/sortorder][%d] updateSortEnvironmentsSpacesOK ", 200)
}

func (o *UpdateSortEnvironmentsSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}