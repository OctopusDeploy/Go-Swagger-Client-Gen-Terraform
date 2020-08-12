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

// GetRunbookSnapshotByProjectAndNameSpacesReader is a Reader for the GetRunbookSnapshotByProjectAndNameSpaces structure.
type GetRunbookSnapshotByProjectAndNameSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunbookSnapshotByProjectAndNameSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunbookSnapshotByProjectAndNameSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRunbookSnapshotByProjectAndNameSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRunbookSnapshotByProjectAndNameSpacesOK creates a GetRunbookSnapshotByProjectAndNameSpacesOK with default headers values
func NewGetRunbookSnapshotByProjectAndNameSpacesOK() *GetRunbookSnapshotByProjectAndNameSpacesOK {
	return &GetRunbookSnapshotByProjectAndNameSpacesOK{}
}

/*GetRunbookSnapshotByProjectAndNameSpacesOK handles this case with default header values.

RunbookSnapshotResource resource returned
*/
type GetRunbookSnapshotByProjectAndNameSpacesOK struct {
	Payload *models.RunbookSnapshotResource
}

func (o *GetRunbookSnapshotByProjectAndNameSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/projects/{id}/runbookSnapshots/{name}][%d] getRunbookSnapshotByProjectAndNameSpacesOK  %+v", 200, o.Payload)
}

func (o *GetRunbookSnapshotByProjectAndNameSpacesOK) GetPayload() *models.RunbookSnapshotResource {
	return o.Payload
}

func (o *GetRunbookSnapshotByProjectAndNameSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RunbookSnapshotResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunbookSnapshotByProjectAndNameSpacesBadRequest creates a GetRunbookSnapshotByProjectAndNameSpacesBadRequest with default headers values
func NewGetRunbookSnapshotByProjectAndNameSpacesBadRequest() *GetRunbookSnapshotByProjectAndNameSpacesBadRequest {
	return &GetRunbookSnapshotByProjectAndNameSpacesBadRequest{}
}

/*GetRunbookSnapshotByProjectAndNameSpacesBadRequest handles this case with default header values.

No id parameter was provided.
No name parameter was provided.
*/
type GetRunbookSnapshotByProjectAndNameSpacesBadRequest struct {
}

func (o *GetRunbookSnapshotByProjectAndNameSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/projects/{id}/runbookSnapshots/{name}][%d] getRunbookSnapshotByProjectAndNameSpacesBadRequest ", 400)
}

func (o *GetRunbookSnapshotByProjectAndNameSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
