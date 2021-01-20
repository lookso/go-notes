package tree

import "fmt"

type Trie struct {
	isEnd bool  //如果要统计个数的话 我觉得这块改为int类型 然后默认是0 代表不是结尾 如果是其他数字 代表出现的次数
	next  [26]*Trie		//26长度的数组来指向下一层的a-z
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{} 	//初始化 根节点为空
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	// v - ‘a’ 的意思是 ascii码相减 减去a的话 a-z 等于 0-25 刚好得到对应相对的index
	// 然后每一层的往下
	for _, v := range word {
		if this.next[v-'a'] == nil {
			this.next[v-'a'] = &Trie{}
		}
		this = this.next[v-'a']
	}
	this.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	// 搜索也类似 往下走，走到哪一层没有的话 就直接返回false
	for _, v := range word {
		fmt.Println(v-'a')
		if this.next[ v-'a'] == nil {
			return false
		} else {
			this = this.next[v-'a']
		}
	}
	return this.isEnd == true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		if this.next[ v-'a'] == nil {
			return false
		} else {
			this = this.next[v-'a']
		}
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

