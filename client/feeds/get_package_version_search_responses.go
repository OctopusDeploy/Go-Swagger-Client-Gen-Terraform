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

// GetPackageVersionSearchReader is a Reader for the GetPackageVersionSearch structure.
type GetPackageVersionSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPackageVersionSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPackageVersionSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPackageVersionSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPackageVersionSearchOK creates a GetPackageVersionSearchOK with default headers values
func NewGetPackageVersionSearchOK() *GetPackageVersionSearchOK {
	return &GetPackageVersionSearchOK{}
}

/*GetPackageVersionSearchOK handles this case with default header values.

ResourceCollection_of_PackageVersionResource resource returned
*/
type GetPackageVersionSearchOK struct {
	Payload *models.ResourceCollectionOfPackageVersionResource
}

func (o *GetPackageVersionSearchOK) Error() string {
	return fmt.Sprintf("[GET /api/feeds/{id}/packages/versions][%d] getPackageVersionSearchOK  %+v", 200, o.Payload)
}

func (o *GetPackageVersionSearchOK) GetPayload() *models.ResourceCollectionOfPackageVersionResource {
	return o.Payload
}

func (o *GetPackageVersionSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionOfPackageVersionResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackageVersionSearchBadRequest creates a GetPackageVersionSearchBadRequest with default headers values
func NewGetPackageVersionSearchBadRequest() *GetPackageVersionSearchBadRequest {
	return &GetPackageVersionSearchBadRequest{}
}

/*GetPackageVersionSearchBadRequest handles this case with default header values.

No id parameter was provided.
No packageId parameter was provided.
*/
type GetPackageVersionSearchBadRequest struct {
}

func (o *GetPackageVersionSearchBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/feeds/{id}/packages/versions][%d] getPackageVersionSearchBadRequest ", 400)
}

func (o *GetPackageVersionSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
