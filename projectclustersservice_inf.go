// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// ProjectClustersService is an interface for [gitlab.Client.ProjectCluster]
type ProjectClustersService interface {
	// ListClusters gets a list of all clusters in a project.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/project_clusters.html#list-project-clusters
	ListClusters(pid interface{}, options ...RequestOptionFunc) ([]*ProjectCluster, *Response, error)
	// GetCluster gets a cluster.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/project_clusters.html#get-a-single-project-cluster
	GetCluster(pid interface{}, cluster int, options ...RequestOptionFunc) (*ProjectCluster, *Response, error)
	// AddCluster adds an existing cluster to the project.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/project_clusters.html#add-existing-cluster-to-project
	AddCluster(pid interface{}, opt *AddClusterOptions, options ...RequestOptionFunc) (*ProjectCluster, *Response, error)
	// EditCluster updates an existing project cluster.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/project_clusters.html#edit-project-cluster
	EditCluster(pid interface{}, cluster int, opt *EditClusterOptions, options ...RequestOptionFunc) (*ProjectCluster, *Response, error)
	// DeleteCluster deletes an existing project cluster.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/project_clusters.html#delete-project-cluster
	DeleteCluster(pid interface{}, cluster int, options ...RequestOptionFunc) (*Response, error)
}
