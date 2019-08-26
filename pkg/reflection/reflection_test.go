package reflection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadMessageDescriptor(suite *testing.T) {
	suite.Run("LoadMessageDescriptor_ShouldReturnMessageDescriptorIfNameFound", func(t *testing.T) {
		msgName := "tutorial.Person"
		md, err := LoadMessageDescriptor("testdata/addressbook.fds", msgName)
		assert.NoError(t, err)
		if assert.NotNil(t, md) {
			assert.Equal(t, msgName, md.GetFullyQualifiedName())
		}
	})

	suite.Run("LoadMessageDescriptor_ShouldReturnErrorIfMessageDescriptorNameNotFound", func(t *testing.T) {
		msgName := "Person"
		md, err := LoadMessageDescriptor("testdata/addressbook.fds", msgName)
		assert.Error(t, err)
		assert.Nil(t, md)
	})

}