package main

func Sum(arr []int )int {
	sum :=0
	for _,num := range arr{
		sum+=num
	}
	return sum
}

func SumAll(arr ...[]int) []int{
	// length:=len(arr)
	// sums := make([]int ,length)
	// for  i, nums:= range arr{
	// 	sums[i]=Sum(nums)
	// }

	// refactoring 

	var sums [] int 
	for _ ,nums := range arr {
		sums = append(sums, Sum(nums))
	}
	return sums

}

func SumOfTails(arr ...[]int)[]int{
	
	var sums []int

	for _ , nums := range arr{
		if(len(nums)==0){
			sums = append(sums, 0)
		}else{

			temp:= nums[1:]
	
			sums = append(sums, Sum(temp))
		}
	}
	
	return sums
}