// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// PlanLimitsService is an interface for [gitlab.Client.PlanLimits]
type PlanLimitsService interface {
	// List the current limits of a plan on the GitLab instance.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/plan_limits.html#get-current-plan-limits
	GetCurrentPlanLimits(opt *GetCurrentPlanLimitsOptions, options ...RequestOptionFunc) (*PlanLimit, *Response, error)
	// ChangePlanLimits modifies the limits of a plan on the GitLab instance.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/plan_limits.html#change-plan-limits
	ChangePlanLimits(opt *ChangePlanLimitOptions, options ...RequestOptionFunc) (*PlanLimit, *Response, error)
}
