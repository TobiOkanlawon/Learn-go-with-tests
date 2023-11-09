package blogposts_test

import (
	"errors"
	blogposts "github.com/TobiOkanlawon/blogposts"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, I always fail")
}

func TestNewBlogPosts(t *testing.T) {

	assertPost := func(t testing.TB, got blogposts.Post, want blogposts.Post) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v, but want %+v", got, want)
		}

	}

	t.Run("returns the correct amount of files", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte("some complicated text")},
			"hello-world2.md": {Data: []byte("Adipiscing diam donec adipiscing tristique risus nec feugiat in fermentum posuere urna nec tincidunt praesent semper feugiat nibh sed pulvinar proin. Hendrerit lectus a molestie lorem ipsum dolor sit amet?")},
		}

		posts, err := blogposts.NewPostsFromFS(fs)

		if err != nil {
			t.Fatalf("returned an error while trying to read fs %v", err)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}
	})

	t.Run("returns an error if FS fails", func(t *testing.T) {

		_, err := blogposts.NewPostsFromFS(StubFailingFS{})

		if err == nil {
			t.Errorf("expected an error while trying to read failing fs, got %s", err)
		}
	})

	t.Run("returns the correct names for each Post", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte("Title: Post 1")},
			"hello-world2.md": {Data: []byte("Title: Post 2")},
		}

		posts, _ := blogposts.NewPostsFromFS(fs)

		got := posts[0]
		want := blogposts.Post{Title: "Post 1"}

		assertPost(t, got, want)

	})
}
