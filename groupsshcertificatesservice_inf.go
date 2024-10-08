// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// GroupSSHCertificatesService is an interface for [gitlab.Client.GroupSSHCertificates]
type GroupSSHCertificatesService interface {
	// ListGroupSSHCertificates gets a list of SSH certificates for a specified
	// group.
	//
	// Gitlab API docs:
	// https://docs.gitlab.com/ee/api/group_ssh_certificates.html#get-all-ssh-certificates-for-a-particular-group
	ListGroupSSHCertificates(gid interface{}, options ...RequestOptionFunc) ([]*GroupSSHCertificate, *Response, error)
	// CreateMemberRole creates a new member role for a specified group.
	//
	// Gitlab API docs:
	// https://docs.gitlab.com/ee/api/group_ssh_certificates.html#create-ssh-certificate
	CreateGroupSSHCertificate(gid interface{}, opt *CreateGroupSSHCertificateOptions, options ...RequestOptionFunc) (*GroupSSHCertificate, *Response, error)
	// DeleteGroupSSHCertificate deletes a SSH certificate from a specified group.
	//
	// Gitlab API docs:
	// https://docs.gitlab.com/ee/api/group_ssh_certificates.html#delete-group-ssh-certificate
	DeleteGroupSSHCertificate(gid interface{}, cert int, options ...RequestOptionFunc) (*Response, error)
}
