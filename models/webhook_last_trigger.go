// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WebhookLastTrigger The webhook policy and last trigger time group by event type.
//
// swagger:model WebhookLastTrigger
type WebhookLastTrigger struct {

	// The creation time of webhook policy.
	// Format: date-time
	CreationTime strfmt.DateTime `json:"creation_time,omitempty"`

	// Whether or not the webhook policy enabled.
	Enabled bool `json:"enabled,omitempty"`

	// The webhook event type.
	EventType string `json:"event_type,omitempty"`

	// The last trigger time of webhook policy.
	// Format: date-time
	LastTriggerTime strfmt.DateTime `json:"last_trigger_time,omitempty"`

	// The webhook policy name.
	PolicyName string `json:"policy_name,omitempty"`
}

// Validate validates this webhook last trigger
func (m *WebhookLastTrigger) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastTriggerTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebhookLastTrigger) validateCreationTime(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationTime) { // not required
		return nil
	}

	if err := validate.FormatOf("creation_time", "body", "date-time", m.CreationTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WebhookLastTrigger) validateLastTriggerTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LastTriggerTime) { // not required
		return nil
	}

	if err := validate.FormatOf("last_trigger_time", "body", "date-time", m.LastTriggerTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebhookLastTrigger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebhookLastTrigger) UnmarshalBinary(b []byte) error {
	var res WebhookLastTrigger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
