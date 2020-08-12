// Code generated by go-swagger; DO NOT EDIT.

package runbooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetRunbookRunPreviewForRunbookReader is a Reader for the GetRunbookRunPreviewForRunbook structure.
type GetRunbookRunPreviewForRunbookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunbookRunPreviewForRunbookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunbookRunPreviewForRunbookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRunbookRunPreviewForRunbookBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRunbookRunPreviewForRunbookOK creates a GetRunbookRunPreviewForRunbookOK with default headers values
func NewGetRunbookRunPreviewForRunbookOK() *GetRunbookRunPreviewForRunbookOK {
	return &GetRunbookRunPreviewForRunbookOK{}
}

/*GetRunbookRunPreviewForRunbookOK handles this case with default header values.

RunbookRunPreviewResource resource returned
*/
type GetRunbookRunPreviewForRunbookOK struct {
	Payload *models.RunbookRunPreviewResource
}

func (o *GetRunbookRunPreviewForRunbookOK) Error() string {
	return fmt.Sprintf("[GET /api/runbooks/{id}/runbookRuns/preview/{environment}][%d] getRunbookRunPreviewForRunbookOK  %+v", 200, o.Payload)
}

func (o *GetRunbookRunPreviewForRunbookOK) GetPayload() *models.RunbookRunPreviewResource {
	return o.Payload
}

func (o *GetRunbookRunPreviewForRunbookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RunbookRunPreviewResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunbookRunPreviewForRunbookBadRequest creates a GetRunbookRunPreviewForRunbookBadRequest with default headers values
func NewGetRunbookRunPreviewForRunbookBadRequest() *GetRunbookRunPreviewForRunbookBadRequest {
	return &GetRunbookRunPreviewForRunbookBadRequest{}
}

/*GetRunbookRunPreviewForRunbookBadRequest handles this case with default header values.

No environment parameter was provided.
No id parameter was provided.
*/
type GetRunbookRunPreviewForRunbookBadRequest struct {
}

func (o *GetRunbookRunPreviewForRunbookBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/runbooks/{id}/runbookRuns/preview/{environment}][%d] getRunbookRunPreviewForRunbookBadRequest ", 400)
}

func (o *GetRunbookRunPreviewForRunbookBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}