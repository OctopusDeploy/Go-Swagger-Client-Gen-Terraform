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

// GetWorkerByIDReader is a Reader for the GetWorkerByID structure.
type GetWorkerByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkerByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkerByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkerByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkerByIDOK creates a GetWorkerByIDOK with default headers values
func NewGetWorkerByIDOK() *GetWorkerByIDOK {
	return &GetWorkerByIDOK{}
}

/*GetWorkerByIDOK handles this case with default header values.

Load WorkerResource by ID
*/
type GetWorkerByIDOK struct {
	Payload *models.WorkerResource
}

func (o *GetWorkerByIDOK) Error() string {
	return fmt.Sprintf("[GET /api/workers/{id}][%d] getWorkerByIdOK  %+v", 200, o.Payload)
}

func (o *GetWorkerByIDOK) GetPayload() *models.WorkerResource {
	return o.Payload
}

func (o *GetWorkerByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkerResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkerByIDBadRequest creates a GetWorkerByIDBadRequest with default headers values
func NewGetWorkerByIDBadRequest() *GetWorkerByIDBadRequest {
	return &GetWorkerByIDBadRequest{}
}

/*GetWorkerByIDBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetWorkerByIDBadRequest struct {
}

func (o *GetWorkerByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/workers/{id}][%d] getWorkerByIdBadRequest ", 400)
}

func (o *GetWorkerByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
