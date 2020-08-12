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

// IndexRunbooksReader is a Reader for the IndexRunbooks structure.
type IndexRunbooksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexRunbooksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIndexRunbooksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIndexRunbooksOK creates a IndexRunbooksOK with default headers values
func NewIndexRunbooksOK() *IndexRunbooksOK {
	return &IndexRunbooksOK{}
}

/*IndexRunbooksOK handles this case with default header values.

ResourceCollection_of_RunbookResource resource returned
*/
type IndexRunbooksOK struct {
	Payload *models.ResourceCollectionOfRunbookResource
}

func (o *IndexRunbooksOK) Error() string {
	return fmt.Sprintf("[GET /api/runbooks][%d] indexRunbooksOK  %+v", 200, o.Payload)
}

func (o *IndexRunbooksOK) GetPayload() *models.ResourceCollectionOfRunbookResource {
	return o.Payload
}

func (o *IndexRunbooksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionOfRunbookResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}