// Code generated by go-swagger; DO NOT EDIT.

package build_information

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetBuildInformationSpacesReader is a Reader for the GetBuildInformationSpaces structure.
type GetBuildInformationSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildInformationSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBuildInformationSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBuildInformationSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBuildInformationSpacesOK creates a GetBuildInformationSpacesOK with default headers values
func NewGetBuildInformationSpacesOK() *GetBuildInformationSpacesOK {
	return &GetBuildInformationSpacesOK{}
}

/*GetBuildInformationSpacesOK handles this case with default header values.

OctopusPackageVersionBuildInformationMappedResource resource returned
*/
type GetBuildInformationSpacesOK struct {
	Payload *models.OctopusPackageVersionBuildInformationMappedResource
}

func (o *GetBuildInformationSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/build-information/{id}][%d] getBuildInformationSpacesOK  %+v", 200, o.Payload)
}

func (o *GetBuildInformationSpacesOK) GetPayload() *models.OctopusPackageVersionBuildInformationMappedResource {
	return o.Payload
}

func (o *GetBuildInformationSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OctopusPackageVersionBuildInformationMappedResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBuildInformationSpacesBadRequest creates a GetBuildInformationSpacesBadRequest with default headers values
func NewGetBuildInformationSpacesBadRequest() *GetBuildInformationSpacesBadRequest {
	return &GetBuildInformationSpacesBadRequest{}
}

/*GetBuildInformationSpacesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetBuildInformationSpacesBadRequest struct {
}

func (o *GetBuildInformationSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/build-information/{id}][%d] getBuildInformationSpacesBadRequest ", 400)
}

func (o *GetBuildInformationSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
