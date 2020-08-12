// Code generated by go-swagger; DO NOT EDIT.

package features_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// UpdateFeaturesConfigurationReader is a Reader for the UpdateFeaturesConfiguration structure.
type UpdateFeaturesConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateFeaturesConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateFeaturesConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateFeaturesConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateFeaturesConfigurationOK creates a UpdateFeaturesConfigurationOK with default headers values
func NewUpdateFeaturesConfigurationOK() *UpdateFeaturesConfigurationOK {
	return &UpdateFeaturesConfigurationOK{}
}

/*UpdateFeaturesConfigurationOK handles this case with default header values.

FeaturesConfigurationResource resource returned
*/
type UpdateFeaturesConfigurationOK struct {
	Payload *models.FeaturesConfigurationResource
}

func (o *UpdateFeaturesConfigurationOK) Error() string {
	return fmt.Sprintf("[PUT /api/featuresconfiguration][%d] updateFeaturesConfigurationOK  %+v", 200, o.Payload)
}

func (o *UpdateFeaturesConfigurationOK) GetPayload() *models.FeaturesConfigurationResource {
	return o.Payload
}

func (o *UpdateFeaturesConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FeaturesConfigurationResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFeaturesConfigurationBadRequest creates a UpdateFeaturesConfigurationBadRequest with default headers values
func NewUpdateFeaturesConfigurationBadRequest() *UpdateFeaturesConfigurationBadRequest {
	return &UpdateFeaturesConfigurationBadRequest{}
}

/*UpdateFeaturesConfigurationBadRequest handles this case with default header values.

Model validation failed.
No request body was supplied.
*/
type UpdateFeaturesConfigurationBadRequest struct {
}

func (o *UpdateFeaturesConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/featuresconfiguration][%d] updateFeaturesConfigurationBadRequest ", 400)
}

func (o *UpdateFeaturesConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}