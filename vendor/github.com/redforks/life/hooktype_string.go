// generated by stringer -type=hookType; DO NOT EDIT

package life

import "fmt"

const _hookType_name = "BeforeStartingBeforeRunningBeforeShutingdownAbort"

var _hookType_index = [...]uint8{0, 14, 27, 44, 49}

func (i hookType) String() string {
	if i < 0 || i+1 >= hookType(len(_hookType_index)) {
		return fmt.Sprintf("hookType(%d)", i)
	}
	return _hookType_name[_hookType_index[i]:_hookType_index[i+1]]
}
