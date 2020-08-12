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

// CreateDeploymentTargetSpacesReader is a Reader for the CreateDeploymentTargetSpaces structure.
type CreateDeploymentTargetSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDeploymentTargetSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDeploymentTargetSpacesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDeploymentTargetSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDeploymentTargetSpacesCreated creates a CreateDeploymentTargetSpacesCreated with default headers values
func NewCreateDeploymentTargetSpacesCreated() *CreateDeploymentTargetSpacesCreated {
	return &CreateDeploymentTargetSpacesCreated{}
}

/*CreateDeploymentTargetSpacesCreated handles this case with default header values.

DeploymentTargetResource Created
*/
type CreateDeploymentTargetSpacesCreated struct {
	Payload *models.DeploymentTargetResource
}

func (o *CreateDeploymentTargetSpacesCreated) Error() string {
	return fmt.Sprintf("[POST /api/{baseSpaceId}/machines][%d] createDeploymentTargetSpacesCreated  %+v", 201, o.Payload)
}

func (o *CreateDeploymentTargetSpacesCreated) GetPayload() *models.DeploymentTargetResource {
	return o.Payload
}

func (o *CreateDeploymentTargetSpacesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentTargetResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDeploymentTargetSpacesBadRequest creates a CreateDeploymentTargetSpacesBadRequest with default headers values
func NewCreateDeploymentTargetSpacesBadRequest() *CreateDeploymentTargetSpacesBadRequest {
	return &CreateDeploymentTargetSpacesBadRequest{}
}

/*CreateDeploymentTargetSpacesBadRequest handles this case with default header values.

Model validation failed.
No request body was supplied.
*/
type CreateDeploymentTargetSpacesBadRequest struct {
}

func (o *CreateDeploymentTargetSpacesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/{baseSpaceId}/machines][%d] createDeploymentTargetSpacesBadRequest ", 400)
}

func (o *CreateDeploymentTargetSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}