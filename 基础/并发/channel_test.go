/**
  @author: wangyingjie
  @since: 2023/4/16
  @desc:
**/

package concurrency

func Stop(stop <-chan bool) {
	//close(stop) //error: 有方向的 channel 不可以被关闭
}
