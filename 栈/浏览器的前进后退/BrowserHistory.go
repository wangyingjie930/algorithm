package 浏览器的前进后退

type BrowserHistory struct {
	urls []string
	top int
	cur int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{urls: []string{homepage}, top: 0, cur: 0}
}


func (this *BrowserHistory) Visit(url string)  {
	this.top = this.cur
	if len(this.urls) > this.top + 1 {
		this.urls[this.top + 1] = url
	}else {
		this.urls = append(this.urls, url)
	}
	this.top++
	this.cur++
}


func (this *BrowserHistory) Back(steps int) string {
	this.cur = this.cur - steps
	if this.cur < 0 {
		this.cur = 0
		return this.urls[0]
	}
	return this.urls[this.cur]
}


func (this *BrowserHistory) Forward(steps int) string {
	this.cur = this.cur + steps
	if this.cur > this.top {
		this.cur = this.top
		return this.urls[this.top]
	}
	return this.urls[this.cur]
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
