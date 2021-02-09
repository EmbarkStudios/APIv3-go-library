// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateCampaignStatus Status of the campaign
// swagger:model updateCampaignStatus
type UpdateCampaignStatus struct {

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this update campaign status
func (m *UpdateCampaignStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateCampaignStatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["suspended","archive","darchive","sent","queued","replicate","replicateTemplate","draft"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateCampaignStatusTypeStatusPropEnum = append(updateCampaignStatusTypeStatusPropEnum, v)
	}
}

const (
	// UpdateCampaignStatusStatusSuspended captures enum value "suspended"
	UpdateCampaignStatusStatusSuspended string = "suspended"
	// UpdateCampaignStatusStatusArchive captures enum value "archive"
	UpdateCampaignStatusStatusArchive string = "archive"
	// UpdateCampaignStatusStatusDarchive captures enum value "darchive"
	UpdateCampaignStatusStatusDarchive string = "darchive"
	// UpdateCampaignStatusStatusSent captures enum value "sent"
	UpdateCampaignStatusStatusSent string = "sent"
	// UpdateCampaignStatusStatusQueued captures enum value "queued"
	UpdateCampaignStatusStatusQueued string = "queued"
	// UpdateCampaignStatusStatusReplicate captures enum value "replicate"
	UpdateCampaignStatusStatusReplicate string = "replicate"
	// UpdateCampaignStatusStatusReplicateTemplate captures enum value "replicateTemplate"
	UpdateCampaignStatusStatusReplicateTemplate string = "replicateTemplate"
	// UpdateCampaignStatusStatusDraft captures enum value "draft"
	UpdateCampaignStatusStatusDraft string = "draft"
)

// prop value enum
func (m *UpdateCampaignStatus) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, updateCampaignStatusTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UpdateCampaignStatus) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateCampaignStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateCampaignStatus) UnmarshalBinary(b []byte) error {
	var res UpdateCampaignStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}