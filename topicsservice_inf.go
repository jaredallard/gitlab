// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// TopicsService is an interface for [gitlab.Client.Topics]
type TopicsService interface {
	// ListTopics returns a list of project topics in the GitLab instance ordered
	// by number of associated projects.
	//
	// GitLab API docs: https://docs.gitlab.com/ee/api/topics.html#list-topics
	ListTopics(opt *ListTopicsOptions, options ...RequestOptionFunc) ([]*Topic, *Response, error)
	// GetTopic gets a project topic by ID.
	//
	// GitLab API docs: https://docs.gitlab.com/ee/api/topics.html#get-a-topic
	GetTopic(topic int, options ...RequestOptionFunc) (*Topic, *Response, error)
	// CreateTopic creates a new project topic.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/topics.html#create-a-project-topic
	CreateTopic(opt *CreateTopicOptions, options ...RequestOptionFunc) (*Topic, *Response, error)
	// UpdateTopic updates a project topic. Only available to administrators.
	//
	// To remove a topic avatar set the TopicAvatar.Filename to an empty string
	// and set TopicAvatar.Image to nil.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/topics.html#update-a-project-topic
	UpdateTopic(topic int, opt *UpdateTopicOptions, options ...RequestOptionFunc) (*Topic, *Response, error)
	// DeleteTopic deletes a project topic. Only available to administrators.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/topics.html#delete-a-project-topic
	DeleteTopic(topic int, options ...RequestOptionFunc) (*Response, error)
}
