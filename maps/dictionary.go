package main

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrDoesNotExist      = DictionaryErr("this word does not exist in the dictionary")
	ErrWordAlreadyExists = DictionaryErr("this word already exists in the dictionary")
	ErrNotFound          = DictionaryErr("this word was not found in the dictionary")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrDoesNotExist
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, ok := d[word]
	if ok {
		return ErrWordAlreadyExists
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	if err != nil {
		return ErrDoesNotExist
	}

	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func (e DictionaryErr) Error() string {
	return string(e)
}
