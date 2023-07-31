package main


func Repeat(character string ,times int) string{
	var res string
	for i :=0 ;i<times;i++{
		res= res+ character;
	}
	return res
}