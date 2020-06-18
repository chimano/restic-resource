package main

import (
	"testing"

	"github.com/chimano/restic-resource/common"
	"github.com/stretchr/testify/assert"
)

func TestCheckFiltering(t *testing.T) {
	c := CheckCommand{}
	versions := []common.Version{
		{VersionID: "1"},
		{VersionID: "2"},
		{VersionID: "3"},
		{VersionID: "4"},
	}
	expectedVersions := []common.Version{
		{VersionID: "3"},
		{VersionID: "4"},
	}
	baseVersion := common.Version{VersionID: "3"}
	filtered, _ := c.keepNewerVersions(versions, baseVersion)
	assert.Equal(t, expectedVersions, filtered)
}
