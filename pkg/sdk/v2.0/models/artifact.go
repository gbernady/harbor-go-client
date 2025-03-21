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

// Artifact artifact
//
// swagger:model Artifact
type Artifact struct {

	// accessories
	Accessories []*Accessory `json:"accessories"`

	// addition links
	AdditionLinks AdditionLinks `json:"addition_links,omitempty"`

	// annotations
	Annotations Annotations `json:"annotations,omitempty"`

	// The artifact_type in the manifest of the artifact
	ArtifactType string `json:"artifact_type,omitempty"`

	// The digest of the artifact
	Digest string `json:"digest,omitempty"`

	// extra attrs
	ExtraAttrs ExtraAttrs `json:"extra_attrs,omitempty"`

	// The digest of the icon
	Icon string `json:"icon,omitempty"`

	// The ID of the artifact
	ID int64 `json:"id,omitempty"`

	// labels
	Labels []*Label `json:"labels"`

	// The manifest media type of the artifact
	ManifestMediaType string `json:"manifest_media_type,omitempty"`

	// The media type of the artifact
	MediaType string `json:"media_type,omitempty"`

	// The ID of the project that the artifact belongs to
	ProjectID int64 `json:"project_id,omitempty"`

	// The latest pull time of the artifact
	// Format: date-time
	PullTime strfmt.DateTime `json:"pull_time,omitempty"`

	// The push time of the artifact
	// Format: date-time
	PushTime strfmt.DateTime `json:"push_time,omitempty"`

	// references
	References []*Reference `json:"references"`

	// The ID of the repository that the artifact belongs to
	RepositoryID int64 `json:"repository_id,omitempty"`

	// The name of the repository that the artifact belongs to
	RepositoryName string `json:"repository_name,omitempty"`

	// The overview of the generating SBOM progress
	SbomOverview *SBOMOverview `json:"sbom_overview,omitempty"`

	// The overview of the scan result.
	ScanOverview ScanOverview `json:"scan_overview,omitempty"`

	// The size of the artifact
	Size int64 `json:"size,omitempty"`

	// tags
	Tags []*Tag `json:"tags"`

	// The type of the artifact, e.g. image, chart, etc
	Type string `json:"type,omitempty"`
}

// Validate validates this artifact
func (m *Artifact) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdditionLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAnnotations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtraAttrs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePullTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePushTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferences(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSbomOverview(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanOverview(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Artifact) validateAccessories(formats strfmt.Registry) error {
	if swag.IsZero(m.Accessories) { // not required
		return nil
	}

	for i := 0; i < len(m.Accessories); i++ {
		if swag.IsZero(m.Accessories[i]) { // not required
			continue
		}

		if m.Accessories[i] != nil {
			if err := m.Accessories[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("accessories" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("accessories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Artifact) validateAdditionLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.AdditionLinks) { // not required
		return nil
	}

	if m.AdditionLinks != nil {
		if err := m.AdditionLinks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addition_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addition_links")
			}
			return err
		}
	}

	return nil
}

func (m *Artifact) validateAnnotations(formats strfmt.Registry) error {
	if swag.IsZero(m.Annotations) { // not required
		return nil
	}

	if m.Annotations != nil {
		if err := m.Annotations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("annotations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("annotations")
			}
			return err
		}
	}

	return nil
}

func (m *Artifact) validateExtraAttrs(formats strfmt.Registry) error {
	if swag.IsZero(m.ExtraAttrs) { // not required
		return nil
	}

	if m.ExtraAttrs != nil {
		if err := m.ExtraAttrs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extra_attrs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("extra_attrs")
			}
			return err
		}
	}

	return nil
}

func (m *Artifact) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Artifact) validatePullTime(formats strfmt.Registry) error {
	if swag.IsZero(m.PullTime) { // not required
		return nil
	}

	if err := validate.FormatOf("pull_time", "body", "date-time", m.PullTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Artifact) validatePushTime(formats strfmt.Registry) error {
	if swag.IsZero(m.PushTime) { // not required
		return nil
	}

	if err := validate.FormatOf("push_time", "body", "date-time", m.PushTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Artifact) validateReferences(formats strfmt.Registry) error {
	if swag.IsZero(m.References) { // not required
		return nil
	}

	for i := 0; i < len(m.References); i++ {
		if swag.IsZero(m.References[i]) { // not required
			continue
		}

		if m.References[i] != nil {
			if err := m.References[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("references" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("references" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Artifact) validateSbomOverview(formats strfmt.Registry) error {
	if swag.IsZero(m.SbomOverview) { // not required
		return nil
	}

	if m.SbomOverview != nil {
		if err := m.SbomOverview.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sbom_overview")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sbom_overview")
			}
			return err
		}
	}

	return nil
}

func (m *Artifact) validateScanOverview(formats strfmt.Registry) error {
	if swag.IsZero(m.ScanOverview) { // not required
		return nil
	}

	if m.ScanOverview != nil {
		if err := m.ScanOverview.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scan_overview")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scan_overview")
			}
			return err
		}
	}

	return nil
}

func (m *Artifact) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this artifact based on the context it is used
func (m *Artifact) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessories(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAdditionLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAnnotations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExtraAttrs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReferences(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSbomOverview(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScanOverview(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Artifact) contextValidateAccessories(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Accessories); i++ {

		if m.Accessories[i] != nil {
			if err := m.Accessories[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("accessories" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("accessories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Artifact) contextValidateAdditionLinks(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AdditionLinks.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("addition_links")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("addition_links")
		}
		return err
	}

	return nil
}

func (m *Artifact) contextValidateAnnotations(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Annotations.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("annotations")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("annotations")
		}
		return err
	}

	return nil
}

func (m *Artifact) contextValidateExtraAttrs(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ExtraAttrs.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("extra_attrs")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("extra_attrs")
		}
		return err
	}

	return nil
}

func (m *Artifact) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Labels); i++ {

		if m.Labels[i] != nil {
			if err := m.Labels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Artifact) contextValidateReferences(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.References); i++ {

		if m.References[i] != nil {
			if err := m.References[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("references" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("references" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Artifact) contextValidateSbomOverview(ctx context.Context, formats strfmt.Registry) error {

	if m.SbomOverview != nil {
		if err := m.SbomOverview.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sbom_overview")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sbom_overview")
			}
			return err
		}
	}

	return nil
}

func (m *Artifact) contextValidateScanOverview(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ScanOverview.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("scan_overview")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("scan_overview")
		}
		return err
	}

	return nil
}

func (m *Artifact) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Artifact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Artifact) UnmarshalBinary(b []byte) error {
	var res Artifact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
