// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// EventsService is an interface for [gitlab.Client.Events]
type EventsService interface {
	// ListCurrentUserContributionEvents gets a list currently authenticated user's events
	//
	// GitLab API docs: https://docs.gitlab.com/ee/api/events.html#list-currently-authenticated-users-events
	ListCurrentUserContributionEvents(opt *ListContributionEventsOptions, options ...RequestOptionFunc) ([]*ContributionEvent, *Response, error)
	// ListProjectVisibleEvents gets the events for the specified project.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/events.html#list-a-projects-visible-events
	ListProjectVisibleEvents(pid interface{}, opt *ListProjectVisibleEventsOptions, options ...RequestOptionFunc) ([]*ProjectEvent, *Response, error)
}
