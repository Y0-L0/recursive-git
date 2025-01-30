package repo

import (
	"github.com/Y0-L0/recursive-git/testutils"
  "github.com/Y0-L0/recursive-git/git/pkg/commit"
  m "github.com/Y0-L0/recursive-git/git/pkg/models"
	"testing"
)

func TestGetObject(t *testing.T) {
	obj, err := getObject(testRepo().base, EXAMPLE_COMMIT.sha)
	testutils.Ok(t, err)

	testutils.Equals(t, EXAMPLE_COMMIT.object, obj)
}
