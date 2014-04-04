package qsort

func quickSort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {//用另外一种表达方式写一下
		for j >= p {
			if values[j] >= temp{ 
				j--
			}
			else{
				values[p] = values[j]
				p = j
			}
		}
		if  i <= p {
			if values[i] <= temp {
				i++
				}
			else{
				values[p] = values[i]
				p = i
			}
		}
	}
	values[p] = temp
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}
