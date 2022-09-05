// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ScanDataExportJob The metadata associated with the scan data export job
//
// swagger:model ScanDataExportJob
type ScanDataExportJob struct {

	// The id of the scan data export job
	ID int64 `json:"id,omitempty"`
}

// Validate validates this scan data export job
func (m *ScanDataExportJob) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ScanDataExportJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScanDataExportJob) UnmarshalBinary(b []byte) error {
	var res ScanDataExportJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
