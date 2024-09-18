// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// ResourceLabelEventsService is an interface for [gitlab.Client.ResourceLabelEvents]
type ResourceLabelEventsService interface {
	// ListIssueLabelEvents retrieves resource label events for the
	// specified project and issue.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/resource_label_events.html#list-project-issue-label-events
	ListIssueLabelEvents(pid interface{}, issue int, opt *ListLabelEventsOptions, options ...RequestOptionFunc) ([]*LabelEvent, *Response, error)
	// GetIssueLabelEvent gets a single issue-label-event.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/resource_label_events.html#get-single-issue-label-event
	GetIssueLabelEvent(pid interface{}, issue int, event int, options ...RequestOptionFunc) (*LabelEvent, *Response, error)
	// ListGroupEpicLabelEvents retrieves resource label events for the specified
	// group and epic.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/resource_label_events.html#list-group-epic-label-events
	ListGroupEpicLabelEvents(gid interface{}, epic int, opt *ListLabelEventsOptions, options ...RequestOptionFunc) ([]*LabelEvent, *Response, error)
	// GetGroupEpicLabelEvent gets a single group epic label event.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/resource_label_events.html#get-single-epic-label-event
	GetGroupEpicLabelEvent(gid interface{}, epic int, event int, options ...RequestOptionFunc) (*LabelEvent, *Response, error)
	// ListMergeRequestsLabelEvents retrieves resource label events for the specified
	// project and merge request.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/resource_label_events.html#list-project-merge-request-label-events
	ListMergeRequestsLabelEvents(pid interface{}, request int, opt *ListLabelEventsOptions, options ...RequestOptionFunc) ([]*LabelEvent, *Response, error)
	// GetMergeRequestLabelEvent gets a single merge request label event.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/resource_label_events.html#get-single-merge-request-label-event
	GetMergeRequestLabelEvent(pid interface{}, request int, event int, options ...RequestOptionFunc) (*LabelEvent, *Response, error)
}
