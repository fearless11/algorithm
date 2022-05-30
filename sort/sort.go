package sort

import "fmt"

// 十种排序算法
// 冒泡排序 O(n^2) O(1) 稳定
// 选择排序 O(n^2) O(1) 不稳定
// 插入排序 O(n^2) O(1) 稳定
// 快速排序 O(nlogn) O(nlogn) 不稳定
// 归并排序 O(nlogn) O(n) 稳定
// 堆排序 O(nlogn) O(1) 不稳定
// 希尔排序 O(nlogn) O(nlogn) 不稳定
// 计数排序 O(n+k) O(n+k) 稳定
// 桶排序 O(n+k) O(n+k) 稳定
// 基数排序 O(n*k) O(n+k) 稳定

// 冒泡排序
// 操作：两两比较，每次把最大的数拍到最后
// T(n)取决于比较次数
// 正序时，比较次数为n-1，交换次数为0，最好T(n)=O(n)
// 逆序时，比较次数为(n-1)+(n-2)+...+2+1 = n(n-1)/2，交换次数3n(n-1)/2，最坏T(n)=O(n^2)
// 乱序时，平均T(n)=O(n^2)
// S(n)=O(1)
// 在排序过程中，元素两两交换时，相同元素的前后顺序并没有改变，因此是稳定排序算法
func BubbleSort(a []int) []int {
	n := len(a)
	for i := 0; i < n-1; i++ {
		for i := 0; i < n-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
	}
	return a
}

// 如果没有交换，说明序列已经有序
func BubbleTagSort(a []int) []int {
	n := len(a)
	tag := 0
	for i := 0; i < n-1 && tag != 0; i++ {
		tag = 0
		for j := 0; j < n-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				tag = 1
			}
		}
	}
	return a
}

// 选择排序
// 操作：在指定位置每次选择剩下最小的数放入
// T(n)取决于比较次数
// 正序时，比较次数(n-1)+(n-2)+(n-3)+...+2+1 = n(n-1)/2，交换N次，最好T(n)=O(n^2)
// 逆序时，比较次数(n-1)+(n-2)+(n-3)+...+2+1 = n(n-1)/2，交换N次，最坏T(n)=O(n^2)
// 乱序时，平均T(n)=O(n^2)
// S(n) = 1
// 在排序过程中，元素的前后位置会变化，因此是不稳定排序算法
func SelectSort(a []int) []int {
	n := len(a)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
	return a
}

// 插入排序
// 操作：在指定的位置每次选择已指定位置内最小的数到最前面
// T(n)取决于交换的次数
// 正序时，比较次数0，交换0次，最好T(n)=O(1)
// 逆序时，比较次数1+2+...+(n-2)+(n-1) = n(n-1)/2，交换1+2+...+(n-2)+(n-1) = n(n-1)/2次，最坏T(n)=O(n^2)
// 乱序时，平均T(n)=O(n^2)
// S(n) = 1
// 在排序过程中，数据是与前面的两两交换，因此是稳定排序算法
func InsertSort(a []int) []int {
	n := len(a)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && a[j-1] > a[j]; j-- {
			a[j-1], a[j] = a[j], a[j-1]
		}
	}
	return a
}

// 快速排序
// 操作：选取第一元素为基准元素；元素排序，比基准大的放右边，比基准小的放左边，一趟后锁定基准元素位置，然后基准两边继续选基准确认位置

func QuickSort(a []int, left int, right int) []int {
	if left < right {
		tag := partion(a, left, right)
		QuickSort(a, left, tag-1)
		QuickSort(a, tag+1, right)
	}
	return a
}

func partion(a []int, left int, right int) int {
	temp := a[left]

	for i := left; i < right; i++ {		
		if a[i] > temp {
			a[i] = a[right]
			right--
		}
	}

	}

	return tag
}

// 希尔排序
// 非稳定的排序算法，也是一种插入排序
// O(n^2)
func shellSort(a []int) {
	fmt.Println("who can tell me")
}

//
