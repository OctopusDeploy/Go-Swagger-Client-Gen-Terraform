// Code generated by go-swagger; DO NOT EDIT.

package environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateSortEnvironmentsReader is a Reader for the UpdateSortEnvironments structure.
type UpdateSortEnvironmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSortEnvironmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSortEnvironmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateSortEnvironmentsOK creates a UpdateSortEnvironmentsOK with default headers values
func NewUpdateSortEnvironmentsOK() *UpdateSortEnvironmentsOK {
	return &UpdateSortEnvironmentsOK{}
}

/*UpdateSortEnvironmentsOK handles this case with default header values.

OK - Default
*/
type UpdateSortEnvironmentsOK struct {
}

func (o *UpdateSortEnvironmentsOK) Error() string {
	return fmt.Sprintf("[PUT /api/environments/sortorder][%d] updateSortEnvironmentsOK ", 200)
}

func (o *UpdateSortEnvironmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
