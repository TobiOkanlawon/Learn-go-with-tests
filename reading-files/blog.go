package blogposts

import (
	"io"
	"io/fs"
)

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")

	if err != nil {
		return []Post{}, err
	}

	var posts []Post

	for _, f := range dir {
		file, err := fileSystem.Open(f.Name())

		if err != nil {
			return nil, err
		}

		fileInBytes, err := io.ReadAll(file)

		if err != nil {
			return nil, err
		}

		post, err := processPost(fileInBytes)

		if err != nil {
			return nil, err
		}

		file.Close()
		posts = append(posts, post)
	}

	return posts, nil
}
