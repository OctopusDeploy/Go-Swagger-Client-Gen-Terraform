// Code generated by go-swagger; DO NOT EDIT.

package workers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// ListAllWorkersReader is a Reader for the ListAllWorkers structure.
type ListAllWorkersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllWorkersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAllWorkersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAllWorkersOK creates a ListAllWorkersOK with default headers values
func NewListAllWorkersOK() *ListAllWorkersOK {
	return &ListAllWorkersOK{}
}

/*ListAllWorkersOK handles this case with default header values.

Load all WorkerResource
*/
type ListAllWorkersOK struct {
	Payload []*models.WorkerResource
}

func (o *ListAllWorkersOK) Error() string {
	return fmt.Sprintf("[GET /api/workers/all][%d] listAllWorkersOK  %+v", 200, o.Payload)
}

func (o *ListAllWorkersOK) GetPayload() []*models.WorkerResource {
	return o.Payload
}

func (o *ListAllWorkersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}