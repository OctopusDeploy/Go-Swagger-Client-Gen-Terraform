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

// GetPackageRepositoryListReader is a Reader for the GetPackageRepositoryList structure.
type GetPackageRepositoryListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPackageRepositoryListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPackageRepositoryListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPackageRepositoryListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPackageRepositoryListOK creates a GetPackageRepositoryListOK with default headers values
func NewGetPackageRepositoryListOK() *GetPackageRepositoryListOK {
	return &GetPackageRepositoryListOK{}
}

/*GetPackageRepositoryListOK handles this case with default header values.

ResourceCollection_of_PackageResource resource returned
*/
type GetPackageRepositoryListOK struct {
	Payload *models.ResourceCollectionOfPackageResource
}

func (o *GetPackageRepositoryListOK) Error() string {
	return fmt.Sprintf("[GET /api/packages][%d] getPackageRepositoryListOK  %+v", 200, o.Payload)
}

func (o *GetPackageRepositoryListOK) GetPayload() *models.ResourceCollectionOfPackageResource {
	return o.Payload
}

func (o *GetPackageRepositoryListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionOfPackageResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackageRepositoryListBadRequest creates a GetPackageRepositoryListBadRequest with default headers values
func NewGetPackageRepositoryListBadRequest() *GetPackageRepositoryListBadRequest {
	return &GetPackageRepositoryListBadRequest{}
}

/*GetPackageRepositoryListBadRequest handles this case with default header values.

Only one of either nuGetPackageId or filter may be specified.
*/
type GetPackageRepositoryListBadRequest struct {
}

func (o *GetPackageRepositoryListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/packages][%d] getPackageRepositoryListBadRequest ", 400)
}

func (o *GetPackageRepositoryListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
