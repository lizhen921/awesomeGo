package slice

import (
	"fmt"
	"testing"
)

/**
使用截取array的形式初始化一个slice，则slice底层就是该数组，那么slice的len为截取的元素的个数，cap为从截取的位置开始到数组的结尾位置。由于slice和array公用一段内存，则修改slice、array都会相互影响
*/
func TestArrayAndSlice(t *testing.T) {

	var strArray = [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var strSlice = strArray[5:6]
	fmt.Println("strArray: ", strArray, "\tlen: ", len(strArray), "\tcap: ", cap(strArray))
	fmt.Println("strSlice: ", strSlice, "\tlen: ", len(strSlice), "\tcap: ", cap(strSlice))

	fmt.Println("修改strSlice的第一个元素为：slice-update")
	strSlice[0] = "slice-update"
	fmt.Println("strArray: ", strArray)
	fmt.Println("strSlice: ", strSlice, "\tlen: ", len(strSlice), "\tcap: ", cap(strSlice))

	fmt.Println("为slice添加一个元素为：slice-add")
	strSlice = append(strSlice, "slice-add")
	fmt.Println("strArray: ", strArray)
	fmt.Println("strSlice: ", strSlice, "\tlen: ", len(strSlice), "\tcap: ", cap(strSlice))

	fmt.Println("修改数组下标为6的元素为：array-update")
	strArray[6] = "array-update"
	fmt.Println("strArray: ", strArray)
	fmt.Println("strSlice: ", strSlice, "\tlen: ", len(strSlice), "\tcap: ", cap(strSlice))

	strSlice = append(strSlice, "slice-add")
	strSlice = append(strSlice, "slice-add")
	strSlice = append(strSlice, "slice-add")
	strSlice = append(strSlice, "slice-add")
	strSlice = append(strSlice, "slice-add")
	fmt.Println("strArray: ", strArray, "\tlen: ", len(strArray), "\tcap: ", cap(strArray))
	fmt.Println("strSlice: ", strSlice, "\tlen: ", len(strSlice), "\tcap: ", cap(strSlice))

}
