package i18n

import (
	"fmt"
	"testing"
)

func TestT(t *testing.T) {
	result, err := T("en", "msg.hello")
	if err != nil {
		fmt.Printf("Fail to T, err: %v", err)
	}
	fmt.Printf("[en] msg.hello => %v", result)

	result, err := T("zh-CN", "msg.hello")
	if err != nil {
		fmt.Printf("Fail to T, err: %v", err)
	}
	fmt.Printf("[zh-CN] msg.hello => %v", result)
}