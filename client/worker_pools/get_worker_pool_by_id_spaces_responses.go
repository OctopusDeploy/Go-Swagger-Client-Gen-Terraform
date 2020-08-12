// Code generated by go-swagger; DO NOT EDIT.

package worker_pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetWorkerPoolByIDSpacesReader is a Reader for the GetWorkerPoolByIDSpaces structure.
type GetWorkerPoolByIDSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkerPoolByIDSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkerPoolByIDSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkerPoolByIDSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkerPoolByIDSpacesOK creates a GetWorkerPoolByIDSpacesOK with default headers values
func NewGetWorkerPoolByIDSpacesOK() *GetWorkerPoolByIDSpacesOK {
	return &GetWorkerPoolByIDSpacesOK{}
}

/*GetWorkerPoolByIDSpacesOK handles this case with default header values.

Load WorkerPoolResource by ID
*/
type GetWorkerPoolByIDSpacesOK struct {
	Payload *models.WorkerPoolResource
}

func (o *GetWorkerPoolByIDSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/workerpools/{id}][%d] getWorkerPoolByIdSpacesOK  %+v", 200, o.Payload)
}

func (o *GetWorkerPoolByIDSpacesOK) GetPayload() *models.WorkerPoolResource {
	return o.Payload
}

func (o *GetWorkerPoolByIDSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkerPoolResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkerPoolByIDSpacesBadRequest creates a GetWorkerPoolByIDSpacesBadRequest with default headers values
func NewGetWorkerPoolByIDSpacesBadRequest() *GetWorkerPoolByIDSpacesBadRequest {
	return &GetWorkerPoolByIDSpacesBadRequest{}
}

/*GetWorkerPoolByIDSpacesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetWorkerPoolByIDSpacesBadRequest struct {
}

func (o *GetWorkerPoolByIDSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/workerpools/{id}][%d] getWorkerPoolByIdSpacesBadRequest ", 400)
}

func (o *GetWorkerPoolByIDSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}