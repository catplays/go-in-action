package concurrent

/**
int list 链表接口
 */
type  IIntList interface {
	// contain a value or not
	Contains (value int) bool
	// insert a value if insert successfully returns true, otherwise return false
	Insert ( value int) bool
	// delete a value
	Delete (value int) bool
	// return length of list
	Len() int
	// traverse list，if f return false, stop
	Range ( f func(value int) bool)
}
