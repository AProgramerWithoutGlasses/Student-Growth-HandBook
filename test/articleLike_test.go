package test

import (
	"studentGrow/service/article"
	"testing"
)

func TestArticleLike(t *testing.T) {
	err := article.Like("221", "20211524201", 0)
	if err != nil {
		t.Fatalf("执行错误")
	}
	t.Logf("执行成功")
}
