// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// ProjectVulnerabilitiesService is an interface for [gitlab.Client.ProjectVulnerabilities]
type ProjectVulnerabilitiesService interface {
	// ListProjectVulnerabilities gets a list of all project vulnerabilities.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/project_vulnerabilities.html#list-project-vulnerabilities
	ListProjectVulnerabilities(pid interface{}, opt *ListProjectVulnerabilitiesOptions, options ...RequestOptionFunc) ([]*ProjectVulnerability, *Response, error)
	// CreateVulnerability creates a new vulnerability on the selected project.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/project_vulnerabilities.html#new-vulnerability
	CreateVulnerability(pid interface{}, opt *CreateVulnerabilityOptions, options ...RequestOptionFunc) (*ProjectVulnerability, *Response, error)
}