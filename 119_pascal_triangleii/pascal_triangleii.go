package lc_0119

func getRow(rowIndex int) []int {
arr:=[]int{1}

for i:=1;i<rowIndex+1;i++{
	//新加个元素
	arr=append(arr, 0)
	//这是精髓
	for j := i; j >= 1; j--{
		arr[j]+=arr[j-1]
	}
}
return arr[:]
}