package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("retrieve a word that exists in our dictionary", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want, "test")
	})

	t.Run("retrieve a word that does not exist in our dictionary", func(t *testing.T) {
		dictionary := Dictionary{}

		query := "NonExistentWord"
		_, err := dictionary.Search(query)

		assertError(t, err, ErrDoesNotExist)
	})
}

func TestAdd(t *testing.T) {
	t.Run("adds a word to the dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		_ = dictionary.Add("test", "meaning test")

		got, _ := dictionary.Search("test")
		want := "meaning test"

		if got != want {
			t.Errorf("got %q, but want %q", got, want)
		}
	})

	t.Run("doesn't override a word that already exists", func(t *testing.T) {
		dictionary := Dictionary{"test": "test text"}

		err := dictionary.Add("test", "some other definition")

		definition, _ := dictionary.Search("test")

		if err == nil {
			t.Fatal("should raise an error")
		}

		assertError(t, err, ErrWordAlreadyExists)

		if definition == "some other definition" {
			t.Error(
				"should not override previous definition")
		}
	})
}

func TestUpdate(t *testing.T) {
	t.Run("updates the definition of a word that exists in the dictionary", func(t *testing.T) {
		dictionary := Dictionary{"test": "initial definition"}

		dictionary.Update("test", "updated definition")
		got, err := dictionary.Search("test")
		if err != nil {
			t.Fatal("did not expect an error")
		}
		want := "updated definition"

		assertStrings(t, got, want, "test")
	})

	t.Run("does not update a word that doesn't exist", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		err := dictionary.Update(word, "definition")

		if err != ErrDoesNotExist {
			t.Fatal("does not return the appropriate error ")
		}
		_, err = dictionary.Search("test")

		assertError(t, err, ErrDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("deletes a word that exists", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "definition"}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)

		if err != ErrDoesNotExist {
			t.Errorf("Expected %q to be deleted", err)
		}
	})
}

func assertStrings(t testing.TB, got, want, given string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, but want %q, when given %q", got, want, given)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("expected error %v, but got %q", want, got)
	}
}
