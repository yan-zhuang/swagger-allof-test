package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// B b
// swagger:model B
type B struct {
	A

	// f3
	// Required: true
	F3 *string `json:"f3"`

	// f4
	// Required: true
	F4 []string `json:"f4"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *B) UnmarshalJSON(raw []byte) error {

	var aO0 A
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.A = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m B) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.A)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this b
func (m *B) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.A.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateF3(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateF4(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *B) validateF3(formats strfmt.Registry) error {

	if err := validate.Required("f3", "body", m.F3); err != nil {
		return err
	}

	return nil
}

func (m *B) validateF4(formats strfmt.Registry) error {

	if err := validate.Required("f4", "body", m.F4); err != nil {
		return err
	}

	return nil
}