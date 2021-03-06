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

// GetInterruptionByIDSpacesReader is a Reader for the GetInterruptionByIDSpaces structure.
type GetInterruptionByIDSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInterruptionByIDSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInterruptionByIDSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInterruptionByIDSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInterruptionByIDSpacesOK creates a GetInterruptionByIDSpacesOK with default headers values
func NewGetInterruptionByIDSpacesOK() *GetInterruptionByIDSpacesOK {
	return &GetInterruptionByIDSpacesOK{}
}

/*GetInterruptionByIDSpacesOK handles this case with default header values.

Load InterruptionResource by ID
*/
type GetInterruptionByIDSpacesOK struct {
	Payload *models.InterruptionResource
}

func (o *GetInterruptionByIDSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/interruptions/{id}][%d] getInterruptionByIdSpacesOK  %+v", 200, o.Payload)
}

func (o *GetInterruptionByIDSpacesOK) GetPayload() *models.InterruptionResource {
	return o.Payload
}

func (o *GetInterruptionByIDSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InterruptionResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInterruptionByIDSpacesBadRequest creates a GetInterruptionByIDSpacesBadRequest with default headers values
func NewGetInterruptionByIDSpacesBadRequest() *GetInterruptionByIDSpacesBadRequest {
	return &GetInterruptionByIDSpacesBadRequest{}
}

/*GetInterruptionByIDSpacesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetInterruptionByIDSpacesBadRequest struct {
}

func (o *GetInterruptionByIDSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/interruptions/{id}][%d] getInterruptionByIdSpacesBadRequest ", 400)
}

func (o *GetInterruptionByIDSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
