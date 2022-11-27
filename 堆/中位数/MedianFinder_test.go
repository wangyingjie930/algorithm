package 中位数

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	t.Logf("%+v", obj.FindMedian())
	obj.AddNum(3)
	t.Logf("%+v", obj.FindMedian())
}
