// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LdapConf The ldap configure properties
//
// swagger:model LdapConf
type LdapConf struct {

	// The base dn of ldap service.
	LdapBaseDn string `json:"ldap_base_dn,omitempty"`

	// The connect timeout of ldap service(second).
	LdapConnectionTimeout int64 `json:"ldap_connection_timeout,omitempty"`

	// The serach filter of ldap service.
	LdapFilter string `json:"ldap_filter,omitempty"`

	// The serach scope of ldap service.
	LdapScope int64 `json:"ldap_scope,omitempty"`

	// The search dn of ldap service.
	LdapSearchDn string `json:"ldap_search_dn,omitempty"`

	// The search password of ldap service.
	LdapSearchPassword string `json:"ldap_search_password,omitempty"`

	// The serach uid from ldap service attributes.
	LdapUID string `json:"ldap_uid,omitempty"`

	// The url of ldap service.
	LdapURL string `json:"ldap_url,omitempty"`

	// Verify Ldap server certificate.
	LdapVerifyCert bool `json:"ldap_verify_cert,omitempty"`
}

// Validate validates this ldap conf
func (m *LdapConf) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LdapConf) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapConf) UnmarshalBinary(b []byte) error {
	var res LdapConf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
