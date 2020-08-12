// Code generated by go-swagger; DO NOT EDIT.

package machines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// IndexDeploymentTargetTasksReader is a Reader for the IndexDeploymentTargetTasks structure.
type IndexDeploymentTargetTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexDeploymentTargetTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIndexDeploymentTargetTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIndexDeploymentTargetTasksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIndexDeploymentTargetTasksOK creates a IndexDeploymentTargetTasksOK with default headers values
func NewIndexDeploymentTargetTasksOK() *IndexDeploymentTargetTasksOK {
	return &IndexDeploymentTargetTasksOK{}
}

/*IndexDeploymentTargetTasksOK handles this case with default header values.

ResourceCollection_of_TaskResource resource returned
*/
type IndexDeploymentTargetTasksOK struct {
	Payload *models.ResourceCollectionOfTaskResource
}

func (o *IndexDeploymentTargetTasksOK) Error() string {
	return fmt.Sprintf("[GET /api/machines/{id}/tasks][%d] indexDeploymentTargetTasksOK  %+v", 200, o.Payload)
}

func (o *IndexDeploymentTargetTasksOK) GetPayload() *models.ResourceCollectionOfTaskResource {
	return o.Payload
}

func (o *IndexDeploymentTargetTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionOfTaskResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIndexDeploymentTargetTasksBadRequest creates a IndexDeploymentTargetTasksBadRequest with default headers values
func NewIndexDeploymentTargetTasksBadRequest() *IndexDeploymentTargetTasksBadRequest {
	return &IndexDeploymentTargetTasksBadRequest{}
}

/*IndexDeploymentTargetTasksBadRequest handles this case with default header values.

No id parameter was provided.
*/
type IndexDeploymentTargetTasksBadRequest struct {
}

func (o *IndexDeploymentTargetTasksBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/machines/{id}/tasks][%d] indexDeploymentTargetTasksBadRequest ", 400)
}

func (o *IndexDeploymentTargetTasksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
