// Copyright (C) 2024 gitlab contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public
// License along with this program. If not, see
// <https://www.gnu.org/licenses/>.
//
// SPDX-License-Identifier: LGPL-3.0

package gitlab_test

import (
	"testing"

	"github.com/1password/gitlab"
	"gotest.tools/v3/assert"
)

func ExampleNewMockClient() {
	t := &testing.T{}

	gl := gitlab.NewMockClient(t)

	// Should be called once w/ the given arguments and return the given
	// result.
	gl.MergeRequestsServiceM.EXPECT().
		GetMergeRequest(1, 1, &gitlab.GetMergeRequestsOptions{}).
		Return(&gitlab.MergeRequest{
			BasicMergeRequest: gitlab.BasicMergeRequest{
				ID: 1,
			},
		}, nil, nil)

	mr, _, err := gl.MergeRequests().GetMergeRequest(1, 1, &gitlab.GetMergeRequestsOptions{})
	assert.NilError(t, err)
	assert.Equal(t, mr.ID, 1)
}
