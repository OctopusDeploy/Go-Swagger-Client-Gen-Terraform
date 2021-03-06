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

// GetRunbookRunPreviewForRunbookSnapshotSpacesReader is a Reader for the GetRunbookRunPreviewForRunbookSnapshotSpaces structure.
type GetRunbookRunPreviewForRunbookSnapshotSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunbookRunPreviewForRunbookSnapshotSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunbookRunPreviewForRunbookSnapshotSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRunbookRunPreviewForRunbookSnapshotSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRunbookRunPreviewForRunbookSnapshotSpacesOK creates a GetRunbookRunPreviewForRunbookSnapshotSpacesOK with default headers values
func NewGetRunbookRunPreviewForRunbookSnapshotSpacesOK() *GetRunbookRunPreviewForRunbookSnapshotSpacesOK {
	return &GetRunbookRunPreviewForRunbookSnapshotSpacesOK{}
}

/*GetRunbookRunPreviewForRunbookSnapshotSpacesOK handles this case with default header values.

RunbookRunPreviewResource resource returned
*/
type GetRunbookRunPreviewForRunbookSnapshotSpacesOK struct {
	Payload *models.RunbookRunPreviewResource
}

func (o *GetRunbookRunPreviewForRunbookSnapshotSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/runbookSnapshots/{id}/runbookRuns/preview/{environment}][%d] getRunbookRunPreviewForRunbookSnapshotSpacesOK  %+v", 200, o.Payload)
}

func (o *GetRunbookRunPreviewForRunbookSnapshotSpacesOK) GetPayload() *models.RunbookRunPreviewResource {
	return o.Payload
}

func (o *GetRunbookRunPreviewForRunbookSnapshotSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RunbookRunPreviewResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunbookRunPreviewForRunbookSnapshotSpacesBadRequest creates a GetRunbookRunPreviewForRunbookSnapshotSpacesBadRequest with default headers values
func NewGetRunbookRunPreviewForRunbookSnapshotSpacesBadRequest() *GetRunbookRunPreviewForRunbookSnapshotSpacesBadRequest {
	return &GetRunbookRunPreviewForRunbookSnapshotSpacesBadRequest{}
}

/*GetRunbookRunPreviewForRunbookSnapshotSpacesBadRequest handles this case with default header values.

No environment parameter was provided.
No id parameter was provided.
*/
type GetRunbookRunPreviewForRunbookSnapshotSpacesBadRequest struct {
}

func (o *GetRunbookRunPreviewForRunbookSnapshotSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/runbookSnapshots/{id}/runbookRuns/preview/{environment}][%d] getRunbookRunPreviewForRunbookSnapshotSpacesBadRequest ", 400)
}

func (o *GetRunbookRunPreviewForRunbookSnapshotSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
