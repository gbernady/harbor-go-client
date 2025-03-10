// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Robot robot
//
// swagger:model Robot
type Robot struct {

	// The creation time of the robot.
	// Format: date-time
	CreationTime strfmt.DateTime `json:"creation_time,omitempty"`

	// The description of the robot
	Description string `json:"description,omitempty"`

	// The disable status of the robot
	Disable bool `json:"disable"`

	// The duration of the robot in days, duration must be either -1(Never) or a positive integer
	Duration *int64 `json:"duration,omitempty"`

	// The editable status of the robot
	Editable bool `json:"editable"`

	// The expiration date of the robot
	ExpiresAt int64 `json:"expires_at,omitempty"`

	// The ID of the robot
	ID int64 `json:"id,omitempty"`

	// The level of the robot, project or system
	Level string `json:"level,omitempty"`

	// The name of the robot
	Name string `json:"name,omitempty"`

	// permissions
	Permissions []*RobotPermission `json:"permissions"`

	// The secret of the robot
	Secret string `json:"secret,omitempty"`

	// The update time of the robot.
	// Format: date-time
	UpdateTime strfmt.DateTime `json:"update_time,omitempty"`
}

// Validate validates this robot
func (m *Robot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Robot) validateCreationTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationTime) { // not required
		return nil
	}

	if err := validate.FormatOf("creation_time", "body", "date-time", m.CreationTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Robot) validatePermissions(formats strfmt.Registry) error {
	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	for i := 0; i < len(m.Permissions); i++ {
		if swag.IsZero(m.Permissions[i]) { // not required
			continue
		}

		if m.Permissions[i] != nil {
			if err := m.Permissions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("permissions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("permissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Robot) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this robot based on the context it is used
func (m *Robot) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePermissions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Robot) contextValidatePermissions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Permissions); i++ {

		if m.Permissions[i] != nil {
			if err := m.Permissions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("permissions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("permissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Robot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Robot) UnmarshalBinary(b []byte) error {
	var res Robot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
