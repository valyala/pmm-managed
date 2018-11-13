// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// InventoryRemoteNode RemoteNode represents a generic remote Node.
// swagger:model inventoryRemoteNode
type InventoryRemoteNode struct {

	// Unique Node identifier.
	ID int64 `json:"id,omitempty"`

	// Unique user-defined Node name.
	Name string `json:"name,omitempty"`
}

// Validate validates this inventory remote node
func (m *InventoryRemoteNode) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InventoryRemoteNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryRemoteNode) UnmarshalBinary(b []byte) error {
	var res InventoryRemoteNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
