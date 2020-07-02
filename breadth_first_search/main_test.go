package breadth_first_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetShortestPath(t *testing.T) {
	graph := map[string][]string{
		"a": {"b", "d"},
		"b": {"a", "c"},
		"c": {"b", "f", "j"},
		"d": {"a", "i"},
		"f": {"c"},
		"i": {"d", "j"},
		"j": {"c", "i"},
	}
	lvl := GetLevel(graph, "a", "b")
	assert.Equal(t, 1, lvl)
	lvl = GetLevel(graph, "a", "i")
	assert.Equal(t, 2, lvl)
	lvl = GetLevel(graph, "a", "j")
	assert.Equal(t, 3, lvl)
}
