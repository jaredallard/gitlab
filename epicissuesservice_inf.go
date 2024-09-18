// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// EpicIssuesService is an interface for [gitlab.Client.EpicIssues]
type EpicIssuesService interface {
	// ListEpicIssues get a list of epic issues.
	//
	// Gitlab API docs:
	// https://docs.gitlab.com/ee/api/epic_issues.html#list-issues-for-an-epic
	ListEpicIssues(gid interface{}, epic int, opt *ListOptions, options ...RequestOptionFunc) ([]*Issue, *Response, error)
	// AssignEpicIssue assigns an existing issue to an epic.
	//
	// Gitlab API Docs:
	// https://docs.gitlab.com/ee/api/epic_issues.html#assign-an-issue-to-the-epic
	AssignEpicIssue(gid interface{}, epic, issue int, options ...RequestOptionFunc) (*EpicIssueAssignment, *Response, error)
	// RemoveEpicIssue removes an issue from an epic.
	//
	// Gitlab API Docs:
	// https://docs.gitlab.com/ee/api/epic_issues.html#remove-an-issue-from-the-epic
	RemoveEpicIssue(gid interface{}, epic, epicIssue int, options ...RequestOptionFunc) (*EpicIssueAssignment, *Response, error)
	// UpdateEpicIssueAssignment moves an issue before or after another issue in an
	// epic issue list.
	//
	// Gitlab API Docs:
	// https://docs.gitlab.com/ee/api/epic_issues.html#update-epic---issue-association
	UpdateEpicIssueAssignment(gid interface{}, epic, epicIssue int, opt *UpdateEpicIsssueAssignmentOptions, options ...RequestOptionFunc) ([]*Issue, *Response, error)
}