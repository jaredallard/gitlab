// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// GroupRepositoryStorageMoveService is an interface for [gitlab.Client.GroupRepositoryStorageMove]
type GroupRepositoryStorageMoveService interface {
	// RetrieveAllStorageMoves retrieves all group repository storage moves
	// accessible by the authenticated user.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/group_repository_storage_moves.html#retrieve-all-group-repository-storage-moves
	RetrieveAllStorageMoves(opts RetrieveAllGroupStorageMovesOptions, options ...RequestOptionFunc) ([]*GroupRepositoryStorageMove, *Response, error)
	// RetrieveAllStorageMovesForGroup retrieves all repository storage moves for
	// a single group accessible by the authenticated user.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/group_repository_storage_moves.html#retrieve-all-repository-storage-moves-for-a-single-group
	RetrieveAllStorageMovesForGroup(group int, opts RetrieveAllGroupStorageMovesOptions, options ...RequestOptionFunc) ([]*GroupRepositoryStorageMove, *Response, error)
	// GetStorageMove gets a single group repository storage move.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/group_repository_storage_moves.html#get-a-single-group-repository-storage-move
	GetStorageMove(repositoryStorage int, options ...RequestOptionFunc) (*GroupRepositoryStorageMove, *Response, error)
	// GetStorageMoveForGroup gets a single repository storage move for a group.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/group_repository_storage_moves.html#get-a-single-repository-storage-move-for-a-group
	GetStorageMoveForGroup(group int, repositoryStorage int, options ...RequestOptionFunc) (*GroupRepositoryStorageMove, *Response, error)
	// ScheduleStorageMoveForGroup schedule a repository to be moved for a group.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/group_repository_storage_moves.html#schedule-a-repository-storage-move-for-a-group
	ScheduleStorageMoveForGroup(group int, opts ScheduleStorageMoveForGroupOptions, options ...RequestOptionFunc) (*GroupRepositoryStorageMove, *Response, error)
	// ScheduleAllStorageMoves schedules all group repositories to be moved.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/group_repository_storage_moves.html#schedule-repository-storage-moves-for-all-groups-on-a-storage-shard
	ScheduleAllStorageMoves(opts ScheduleAllGroupStorageMovesOptions, options ...RequestOptionFunc) (*Response, error)
}
