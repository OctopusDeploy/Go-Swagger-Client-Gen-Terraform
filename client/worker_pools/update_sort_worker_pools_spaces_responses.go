// Code generated by go-swagger; DO NOT EDIT.

package worker_pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateSortWorkerPoolsSpacesReader is a Reader for the UpdateSortWorkerPoolsSpaces structure.
type UpdateSortWorkerPoolsSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSortWorkerPoolsSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSortWorkerPoolsSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateSortWorkerPoolsSpacesOK creates a UpdateSortWorkerPoolsSpacesOK with default headers values
func NewUpdateSortWorkerPoolsSpacesOK() *UpdateSortWorkerPoolsSpacesOK {
	return &UpdateSortWorkerPoolsSpacesOK{}
}

/*UpdateSortWorkerPoolsSpacesOK handles this case with default header values.

OK - Default
*/
type UpdateSortWorkerPoolsSpacesOK struct {
}

func (o *UpdateSortWorkerPoolsSpacesOK) Error() string {
	return fmt.Sprintf("[PUT /api/{baseSpaceId}/workerpools/sortorder][%d] updateSortWorkerPoolsSpacesOK ", 200)
}

func (o *UpdateSortWorkerPoolsSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
