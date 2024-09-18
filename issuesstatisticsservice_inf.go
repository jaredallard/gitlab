// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// IssuesStatisticsService is an interface for [gitlab.Client.IssuesStatistics]
type IssuesStatisticsService interface {
	// GetIssuesStatistics gets issues statistics on all issues the authenticated
	// user has access to.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/issues_statistics.html#get-issues-statistics
	GetIssuesStatistics(opt *GetIssuesStatisticsOptions, options ...RequestOptionFunc) (*IssuesStatistics, *Response, error)
	// GetGroupIssuesStatistics gets issues count statistics for given group.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/issues_statistics.html#get-group-issues-statistics
	GetGroupIssuesStatistics(gid interface{}, opt *GetGroupIssuesStatisticsOptions, options ...RequestOptionFunc) (*IssuesStatistics, *Response, error)
	// GetProjectIssuesStatistics gets issues count statistics for given project.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/issues_statistics.html#get-project-issues-statistics
	GetProjectIssuesStatistics(pid interface{}, opt *GetProjectIssuesStatisticsOptions, options ...RequestOptionFunc) (*IssuesStatistics, *Response, error)
}
