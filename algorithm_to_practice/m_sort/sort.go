package main

import "fmt"

// 排序算法集合

// 1.冒泡排序
// 2.插入排序
// 3.选择排序

// 冒泡排序
//每次冒泡操作都会对相邻的两个元素进行比较，看是否满足大小关系要求。如果不满足就让它俩互换
func bubblingSort(arr [8]int) [8]int {
	count := len(arr)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < count-1; j++ {
			if arr[j] > arr[j+1] { // 判断大小决定是否替换
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
		count-- //这是最关键的，每次处理后，最大的元素都会在最右边，因此后续要处理的数组长度均减一
	}
	return arr
}

// 插入排序
//插入算法的核心思想是取未排序区间中的元素，在已排序区间中找到合适的插入位置将其插入，并保证已排序区间数据一直有序。重复这个过程，直到未排序区间中元素为空，算法结束
func insertionSort(arr [8]int) [8]int {
	//[3, 6, 7], [5, 1, 2, 10, 4]
	//[3, 5, 6, 7],[1, 2, 10, 4]
	//[1，3, 5, 6, 7]，[2, 10, 4]
	// 假设第一个元素已经是已排序区间的第一个值
	hasSortNum := 1             // arr[hasSortNum] 就是新的要处理的元素, hasSortNum 刚好是未排序区间的第一个元素下标
	for hasSortNum < len(arr) { // 说明都排完了
		// 遍历一下已排序区间，插入到合适的位置，并且搬运数据
		for k, _ := range arr[:hasSortNum] {
			// 找到要插入的位置
			if arr[k] > arr[hasSortNum] { // 此时 arr[k]就是要插入的位置
				// 搬运数据，需要将 k 位置给插入的元素，其他元素都后移，直到arr[hasSortNum]元素
				// k=2 hasSortNum=3 需要搬运2个位置
				// k=1 hasSortNum=4,需要搬运4个位置 (hasSortNum-k+1)
				// 先保留要被代替的元素
				temp := arr[hasSortNum]
				for i := hasSortNum; i > k; i-- { // k+1 到 hasSortNum 位置都要动
					arr[i] = arr[i-1]
				}
				arr[k] = temp
			}
		}
		hasSortNum++
	}
	return arr
}

// 选择排序
//选择排序算法的实现思路有点类似插入排序，也分已排序区间和未排序区间。但是选择排序每次会从未排序区间中找到最小的元素，将其放到已排序区间的末尾
func selectionSort(arr [8]int) [8]int {
	//[3, 6, 7, 5, 1, 2, 10, 4]
	//[1, 6, 7, 5 , 3, 2, 10, 4]
	hasSortNum := 0 //一开始已排区间为空
	for hasSortNum < len(arr) {
		minValue := arr[hasSortNum] // 取未排第一个元素作为比较值
		minValueIndex := hasSortNum
		for i := hasSortNum + 1; i < len(arr); i++ {
			// 遍历未排序区间，找出最小值
			if arr[i] < minValue {
				minValue = arr[i]
				minValueIndex = i
			}
		}
		// 替换两个的位置
		temp := arr[hasSortNum]
		arr[hasSortNum] = minValue
		arr[minValueIndex] = temp
		hasSortNum++
	}

	return arr
}

func main() {
	arr := [8]int{3, 6, 7, 5, 1, 2, 10, 14}
	fmt.Println(bubblingSort(arr))
	fmt.Println(insertionSort(arr))
	fmt.Println(selectionSort(arr))
}
