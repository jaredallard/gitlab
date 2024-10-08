// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// SettingsService is an interface for [gitlab.Client.Settings]
type SettingsService interface {
	// GetSettings gets the current application settings.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/settings.html#get-current-application-settings
	GetSettings(options ...RequestOptionFunc) (*Settings, *Response, error)
	// UpdateSettings updates the application settings.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/settings.html#change-application-settings
	UpdateSettings(opt *UpdateSettingsOptions, options ...RequestOptionFunc) (*Settings, *Response, error)
}
