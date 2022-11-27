package 浏览器的前进后退

import "testing"

func TestBrowserHistory_Visit(t *testing.T) {
	browserHistory := Constructor("leetcode.com")
	browserHistory.Visit("google.com")
	browserHistory.Visit("facebook.com")

	t.Log(browserHistory.Back(1))
	t.Log(browserHistory.Back(1))
	t.Log(browserHistory.Forward(2))
	t.Log(browserHistory.Back(2))

	browserHistory.Visit("linkedin.com")
	t.Log(browserHistory.Back(1))
	t.Log(browserHistory.Forward(2))
}
