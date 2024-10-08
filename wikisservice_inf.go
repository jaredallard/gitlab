// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// WikisService is an interface for [gitlab.Client.Wikis]
type WikisService interface {
	// ListWikis lists all pages of the wiki of the given project id.
	// When with_content is set, it also returns the content of the pages.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/wikis.html#list-wiki-pages
	ListWikis(pid interface{}, opt *ListWikisOptions, options ...RequestOptionFunc) ([]*Wiki, *Response, error)
	// GetWikiPage gets a wiki page for a given project.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/wikis.html#get-a-wiki-page
	GetWikiPage(pid interface{}, slug string, opt *GetWikiPageOptions, options ...RequestOptionFunc) (*Wiki, *Response, error)
	// CreateWikiPage creates a new wiki page for the given repository with
	// the given title, slug, and content.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/wikis.html#create-a-new-wiki-page
	CreateWikiPage(pid interface{}, opt *CreateWikiPageOptions, options ...RequestOptionFunc) (*Wiki, *Response, error)
	// EditWikiPage Updates an existing wiki page. At least one parameter is
	// required to update the wiki page.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/wikis.html#edit-an-existing-wiki-page
	EditWikiPage(pid interface{}, slug string, opt *EditWikiPageOptions, options ...RequestOptionFunc) (*Wiki, *Response, error)
	// DeleteWikiPage deletes a wiki page with a given slug.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/wikis.html#delete-a-wiki-page
	DeleteWikiPage(pid interface{}, slug string, options ...RequestOptionFunc) (*Response, error)
}
