// Code generated by go-swagger; DO NOT EDIT.

package channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateVersionRuleTestSpacesReader is a Reader for the CreateVersionRuleTestSpaces structure.
type CreateVersionRuleTestSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVersionRuleTestSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateVersionRuleTestSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateVersionRuleTestSpacesOK creates a CreateVersionRuleTestSpacesOK with default headers values
func NewCreateVersionRuleTestSpacesOK() *CreateVersionRuleTestSpacesOK {
	return &CreateVersionRuleTestSpacesOK{}
}

/*CreateVersionRuleTestSpacesOK handles this case with default header values.

OK - Default
*/
type CreateVersionRuleTestSpacesOK struct {
}

func (o *CreateVersionRuleTestSpacesOK) Error() string {
	return fmt.Sprintf("[POST /api/{baseSpaceId}/channels/rule-test][%d] createVersionRuleTestSpacesOK ", 200)
}

func (o *CreateVersionRuleTestSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
