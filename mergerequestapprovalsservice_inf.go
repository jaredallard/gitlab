// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// MergeRequestApprovalsService is an interface for [gitlab.Client.MergeRequestApprovals]
type MergeRequestApprovalsService interface {
	// ApproveMergeRequest approves a merge request on GitLab. If a non-empty sha
	// is provided then it must match the sha at the HEAD of the MR.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/merge_request_approvals.html#approve-merge-request
	ApproveMergeRequest(pid interface{}, mr int, opt *ApproveMergeRequestOptions, options ...RequestOptionFunc) (*MergeRequestApprovals, *Response, error)
	// UnapproveMergeRequest unapproves a previously approved merge request on GitLab.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/merge_request_approvals.html#unapprove-merge-request
	UnapproveMergeRequest(pid interface{}, mr int, options ...RequestOptionFunc) (*Response, error)
	// ResetApprovalsOfMergeRequest clear all approvals of merge request on GitLab.
	// Available only for bot users based on project or group tokens.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/merge_request_approvals.html#reset-approvals-of-a-merge-request
	ResetApprovalsOfMergeRequest(pid interface{}, mr int, options ...RequestOptionFunc) (*Response, error)
	// GetConfiguration shows information about single merge request approvals
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/merge_request_approvals.html#get-configuration-1
	GetConfiguration(pid interface{}, mr int, options ...RequestOptionFunc) (*MergeRequestApprovals, *Response, error)
	// ChangeApprovalConfiguration updates the approval configuration of a merge request.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/merge_request_approvals.html#change-approval-configuration-deprecated
	ChangeApprovalConfiguration(pid interface{}, mergeRequest int, opt *ChangeMergeRequestApprovalConfigurationOptions, options ...RequestOptionFunc) (*MergeRequest, *Response, error)
	// ChangeAllowedApprovers updates the approvers for a merge request.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/merge_request_approvals.html#change-allowed-approvers-for-merge-request
	ChangeAllowedApprovers(pid interface{}, mergeRequest int, opt *ChangeMergeRequestAllowedApproversOptions, options ...RequestOptionFunc) (*MergeRequest, *Response, error)
	// GetApprovalRules requests information about a merge request’s approval rules
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/merge_request_approvals.html#get-merge-request-level-rules
	GetApprovalRules(pid interface{}, mergeRequest int, options ...RequestOptionFunc) ([]*MergeRequestApprovalRule, *Response, error)
	// GetApprovalState requests information about a merge request’s approval state
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/merge_request_approvals.html#get-the-approval-state-of-merge-requests
	GetApprovalState(pid interface{}, mergeRequest int, options ...RequestOptionFunc) (*MergeRequestApprovalState, *Response, error)
	// CreateApprovalRule creates a new MR level approval rule.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/merge_request_approvals.html#create-merge-request-level-rule
	CreateApprovalRule(pid interface{}, mergeRequest int, opt *CreateMergeRequestApprovalRuleOptions, options ...RequestOptionFunc) (*MergeRequestApprovalRule, *Response, error)
	// UpdateApprovalRule updates an existing approval rule with new options.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/merge_request_approvals.html#update-merge-request-level-rule
	UpdateApprovalRule(pid interface{}, mergeRequest int, approvalRule int, opt *UpdateMergeRequestApprovalRuleOptions, options ...RequestOptionFunc) (*MergeRequestApprovalRule, *Response, error)
	// DeleteApprovalRule deletes a mr level approval rule.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/merge_request_approvals.html#delete-merge-request-level-rule
	DeleteApprovalRule(pid interface{}, mergeRequest int, approvalRule int, options ...RequestOptionFunc) (*Response, error)
}
