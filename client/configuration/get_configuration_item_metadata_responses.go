// Code generated by go-swagger; DO NOT EDIT.

package configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetConfigurationItemMetadataReader is a Reader for the GetConfigurationItemMetadata structure.
type GetConfigurationItemMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfigurationItemMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConfigurationItemMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConfigurationItemMetadataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConfigurationItemMetadataOK creates a GetConfigurationItemMetadataOK with default headers values
func NewGetConfigurationItemMetadataOK() *GetConfigurationItemMetadataOK {
	return &GetConfigurationItemMetadataOK{}
}

/*GetConfigurationItemMetadataOK handles this case with default header values.

Metadata resource returned
*/
type GetConfigurationItemMetadataOK struct {
	Payload *models.Metadata
}

func (o *GetConfigurationItemMetadataOK) Error() string {
	return fmt.Sprintf("[GET /api/configuration/{id}/metadata][%d] getConfigurationItemMetadataOK  %+v", 200, o.Payload)
}

func (o *GetConfigurationItemMetadataOK) GetPayload() *models.Metadata {
	return o.Payload
}

func (o *GetConfigurationItemMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Metadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationItemMetadataBadRequest creates a GetConfigurationItemMetadataBadRequest with default headers values
func NewGetConfigurationItemMetadataBadRequest() *GetConfigurationItemMetadataBadRequest {
	return &GetConfigurationItemMetadataBadRequest{}
}

/*GetConfigurationItemMetadataBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetConfigurationItemMetadataBadRequest struct {
}

func (o *GetConfigurationItemMetadataBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/configuration/{id}/metadata][%d] getConfigurationItemMetadataBadRequest ", 400)
}

func (o *GetConfigurationItemMetadataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
