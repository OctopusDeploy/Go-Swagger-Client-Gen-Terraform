// Code generated by go-swagger; DO NOT EDIT.

package tag_sets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateSortTagSetsSpacesReader is a Reader for the UpdateSortTagSetsSpaces structure.
type UpdateSortTagSetsSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSortTagSetsSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSortTagSetsSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateSortTagSetsSpacesOK creates a UpdateSortTagSetsSpacesOK with default headers values
func NewUpdateSortTagSetsSpacesOK() *UpdateSortTagSetsSpacesOK {
	return &UpdateSortTagSetsSpacesOK{}
}

/*UpdateSortTagSetsSpacesOK handles this case with default header values.

OK - Default
*/
type UpdateSortTagSetsSpacesOK struct {
}

func (o *UpdateSortTagSetsSpacesOK) Error() string {
	return fmt.Sprintf("[PUT /api/{baseSpaceId}/tagsets/sortorder][%d] updateSortTagSetsSpacesOK ", 200)
}

func (o *UpdateSortTagSetsSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
