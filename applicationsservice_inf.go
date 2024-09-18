// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// ApplicationsService is an interface for [gitlab.Client.Applications]
type ApplicationsService interface {
	// CreateApplication creates a new application owned by the authenticated user.
	//
	// Gitlab API docs : https://docs.gitlab.com/ee/api/applications.html#create-an-application
	CreateApplication(opt *CreateApplicationOptions, options ...RequestOptionFunc) (*Application, *Response, error)
	// ListApplications get a list of administrables applications by the authenticated user
	//
	// Gitlab API docs : https://docs.gitlab.com/ee/api/applications.html#list-all-applications
	ListApplications(opt *ListApplicationsOptions, options ...RequestOptionFunc) ([]*Application, *Response, error)
	// DeleteApplication removes a specific application.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/applications.html#delete-an-application
	DeleteApplication(application int, options ...RequestOptionFunc) (*Response, error)
}