package main

type LowercaseTrie struct {
	nodes       map[byte]*LowercaseTrie
	branchCount int  // Number of words that branch out of this node
	eow         bool // This node marks the end of a word
}

func MakeLowercaseTrie() *LowercaseTrie {
	return &LowercaseTrie{
		nodes:       make(map[byte]*LowercaseTrie),
		branchCount: 0,
		eow:         false,
	}
}

func (t *LowercaseTrie) LongestCommonPrefix() string {
	if t.branchCount != 1 {
		return ""
	}

	for k, v := range t.nodes {
		if v != nil {
			if k == '\x00' {
				return "" + v.LongestCommonPrefix()
			} else if t.eow {
				return string(k)
			}

			return string(k) + v.LongestCommonPrefix()
		}
	}
	return ""
}

func (t *LowercaseTrie) Put(v string) error {
	if v == "" { // someone adds an empty word to the trie
		if t.nodes['\x00'] == nil {
			t.nodes['\x00'] = MakeLowercaseTrie()
			t.branchCount++
		}
		t.nodes['\x00'].eow = true
		return nil
	}

	var (
		head = v[0]
		tail = v[1:]
	)

	if c := t.nodes[head]; c == nil {
		t.nodes[v[0]] = MakeLowercaseTrie()
		t.branchCount++
	}

	if tail != "" {
		return t.nodes[head].Put(tail)
	}

	t.eow = true // This is the last letter
	return nil
}
