// Code generated by go-swagger; DO NOT EDIT.

package reporting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetDeploymentsByProjectReportReader is a Reader for the GetDeploymentsByProjectReport structure.
type GetDeploymentsByProjectReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentsByProjectReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentsByProjectReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDeploymentsByProjectReportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeploymentsByProjectReportOK creates a GetDeploymentsByProjectReportOK with default headers values
func NewGetDeploymentsByProjectReportOK() *GetDeploymentsByProjectReportOK {
	return &GetDeploymentsByProjectReportOK{}
}

/*GetDeploymentsByProjectReportOK handles this case with default header values.

IEnumerable_of_ReportDeploymentCountOverTimeResource resource returned
*/
type GetDeploymentsByProjectReportOK struct {
	Payload []*models.ReportDeploymentCountOverTimeResource
}

func (o *GetDeploymentsByProjectReportOK) Error() string {
	return fmt.Sprintf("[GET /api/reporting/deployments-counted-by-week][%d] getDeploymentsByProjectReportOK  %+v", 200, o.Payload)
}

func (o *GetDeploymentsByProjectReportOK) GetPayload() []*models.ReportDeploymentCountOverTimeResource {
	return o.Payload
}

func (o *GetDeploymentsByProjectReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentsByProjectReportBadRequest creates a GetDeploymentsByProjectReportBadRequest with default headers values
func NewGetDeploymentsByProjectReportBadRequest() *GetDeploymentsByProjectReportBadRequest {
	return &GetDeploymentsByProjectReportBadRequest{}
}

/*GetDeploymentsByProjectReportBadRequest handles this case with default header values.

A projectIds query string parameter is required.
*/
type GetDeploymentsByProjectReportBadRequest struct {
}

func (o *GetDeploymentsByProjectReportBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/reporting/deployments-counted-by-week][%d] getDeploymentsByProjectReportBadRequest ", 400)
}

func (o *GetDeploymentsByProjectReportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}