package lc_0003
/*
******************************************************************
  * @brief  用stack的解法，始终保持stack大小=不重复字符数量  
  * @param  s string
  * @ret    max int
  * @author    Troy
  * @date      2017/10/19 23:10
******************************************************************
*/
func lengthOfLongestSubstring(s string) int {
	if len(s)<=1 {
		return len(s)
	}
	var (
		max  int  //最大长度
		lo int    //栈低地址
		hi int    //栈高地址
	)
	for i := 0; i < len(s); i++ {
		c:=s[i] //c表示这个字符的实际值
		rep:=-1 //保存位置 -1表示未出现
		for j := lo; j < hi; j++ {//此过程将此值和之前保存的不重复值比较hi次
			if s[j] == c {
				rep=j  //出现重复记录位置，并跳出
				break
			}
		}
		hi++
		if rep == -1 {
			continue   //这个值第一次出现
		}else{
			//得到了一个重复值，此次累加的hi必须弹出
			if max < hi-lo-1 {
				max=hi-lo-1
				//假定输入的是ascii字符 不重复的值可能<=256
				if max >= 256 {
					return max
				}
			}
			lo=rep+1
		}
	}
	if max<hi-lo{
		max=hi-lo
	}
	return max


}