package backlogcli

import (
	"fmt"

	"github.com/tmcna/backlog/pkg/client"
)

// CommentAdd function comment add subcommand.
func CommentAdd(issueKey string, content string) error {
	cfg, err := client.NewConfig()
	if err != nil {
		err = fmt.Errorf("configuration error, %s", err)
		return err
	}

	// Create request parameters for Backlog API.
	q := client.NewCommentRequest()
	q.Content(content)

	// Commentオブジェクトを作成し、リクエストパラメーターを設定する。
	comment := client.NewComment(cfg.Space, cfg.APIKey, issueKey)
	_, err = comment.Add(q)
	if err != nil {
		return err
	}

	return nil
}
