// Code generated by go-swagger; DO NOT EDIT.

package runbook_snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetRunbookRunPreviewForRunbookSnapshot1Reader is a Reader for the GetRunbookRunPreviewForRunbookSnapshot1 structure.
type GetRunbookRunPreviewForRunbookSnapshot1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunbookRunPreviewForRunbookSnapshot1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunbookRunPreviewForRunbookSnapshot1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRunbookRunPreviewForRunbookSnapshot1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRunbookRunPreviewForRunbookSnapshot1OK creates a GetRunbookRunPreviewForRunbookSnapshot1OK with default headers values
func NewGetRunbookRunPreviewForRunbookSnapshot1OK() *GetRunbookRunPreviewForRunbookSnapshot1OK {
	return &GetRunbookRunPreviewForRunbookSnapshot1OK{}
}

/*GetRunbookRunPreviewForRunbookSnapshot1OK handles this case with default header values.

RunbookRunPreviewResource resource returned
*/
type GetRunbookRunPreviewForRunbookSnapshot1OK struct {
	Payload *models.RunbookRunPreviewResource
}

func (o *GetRunbookRunPreviewForRunbookSnapshot1OK) Error() string {
	return fmt.Sprintf("[GET /api/runbookSnapshots/{id}/runbookRuns/preview/{environment}/{tenant}][%d] getRunbookRunPreviewForRunbookSnapshot1OK  %+v", 200, o.Payload)
}

func (o *GetRunbookRunPreviewForRunbookSnapshot1OK) GetPayload() *models.RunbookRunPreviewResource {
	return o.Payload
}

func (o *GetRunbookRunPreviewForRunbookSnapshot1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RunbookRunPreviewResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunbookRunPreviewForRunbookSnapshot1BadRequest creates a GetRunbookRunPreviewForRunbookSnapshot1BadRequest with default headers values
func NewGetRunbookRunPreviewForRunbookSnapshot1BadRequest() *GetRunbookRunPreviewForRunbookSnapshot1BadRequest {
	return &GetRunbookRunPreviewForRunbookSnapshot1BadRequest{}
}

/*GetRunbookRunPreviewForRunbookSnapshot1BadRequest handles this case with default header values.

No environment parameter was provided.
No id parameter was provided.
*/
type GetRunbookRunPreviewForRunbookSnapshot1BadRequest struct {
}

func (o *GetRunbookRunPreviewForRunbookSnapshot1BadRequest) Error() string {
	return fmt.Sprintf("[GET /api/runbookSnapshots/{id}/runbookRuns/preview/{environment}/{tenant}][%d] getRunbookRunPreviewForRunbookSnapshot1BadRequest ", 400)
}

func (o *GetRunbookRunPreviewForRunbookSnapshot1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
