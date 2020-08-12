// Code generated by go-swagger; DO NOT EDIT.

package runbook_runs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetRunbookRunByIDReader is a Reader for the GetRunbookRunByID structure.
type GetRunbookRunByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunbookRunByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunbookRunByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRunbookRunByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRunbookRunByIDOK creates a GetRunbookRunByIDOK with default headers values
func NewGetRunbookRunByIDOK() *GetRunbookRunByIDOK {
	return &GetRunbookRunByIDOK{}
}

/*GetRunbookRunByIDOK handles this case with default header values.

Load RunbookRunResource by ID
*/
type GetRunbookRunByIDOK struct {
	Payload *models.RunbookRunResource
}

func (o *GetRunbookRunByIDOK) Error() string {
	return fmt.Sprintf("[GET /api/runbookRuns/{id}][%d] getRunbookRunByIdOK  %+v", 200, o.Payload)
}

func (o *GetRunbookRunByIDOK) GetPayload() *models.RunbookRunResource {
	return o.Payload
}

func (o *GetRunbookRunByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RunbookRunResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunbookRunByIDBadRequest creates a GetRunbookRunByIDBadRequest with default headers values
func NewGetRunbookRunByIDBadRequest() *GetRunbookRunByIDBadRequest {
	return &GetRunbookRunByIDBadRequest{}
}

/*GetRunbookRunByIDBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetRunbookRunByIDBadRequest struct {
}

func (o *GetRunbookRunByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/runbookRuns/{id}][%d] getRunbookRunByIdBadRequest ", 400)
}

func (o *GetRunbookRunByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
