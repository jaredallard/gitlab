// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// SnippetsService is an interface for [gitlab.Client.Snippets]
type SnippetsService interface {
	// ListSnippets gets a list of snippets.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/snippets.html#list-all-snippets-for-a-user
	ListSnippets(opt *ListSnippetsOptions, options ...RequestOptionFunc) ([]*Snippet, *Response, error)
	// GetSnippet gets a single snippet
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/snippets.html#get-a-single-snippet
	GetSnippet(snippet int, options ...RequestOptionFunc) (*Snippet, *Response, error)
	// SnippetContent gets a single snippet’s raw contents.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/snippets.html#single-snippet-contents
	SnippetContent(snippet int, options ...RequestOptionFunc) ([]byte, *Response, error)
	// SnippetFileContent returns the raw file content as plain text.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/snippets.html#snippet-repository-file-content
	SnippetFileContent(snippet int, ref, filename string, options ...RequestOptionFunc) ([]byte, *Response, error)
	// CreateSnippet creates a new snippet. The user must have permission
	// to create new snippets.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/snippets.html#create-new-snippet
	CreateSnippet(opt *CreateSnippetOptions, options ...RequestOptionFunc) (*Snippet, *Response, error)
	// UpdateSnippet updates an existing snippet. The user must have
	// permission to change an existing snippet.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/snippets.html#update-snippet
	UpdateSnippet(snippet int, opt *UpdateSnippetOptions, options ...RequestOptionFunc) (*Snippet, *Response, error)
	// DeleteSnippet deletes an existing snippet. This is an idempotent
	// function and deleting a non-existent snippet still returns a 200 OK status
	// code.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/snippets.html#delete-snippet
	DeleteSnippet(snippet int, options ...RequestOptionFunc) (*Response, error)
	// ExploreSnippets gets the list of public snippets.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/snippets.html#list-all-public-snippets
	ExploreSnippets(opt *ExploreSnippetsOptions, options ...RequestOptionFunc) ([]*Snippet, *Response, error)
	// ListAllSnippets gets all snippets the current user has access to.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/snippets.html#list-all-snippets
	ListAllSnippets(opt *ListAllSnippetsOptions, options ...RequestOptionFunc) ([]*Snippet, *Response, error)
}
