// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDcimConsolePortsListParams creates a new DcimConsolePortsListParams object
// with the default values initialized.
func NewDcimConsolePortsListParams() *DcimConsolePortsListParams {
	var ()
	return &DcimConsolePortsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsolePortsListParamsWithTimeout creates a new DcimConsolePortsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimConsolePortsListParamsWithTimeout(timeout time.Duration) *DcimConsolePortsListParams {
	var ()
	return &DcimConsolePortsListParams{

		timeout: timeout,
	}
}

// NewDcimConsolePortsListParamsWithContext creates a new DcimConsolePortsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimConsolePortsListParamsWithContext(ctx context.Context) *DcimConsolePortsListParams {
	var ()
	return &DcimConsolePortsListParams{

		Context: ctx,
	}
}

// NewDcimConsolePortsListParamsWithHTTPClient creates a new DcimConsolePortsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimConsolePortsListParamsWithHTTPClient(client *http.Client) *DcimConsolePortsListParams {
	var ()
	return &DcimConsolePortsListParams{
		HTTPClient: client,
	}
}

/*DcimConsolePortsListParams contains all the parameters to send to the API endpoint
for the dcim console ports list operation typically these are written to a http.Request
*/
type DcimConsolePortsListParams struct {

	/*Device*/
	Device *string
	/*DeviceID*/
	DeviceID *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Name*/
	Name *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Tag*/
	Tag *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim console ports list params
func (o *DcimConsolePortsListParams) WithTimeout(timeout time.Duration) *DcimConsolePortsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console ports list params
func (o *DcimConsolePortsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console ports list params
func (o *DcimConsolePortsListParams) WithContext(ctx context.Context) *DcimConsolePortsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console ports list params
func (o *DcimConsolePortsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console ports list params
func (o *DcimConsolePortsListParams) WithHTTPClient(client *http.Client) *DcimConsolePortsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console ports list params
func (o *DcimConsolePortsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevice adds the device to the dcim console ports list params
func (o *DcimConsolePortsListParams) WithDevice(device *string) *DcimConsolePortsListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the dcim console ports list params
func (o *DcimConsolePortsListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the dcim console ports list params
func (o *DcimConsolePortsListParams) WithDeviceID(deviceID *string) *DcimConsolePortsListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the dcim console ports list params
func (o *DcimConsolePortsListParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithLimit adds the limit to the dcim console ports list params
func (o *DcimConsolePortsListParams) WithLimit(limit *int64) *DcimConsolePortsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim console ports list params
func (o *DcimConsolePortsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim console ports list params
func (o *DcimConsolePortsListParams) WithName(name *string) *DcimConsolePortsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim console ports list params
func (o *DcimConsolePortsListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the dcim console ports list params
func (o *DcimConsolePortsListParams) WithOffset(offset *int64) *DcimConsolePortsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim console ports list params
func (o *DcimConsolePortsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithTag adds the tag to the dcim console ports list params
func (o *DcimConsolePortsListParams) WithTag(tag *string) *DcimConsolePortsListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the dcim console ports list params
func (o *DcimConsolePortsListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsolePortsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Device != nil {

		// query param device
		var qrDevice string
		if o.Device != nil {
			qrDevice = *o.Device
		}
		qDevice := qrDevice
		if qDevice != "" {
			if err := r.SetQueryParam("device", qDevice); err != nil {
				return err
			}
		}

	}

	if o.DeviceID != nil {

		// query param device_id
		var qrDeviceID string
		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := qrDeviceID
		if qDeviceID != "" {
			if err := r.SetQueryParam("device_id", qDeviceID); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
