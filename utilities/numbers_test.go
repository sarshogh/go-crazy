package utilities

import (
	"fmt"
	"testing"
)

var isEqualTestData = []struct {
	name         string
	nums1        []int
	nums2        []int
	expectResult bool
}{
	{"testing IsEqual_1", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}, true},
	{"testing IsEqual_2", []int{1, 2, 3, 4}, []int{5, 6, 7, 8}, false},
	{"testing IsEqual_3", []int{1, 2, 3, 4}, []int{1, 2, 3, 4, 5}, false},
}

func Test_IsEqual(t *testing.T) {

	for _, item := range isEqualTestData {
		fmt.Println(item.name)
		t.Run(item.name, func(t *testing.T) {

			result := IsEqual(item.nums1, item.nums2)
			if result != item.expectResult {
				t.Errorf("IsEqual returns %t, while want %t", result, item.expectResult)
			}
		})
	}
}
