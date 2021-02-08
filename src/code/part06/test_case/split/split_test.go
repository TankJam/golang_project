package split

import (
	"reflect"
	"testing"
)

// 测试
//func TestSplit(t *testing.T){
//	// 调用Split想要返回的结果
//	got := Split("我爱你", "爱")
//	// 想要测试返回的结果
//	want := []string{"我", "你"}
//	// 反射断言，比较got与want的结果是否一致，若一致返回true
//	if !reflect.DeepEqual(got, want){
//		t.Errorf("want: %v got:%v", want, got)
//	}
//}

// 测试组，创建多个结构体示例进行测试，更加全面
func TestSplit(t *testing.T){
	type test struct{
		input string
		sep string
		want []string
	}
	tests := map[string]test{
		"simple": test{input: "我爱你", sep:"爱", want: []string{"我", "你"}},
		"multi sep": test{input: "a:b:c", sep:":", want: []string{"a", "b", "c"}},
		"multi sep2": test{input: "abcd", sep:"bc", want: []string{"a", "d"}},
		"chinese": test{input: "上海有上又有海", sep:"上", want: []string{"", "海有", "又有海"}},
	}

	// 反射断言，比较got与want的结果是否一致，若一致返回true
	for name, tc := range tests{
		// 测试组需要通过t.Run来执行
		t.Run(name, func(t *testing.T){
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want){
				t.Errorf("name: %s failed, want: %v got:%v", name, tc.want, got)
			}
		})
	}

}

// 性能基准测试
func BenchmarkSplit(b *testing.B) {
	// b.N 不是固定的数
	for i:=0; i<b.N; i++ {
		Split("上海有上又有海", "沙")
	}
}