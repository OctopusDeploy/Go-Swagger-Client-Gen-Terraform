// Code generated by go-swagger; DO NOT EDIT.

package interruptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetInterruptionByIDReader is a Reader for the GetInterruptionByID structure.
type GetInterruptionByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInterruptionByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInterruptionByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInterruptionByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInterruptionByIDOK creates a GetInterruptionByIDOK with default headers values
func NewGetInterruptionByIDOK() *GetInterruptionByIDOK {
	return &GetInterruptionByIDOK{}
}

/*GetInterruptionByIDOK handles this case with default header values.

Load InterruptionResource by ID
*/
type GetInterruptionByIDOK struct {
	Payload *models.InterruptionResource
}

func (o *GetInterruptionByIDOK) Error() string {
	return fmt.Sprintf("[GET /api/interruptions/{id}][%d] getInterruptionByIdOK  %+v", 200, o.Payload)
}

func (o *GetInterruptionByIDOK) GetPayload() *models.InterruptionResource {
	return o.Payload
}

func (o *GetInterruptionByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InterruptionResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInterruptionByIDBadRequest creates a GetInterruptionByIDBadRequest with default headers values
func NewGetInterruptionByIDBadRequest() *GetInterruptionByIDBadRequest {
	return &GetInterruptionByIDBadRequest{}
}

/*GetInterruptionByIDBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetInterruptionByIDBadRequest struct {
}

func (o *GetInterruptionByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/interruptions/{id}][%d] getInterruptionByIdBadRequest ", 400)
}

func (o *GetInterruptionByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}