package client

import (
	"fmt"
	"testing"
)

func TestCommentList(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	comment := NewComment(space, apiKey, "TESTTOOL-64")
	r, err := comment.List()
	comment.PrintCSV(r)
}

func TestCommentAddDelete(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	q := NewCommentRequest()
	q.Content("コメント登録のテスト")

	comment := NewComment(space, apiKey, "TESTTOOL-64")
	r, err := comment.Add(q)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Add comment:%d\n", r.ID)

	r, err = comment.Delete(r.ID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Delete comment:%d\n", r.ID)

}

func TestCommentUpdate(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	// Backlog APIのリクエストパラメーターを作成する。
	q := NewCommentRequest()
	q.Content("コメント情報を更新します。6")
	comment := NewComment(space, apiKey, "TESTTOOL-64")
	id := "13562471"
	_, err = comment.Update(id, q)
	if err != nil {
		t.Fatal(err)
	}
}
