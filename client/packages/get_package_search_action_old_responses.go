// Code generated by go-swagger; DO NOT EDIT.

package packages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetPackageSearchActionOldReader is a Reader for the GetPackageSearchActionOld structure.
type GetPackageSearchActionOldReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPackageSearchActionOldReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPackageSearchActionOldOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPackageSearchActionOldBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPackageSearchActionOldOK creates a GetPackageSearchActionOldOK with default headers values
func NewGetPackageSearchActionOldOK() *GetPackageSearchActionOldOK {
	return &GetPackageSearchActionOldOK{}
}

/*GetPackageSearchActionOldOK handles this case with default header values.

IEnumerable_of_PackageResource resource returned
*/
type GetPackageSearchActionOldOK struct {
	Payload []*models.PackageResource
}

func (o *GetPackageSearchActionOldOK) Error() string {
	return fmt.Sprintf("[GET /api/feeds/{id}/packages][%d] getPackageSearchActionOldOK  %+v", 200, o.Payload)
}

func (o *GetPackageSearchActionOldOK) GetPayload() []*models.PackageResource {
	return o.Payload
}

func (o *GetPackageSearchActionOldOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackageSearchActionOldBadRequest creates a GetPackageSearchActionOldBadRequest with default headers values
func NewGetPackageSearchActionOldBadRequest() *GetPackageSearchActionOldBadRequest {
	return &GetPackageSearchActionOldBadRequest{}
}

/*GetPackageSearchActionOldBadRequest handles this case with default header values.

No id parameter was provided.
The 'versionRange' parameter was not a valid NuGet version-range (see http://g.octopushq.com/NuGetVersioning)
*/
type GetPackageSearchActionOldBadRequest struct {
}

func (o *GetPackageSearchActionOldBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/feeds/{id}/packages][%d] getPackageSearchActionOldBadRequest ", 400)
}

func (o *GetPackageSearchActionOldBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}