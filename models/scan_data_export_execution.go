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

// ScanDataExportExecution The replication execution
//
// swagger:model ScanDataExportExecution
type ScanDataExportExecution struct {

	// The end time
	// Format: date-time
	EndTime strfmt.DateTime `json:"end_time,omitempty"`

	// Indicates whether the export artifact is present in registry
	FilePresent bool `json:"file_present"`

	// The ID of the execution
	ID int64 `json:"id,omitempty"`

	// The start time
	// Format: date-time
	StartTime strfmt.DateTime `json:"start_time,omitempty"`

	// The status of the execution
	Status string `json:"status,omitempty"`

	// The status text
	StatusText string `json:"status_text"`

	// The trigger mode
	Trigger string `json:"trigger,omitempty"`

	// The ID if the user triggering the export job
	UserID int64 `json:"user_id,omitempty"`

	// The name of the user triggering the job
	UserName string `json:"user_name"`
}

// Validate validates this scan data export execution
func (m *ScanDataExportExecution) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScanDataExportExecution) validateEndTime(formats strfmt.Registry) error {

	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("end_time", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ScanDataExportExecution) validateStartTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("start_time", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScanDataExportExecution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScanDataExportExecution) UnmarshalBinary(b []byte) error {
	var res ScanDataExportExecution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
