// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// GroupBadgesService is an interface for [gitlab.Client.GroupBadges]
type GroupBadgesService interface {
	// ListGroupBadges gets a list of a group badges.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/group_badges.html#list-all-badges-of-a-group
	ListGroupBadges(gid interface{}, opt *ListGroupBadgesOptions, options ...RequestOptionFunc) ([]*GroupBadge, *Response, error)
	// GetGroupBadge gets a group badge.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/group_badges.html#get-a-badge-of-a-group
	GetGroupBadge(gid interface{}, badge int, options ...RequestOptionFunc) (*GroupBadge, *Response, error)
	// AddGroupBadge adds a badge to a group.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/group_badges.html#add-a-badge-to-a-group
	AddGroupBadge(gid interface{}, opt *AddGroupBadgeOptions, options ...RequestOptionFunc) (*GroupBadge, *Response, error)
	// EditGroupBadge updates a badge of a group.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/group_badges.html#edit-a-badge-of-a-group
	EditGroupBadge(gid interface{}, badge int, opt *EditGroupBadgeOptions, options ...RequestOptionFunc) (*GroupBadge, *Response, error)
	// DeleteGroupBadge removes a badge from a group.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/group_badges.html#remove-a-badge-from-a-group
	DeleteGroupBadge(gid interface{}, badge int, options ...RequestOptionFunc) (*Response, error)
	// PreviewGroupBadge returns how the link_url and image_url final URLs would be after
	// resolving the placeholder interpolation.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/group_badges.html#preview-a-badge-from-a-group
	PreviewGroupBadge(gid interface{}, opt *GroupBadgePreviewOptions, options ...RequestOptionFunc) (*GroupBadge, *Response, error)
}