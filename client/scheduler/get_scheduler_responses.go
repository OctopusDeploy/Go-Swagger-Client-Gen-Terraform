// Code generated by go-swagger; DO NOT EDIT.

package scheduler

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetSchedulerReader is a Reader for the GetScheduler structure.
type GetSchedulerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSchedulerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSchedulerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSchedulerOK creates a GetSchedulerOK with default headers values
func NewGetSchedulerOK() *GetSchedulerOK {
	return &GetSchedulerOK{}
}

/*GetSchedulerOK handles this case with default header values.

SchedulerStatusResource resource returned
*/
type GetSchedulerOK struct {
	Payload *models.SchedulerStatusResource
}

func (o *GetSchedulerOK) Error() string {
	return fmt.Sprintf("[GET /api/scheduler][%d] getSchedulerOK  %+v", 200, o.Payload)
}

func (o *GetSchedulerOK) GetPayload() *models.SchedulerStatusResource {
	return o.Payload
}

func (o *GetSchedulerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SchedulerStatusResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
