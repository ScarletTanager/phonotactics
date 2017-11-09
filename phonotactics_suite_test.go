package phonotactics_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPhonotactics(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Phonotactics Suite")
}
