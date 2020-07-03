package dijkstra_algo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
_______c_____
|6     |	|1
a	   |3	d
|2	   |	|
|_____ b____|5
*/

var testData = map[Name]Neighbors{
	"a": {
		"b": 2,
		"c": 6,
	},
	"b": {
		"a": 2,
		"c": 3,
		"d": 5,
	},
	"c": {
		"a": 6,
		"b": 3,
		"d": 1,
	},
	"d": {
		"c": 1,
		"b": 5,
	},
}

func TestGetPath(t *testing.T) {
	r := GetPath(testData, "a")
	assert.Equal(t, 2, int(r["b"]))
	assert.Equal(t, 5, int(r["c"]))
	assert.Equal(t, 6, int(r["d"]))
}
