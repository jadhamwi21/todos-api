package tests

import (
	"fmt"
	"testing"
	"todos-api/internal/validation"
)

func TestTagFormatter(t *testing.T) {
	const mockField = "field"
	tagsMap := map[string]string{"required": fmt.Sprintf("%v is required", mockField)}

	tagsList := []string{"required"}
	for _, tag := range tagsList {
		got := validation.TagFormatter(mockField, tag)
		if tagsMap[tag] != got {
			t.Errorf("tag %v incorrect format", tag)
		}
	}
}
