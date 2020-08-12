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

// ListAllTagSetsReader is a Reader for the ListAllTagSets structure.
type ListAllTagSetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllTagSetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAllTagSetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAllTagSetsOK creates a ListAllTagSetsOK with default headers values
func NewListAllTagSetsOK() *ListAllTagSetsOK {
	return &ListAllTagSetsOK{}
}

/*ListAllTagSetsOK handles this case with default header values.

Load all TagSetResource
*/
type ListAllTagSetsOK struct {
	Payload []*models.TagSetResource
}

func (o *ListAllTagSetsOK) Error() string {
	return fmt.Sprintf("[GET /api/tagsets/all][%d] listAllTagSetsOK  %+v", 200, o.Payload)
}

func (o *ListAllTagSetsOK) GetPayload() []*models.TagSetResource {
	return o.Payload
}

func (o *ListAllTagSetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}