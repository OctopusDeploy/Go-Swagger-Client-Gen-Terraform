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

// GetDiscoverWorkerReader is a Reader for the GetDiscoverWorker structure.
type GetDiscoverWorkerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDiscoverWorkerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDiscoverWorkerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDiscoverWorkerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDiscoverWorkerOK creates a GetDiscoverWorkerOK with default headers values
func NewGetDiscoverWorkerOK() *GetDiscoverWorkerOK {
	return &GetDiscoverWorkerOK{}
}

/*GetDiscoverWorkerOK handles this case with default header values.

MachineResource resource returned
*/
type GetDiscoverWorkerOK struct {
	Payload *models.MachineResource
}

func (o *GetDiscoverWorkerOK) Error() string {
	return fmt.Sprintf("[GET /api/workers/discover][%d] getDiscoverWorkerOK  %+v", 200, o.Payload)
}

func (o *GetDiscoverWorkerOK) GetPayload() *models.MachineResource {
	return o.Payload
}

func (o *GetDiscoverWorkerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MachineResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDiscoverWorkerBadRequest creates a GetDiscoverWorkerBadRequest with default headers values
func NewGetDiscoverWorkerBadRequest() *GetDiscoverWorkerBadRequest {
	return &GetDiscoverWorkerBadRequest{}
}

/*GetDiscoverWorkerBadRequest handles this case with default header values.

The hostname of the machine to discover must be supplied as URI parameter named 'host'.
There was a controlled failure.
*/
type GetDiscoverWorkerBadRequest struct {
}

func (o *GetDiscoverWorkerBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/workers/discover][%d] getDiscoverWorkerBadRequest ", 400)
}

func (o *GetDiscoverWorkerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
