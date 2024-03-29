package sort

/**
 * 冒泡排序
 * 时间复杂度：
 *  * 最好：O(n)
 *  * 最坏：O(n^2)
 *  * 平均：O(n^2)
 * 空间复杂度：O(1)
 * 稳定性：稳定
 * @param SortInterface data 源数据
 * @param int start 需排序起始位置
 * @param int end 需排序末尾位置
 */
func BubbleSort(data SortInterface, start, end int) {
	if data.Len() <= 1 {
		return
	}

	for i := end; i >= start; i-- {
		for j := start; j < i; j++ {
			if data.Less(j, j+1) {
				data.Swap(j, j+1)
			}
		}
	}
}
