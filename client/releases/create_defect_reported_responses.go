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

// CreateDefectReportedReader is a Reader for the CreateDefectReported structure.
type CreateDefectReportedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDefectReportedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDefectReportedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDefectReportedBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDefectReportedOK creates a CreateDefectReportedOK with default headers values
func NewCreateDefectReportedOK() *CreateDefectReportedOK {
	return &CreateDefectReportedOK{}
}

/*CreateDefectReportedOK handles this case with default header values.

DefectResource resource returned
*/
type CreateDefectReportedOK struct {
	Payload *models.DefectResource
}

func (o *CreateDefectReportedOK) Error() string {
	return fmt.Sprintf("[POST /api/releases/{id}/defects][%d] createDefectReportedOK  %+v", 200, o.Payload)
}

func (o *CreateDefectReportedOK) GetPayload() *models.DefectResource {
	return o.Payload
}

func (o *CreateDefectReportedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefectResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDefectReportedBadRequest creates a CreateDefectReportedBadRequest with default headers values
func NewCreateDefectReportedBadRequest() *CreateDefectReportedBadRequest {
	return &CreateDefectReportedBadRequest{}
}

/*CreateDefectReportedBadRequest handles this case with default header values.

An unresolved defect already exists for this release.
No id parameter was provided.
No request body was supplied.
You must provide a description.
*/
type CreateDefectReportedBadRequest struct {
}

func (o *CreateDefectReportedBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/releases/{id}/defects][%d] createDefectReportedBadRequest ", 400)
}

func (o *CreateDefectReportedBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}