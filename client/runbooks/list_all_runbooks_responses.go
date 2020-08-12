// Code generated by go-swagger; DO NOT EDIT.

package runbooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// ListAllRunbooksReader is a Reader for the ListAllRunbooks structure.
type ListAllRunbooksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllRunbooksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAllRunbooksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAllRunbooksOK creates a ListAllRunbooksOK with default headers values
func NewListAllRunbooksOK() *ListAllRunbooksOK {
	return &ListAllRunbooksOK{}
}

/*ListAllRunbooksOK handles this case with default header values.

Load all RunbookResource
*/
type ListAllRunbooksOK struct {
	Payload []*models.RunbookResource
}

func (o *ListAllRunbooksOK) Error() string {
	return fmt.Sprintf("[GET /api/runbooks/all][%d] listAllRunbooksOK  %+v", 200, o.Payload)
}

func (o *ListAllRunbooksOK) GetPayload() []*models.RunbookResource {
	return o.Payload
}

func (o *ListAllRunbooksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
