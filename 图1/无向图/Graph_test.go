package 无向图

import (
	"testing"
)

func TestNewGraph(t *testing.T) {
	g := NewGraph([][]int{{1, 2, 1}, {2, 3, 2}, {3, 4, 3}})
	g.PrintGraph()
}
