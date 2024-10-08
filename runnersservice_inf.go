// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// RunnersService is an interface for [gitlab.Client.Runners]
type RunnersService interface {
	// ListRunners gets a list of runners accessible by the authenticated user.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/runners.html#list-owned-runners
	ListRunners(opt *ListRunnersOptions, options ...RequestOptionFunc) ([]*Runner, *Response, error)
	// ListAllRunners gets a list of all runners in the GitLab instance. Access is
	// restricted to users with admin privileges.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/runners.html#list-all-runners
	ListAllRunners(opt *ListRunnersOptions, options ...RequestOptionFunc) ([]*Runner, *Response, error)
	// GetRunnerDetails returns details for given runner.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/runners.html#get-runners-details
	GetRunnerDetails(rid interface{}, options ...RequestOptionFunc) (*RunnerDetails, *Response, error)
	// UpdateRunnerDetails updates details for a given runner.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/runners.html#update-runners-details
	UpdateRunnerDetails(rid interface{}, opt *UpdateRunnerDetailsOptions, options ...RequestOptionFunc) (*RunnerDetails, *Response, error)
	// RemoveRunner removes a runner.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/runners.html#delete-a-runner
	RemoveRunner(rid interface{}, options ...RequestOptionFunc) (*Response, error)
	// ListRunnerJobs gets a list of jobs that are being processed or were processed by specified Runner.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/runners.html#list-runners-jobs
	ListRunnerJobs(rid interface{}, opt *ListRunnerJobsOptions, options ...RequestOptionFunc) ([]*Job, *Response, error)
	// ListProjectRunners gets a list of runners accessible by the authenticated user.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/runners.html#list-projects-runners
	ListProjectRunners(pid interface{}, opt *ListProjectRunnersOptions, options ...RequestOptionFunc) ([]*Runner, *Response, error)
	// EnableProjectRunner enables an available specific runner in the project.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/runners.html#enable-a-runner-in-project
	EnableProjectRunner(pid interface{}, opt *EnableProjectRunnerOptions, options ...RequestOptionFunc) (*Runner, *Response, error)
	// DisableProjectRunner disables a specific runner from project.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/runners.html#disable-a-runner-from-project
	DisableProjectRunner(pid interface{}, runner int, options ...RequestOptionFunc) (*Response, error)
	// ListGroupsRunners lists all runners (specific and shared) available in the
	// group as well it’s ancestor groups. Shared runners are listed if at least one
	// shared runner is defined.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/runners.html#list-groups-runners
	ListGroupsRunners(gid interface{}, opt *ListGroupsRunnersOptions, options ...RequestOptionFunc) ([]*Runner, *Response, error)
	// RegisterNewRunner registers a new Runner for the instance.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/runners.html#register-a-new-runner
	RegisterNewRunner(opt *RegisterNewRunnerOptions, options ...RequestOptionFunc) (*Runner, *Response, error)
	// DeleteRegisteredRunner deletes a Runner by Token.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/runners.html#delete-a-runner-by-authentication-token
	DeleteRegisteredRunner(opt *DeleteRegisteredRunnerOptions, options ...RequestOptionFunc) (*Response, error)
	// DeleteRegisteredRunnerByID deletes a runner by ID.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/runners.html#delete-a-runner-by-id
	DeleteRegisteredRunnerByID(rid int, options ...RequestOptionFunc) (*Response, error)
	// VerifyRegisteredRunner registers a new runner for the instance.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/runners.html#verify-authentication-for-a-registered-runner
	VerifyRegisteredRunner(opt *VerifyRegisteredRunnerOptions, options ...RequestOptionFunc) (*Response, error)
	// ResetInstanceRunnerRegistrationToken resets the instance runner registration
	// token.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/runners.html#reset-instances-runner-registration-token
	ResetInstanceRunnerRegistrationToken(options ...RequestOptionFunc) (*RunnerRegistrationToken, *Response, error)
	// ResetGroupRunnerRegistrationToken resets a group's runner registration token.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/runners.html#reset-groups-runner-registration-token
	ResetGroupRunnerRegistrationToken(gid interface{}, options ...RequestOptionFunc) (*RunnerRegistrationToken, *Response, error)
	// ResetGroupRunnerRegistrationToken resets a projects's runner registration token.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/runners.html#reset-projects-runner-registration-token
	ResetProjectRunnerRegistrationToken(pid interface{}, options ...RequestOptionFunc) (*RunnerRegistrationToken, *Response, error)
	// ResetRunnerAuthenticationToken resets a runner's authentication token.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/runners.html#reset-runners-authentication-token-by-using-the-runner-id
	ResetRunnerAuthenticationToken(rid int, options ...RequestOptionFunc) (*RunnerAuthenticationToken, *Response, error)
}
