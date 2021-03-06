// Code generated by go-swagger; DO NOT EDIT.

package runbook_runs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteRunbookRunReader is a Reader for the DeleteRunbookRun structure.
type DeleteRunbookRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRunbookRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRunbookRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRunbookRunBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRunbookRunOK creates a DeleteRunbookRunOK with default headers values
func NewDeleteRunbookRunOK() *DeleteRunbookRunOK {
	return &DeleteRunbookRunOK{}
}

/*DeleteRunbookRunOK handles this case with default header values.

OK
*/
type DeleteRunbookRunOK struct {
}

func (o *DeleteRunbookRunOK) Error() string {
	return fmt.Sprintf("[DELETE /api/runbookRuns/{id}][%d] deleteRunbookRunOK ", 200)
}

func (o *DeleteRunbookRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRunbookRunBadRequest creates a DeleteRunbookRunBadRequest with default headers values
func NewDeleteRunbookRunBadRequest() *DeleteRunbookRunBadRequest {
	return &DeleteRunbookRunBadRequest{}
}

/*DeleteRunbookRunBadRequest handles this case with default header values.

BadRequest
No id parameter was provided.
This resource is read-only and cannot be deleted.
*/
type DeleteRunbookRunBadRequest struct {
}

func (o *DeleteRunbookRunBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/runbookRuns/{id}][%d] deleteRunbookRunBadRequest ", 400)
}

func (o *DeleteRunbookRunBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
