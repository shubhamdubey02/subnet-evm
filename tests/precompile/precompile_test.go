// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package precompile

import (
	"os"
	"testing"

	ginkgo "github.com/onsi/ginkgo/v2"

	// Import the solidity package, so that ginkgo maps out the tests declared within the package
	"github.com/shubhamdubey02/subnet-evm/tests/precompile/solidity"
)

func TestE2E(t *testing.T) {
	if basePath := os.Getenv("TEST_SOURCE_ROOT"); basePath != "" {
		os.Chdir(basePath)
	}
	solidity.RegisterAsyncTests()
	ginkgo.RunSpecs(t, "subnet-evm precompile ginkgo test suite")
}
