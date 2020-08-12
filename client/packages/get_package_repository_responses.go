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

// GetPackageRepositoryReader is a Reader for the GetPackageRepository structure.
type GetPackageRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPackageRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPackageRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPackageRepositoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPackageRepositoryOK creates a GetPackageRepositoryOK with default headers values
func NewGetPackageRepositoryOK() *GetPackageRepositoryOK {
	return &GetPackageRepositoryOK{}
}

/*GetPackageRepositoryOK handles this case with default header values.

PackageFromBuiltInFeedResource resource returned
*/
type GetPackageRepositoryOK struct {
	Payload *models.PackageFromBuiltInFeedResource
}

func (o *GetPackageRepositoryOK) Error() string {
	return fmt.Sprintf("[GET /api/packages/{id}][%d] getPackageRepositoryOK  %+v", 200, o.Payload)
}

func (o *GetPackageRepositoryOK) GetPayload() *models.PackageFromBuiltInFeedResource {
	return o.Payload
}

func (o *GetPackageRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PackageFromBuiltInFeedResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackageRepositoryBadRequest creates a GetPackageRepositoryBadRequest with default headers values
func NewGetPackageRepositoryBadRequest() *GetPackageRepositoryBadRequest {
	return &GetPackageRepositoryBadRequest{}
}

/*GetPackageRepositoryBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetPackageRepositoryBadRequest struct {
}

func (o *GetPackageRepositoryBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/packages/{id}][%d] getPackageRepositoryBadRequest ", 400)
}

func (o *GetPackageRepositoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}