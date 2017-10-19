package lc_0012

//查表法 每个数组内的内容为N~9N这样，
//比如888这样的数 某一位上mod10是他本身 这样就能定位到某个数
func IntToRoman(num int) string {
	var M = []string{"", "M", "MM", "MMM"}
	var C = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	var X = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	var I = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	return M[num/1000] + C[(num/100)%10] + X[(num/10)%10] + I[num%10]
}
