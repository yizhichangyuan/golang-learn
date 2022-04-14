package storage2

import (
	"strings"
	"testing"
)

func TestCheckQuota(t *testing.T) {
	// 在该测试函数完成后进行重置，以防影响其他测试函数
	saved := notifyUser
	defer func() {
		notifyUser = saved
	}()
	var notifiedUser, notifiedMsg string
	// 通过检查上述两个变量是否为空，测试是否到了发邮件的流程，或者检查测试的信息是否正确
	// 使用函数值的好处是可以替换而不需要修改原函数内部逻辑
	notifyUser = func(username, msg string) {
		notifiedUser, notifiedMsg = username, msg
	}

	const user = "joe@example.org"
	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}
	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s", notifiedUser, user)
	}
	const wantSubString = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubString) {
		t.Errorf("unexpected notification message <<%s>>, "+
			"want substring %q", notifiedMsg, wantSubString)
	}
}
