// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// ImportService is an interface for [gitlab.Client.Import]
type ImportService interface {
	// Import a repository from GitHub.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/import.html#import-repository-from-github
	ImportRepositoryFromGitHub(opt *ImportRepositoryFromGitHubOptions, options ...RequestOptionFunc) (*GitHubImport, *Response, error)
	// Cancel an import of a repository from GitHub.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/import.html#cancel-github-project-import
	CancelGitHubProjectImport(opt *CancelGitHubProjectImportOptions, options ...RequestOptionFunc) (*CancelledGitHubImport, *Response, error)
	// Import personal GitHub Gists into personal GitLab Snippets.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/import.html#import-github-gists-into-gitlab-snippets
	ImportGitHubGistsIntoGitLabSnippets(opt *ImportGitHubGistsIntoGitLabSnippetsOptions, options ...RequestOptionFunc) (*Response, error)
	// Import a repository from Bitbucket Server.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/import.html#import-repository-from-bitbucket-server
	ImportRepositoryFromBitbucketServer(opt *ImportRepositoryFromBitbucketServerOptions, options ...RequestOptionFunc) (*BitbucketServerImport, *Response, error)
	// Import a repository from Bitbucket Cloud.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/import.html#import-repository-from-bitbucket-cloud
	ImportRepositoryFromBitbucketCloud(opt *ImportRepositoryFromBitbucketCloudOptions, options ...RequestOptionFunc) (*BitbucketCloudImport, *Response, error)
}
