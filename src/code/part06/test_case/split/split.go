package split

import "strings"

// split.go
// Split 将s按照sep进行分割，返回一个字符串的切片
// Split("我爱你","爱") => ["", "你"]
func Split(s, sep string)(ret[] string) {
	// 性能优化，减少内存申请
	// 让它不要重复做内存的申请，做了内存的初始化
	ret = make([]string, 0, strings.Count(s, sep) + 1)

	idx := strings.Index(s, sep)
	for idx > -1{
		// 每次做了内存申请，通过测试来优化代码
		ret = append(ret, s[:idx])
		// 根据字节来切割，因为有中文，所以需要用len来获取它的长度
		s = s[idx+len(sep):]
		idx = strings.Index(s, sep)

	}
	ret = append(ret, s)
	return
}
