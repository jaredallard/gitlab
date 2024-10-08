// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// MetadataService is an interface for [gitlab.Client.Metadata]
type MetadataService interface {
	// GetMetadata gets a GitLab server instance meteadata.
	//
	// GitLab API docs: https://docs.gitlab.com/ee/api/metadata.html
	GetMetadata(options ...RequestOptionFunc) (*Metadata, *Response, error)
}
