// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// DiscussionsService is an interface for [gitlab.Client.Discussions]
type DiscussionsService interface {
	// ListIssueDiscussions gets a list of all discussions for a single
	// issue.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#list-project-issue-discussion-items
	ListIssueDiscussions(pid interface{}, issue int, opt *ListIssueDiscussionsOptions, options ...RequestOptionFunc) ([]*Discussion, *Response, error)
	// GetIssueDiscussion returns a single discussion for a specific project issue.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#get-single-issue-discussion-item
	GetIssueDiscussion(pid interface{}, issue int, discussion string, options ...RequestOptionFunc) (*Discussion, *Response, error)
	// CreateIssueDiscussion creates a new discussion to a single project issue.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#create-new-issue-thread
	CreateIssueDiscussion(pid interface{}, issue int, opt *CreateIssueDiscussionOptions, options ...RequestOptionFunc) (*Discussion, *Response, error)
	// AddIssueDiscussionNote creates a new discussion to a single project issue.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#add-note-to-existing-issue-thread
	AddIssueDiscussionNote(pid interface{}, issue int, discussion string, opt *AddIssueDiscussionNoteOptions, options ...RequestOptionFunc) (*Note, *Response, error)
	// UpdateIssueDiscussionNote modifies existing discussion of an issue.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#modify-existing-issue-thread-note
	UpdateIssueDiscussionNote(pid interface{}, issue int, discussion string, note int, opt *UpdateIssueDiscussionNoteOptions, options ...RequestOptionFunc) (*Note, *Response, error)
	// DeleteIssueDiscussionNote deletes an existing discussion of an issue.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#delete-an-issue-thread-note
	DeleteIssueDiscussionNote(pid interface{}, issue int, discussion string, note int, options ...RequestOptionFunc) (*Response, error)
	// ListSnippetDiscussions gets a list of all discussions for a single
	// snippet. Snippet discussions are comments users can post to a snippet.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#list-project-snippet-discussion-items
	ListSnippetDiscussions(pid interface{}, snippet int, opt *ListSnippetDiscussionsOptions, options ...RequestOptionFunc) ([]*Discussion, *Response, error)
	// GetSnippetDiscussion returns a single discussion for a given snippet.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#get-single-snippet-discussion-item
	GetSnippetDiscussion(pid interface{}, snippet int, discussion string, options ...RequestOptionFunc) (*Discussion, *Response, error)
	// CreateSnippetDiscussion creates a new discussion for a single snippet.
	// Snippet discussions are comments users can post to a snippet.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#create-new-snippet-thread
	CreateSnippetDiscussion(pid interface{}, snippet int, opt *CreateSnippetDiscussionOptions, options ...RequestOptionFunc) (*Discussion, *Response, error)
	// AddSnippetDiscussionNote creates a new discussion to a single project
	// snippet.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#add-note-to-existing-snippet-thread
	AddSnippetDiscussionNote(pid interface{}, snippet int, discussion string, opt *AddSnippetDiscussionNoteOptions, options ...RequestOptionFunc) (*Note, *Response, error)
	// UpdateSnippetDiscussionNote modifies existing discussion of a snippet.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#modify-existing-snippet-thread-note
	UpdateSnippetDiscussionNote(pid interface{}, snippet int, discussion string, note int, opt *UpdateSnippetDiscussionNoteOptions, options ...RequestOptionFunc) (*Note, *Response, error)
	// DeleteSnippetDiscussionNote deletes an existing discussion of a snippet.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#delete-a-snippet-thread-note
	DeleteSnippetDiscussionNote(pid interface{}, snippet int, discussion string, note int, options ...RequestOptionFunc) (*Response, error)
	// ListGroupEpicDiscussions gets a list of all discussions for a single
	// epic. Epic discussions are comments users can post to a epic.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#list-group-epic-discussion-items
	ListGroupEpicDiscussions(gid interface{}, epic int, opt *ListGroupEpicDiscussionsOptions, options ...RequestOptionFunc) ([]*Discussion, *Response, error)
	// GetEpicDiscussion returns a single discussion for a given epic.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#get-single-epic-discussion-item
	GetEpicDiscussion(gid interface{}, epic int, discussion string, options ...RequestOptionFunc) (*Discussion, *Response, error)
	// CreateEpicDiscussion creates a new discussion for a single epic. Epic
	// discussions are comments users can post to a epic.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#create-new-epic-thread
	CreateEpicDiscussion(gid interface{}, epic int, opt *CreateEpicDiscussionOptions, options ...RequestOptionFunc) (*Discussion, *Response, error)
	// AddEpicDiscussionNote creates a new discussion to a single project epic.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#add-note-to-existing-epic-thread
	AddEpicDiscussionNote(gid interface{}, epic int, discussion string, opt *AddEpicDiscussionNoteOptions, options ...RequestOptionFunc) (*Note, *Response, error)
	// UpdateEpicDiscussionNote modifies existing discussion of a epic.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#modify-existing-epic-thread-note
	UpdateEpicDiscussionNote(gid interface{}, epic int, discussion string, note int, opt *UpdateEpicDiscussionNoteOptions, options ...RequestOptionFunc) (*Note, *Response, error)
	// DeleteEpicDiscussionNote deletes an existing discussion of a epic.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#delete-an-epic-thread-note
	DeleteEpicDiscussionNote(gid interface{}, epic int, discussion string, note int, options ...RequestOptionFunc) (*Response, error)
	// ListMergeRequestDiscussions gets a list of all discussions for a single
	// merge request.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#list-project-merge-request-discussion-items
	ListMergeRequestDiscussions(pid interface{}, mergeRequest int, opt *ListMergeRequestDiscussionsOptions, options ...RequestOptionFunc) ([]*Discussion, *Response, error)
	// GetMergeRequestDiscussion returns a single discussion for a given merge
	// request.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#get-single-merge-request-discussion-item
	GetMergeRequestDiscussion(pid interface{}, mergeRequest int, discussion string, options ...RequestOptionFunc) (*Discussion, *Response, error)
	// CreateMergeRequestDiscussion creates a new discussion for a single merge
	// request.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#create-new-merge-request-thread
	CreateMergeRequestDiscussion(pid interface{}, mergeRequest int, opt *CreateMergeRequestDiscussionOptions, options ...RequestOptionFunc) (*Discussion, *Response, error)
	// ResolveMergeRequestDiscussion resolves/unresolves whole discussion of a merge
	// request.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#resolve-a-merge-request-thread
	ResolveMergeRequestDiscussion(pid interface{}, mergeRequest int, discussion string, opt *ResolveMergeRequestDiscussionOptions, options ...RequestOptionFunc) (*Discussion, *Response, error)
	// AddMergeRequestDiscussionNote creates a new discussion to a single project
	// merge request.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#add-note-to-existing-merge-request-thread
	AddMergeRequestDiscussionNote(pid interface{}, mergeRequest int, discussion string, opt *AddMergeRequestDiscussionNoteOptions, options ...RequestOptionFunc) (*Note, *Response, error)
	// UpdateMergeRequestDiscussionNote modifies existing discussion of a merge
	// request.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#modify-an-existing-merge-request-thread-note
	UpdateMergeRequestDiscussionNote(pid interface{}, mergeRequest int, discussion string, note int, opt *UpdateMergeRequestDiscussionNoteOptions, options ...RequestOptionFunc) (*Note, *Response, error)
	// DeleteMergeRequestDiscussionNote deletes an existing discussion of a merge
	// request.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#delete-a-merge-request-thread-note
	DeleteMergeRequestDiscussionNote(pid interface{}, mergeRequest int, discussion string, note int, options ...RequestOptionFunc) (*Response, error)
	// ListCommitDiscussions gets a list of all discussions for a single
	// commit.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#list-project-commit-discussion-items
	ListCommitDiscussions(pid interface{}, commit string, opt *ListCommitDiscussionsOptions, options ...RequestOptionFunc) ([]*Discussion, *Response, error)
	// GetCommitDiscussion returns a single discussion for a specific project
	// commit.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#get-single-commit-discussion-item
	GetCommitDiscussion(pid interface{}, commit string, discussion string, options ...RequestOptionFunc) (*Discussion, *Response, error)
	// CreateCommitDiscussion creates a new discussion to a single project commit.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#create-new-commit-thread
	CreateCommitDiscussion(pid interface{}, commit string, opt *CreateCommitDiscussionOptions, options ...RequestOptionFunc) (*Discussion, *Response, error)
	// AddCommitDiscussionNote creates a new discussion to a single project commit.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#add-note-to-existing-commit-thread
	AddCommitDiscussionNote(pid interface{}, commit string, discussion string, opt *AddCommitDiscussionNoteOptions, options ...RequestOptionFunc) (*Note, *Response, error)
	// UpdateCommitDiscussionNote modifies existing discussion of an commit.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#modify-an-existing-commit-thread-note
	UpdateCommitDiscussionNote(pid interface{}, commit string, discussion string, note int, opt *UpdateCommitDiscussionNoteOptions, options ...RequestOptionFunc) (*Note, *Response, error)
	// DeleteCommitDiscussionNote deletes an existing discussion of an commit.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/discussions.html#delete-a-commit-thread-note
	DeleteCommitDiscussionNote(pid interface{}, commit string, discussion string, note int, options ...RequestOptionFunc) (*Response, error)
}