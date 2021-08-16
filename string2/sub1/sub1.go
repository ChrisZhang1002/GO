package sub1

var s1, s2 []string
var n int

//var s2=[] string

func Sub1() {

	s1 = []string{"北京", "上海", "深圳"}
	s2 = []string{"北京", "上海", "深圳"}

	n := 12
	p := &n
	println(s1[0])
	println(s2[1])

	println(p)
	println(*p)

}
