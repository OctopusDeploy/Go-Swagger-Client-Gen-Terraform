// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetEventByIDSpacesReader is a Reader for the GetEventByIDSpaces structure.
type GetEventByIDSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventByIDSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEventByIDSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEventByIDSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEventByIDSpacesOK creates a GetEventByIDSpacesOK with default headers values
func NewGetEventByIDSpacesOK() *GetEventByIDSpacesOK {
	return &GetEventByIDSpacesOK{}
}

/*GetEventByIDSpacesOK handles this case with default header values.

Load EventResource by ID
*/
type GetEventByIDSpacesOK struct {
	Payload *models.EventResource
}

func (o *GetEventByIDSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/events/{id}][%d] getEventByIdSpacesOK  %+v", 200, o.Payload)
}

func (o *GetEventByIDSpacesOK) GetPayload() *models.EventResource {
	return o.Payload
}

func (o *GetEventByIDSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EventResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventByIDSpacesBadRequest creates a GetEventByIDSpacesBadRequest with default headers values
func NewGetEventByIDSpacesBadRequest() *GetEventByIDSpacesBadRequest {
	return &GetEventByIDSpacesBadRequest{}
}

/*GetEventByIDSpacesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetEventByIDSpacesBadRequest struct {
}

func (o *GetEventByIDSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/events/{id}][%d] getEventByIdSpacesBadRequest ", 400)
}

func (o *GetEventByIDSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
