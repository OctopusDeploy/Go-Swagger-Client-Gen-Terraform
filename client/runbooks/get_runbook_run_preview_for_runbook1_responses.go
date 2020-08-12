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

// GetRunbookRunPreviewForRunbook1Reader is a Reader for the GetRunbookRunPreviewForRunbook1 structure.
type GetRunbookRunPreviewForRunbook1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunbookRunPreviewForRunbook1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunbookRunPreviewForRunbook1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRunbookRunPreviewForRunbook1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRunbookRunPreviewForRunbook1OK creates a GetRunbookRunPreviewForRunbook1OK with default headers values
func NewGetRunbookRunPreviewForRunbook1OK() *GetRunbookRunPreviewForRunbook1OK {
	return &GetRunbookRunPreviewForRunbook1OK{}
}

/*GetRunbookRunPreviewForRunbook1OK handles this case with default header values.

RunbookRunPreviewResource resource returned
*/
type GetRunbookRunPreviewForRunbook1OK struct {
	Payload *models.RunbookRunPreviewResource
}

func (o *GetRunbookRunPreviewForRunbook1OK) Error() string {
	return fmt.Sprintf("[GET /api/runbooks/{id}/runbookRuns/preview/{environment}/{tenant}][%d] getRunbookRunPreviewForRunbook1OK  %+v", 200, o.Payload)
}

func (o *GetRunbookRunPreviewForRunbook1OK) GetPayload() *models.RunbookRunPreviewResource {
	return o.Payload
}

func (o *GetRunbookRunPreviewForRunbook1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RunbookRunPreviewResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunbookRunPreviewForRunbook1BadRequest creates a GetRunbookRunPreviewForRunbook1BadRequest with default headers values
func NewGetRunbookRunPreviewForRunbook1BadRequest() *GetRunbookRunPreviewForRunbook1BadRequest {
	return &GetRunbookRunPreviewForRunbook1BadRequest{}
}

/*GetRunbookRunPreviewForRunbook1BadRequest handles this case with default header values.

No environment parameter was provided.
No id parameter was provided.
*/
type GetRunbookRunPreviewForRunbook1BadRequest struct {
}

func (o *GetRunbookRunPreviewForRunbook1BadRequest) Error() string {
	return fmt.Sprintf("[GET /api/runbooks/{id}/runbookRuns/preview/{environment}/{tenant}][%d] getRunbookRunPreviewForRunbook1BadRequest ", 400)
}

func (o *GetRunbookRunPreviewForRunbook1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}