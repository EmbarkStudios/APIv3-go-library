// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetSMSCampaignAllOf1Recipients get Sms campaign all of1 recipients
// swagger:model getSmsCampaignAllOf1Recipients
type GetSMSCampaignAllOf1Recipients struct {
	GetCampaignRecipients
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GetSMSCampaignAllOf1Recipients) UnmarshalJSON(raw []byte) error {

	var aO0 GetCampaignRecipients
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.GetCampaignRecipients = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GetSMSCampaignAllOf1Recipients) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.GetCampaignRecipients)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get Sms campaign all of1 recipients
func (m *GetSMSCampaignAllOf1Recipients) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.GetCampaignRecipients.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetSMSCampaignAllOf1Recipients) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSMSCampaignAllOf1Recipients) UnmarshalBinary(b []byte) error {
	var res GetSMSCampaignAllOf1Recipients
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
