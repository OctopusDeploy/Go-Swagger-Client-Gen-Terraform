// Code generated by go-swagger; DO NOT EDIT.

package progression

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetRunbookTaskRunDashboardItemsReader is a Reader for the GetRunbookTaskRunDashboardItems structure.
type GetRunbookTaskRunDashboardItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunbookTaskRunDashboardItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunbookTaskRunDashboardItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRunbookTaskRunDashboardItemsOK creates a GetRunbookTaskRunDashboardItemsOK with default headers values
func NewGetRunbookTaskRunDashboardItemsOK() *GetRunbookTaskRunDashboardItemsOK {
	return &GetRunbookTaskRunDashboardItemsOK{}
}

/*GetRunbookTaskRunDashboardItemsOK handles this case with default header values.

ResourceCollection_of_RunbooksDashboardItemResource resource returned
*/
type GetRunbookTaskRunDashboardItemsOK struct {
	Payload *models.ResourceCollectionOfRunbooksDashboardItemResource
}

func (o *GetRunbookTaskRunDashboardItemsOK) Error() string {
	return fmt.Sprintf("[GET /api/progression/runbooks/taskRuns][%d] getRunbookTaskRunDashboardItemsOK  %+v", 200, o.Payload)
}

func (o *GetRunbookTaskRunDashboardItemsOK) GetPayload() *models.ResourceCollectionOfRunbooksDashboardItemResource {
	return o.Payload
}

func (o *GetRunbookTaskRunDashboardItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionOfRunbooksDashboardItemResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
