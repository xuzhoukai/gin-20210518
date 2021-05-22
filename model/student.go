package model

type StudentFunc func(*Student)

type Student struct {
	Id   int64  `bson:"id",json:"id"`
	Name string `bson:"name,omitempty" json:"name,omitempty"`
	Age  int    `bson:"age,omitempty" json:"age,omitempty"`
}

func NewStudent() *Student {
	s := &Student{
		Id:   12,
		Name: "kaixin",
		Age:  4,
	}
	return s
}

func NewStudent2(opts ...StudentFunc) *Student {
	stu := NewStudent()
	for _, o := range opts {
		o(stu)
	}
	return stu
}

func WithName(name string) StudentFunc {
	return func(s *Student) {
		s.Name = name
	}
}

func WithAge(age int) StudentFunc {
	return func(s *Student) {
		s.Age = age
	}
}
