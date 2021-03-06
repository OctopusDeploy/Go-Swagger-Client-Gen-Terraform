// Code generated by go-swagger; DO NOT EDIT.

package runbook_processes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetRunbookSnapshotTemplateForRunbookStepSpacesReader is a Reader for the GetRunbookSnapshotTemplateForRunbookStepSpaces structure.
type GetRunbookSnapshotTemplateForRunbookStepSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunbookSnapshotTemplateForRunbookStepSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunbookSnapshotTemplateForRunbookStepSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRunbookSnapshotTemplateForRunbookStepSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRunbookSnapshotTemplateForRunbookStepSpacesOK creates a GetRunbookSnapshotTemplateForRunbookStepSpacesOK with default headers values
func NewGetRunbookSnapshotTemplateForRunbookStepSpacesOK() *GetRunbookSnapshotTemplateForRunbookStepSpacesOK {
	return &GetRunbookSnapshotTemplateForRunbookStepSpacesOK{}
}

/*GetRunbookSnapshotTemplateForRunbookStepSpacesOK handles this case with default header values.

RunbookSnapshotTemplateResource resource returned
*/
type GetRunbookSnapshotTemplateForRunbookStepSpacesOK struct {
	Payload *models.RunbookSnapshotTemplateResource
}

func (o *GetRunbookSnapshotTemplateForRunbookStepSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/runbookProcesses/{id}/runbookSnapshotTemplate][%d] getRunbookSnapshotTemplateForRunbookStepSpacesOK  %+v", 200, o.Payload)
}

func (o *GetRunbookSnapshotTemplateForRunbookStepSpacesOK) GetPayload() *models.RunbookSnapshotTemplateResource {
	return o.Payload
}

func (o *GetRunbookSnapshotTemplateForRunbookStepSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RunbookSnapshotTemplateResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunbookSnapshotTemplateForRunbookStepSpacesBadRequest creates a GetRunbookSnapshotTemplateForRunbookStepSpacesBadRequest with default headers values
func NewGetRunbookSnapshotTemplateForRunbookStepSpacesBadRequest() *GetRunbookSnapshotTemplateForRunbookStepSpacesBadRequest {
	return &GetRunbookSnapshotTemplateForRunbookStepSpacesBadRequest{}
}

/*GetRunbookSnapshotTemplateForRunbookStepSpacesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetRunbookSnapshotTemplateForRunbookStepSpacesBadRequest struct {
}

func (o *GetRunbookSnapshotTemplateForRunbookStepSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/runbookProcesses/{id}/runbookSnapshotTemplate][%d] getRunbookSnapshotTemplateForRunbookStepSpacesBadRequest ", 400)
}

func (o *GetRunbookSnapshotTemplateForRunbookStepSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
