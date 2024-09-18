// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// ProjectVariablesService is an interface for [gitlab.Client.ProjectVariables]
type ProjectVariablesService interface {
	// ListVariables gets a list of all variables in a project.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/project_level_variables.html#list-project-variables
	ListVariables(pid interface{}, opt *ListProjectVariablesOptions, options ...RequestOptionFunc) ([]*ProjectVariable, *Response, error)
	// GetVariable gets a variable.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/project_level_variables.html#get-a-single-variable
	GetVariable(pid interface{}, key string, opt *GetProjectVariableOptions, options ...RequestOptionFunc) (*ProjectVariable, *Response, error)
	// CreateVariable creates a new project variable.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/project_level_variables.html#create-a-variable
	CreateVariable(pid interface{}, opt *CreateProjectVariableOptions, options ...RequestOptionFunc) (*ProjectVariable, *Response, error)
	// UpdateVariable updates a project's variable.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/project_level_variables.html#update-a-variable
	UpdateVariable(pid interface{}, key string, opt *UpdateProjectVariableOptions, options ...RequestOptionFunc) (*ProjectVariable, *Response, error)
	// RemoveVariable removes a project's variable.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/project_level_variables.html#delete-a-variable
	RemoveVariable(pid interface{}, key string, opt *RemoveProjectVariableOptions, options ...RequestOptionFunc) (*Response, error)
}