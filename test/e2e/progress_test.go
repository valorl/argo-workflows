//go:build functional
// +build functional

package e2e

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	wfv1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"github.com/argoproj/argo-workflows/v3/test/e2e/fixtures"
)

type ProgressSuite struct {
	fixtures.E2ESuite
}

func (s *ProgressSuite) TestDefaultProgress() {
	s.Given().
		Workflow("@testdata/basic-workflow.yaml").
		When().
		SubmitWorkflow().
		WaitForWorkflow(fixtures.ToBeSucceeded).
		Then().
		ExpectWorkflow(func(t *testing.T, metadata *metav1.ObjectMeta, status *wfv1.WorkflowStatus) {
			assert.Equal(t, wfv1.Progress("1/1"), status.Progress)
			assert.Equal(t, wfv1.Progress("1/1"), status.Nodes[metadata.Name].Progress)
		})
}

func TestProgressSuite(t *testing.T) {
	suite.Run(t, new(ProgressSuite))
}
