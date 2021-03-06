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

// GetBuildInformationListSpacesReader is a Reader for the GetBuildInformationListSpaces structure.
type GetBuildInformationListSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildInformationListSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBuildInformationListSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBuildInformationListSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBuildInformationListSpacesOK creates a GetBuildInformationListSpacesOK with default headers values
func NewGetBuildInformationListSpacesOK() *GetBuildInformationListSpacesOK {
	return &GetBuildInformationListSpacesOK{}
}

/*GetBuildInformationListSpacesOK handles this case with default header values.

ResourceCollection_of_OctopusPackageVersionBuildInformationMappedResource resource returned
*/
type GetBuildInformationListSpacesOK struct {
	Payload *models.ResourceCollectionOfOctopusPackageVersionBuildInformationMappedResource
}

func (o *GetBuildInformationListSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/build-information][%d] getBuildInformationListSpacesOK  %+v", 200, o.Payload)
}

func (o *GetBuildInformationListSpacesOK) GetPayload() *models.ResourceCollectionOfOctopusPackageVersionBuildInformationMappedResource {
	return o.Payload
}

func (o *GetBuildInformationListSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionOfOctopusPackageVersionBuildInformationMappedResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBuildInformationListSpacesBadRequest creates a GetBuildInformationListSpacesBadRequest with default headers values
func NewGetBuildInformationListSpacesBadRequest() *GetBuildInformationListSpacesBadRequest {
	return &GetBuildInformationListSpacesBadRequest{}
}

/*GetBuildInformationListSpacesBadRequest handles this case with default header values.

Only one of either packageId or filter may be specified.
*/
type GetBuildInformationListSpacesBadRequest struct {
}

func (o *GetBuildInformationListSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/build-information][%d] getBuildInformationListSpacesBadRequest ", 400)
}

func (o *GetBuildInformationListSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
