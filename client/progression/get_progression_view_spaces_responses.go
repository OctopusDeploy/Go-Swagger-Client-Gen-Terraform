// Code generated by go-swagger; DO NOT EDIT.

package progression

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetProgressionViewSpacesReader is a Reader for the GetProgressionViewSpaces structure.
type GetProgressionViewSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProgressionViewSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProgressionViewSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProgressionViewSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProgressionViewSpacesOK creates a GetProgressionViewSpacesOK with default headers values
func NewGetProgressionViewSpacesOK() *GetProgressionViewSpacesOK {
	return &GetProgressionViewSpacesOK{}
}

/*GetProgressionViewSpacesOK handles this case with default header values.

ProgressionResource resource returned
*/
type GetProgressionViewSpacesOK struct {
	Payload *models.ProgressionResource
}

func (o *GetProgressionViewSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/progression/{id}][%d] getProgressionViewSpacesOK  %+v", 200, o.Payload)
}

func (o *GetProgressionViewSpacesOK) GetPayload() *models.ProgressionResource {
	return o.Payload
}

func (o *GetProgressionViewSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProgressionResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProgressionViewSpacesBadRequest creates a GetProgressionViewSpacesBadRequest with default headers values
func NewGetProgressionViewSpacesBadRequest() *GetProgressionViewSpacesBadRequest {
	return &GetProgressionViewSpacesBadRequest{}
}

/*GetProgressionViewSpacesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetProgressionViewSpacesBadRequest struct {
}

func (o *GetProgressionViewSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/progression/{id}][%d] getProgressionViewSpacesBadRequest ", 400)
}

func (o *GetProgressionViewSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
