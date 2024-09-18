// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// DeploymentsService is an interface for [gitlab.Client.Deployments]
type DeploymentsService interface {
	// ListProjectDeployments gets a list of deployments in a project.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/deployments.html#list-project-deployments
	ListProjectDeployments(pid interface{}, opts *ListProjectDeploymentsOptions, options ...RequestOptionFunc) ([]*Deployment, *Response, error)
	// GetProjectDeployment get a deployment for a project.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/deployments.html#get-a-specific-deployment
	GetProjectDeployment(pid interface{}, deployment int, options ...RequestOptionFunc) (*Deployment, *Response, error)
	// CreateProjectDeployment creates a project deployment.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/deployments.html#create-a-deployment
	CreateProjectDeployment(pid interface{}, opt *CreateProjectDeploymentOptions, options ...RequestOptionFunc) (*Deployment, *Response, error)
	// UpdateProjectDeployment updates a project deployment.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/deployments.html#update-a-deployment
	UpdateProjectDeployment(pid interface{}, deployment int, opt *UpdateProjectDeploymentOptions, options ...RequestOptionFunc) (*Deployment, *Response, error)
	// ApproveOrRejectProjectDeployment approve or reject a blocked deployment.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/deployments.html#approve-or-reject-a-blocked-deployment
	ApproveOrRejectProjectDeployment(pid interface{}, deployment int, opt *ApproveOrRejectProjectDeploymentOptions, options ...RequestOptionFunc) (*Response, error)
	// DeleteProjectDeployment delete a project deployment.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/deployments.html#delete-a-specific-deployment
	DeleteProjectDeployment(pid interface{}, deployment int, options ...RequestOptionFunc) (*Response, error)
}