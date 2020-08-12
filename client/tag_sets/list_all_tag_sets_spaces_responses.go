// Code generated by go-swagger; DO NOT EDIT.

package tag_sets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// ListAllTagSetsSpacesReader is a Reader for the ListAllTagSetsSpaces structure.
type ListAllTagSetsSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllTagSetsSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAllTagSetsSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAllTagSetsSpacesOK creates a ListAllTagSetsSpacesOK with default headers values
func NewListAllTagSetsSpacesOK() *ListAllTagSetsSpacesOK {
	return &ListAllTagSetsSpacesOK{}
}

/*ListAllTagSetsSpacesOK handles this case with default header values.

Load all TagSetResource
*/
type ListAllTagSetsSpacesOK struct {
	Payload []*models.TagSetResource
}

func (o *ListAllTagSetsSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/tagsets/all][%d] listAllTagSetsSpacesOK  %+v", 200, o.Payload)
}

func (o *ListAllTagSetsSpacesOK) GetPayload() []*models.TagSetResource {
	return o.Payload
}

func (o *ListAllTagSetsSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
