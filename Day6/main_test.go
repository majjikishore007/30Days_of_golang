package main

import "testing"

func TestSearch(t *testing.T)  {
	dictionary := Dictionary{"test":"this is just a test"}


	t.Run("Known word",func(t *testing.T) {
		got ,_ := dictionary.Search("test")
		want := "this is just a test"
		AssertionStrings(t,got,want)
	})
	
	t.Run("Unknown word",func(t *testing.T) {
		_,err := dictionary.Search("hello")
		want := DictionaryErr(ErrNotFound)
		AssertError(t,err,want)
	})
	
}

func TestAdd(t *testing.T){
	
	t.Run("new key-value",func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		value := "this is just a test"
		err:=dictionary.Add(key,value)
		AssertError(t,err,nil)
		assertDefinition(t,dictionary,key,value)
	})
	t.Run("existing word",func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dictionary := Dictionary{key:value}
		err:=dictionary.Add(key,value)

		AssertError(t,err,DictionaryErr(ErrWordExists))
		assertDefinition(t,dictionary,key,value)

		
	})
}

func TestUpdate(t *testing.T) {
    key := "test"
    value := "this is just a test"
    dictionary := Dictionary{key: value}
    newValue := "new definition"

    dictionary.Update(key, value)

    assertDefinition(t, dictionary, key, newValue)
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
    t.Helper()

    got, err := dictionary.Search(word)
    if err != nil {
        t.Fatal("should find added word:", err)
    }

    if definition != got {
        t.Errorf("got %q want %q", got, definition)
    }
}

func AssertError(t *testing.T,got, want error)  {
	t.Helper()

    if got != want {
        t.Errorf("got error %q want %q", got, want)
    }
}

func AssertionStrings(t *testing.T,got,want string)  {
	t.Helper()
	if got != want {
		t.Errorf("%q got %q want", got, want)
	}
}