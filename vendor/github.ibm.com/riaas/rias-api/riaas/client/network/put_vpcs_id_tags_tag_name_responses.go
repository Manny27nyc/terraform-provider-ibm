// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// PutVpcsIDTagsTagNameReader is a Reader for the PutVpcsIDTagsTagName structure.
type PutVpcsIDTagsTagNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutVpcsIDTagsTagNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPutVpcsIDTagsTagNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutVpcsIDTagsTagNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutVpcsIDTagsTagNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutVpcsIDTagsTagNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewPutVpcsIDTagsTagNameNoContent creates a PutVpcsIDTagsTagNameNoContent with default headers values
func NewPutVpcsIDTagsTagNameNoContent() *PutVpcsIDTagsTagNameNoContent {
	return &PutVpcsIDTagsTagNameNoContent{}
}

/*PutVpcsIDTagsTagNameNoContent handles this case with default header values.

error
*/
type PutVpcsIDTagsTagNameNoContent struct {
	Payload *models.Riaaserror
}

func (o *PutVpcsIDTagsTagNameNoContent) Error() string {
	return fmt.Sprintf("[PUT /vpcs/{id}/tags/{tag_name}][%d] putVpcsIdTagsTagNameNoContent  %+v", 204, o.Payload)
}

func (o *PutVpcsIDTagsTagNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVpcsIDTagsTagNameBadRequest creates a PutVpcsIDTagsTagNameBadRequest with default headers values
func NewPutVpcsIDTagsTagNameBadRequest() *PutVpcsIDTagsTagNameBadRequest {
	return &PutVpcsIDTagsTagNameBadRequest{}
}

/*PutVpcsIDTagsTagNameBadRequest handles this case with default header values.

error
*/
type PutVpcsIDTagsTagNameBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PutVpcsIDTagsTagNameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /vpcs/{id}/tags/{tag_name}][%d] putVpcsIdTagsTagNameBadRequest  %+v", 400, o.Payload)
}

func (o *PutVpcsIDTagsTagNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVpcsIDTagsTagNameNotFound creates a PutVpcsIDTagsTagNameNotFound with default headers values
func NewPutVpcsIDTagsTagNameNotFound() *PutVpcsIDTagsTagNameNotFound {
	return &PutVpcsIDTagsTagNameNotFound{}
}

/*PutVpcsIDTagsTagNameNotFound handles this case with default header values.

error
*/
type PutVpcsIDTagsTagNameNotFound struct {
	Payload *models.Riaaserror
}

func (o *PutVpcsIDTagsTagNameNotFound) Error() string {
	return fmt.Sprintf("[PUT /vpcs/{id}/tags/{tag_name}][%d] putVpcsIdTagsTagNameNotFound  %+v", 404, o.Payload)
}

func (o *PutVpcsIDTagsTagNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVpcsIDTagsTagNameDefault creates a PutVpcsIDTagsTagNameDefault with default headers values
func NewPutVpcsIDTagsTagNameDefault(code int) *PutVpcsIDTagsTagNameDefault {
	return &PutVpcsIDTagsTagNameDefault{
		_statusCode: code,
	}
}

/*PutVpcsIDTagsTagNameDefault handles this case with default header values.

unexpectederror
*/
type PutVpcsIDTagsTagNameDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the put vpcs ID tags tag name default response
func (o *PutVpcsIDTagsTagNameDefault) Code() int {
	return o._statusCode
}

func (o *PutVpcsIDTagsTagNameDefault) Error() string {
	return fmt.Sprintf("[PUT /vpcs/{id}/tags/{tag_name}][%d] PutVpcsIDTagsTagName default  %+v", o._statusCode, o.Payload)
}

func (o *PutVpcsIDTagsTagNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}