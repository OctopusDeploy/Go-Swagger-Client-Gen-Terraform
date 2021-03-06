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

// GetReleaseByIDReader is a Reader for the GetReleaseByID structure.
type GetReleaseByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReleaseByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReleaseByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetReleaseByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReleaseByIDOK creates a GetReleaseByIDOK with default headers values
func NewGetReleaseByIDOK() *GetReleaseByIDOK {
	return &GetReleaseByIDOK{}
}

/*GetReleaseByIDOK handles this case with default header values.

Load ReleaseResource by ID
*/
type GetReleaseByIDOK struct {
	Payload *models.ReleaseResource
}

func (o *GetReleaseByIDOK) Error() string {
	return fmt.Sprintf("[GET /api/releases/{id}][%d] getReleaseByIdOK  %+v", 200, o.Payload)
}

func (o *GetReleaseByIDOK) GetPayload() *models.ReleaseResource {
	return o.Payload
}

func (o *GetReleaseByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReleaseResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReleaseByIDBadRequest creates a GetReleaseByIDBadRequest with default headers values
func NewGetReleaseByIDBadRequest() *GetReleaseByIDBadRequest {
	return &GetReleaseByIDBadRequest{}
}

/*GetReleaseByIDBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetReleaseByIDBadRequest struct {
}

func (o *GetReleaseByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/releases/{id}][%d] getReleaseByIdBadRequest ", 400)
}

func (o *GetReleaseByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
