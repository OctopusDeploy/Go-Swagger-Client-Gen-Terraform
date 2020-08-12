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

// GetWorkerConnectionStatusReader is a Reader for the GetWorkerConnectionStatus structure.
type GetWorkerConnectionStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkerConnectionStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkerConnectionStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkerConnectionStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkerConnectionStatusOK creates a GetWorkerConnectionStatusOK with default headers values
func NewGetWorkerConnectionStatusOK() *GetWorkerConnectionStatusOK {
	return &GetWorkerConnectionStatusOK{}
}

/*GetWorkerConnectionStatusOK handles this case with default header values.

MachineConnectionStatus resource returned
*/
type GetWorkerConnectionStatusOK struct {
	Payload *models.MachineConnectionStatus
}

func (o *GetWorkerConnectionStatusOK) Error() string {
	return fmt.Sprintf("[GET /api/workers/{id}/connection][%d] getWorkerConnectionStatusOK  %+v", 200, o.Payload)
}

func (o *GetWorkerConnectionStatusOK) GetPayload() *models.MachineConnectionStatus {
	return o.Payload
}

func (o *GetWorkerConnectionStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MachineConnectionStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkerConnectionStatusBadRequest creates a GetWorkerConnectionStatusBadRequest with default headers values
func NewGetWorkerConnectionStatusBadRequest() *GetWorkerConnectionStatusBadRequest {
	return &GetWorkerConnectionStatusBadRequest{}
}

/*GetWorkerConnectionStatusBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetWorkerConnectionStatusBadRequest struct {
}

func (o *GetWorkerConnectionStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/workers/{id}/connection][%d] getWorkerConnectionStatusBadRequest ", 400)
}

func (o *GetWorkerConnectionStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}