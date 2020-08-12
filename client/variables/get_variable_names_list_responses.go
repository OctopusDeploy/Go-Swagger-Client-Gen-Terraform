// Code generated by go-swagger; DO NOT EDIT.

package variables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetVariableNamesListReader is a Reader for the GetVariableNamesList structure.
type GetVariableNamesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVariableNamesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVariableNamesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVariableNamesListOK creates a GetVariableNamesListOK with default headers values
func NewGetVariableNamesListOK() *GetVariableNamesListOK {
	return &GetVariableNamesListOK{}
}

/*GetVariableNamesListOK handles this case with default header values.

OK - Default
*/
type GetVariableNamesListOK struct {
}

func (o *GetVariableNamesListOK) Error() string {
	return fmt.Sprintf("[GET /api/variables/names][%d] getVariableNamesListOK ", 200)
}

func (o *GetVariableNamesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}