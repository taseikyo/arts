/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-24 15:39:07
 * @link    github.com/taseikyo
 */

type CQueue struct {
	BookList []int
}

func Constructor() CQueue {
	return CQueue{
		BookList: []int{},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.BookList = append(this.BookList, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.BookList) == 0 {
		return -1
	}
	book := this.BookList[0]
	this.BookList = this.BookList[1:]
	return book
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
