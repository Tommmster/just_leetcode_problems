package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	trie := MakeLowercaseTrie()

	for _, e := range strs {
		_ = trie.Put(e)
	}

	return trie.LongestCommonPrefix()

}
