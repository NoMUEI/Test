package Maps

import (
	"testing"
)

//func TestSearch(t *testing.T) {
//	dictionary := map[string]string{"test":"this is just a test"}
//
//	got := Search(dictionary, "test")
//	want := "this is just a test"
//
//	assertStrings(t,got,want)
//}
//
////自定义函数判断函数类似于assert
func assertStrings(t *testing.T,got,want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}


func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test":"this is just a test"}
	t.Run("know word", func(t *testing.T){
		got,_ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got,want)
	})
	t.Run("unknow word", func(t *testing.T){
		_, err := dictionary.Search("unknow")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("get an error")
		}

		assertStrings(t, err.Error(), want)
	})


}
