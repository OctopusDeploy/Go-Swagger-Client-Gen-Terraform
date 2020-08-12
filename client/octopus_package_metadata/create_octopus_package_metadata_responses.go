// Code generated by go-swagger; DO NOT EDIT.

package octopus_package_metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// CreateOctopusPackageMetadataReader is a Reader for the CreateOctopusPackageMetadata structure.
type CreateOctopusPackageMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOctopusPackageMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOctopusPackageMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateOctopusPackageMetadataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateOctopusPackageMetadataConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateOctopusPackageMetadataOK creates a CreateOctopusPackageMetadataOK with default headers values
func NewCreateOctopusPackageMetadataOK() *CreateOctopusPackageMetadataOK {
	return &CreateOctopusPackageMetadataOK{}
}

/*CreateOctopusPackageMetadataOK handles this case with default header values.

OctopusPackageMetadataMappedResource resource returned
*/
type CreateOctopusPackageMetadataOK struct {
	Payload *models.OctopusPackageMetadataMappedResource
}

func (o *CreateOctopusPackageMetadataOK) Error() string {
	return fmt.Sprintf("[POST /api/package-metadata][%d] createOctopusPackageMetadataOK  %+v", 200, o.Payload)
}

func (o *CreateOctopusPackageMetadataOK) GetPayload() *models.OctopusPackageMetadataMappedResource {
	return o.Payload
}

func (o *CreateOctopusPackageMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OctopusPackageMetadataMappedResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOctopusPackageMetadataBadRequest creates a CreateOctopusPackageMetadataBadRequest with default headers values
func NewCreateOctopusPackageMetadataBadRequest() *CreateOctopusPackageMetadataBadRequest {
	return &CreateOctopusPackageMetadataBadRequest{}
}

/*CreateOctopusPackageMetadataBadRequest handles this case with default header values.

Invalid package metadata
*/
type CreateOctopusPackageMetadataBadRequest struct {
}

func (o *CreateOctopusPackageMetadataBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/package-metadata][%d] createOctopusPackageMetadataBadRequest ", 400)
}

func (o *CreateOctopusPackageMetadataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOctopusPackageMetadataConflict creates a CreateOctopusPackageMetadataConflict with default headers values
func NewCreateOctopusPackageMetadataConflict() *CreateOctopusPackageMetadataConflict {
	return &CreateOctopusPackageMetadataConflict{}
}

/*CreateOctopusPackageMetadataConflict handles this case with default header values.

Metadata for the specified Package ID and version already exists. To proceed anyway, specify an overwriteMode of OverwriteExisting or IgnoreIfExists.
*/
type CreateOctopusPackageMetadataConflict struct {
}

func (o *CreateOctopusPackageMetadataConflict) Error() string {
	return fmt.Sprintf("[POST /api/package-metadata][%d] createOctopusPackageMetadataConflict ", 409)
}

func (o *CreateOctopusPackageMetadataConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
