// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// GroupVariablesService is an interface for [gitlab.Client.GroupVariables]
type GroupVariablesService interface {
	// ListVariables gets a list of all variables for a group.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/group_level_variables.html#list-group-variables
	ListVariables(gid interface{}, opt *ListGroupVariablesOptions, options ...RequestOptionFunc) ([]*GroupVariable, *Response, error)
	// GetVariable gets a variable.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/group_level_variables.html#show-variable-details
	GetVariable(gid interface{}, key string, opt *GetGroupVariableOptions, options ...RequestOptionFunc) (*GroupVariable, *Response, error)
	// CreateVariable creates a new group variable.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/group_level_variables.html#create-variable
	CreateVariable(gid interface{}, opt *CreateGroupVariableOptions, options ...RequestOptionFunc) (*GroupVariable, *Response, error)
	// UpdateVariable updates the position of an existing
	// group issue board list.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/group_level_variables.html#update-variable
	UpdateVariable(gid interface{}, key string, opt *UpdateGroupVariableOptions, options ...RequestOptionFunc) (*GroupVariable, *Response, error)
	// RemoveVariable removes a group's variable.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/group_level_variables.html#remove-variable
	RemoveVariable(gid interface{}, key string, options ...RequestOptionFunc) (*Response, error)
}