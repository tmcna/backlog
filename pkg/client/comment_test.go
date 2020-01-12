package client

import (
	"fmt"
	"testing"
)

func TestComment_List(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	comment := NewComment(cfg.Space, cfg.APIKey, "TESTTOOL-64")
	r, err := comment.List()
	comment.PrintCSV(r)
}

func TestComment_AddDelete(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	q := NewCommentRequest()
	q.Content("コメント登録のテスト")

	comment := NewComment(cfg.Space, cfg.APIKey, "TESTTOOL-64")
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

func TestComment_Update(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	// Backlog APIのリクエストパラメーターを作成する。
	q := NewCommentRequest()
	q.Content("コメント情報を更新します。6")
	comment := NewComment(cfg.Space, cfg.APIKey, "TESTTOOL-64")
	id := "13562471"
	_, err = comment.Update(id, q)
	if err != nil {
		t.Fatal(err)
	}
}
