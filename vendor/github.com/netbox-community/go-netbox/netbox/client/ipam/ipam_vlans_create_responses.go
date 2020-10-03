// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// IpamVlansCreateReader is a Reader for the IpamVlansCreate structure.
type IpamVlansCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlansCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamVlansCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVlansCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVlansCreateCreated creates a IpamVlansCreateCreated with default headers values
func NewIpamVlansCreateCreated() *IpamVlansCreateCreated {
	return &IpamVlansCreateCreated{}
}

/*IpamVlansCreateCreated handles this case with default header values.

IpamVlansCreateCreated ipam vlans create created
*/
type IpamVlansCreateCreated struct {
	Payload *models.VLAN
}

func (o *IpamVlansCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/vlans/][%d] ipamVlansCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamVlansCreateCreated) GetPayload() *models.VLAN {
	return o.Payload
}

func (o *IpamVlansCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLAN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamVlansCreateDefault creates a IpamVlansCreateDefault with default headers values
func NewIpamVlansCreateDefault(code int) *IpamVlansCreateDefault {
	return &IpamVlansCreateDefault{
		_statusCode: code,
	}
}

/*IpamVlansCreateDefault handles this case with default header values.

IpamVlansCreateDefault ipam vlans create default
*/
type IpamVlansCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam vlans create default response
func (o *IpamVlansCreateDefault) Code() int {
	return o._statusCode
}

func (o *IpamVlansCreateDefault) Error() string {
	return fmt.Sprintf("[POST /ipam/vlans/][%d] ipam_vlans_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVlansCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVlansCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
