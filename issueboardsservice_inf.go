// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// IssueBoardsService is an interface for [gitlab.Client.Boards]
type IssueBoardsService interface {
	// CreateIssueBoard creates a new issue board.
	//
	// GitLab API docs: https://docs.gitlab.com/ee/api/boards.html#create-an-issue-board
	CreateIssueBoard(pid interface{}, opt *CreateIssueBoardOptions, options ...RequestOptionFunc) (*IssueBoard, *Response, error)
	// UpdateIssueBoard update an issue board.
	//
	// GitLab API docs: https://docs.gitlab.com/ee/api/boards.html#update-an-issue-board
	UpdateIssueBoard(pid interface{}, board int, opt *UpdateIssueBoardOptions, options ...RequestOptionFunc) (*IssueBoard, *Response, error)
	// DeleteIssueBoard deletes an issue board.
	//
	// GitLab API docs: https://docs.gitlab.com/ee/api/boards.html#delete-an-issue-board
	DeleteIssueBoard(pid interface{}, board int, options ...RequestOptionFunc) (*Response, error)
	// ListIssueBoards gets a list of all issue boards in a project.
	//
	// GitLab API docs: https://docs.gitlab.com/ee/api/boards.html#list-project-issue-boards
	ListIssueBoards(pid interface{}, opt *ListIssueBoardsOptions, options ...RequestOptionFunc) ([]*IssueBoard, *Response, error)
	// GetIssueBoard gets a single issue board of a project.
	//
	// GitLab API docs: https://docs.gitlab.com/ee/api/boards.html#show-a-single-issue-board
	GetIssueBoard(pid interface{}, board int, options ...RequestOptionFunc) (*IssueBoard, *Response, error)
	// GetIssueBoardLists gets a list of the issue board's lists. Does not include
	// backlog and closed lists.
	//
	// GitLab API docs: https://docs.gitlab.com/ee/api/boards.html#list-board-lists-in-a-project-issue-board
	GetIssueBoardLists(pid interface{}, board int, opt *GetIssueBoardListsOptions, options ...RequestOptionFunc) ([]*BoardList, *Response, error)
	// GetIssueBoardList gets a single issue board list.
	//
	// GitLab API docs: https://docs.gitlab.com/ee/api/boards.html#show-a-single-board-list
	GetIssueBoardList(pid interface{}, board, list int, options ...RequestOptionFunc) (*BoardList, *Response, error)
	// CreateIssueBoardList creates a new issue board list.
	//
	// GitLab API docs: https://docs.gitlab.com/ee/api/boards.html#create-a-board-list
	CreateIssueBoardList(pid interface{}, board int, opt *CreateIssueBoardListOptions, options ...RequestOptionFunc) (*BoardList, *Response, error)
	// UpdateIssueBoardList updates the position of an existing issue board list.
	//
	// GitLab API docs: https://docs.gitlab.com/ee/api/boards.html#reorder-a-list-in-a-board
	UpdateIssueBoardList(pid interface{}, board, list int, opt *UpdateIssueBoardListOptions, options ...RequestOptionFunc) (*BoardList, *Response, error)
	// DeleteIssueBoardList soft deletes an issue board list. Only for admins and
	// project owners.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/boards.html#delete-a-board-list-from-a-board
	DeleteIssueBoardList(pid interface{}, board, list int, options ...RequestOptionFunc) (*Response, error)
}
