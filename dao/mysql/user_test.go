package dao_mysql

import "testing"

func Test_insertRowDemo(t *testing.T) {
	insertRowDemo()
}

func Test_queryRowDemo(t *testing.T) {
	queryRowDemo()
}

func Test_queryMultiRowDemo(t *testing.T) {
	queryMultiRowDemo()
}

func Test_updateRowDemo(t *testing.T) {
	updateRowDemo()
}

func Test_deleteRowDemo(t *testing.T) {
	deleteRowDemo()
}

func BenchmarkSprintf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queryRowDemo()
	}
}
