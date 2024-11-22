// (c) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package deployerallowlist

import (
	"testing"

	"github.com/shubhamdubey02/subnet/core/state"
	"github.com/shubhamdubey02/subnet/precompile/allowlist"
)

func TestContractDeployerAllowListRun(t *testing.T) {
	allowlist.RunPrecompileWithAllowListTests(t, Module, state.NewTestStateDB, nil)
}

func BenchmarkContractDeployerAllowList(b *testing.B) {
	allowlist.BenchPrecompileWithAllowList(b, Module, state.NewTestStateDB, nil)
}