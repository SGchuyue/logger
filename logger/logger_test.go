package mylog_test

import "testing"

func TestSimpleHttpGet(t *testing.T) {
	ge,err := SimpleHttpGet("https://www.liwenzhou.com/posts/Go/zap/")

}
