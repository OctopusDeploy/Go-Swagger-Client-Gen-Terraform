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

// GetRunbooksProgressionViewReader is a Reader for the GetRunbooksProgressionView structure.
type GetRunbooksProgressionViewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunbooksProgressionViewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunbooksProgressionViewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRunbooksProgressionViewBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRunbooksProgressionViewOK creates a GetRunbooksProgressionViewOK with default headers values
func NewGetRunbooksProgressionViewOK() *GetRunbooksProgressionViewOK {
	return &GetRunbooksProgressionViewOK{}
}

/*GetRunbooksProgressionViewOK handles this case with default header values.

RunbooksProgressionResource resource returned
*/
type GetRunbooksProgressionViewOK struct {
	Payload *models.RunbooksProgressionResource
}

func (o *GetRunbooksProgressionViewOK) Error() string {
	return fmt.Sprintf("[GET /api/progression/runbooks/{id}][%d] getRunbooksProgressionViewOK  %+v", 200, o.Payload)
}

func (o *GetRunbooksProgressionViewOK) GetPayload() *models.RunbooksProgressionResource {
	return o.Payload
}

func (o *GetRunbooksProgressionViewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RunbooksProgressionResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunbooksProgressionViewBadRequest creates a GetRunbooksProgressionViewBadRequest with default headers values
func NewGetRunbooksProgressionViewBadRequest() *GetRunbooksProgressionViewBadRequest {
	return &GetRunbooksProgressionViewBadRequest{}
}

/*GetRunbooksProgressionViewBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetRunbooksProgressionViewBadRequest struct {
}

func (o *GetRunbooksProgressionViewBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/progression/runbooks/{id}][%d] getRunbooksProgressionViewBadRequest ", 400)
}

func (o *GetRunbooksProgressionViewBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
