// Code generated by go-swagger; DO NOT EDIT.

package releases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// CreateDefectReportedSpacesReader is a Reader for the CreateDefectReportedSpaces structure.
type CreateDefectReportedSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDefectReportedSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDefectReportedSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDefectReportedSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDefectReportedSpacesOK creates a CreateDefectReportedSpacesOK with default headers values
func NewCreateDefectReportedSpacesOK() *CreateDefectReportedSpacesOK {
	return &CreateDefectReportedSpacesOK{}
}

/*CreateDefectReportedSpacesOK handles this case with default header values.

DefectResource resource returned
*/
type CreateDefectReportedSpacesOK struct {
	Payload *models.DefectResource
}

func (o *CreateDefectReportedSpacesOK) Error() string {
	return fmt.Sprintf("[POST /api/{baseSpaceId}/releases/{id}/defects][%d] createDefectReportedSpacesOK  %+v", 200, o.Payload)
}

func (o *CreateDefectReportedSpacesOK) GetPayload() *models.DefectResource {
	return o.Payload
}

func (o *CreateDefectReportedSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefectResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDefectReportedSpacesBadRequest creates a CreateDefectReportedSpacesBadRequest with default headers values
func NewCreateDefectReportedSpacesBadRequest() *CreateDefectReportedSpacesBadRequest {
	return &CreateDefectReportedSpacesBadRequest{}
}

/*CreateDefectReportedSpacesBadRequest handles this case with default header values.

An unresolved defect already exists for this release.
No id parameter was provided.
No request body was supplied.
You must provide a description.
*/
type CreateDefectReportedSpacesBadRequest struct {
}

func (o *CreateDefectReportedSpacesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/{baseSpaceId}/releases/{id}/defects][%d] createDefectReportedSpacesBadRequest ", 400)
}

func (o *CreateDefectReportedSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}