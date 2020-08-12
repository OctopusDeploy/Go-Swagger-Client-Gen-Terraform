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

// IndexRunbookRunbookSnapshotsSpacesReader is a Reader for the IndexRunbookRunbookSnapshotsSpaces structure.
type IndexRunbookRunbookSnapshotsSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexRunbookRunbookSnapshotsSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIndexRunbookRunbookSnapshotsSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIndexRunbookRunbookSnapshotsSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIndexRunbookRunbookSnapshotsSpacesOK creates a IndexRunbookRunbookSnapshotsSpacesOK with default headers values
func NewIndexRunbookRunbookSnapshotsSpacesOK() *IndexRunbookRunbookSnapshotsSpacesOK {
	return &IndexRunbookRunbookSnapshotsSpacesOK{}
}

/*IndexRunbookRunbookSnapshotsSpacesOK handles this case with default header values.

ResourceCollection_of_RunbookSnapshotResource resource returned
*/
type IndexRunbookRunbookSnapshotsSpacesOK struct {
	Payload *models.ResourceCollectionOfRunbookSnapshotResource
}

func (o *IndexRunbookRunbookSnapshotsSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/runbooks/{id}/runbookSnapshots][%d] indexRunbookRunbookSnapshotsSpacesOK  %+v", 200, o.Payload)
}

func (o *IndexRunbookRunbookSnapshotsSpacesOK) GetPayload() *models.ResourceCollectionOfRunbookSnapshotResource {
	return o.Payload
}

func (o *IndexRunbookRunbookSnapshotsSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionOfRunbookSnapshotResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIndexRunbookRunbookSnapshotsSpacesBadRequest creates a IndexRunbookRunbookSnapshotsSpacesBadRequest with default headers values
func NewIndexRunbookRunbookSnapshotsSpacesBadRequest() *IndexRunbookRunbookSnapshotsSpacesBadRequest {
	return &IndexRunbookRunbookSnapshotsSpacesBadRequest{}
}

/*IndexRunbookRunbookSnapshotsSpacesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type IndexRunbookRunbookSnapshotsSpacesBadRequest struct {
}

func (o *IndexRunbookRunbookSnapshotsSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/runbooks/{id}/runbookSnapshots][%d] indexRunbookRunbookSnapshotsSpacesBadRequest ", 400)
}

func (o *IndexRunbookRunbookSnapshotsSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}