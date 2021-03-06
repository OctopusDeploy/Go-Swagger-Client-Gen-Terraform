// Code generated by go-swagger; DO NOT EDIT.

package feeds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetPackageVersionSearchSpacesReader is a Reader for the GetPackageVersionSearchSpaces structure.
type GetPackageVersionSearchSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPackageVersionSearchSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPackageVersionSearchSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPackageVersionSearchSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPackageVersionSearchSpacesOK creates a GetPackageVersionSearchSpacesOK with default headers values
func NewGetPackageVersionSearchSpacesOK() *GetPackageVersionSearchSpacesOK {
	return &GetPackageVersionSearchSpacesOK{}
}

/*GetPackageVersionSearchSpacesOK handles this case with default header values.

ResourceCollection_of_PackageVersionResource resource returned
*/
type GetPackageVersionSearchSpacesOK struct {
	Payload *models.ResourceCollectionOfPackageVersionResource
}

func (o *GetPackageVersionSearchSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/feeds/{id}/packages/versions][%d] getPackageVersionSearchSpacesOK  %+v", 200, o.Payload)
}

func (o *GetPackageVersionSearchSpacesOK) GetPayload() *models.ResourceCollectionOfPackageVersionResource {
	return o.Payload
}

func (o *GetPackageVersionSearchSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionOfPackageVersionResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackageVersionSearchSpacesBadRequest creates a GetPackageVersionSearchSpacesBadRequest with default headers values
func NewGetPackageVersionSearchSpacesBadRequest() *GetPackageVersionSearchSpacesBadRequest {
	return &GetPackageVersionSearchSpacesBadRequest{}
}

/*GetPackageVersionSearchSpacesBadRequest handles this case with default header values.

No id parameter was provided.
No packageId parameter was provided.
*/
type GetPackageVersionSearchSpacesBadRequest struct {
}

func (o *GetPackageVersionSearchSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/feeds/{id}/packages/versions][%d] getPackageVersionSearchSpacesBadRequest ", 400)
}

func (o *GetPackageVersionSearchSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
