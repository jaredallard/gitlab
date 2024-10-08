// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// ManagedLicensesService is an interface for [gitlab.Client.ManagedLicenses]
type ManagedLicensesService interface {
	// ListManagedLicenses returns a list of managed licenses from a project.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/managed_licenses.html#list-managed-licenses
	ListManagedLicenses(pid interface{}, options ...RequestOptionFunc) ([]*ManagedLicense, *Response, error)
	// GetManagedLicense returns an existing managed license.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/managed_licenses.html#show-an-existing-managed-license
	GetManagedLicense(pid, mlid interface{}, options ...RequestOptionFunc) (*ManagedLicense, *Response, error)
	// AddManagedLicense adds a managed license to a project.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/managed_licenses.html#create-a-new-managed-license
	AddManagedLicense(pid interface{}, opt *AddManagedLicenseOptions, options ...RequestOptionFunc) (*ManagedLicense, *Response, error)
	// DeleteManagedLicense deletes a managed license with a given ID.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/managed_licenses.html#delete-a-managed-license
	DeleteManagedLicense(pid, mlid interface{}, options ...RequestOptionFunc) (*Response, error)
	// EditManagedLicense updates an existing managed license with a new approval
	// status.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/managed_licenses.html#edit-an-existing-managed-license
	EditManagedLicense(pid, mlid interface{}, opt *EditManagedLicenceOptions, options ...RequestOptionFunc) (*ManagedLicense, *Response, error)
}
