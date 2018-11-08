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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Interface Interface
// swagger:model Interface
type Interface struct {

	// circuit termination
	CircuitTermination *InterfaceCircuitTermination `json:"circuit_termination,omitempty"`

	// Description
	// Max Length: 100
	Description string `json:"description,omitempty"`

	// device
	// Required: true
	Device *NestedDevice `json:"device"`

	// Enabled
	Enabled bool `json:"enabled,omitempty"`

	// form factor
	FormFactor *InterfaceFormFactor `json:"form_factor,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Interface connection
	// Read Only: true
	InterfaceConnection string `json:"interface_connection,omitempty"`

	// Is connected
	// Read Only: true
	IsConnected string `json:"is_connected,omitempty"`

	// lag
	Lag *NestedInterface `json:"lag,omitempty"`

	// MAC Address
	MacAddress string `json:"mac_address,omitempty"`

	// OOB Management
	//
	// This interface is used only for out-of-band management
	MgmtOnly bool `json:"mgmt_only,omitempty"`

	// mode
	Mode *InterfaceMode `json:"mode,omitempty"`

	// MTU
	// Maximum: 65536
	// Minimum: 1
	Mtu int64 `json:"mtu,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// tagged vlans
	// Unique: true
	TaggedVlans []int64 `json:"tagged_vlans"`

	// Tags
	Tags string `json:"tags,omitempty"`

	// untagged vlan
	UntaggedVlan *InterfaceVLAN `json:"untagged_vlan,omitempty"`
}

// Validate validates this interface
func (m *Interface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCircuitTermination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormFactor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtu(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaggedVlans(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUntaggedVlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Interface) validateCircuitTermination(formats strfmt.Registry) error {

	if swag.IsZero(m.CircuitTermination) { // not required
		return nil
	}

	if m.CircuitTermination != nil {
		if err := m.CircuitTermination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("circuit_termination")
			}
			return err
		}
	}

	return nil
}

func (m *Interface) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 100); err != nil {
		return err
	}

	return nil
}

func (m *Interface) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *Interface) validateFormFactor(formats strfmt.Registry) error {

	if swag.IsZero(m.FormFactor) { // not required
		return nil
	}

	if m.FormFactor != nil {
		if err := m.FormFactor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("form_factor")
			}
			return err
		}
	}

	return nil
}

func (m *Interface) validateLag(formats strfmt.Registry) error {

	if swag.IsZero(m.Lag) { // not required
		return nil
	}

	if m.Lag != nil {
		if err := m.Lag.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lag")
			}
			return err
		}
	}

	return nil
}

func (m *Interface) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	if m.Mode != nil {
		if err := m.Mode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode")
			}
			return err
		}
	}

	return nil
}

func (m *Interface) validateMtu(formats strfmt.Registry) error {

	if swag.IsZero(m.Mtu) { // not required
		return nil
	}

	if err := validate.MinimumInt("mtu", "body", int64(m.Mtu), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("mtu", "body", int64(m.Mtu), 65536, false); err != nil {
		return err
	}

	return nil
}

func (m *Interface) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 64); err != nil {
		return err
	}

	return nil
}

func (m *Interface) validateTaggedVlans(formats strfmt.Registry) error {

	if swag.IsZero(m.TaggedVlans) { // not required
		return nil
	}

	if err := validate.UniqueItems("tagged_vlans", "body", m.TaggedVlans); err != nil {
		return err
	}

	return nil
}

func (m *Interface) validateUntaggedVlan(formats strfmt.Registry) error {

	if swag.IsZero(m.UntaggedVlan) { // not required
		return nil
	}

	if m.UntaggedVlan != nil {
		if err := m.UntaggedVlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("untagged_vlan")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Interface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Interface) UnmarshalBinary(b []byte) error {
	var res Interface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InterfaceFormFactor Form factor
// swagger:model InterfaceFormFactor
type InterfaceFormFactor struct {

	// label
	// Required: true
	Label *string `json:"label"`

	// value
	// Required: true
	Value *int64 `json:"value"`
}

// Validate validates this interface form factor
func (m *InterfaceFormFactor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InterfaceFormFactor) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("form_factor"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *InterfaceFormFactor) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("form_factor"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InterfaceFormFactor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceFormFactor) UnmarshalBinary(b []byte) error {
	var res InterfaceFormFactor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InterfaceMode Mode
// swagger:model InterfaceMode
type InterfaceMode struct {

	// label
	// Required: true
	Label *string `json:"label"`

	// value
	// Required: true
	Value *int64 `json:"value"`
}

// Validate validates this interface mode
func (m *InterfaceMode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InterfaceMode) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("mode"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *InterfaceMode) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("mode"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InterfaceMode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceMode) UnmarshalBinary(b []byte) error {
	var res InterfaceMode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
