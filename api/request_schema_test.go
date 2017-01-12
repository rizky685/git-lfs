package api_test

import (
	"testing"

	"github.com/git-lfs/git-lfs/api"
	"github.com/stretchr/testify/assert"
)

// AssertRequestSchema encapsulates a single assertion of equality against two
// generated RequestSchema instances.
//
// This assertion is meant only to test that the request schema generated by an
// API service matches what we expect it to be. It does not make use of the
// *api.Client, any particular lifecycle, or spin up a test server. All of that
// behavior is tested at a higher strata in the client/lifecycle tests.
//
//   - t is the *testing.T used to preform the assertion.
//   - Expected is the *api.RequestSchema that we expected to be generated.
//   - Got is the *api.RequestSchema that was generated by a service.
func AssertRequestSchema(t *testing.T, expected, got *api.RequestSchema) {
	assert.Equal(t, expected, got)
}