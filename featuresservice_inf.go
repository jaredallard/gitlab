// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// FeaturesService is an interface for [gitlab.Client.Features]
type FeaturesService interface {
	// ListFeatures gets a list of feature flags
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/features.html#list-all-features
	ListFeatures(options ...RequestOptionFunc) ([]*Feature, *Response, error)
	// SetFeatureFlag sets or creates a feature flag gate
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/features.html#set-or-create-a-feature
	SetFeatureFlag(name string, value interface{}, options ...RequestOptionFunc) (*Feature, *Response, error)
}