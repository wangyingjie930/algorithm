package 位图

type BitMap struct {
	bits []byte
	digits int //位数
}

func NewBitMap(digits int) *BitMap {
	//1个byte的位数为8bit, 所以申请的byte数组所需的索引树为 digits / 8 + 1
	bits := make([]byte, digits / 8 + 1)
	return &BitMap{digits: digits, bits: bits}
}

func (b *BitMap) get(num int) bool {
	return b.bits[num / 8] & (1 << (num % 8)) > 0
}

func (b *BitMap) set(num int) {
	b.bits[num / 8] = b.bits[num / 8] | (1 << (num % 8))
}

func (b *BitMap) remove(num int) {
	b.bits[num / 8] = b.bits[num / 8] & ^(1 << (num % 8))
}


