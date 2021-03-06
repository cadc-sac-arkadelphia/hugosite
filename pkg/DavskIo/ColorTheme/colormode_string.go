// Code generated by "stringer -type=ColorMode,ColorBlind"; DO NOT EDIT.

package ColorTheme

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[LightMode-0]
	_ = x[DarkMode-1]
}

const _ColorMode_name = "LightModeDarkMode"

var _ColorMode_index = [...]uint8{0, 9, 17}

func (i ColorMode) String() string {
	if i < 0 || i >= ColorMode(len(_ColorMode_index)-1) {
		return "ColorMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ColorMode_name[_ColorMode_index[i]:_ColorMode_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Normal-0]
	_ = x[Achromatopsia-1]
	_ = x[Protanopia-2]
	_ = x[Deuteranopia-3]
	_ = x[Tritanopia-4]
	_ = x[Protanomaly-5]
	_ = x[Deuteranomaly-6]
	_ = x[Tritanomaly-7]
}

const _ColorBlind_name = "NormalAchromatopsiaProtanopiaDeuteranopiaTritanopiaProtanomalyDeuteranomalyTritanomaly"

var _ColorBlind_index = [...]uint8{0, 6, 19, 29, 41, 51, 62, 75, 86}

func (i ColorBlind) String() string {
	if i < 0 || i >= ColorBlind(len(_ColorBlind_index)-1) {
		return "ColorBlind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ColorBlind_name[_ColorBlind_index[i]:_ColorBlind_index[i+1]]
}
