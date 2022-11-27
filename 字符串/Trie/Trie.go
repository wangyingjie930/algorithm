package Trie

type Trie struct {
	val   byte
	values [26]*Trie
	isEnd bool
}


func Constructor() Trie {
	trie := Trie{val: '/', isEnd: true}
	return trie
}


func (this *Trie) Insert(word string)  {
	node := this
	for _, val := range word {
		if node.values[val - 'a'] == nil {
			node.values[val - 'a'] = &Trie{val: byte(val), isEnd: false}
		}
		node = node.values[val - 'a']
	}
	node.isEnd = true
}


func (this *Trie) Search(word string) bool {
	node := this
	for _, val := range word {
		if node.values[val - 'a'] == nil {
			return false
		}
		node = node.values[val - 'a']
	}
	return node.isEnd
}


func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, val := range prefix {
		if node.values[val - 'a'] == nil {
			return false
		}
		node = node.values[val - 'a']
	}
	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
