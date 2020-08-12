// Code generated by go-swagger; DO NOT EDIT.

package runbook_processes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// IndexRunbookProcessesReader is a Reader for the IndexRunbookProcesses structure.
type IndexRunbookProcessesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexRunbookProcessesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIndexRunbookProcessesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIndexRunbookProcessesOK creates a IndexRunbookProcessesOK with default headers values
func NewIndexRunbookProcessesOK() *IndexRunbookProcessesOK {
	return &IndexRunbookProcessesOK{}
}

/*IndexRunbookProcessesOK handles this case with default header values.

ResourceCollection_of_RunbookProcessResource resource returned
*/
type IndexRunbookProcessesOK struct {
	Payload *models.ResourceCollectionOfRunbookProcessResource
}

func (o *IndexRunbookProcessesOK) Error() string {
	return fmt.Sprintf("[GET /api/runbookProcesses][%d] indexRunbookProcessesOK  %+v", 200, o.Payload)
}

func (o *IndexRunbookProcessesOK) GetPayload() *models.ResourceCollectionOfRunbookProcessResource {
	return o.Payload
}

func (o *IndexRunbookProcessesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionOfRunbookProcessResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}