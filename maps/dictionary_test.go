package maps

import "testing"

func TestSearch(t *testing.T) {
	assertString := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}
	assertErr := func(t *testing.T, got, want error) {
		t.Helper()
		if got != want {
			t.Errorf("expected error '%s' but didn't get any", want)
		}
	}
	assertNoErr := func(t *testing.T, err error) {
		t.Helper()
		if err != nil {
			t.Errorf("unexpected error '%s'", err)
		}
	}

	t.Run("Known word", func(t *testing.T) {
		dictionary := Dictionary{
			"test": "this is just a test",
		}
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertString(t, got, want)
	})

	t.Run("Unknown word", func(t *testing.T) {
		dictionary := Dictionary{}
		_, err := dictionary.Search("test")
		assertErr(t, err, ErrKeyNotFound)
	})

	t.Run("Add word", func(t *testing.T) {
		dictionary := Dictionary{
			"test": "this is just a test",
		}
		dictionary.Add("hello", "hello, Zahid")
		got, err := dictionary.Search("hello")
		want := "hello, Zahid"
		assertNoErr(t, err)
		assertString(t, got, want)
	})

	t.Run("Add duplicate key", func(t *testing.T) {
		dictionary := Dictionary{
			"test": "this is just a test",
		}
		err := dictionary.Add("test", "new test")
		assertErr(t, err, ErrDuplicateKey)
		got, err := dictionary.Search("test")
		want := "this is just a test"
		assertString(t, got, want)
	})

	t.Run("Update word", func(t *testing.T) {
		dictionary := Dictionary{
			"test": "this is just a test",
		}
		err := dictionary.Update("test", "new test")
		assertNoErr(t, err)
		got, err := dictionary.Search("test")
		want := "new test"
		assertString(t, got, want)
	})

	t.Run("Update non existing key", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("test", "new test")
		assertErr(t, err, ErrKeyNotFound)
	})

	t.Run("Delete a key-word pair", func(t *testing.T) {
		dictionary := Dictionary{
			"test": "this is just a test",
		}
		err := dictionary.Delete("test")
		assertNoErr(t, err)
		_, err = dictionary.Search("test")
		assertErr(t, err, ErrKeyNotFound)
	})

	t.Run("Delete an unknown key-word pair", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Delete("test")
		assertErr(t, err, ErrKeyNotFound)
	})
}
