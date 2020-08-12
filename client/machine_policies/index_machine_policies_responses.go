// Code generated by go-swagger; DO NOT EDIT.

package machine_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// IndexMachinePoliciesReader is a Reader for the IndexMachinePolicies structure.
type IndexMachinePoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexMachinePoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIndexMachinePoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIndexMachinePoliciesOK creates a IndexMachinePoliciesOK with default headers values
func NewIndexMachinePoliciesOK() *IndexMachinePoliciesOK {
	return &IndexMachinePoliciesOK{}
}

/*IndexMachinePoliciesOK handles this case with default header values.

ResourceCollection_of_MachinePolicyResource resource returned
*/
type IndexMachinePoliciesOK struct {
	Payload *models.ResourceCollectionOfMachinePolicyResource
}

func (o *IndexMachinePoliciesOK) Error() string {
	return fmt.Sprintf("[GET /api/machinepolicies][%d] indexMachinePoliciesOK  %+v", 200, o.Payload)
}

func (o *IndexMachinePoliciesOK) GetPayload() *models.ResourceCollectionOfMachinePolicyResource {
	return o.Payload
}

func (o *IndexMachinePoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionOfMachinePolicyResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
