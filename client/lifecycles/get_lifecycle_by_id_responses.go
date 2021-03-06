// Code generated by go-swagger; DO NOT EDIT.

package lifecycles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetLifecycleByIDReader is a Reader for the GetLifecycleByID structure.
type GetLifecycleByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLifecycleByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLifecycleByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLifecycleByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLifecycleByIDOK creates a GetLifecycleByIDOK with default headers values
func NewGetLifecycleByIDOK() *GetLifecycleByIDOK {
	return &GetLifecycleByIDOK{}
}

/*GetLifecycleByIDOK handles this case with default header values.

Load LifecycleResource by ID
*/
type GetLifecycleByIDOK struct {
	Payload *models.LifecycleResource
}

func (o *GetLifecycleByIDOK) Error() string {
	return fmt.Sprintf("[GET /api/lifecycles/{id}][%d] getLifecycleByIdOK  %+v", 200, o.Payload)
}

func (o *GetLifecycleByIDOK) GetPayload() *models.LifecycleResource {
	return o.Payload
}

func (o *GetLifecycleByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LifecycleResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLifecycleByIDBadRequest creates a GetLifecycleByIDBadRequest with default headers values
func NewGetLifecycleByIDBadRequest() *GetLifecycleByIDBadRequest {
	return &GetLifecycleByIDBadRequest{}
}

/*GetLifecycleByIDBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetLifecycleByIDBadRequest struct {
}

func (o *GetLifecycleByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/lifecycles/{id}][%d] getLifecycleByIdBadRequest ", 400)
}

func (o *GetLifecycleByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
