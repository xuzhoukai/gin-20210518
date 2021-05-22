package dao_mongo

import (
	"gin-20210518/model"
	"testing"
)

func Test_studentCollection_FindAll(t *testing.T) {
	all, err := StudentColl.FindAll()
	if err != nil {
		t.Errorf("fail to find student err:%v", err)
	}
	t.Logf("len:%v", len(all))
	for _, v := range all {
		t.Logf("name:%v", v.Name)
	}
}

func Test_studentCollection_Find(t *testing.T) {
	stu, err := StudentColl.Find(1)
	if err != nil {
		t.Errorf("err:%v", err)
	}
	t.Logf("name:%v", stu.Name)
}

func Test_studentCollection_Add(t *testing.T) {
	stu := model.NewStudent()
	err := StudentColl.Add(stu)
	if err != nil {
		t.Logf("add err:%v", err)
	}
}

func Test_studentCollection_Delete(t *testing.T) {
	err := StudentColl.Delete(12)
	if err != nil {
		t.Errorf("delete err:%v", err)
	}
}

func BenchmarkSprintf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stu, err := StudentColl.Find(1)
		if err != nil {
			b.Errorf("err:%v", err)
		}
		b.Logf("name:%v", stu.Name)
	}
}
