package blogposts

import (
	"errors"
)

type Post struct {
	Title string
}

func processPost(postData []byte) (Post, error) {
	postDataAsString := string(postData)
	title, err := extractTitle(postDataAsString)
	if err != nil {
		return Post{}, err
	}
	return Post{Title: title}, nil
}

func extractTitle(postData string) (string, error) {

	if len(postData) < 8 {
		return "", errors.New("no title")
	}
	return postData[7:], nil
}
