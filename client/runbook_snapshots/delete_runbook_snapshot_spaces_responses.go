// Code generated by go-swagger; DO NOT EDIT.

package runbook_snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteRunbookSnapshotSpacesReader is a Reader for the DeleteRunbookSnapshotSpaces structure.
type DeleteRunbookSnapshotSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRunbookSnapshotSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRunbookSnapshotSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRunbookSnapshotSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRunbookSnapshotSpacesOK creates a DeleteRunbookSnapshotSpacesOK with default headers values
func NewDeleteRunbookSnapshotSpacesOK() *DeleteRunbookSnapshotSpacesOK {
	return &DeleteRunbookSnapshotSpacesOK{}
}

/*DeleteRunbookSnapshotSpacesOK handles this case with default header values.

OK
*/
type DeleteRunbookSnapshotSpacesOK struct {
}

func (o *DeleteRunbookSnapshotSpacesOK) Error() string {
	return fmt.Sprintf("[DELETE /api/{baseSpaceId}/runbookSnapshots/{id}][%d] deleteRunbookSnapshotSpacesOK ", 200)
}

func (o *DeleteRunbookSnapshotSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRunbookSnapshotSpacesBadRequest creates a DeleteRunbookSnapshotSpacesBadRequest with default headers values
func NewDeleteRunbookSnapshotSpacesBadRequest() *DeleteRunbookSnapshotSpacesBadRequest {
	return &DeleteRunbookSnapshotSpacesBadRequest{}
}

/*DeleteRunbookSnapshotSpacesBadRequest handles this case with default header values.

BadRequest
No id parameter was provided.
This resource is read-only and cannot be deleted.
*/
type DeleteRunbookSnapshotSpacesBadRequest struct {
}

func (o *DeleteRunbookSnapshotSpacesBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/{baseSpaceId}/runbookSnapshots/{id}][%d] deleteRunbookSnapshotSpacesBadRequest ", 400)
}

func (o *DeleteRunbookSnapshotSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
