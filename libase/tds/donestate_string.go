// Code generated by "stringer -type=DoneState"; DO NOT EDIT.

package tds

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TDS_DONE_FINAL-0]
}

const _DoneState_name = "TDS_DONE_FINAL"

var _DoneState_index = [...]uint8{0, 14}

func (i DoneState) String() string {
	if i >= DoneState(len(_DoneState_index)-1) {
		return "DoneState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DoneState_name[_DoneState_index[i]:_DoneState_index[i+1]]
}
