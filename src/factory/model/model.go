package model

type Student struct{
	Name string
	Score int
}

func SetStu(n string, s int) *Student{
	return &Student {
		Name : n,
		Score : s,
	}
}

func(st Student) Setst(n string, s int) (Student){
	 st.Name = n
	 st.Score = s 
	 return st
}