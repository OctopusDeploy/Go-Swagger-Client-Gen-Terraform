// Code generated by go-swagger; DO NOT EDIT.

package environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetDeploymentEnvironmentSettingsMetadataSpacesReader is a Reader for the GetDeploymentEnvironmentSettingsMetadataSpaces structure.
type GetDeploymentEnvironmentSettingsMetadataSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentEnvironmentSettingsMetadataSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentEnvironmentSettingsMetadataSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDeploymentEnvironmentSettingsMetadataSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeploymentEnvironmentSettingsMetadataSpacesOK creates a GetDeploymentEnvironmentSettingsMetadataSpacesOK with default headers values
func NewGetDeploymentEnvironmentSettingsMetadataSpacesOK() *GetDeploymentEnvironmentSettingsMetadataSpacesOK {
	return &GetDeploymentEnvironmentSettingsMetadataSpacesOK{}
}

/*GetDeploymentEnvironmentSettingsMetadataSpacesOK handles this case with default header values.

List_of_DeploymentEnvironmentSettingsMetadata resource returned
*/
type GetDeploymentEnvironmentSettingsMetadataSpacesOK struct {
	Payload []*models.DeploymentEnvironmentSettingsMetadata
}

func (o *GetDeploymentEnvironmentSettingsMetadataSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/environments/{id}/metadata][%d] getDeploymentEnvironmentSettingsMetadataSpacesOK  %+v", 200, o.Payload)
}

func (o *GetDeploymentEnvironmentSettingsMetadataSpacesOK) GetPayload() []*models.DeploymentEnvironmentSettingsMetadata {
	return o.Payload
}

func (o *GetDeploymentEnvironmentSettingsMetadataSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentEnvironmentSettingsMetadataSpacesBadRequest creates a GetDeploymentEnvironmentSettingsMetadataSpacesBadRequest with default headers values
func NewGetDeploymentEnvironmentSettingsMetadataSpacesBadRequest() *GetDeploymentEnvironmentSettingsMetadataSpacesBadRequest {
	return &GetDeploymentEnvironmentSettingsMetadataSpacesBadRequest{}
}

/*GetDeploymentEnvironmentSettingsMetadataSpacesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetDeploymentEnvironmentSettingsMetadataSpacesBadRequest struct {
}

func (o *GetDeploymentEnvironmentSettingsMetadataSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/environments/{id}/metadata][%d] getDeploymentEnvironmentSettingsMetadataSpacesBadRequest ", 400)
}

func (o *GetDeploymentEnvironmentSettingsMetadataSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
