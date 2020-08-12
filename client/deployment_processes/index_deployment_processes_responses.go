// Code generated by go-swagger; DO NOT EDIT.

package deployment_processes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// IndexDeploymentProcessesReader is a Reader for the IndexDeploymentProcesses structure.
type IndexDeploymentProcessesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexDeploymentProcessesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIndexDeploymentProcessesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIndexDeploymentProcessesOK creates a IndexDeploymentProcessesOK with default headers values
func NewIndexDeploymentProcessesOK() *IndexDeploymentProcessesOK {
	return &IndexDeploymentProcessesOK{}
}

/*IndexDeploymentProcessesOK handles this case with default header values.

ResourceCollection_of_DeploymentProcessResource resource returned
*/
type IndexDeploymentProcessesOK struct {
	Payload *models.ResourceCollectionOfDeploymentProcessResource
}

func (o *IndexDeploymentProcessesOK) Error() string {
	return fmt.Sprintf("[GET /api/deploymentprocesses][%d] indexDeploymentProcessesOK  %+v", 200, o.Payload)
}

func (o *IndexDeploymentProcessesOK) GetPayload() *models.ResourceCollectionOfDeploymentProcessResource {
	return o.Payload
}

func (o *IndexDeploymentProcessesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionOfDeploymentProcessResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
