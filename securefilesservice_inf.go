// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

import (
	"io"
)

// SecureFilesService is an interface for [gitlab.Client.SecureFiles]
type SecureFilesService interface {
	// ListProjectSecureFiles gets a list of secure files in a project.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/api/secure_files/#list-project-secure-files
	ListProjectSecureFiles(pid interface{}, opt *ListProjectSecureFilesOptions, options ...RequestOptionFunc) ([]*SecureFile, *Response, error)
	// ShowSecureFileDetails gets the details of a specific secure file in a project.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/api/secure_files/#show-secure-file-details
	ShowSecureFileDetails(pid interface{}, id int, options ...RequestOptionFunc) (*SecureFile, *Response, error)
	// CreateSecureFile creates a new secure file.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/api/secure_files/#create-secure-file
	CreateSecureFile(pid interface{}, content io.Reader, filename string, options ...RequestOptionFunc) (*SecureFile, *Response, error)
	// DownloadSecureFile downloads the contents of a project's secure file.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/api/secure_files/#download-secure-file
	DownloadSecureFile(pid interface{}, id int, options ...RequestOptionFunc) (io.Reader, *Response, error)
	// RemoveSecureFile removes a project's secure file.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/api/secure_files/#remove-secure-file
	RemoveSecureFile(pid interface{}, id int, options ...RequestOptionFunc) (*Response, error)
}
