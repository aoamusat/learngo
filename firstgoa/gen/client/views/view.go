// Code generated by goa v3.16.1, DO NOT EDIT.
//
// client views
//
// Command:
// $ goa gen learngo.io/firstgoa/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// ClientManagement is the viewed result type that is projected based on a view.
type ClientManagement struct {
	// Type to project
	Projected *ClientManagementView
	// View to render
	View string
}

// ClientManagementCollection is the viewed result type that is projected based
// on a view.
type ClientManagementCollection struct {
	// Type to project
	Projected ClientManagementCollectionView
	// View to render
	View string
}

// ClientManagementView is a type that runs validations on a projected type.
type ClientManagementView struct {
	// ID is the unique id of the Client.
	ClientID *string
	// Name of the Client
	ClientName *string
}

// ClientManagementCollectionView is a type that runs validations on a
// projected type.
type ClientManagementCollectionView []*ClientManagementView

var (
	// ClientManagementMap is a map indexing the attribute names of
	// ClientManagement by view name.
	ClientManagementMap = map[string][]string{
		"default": {
			"ClientID",
			"ClientName",
		},
	}
	// ClientManagementCollectionMap is a map indexing the attribute names of
	// ClientManagementCollection by view name.
	ClientManagementCollectionMap = map[string][]string{
		"default": {
			"ClientID",
			"ClientName",
		},
	}
)

// ValidateClientManagement runs the validations defined on the viewed result
// type ClientManagement.
func ValidateClientManagement(result *ClientManagement) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateClientManagementView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateClientManagementCollection runs the validations defined on the
// viewed result type ClientManagementCollection.
func ValidateClientManagementCollection(result ClientManagementCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateClientManagementCollectionView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateClientManagementView runs the validations defined on
// ClientManagementView using the "default" view.
func ValidateClientManagementView(result *ClientManagementView) (err error) {
	if result.ClientID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ClientID", "result"))
	}
	if result.ClientName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ClientName", "result"))
	}
	return
}

// ValidateClientManagementCollectionView runs the validations defined on
// ClientManagementCollectionView using the "default" view.
func ValidateClientManagementCollectionView(result ClientManagementCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateClientManagementView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}
