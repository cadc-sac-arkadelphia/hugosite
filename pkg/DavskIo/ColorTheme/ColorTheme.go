////////////////////////////////////////////////////////////////////////////////
// ColorTheme.go                                                               /
//                                                                             /
// by david                                                                    /
// on 12/9/19, 6:00 PM                                                         /
// for hugosite                                                                /
//                                                                             /
// Copyright © 2019. Davsk Ltd. Co.                                            /
// All Rights Reserved. Licensed under the MIT License.                        /
//                                                                             /
// Permission is hereby granted, free of charge, to any person obtaining a     /
// copy of this software and associated documentation files (the "Software"),  /
// to deal in the Software without restriction, including without limitation   /
// the rights to use, copy, modify, merge, publish, distribute, sublicense,    /
// and/or sell copies of the Software, and to permit persons to whom the       /
// Software is furnished to do so, subject to the following conditions:        /
//                                                                             /
// The above copyright notice and this permission notice shall be included in  /
// all copies or substantial portions of the Software.                         /
//                                                                             /
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS     /
// OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, /
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE /
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER      /
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING     /
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER         /
// DEALINGS IN THE SOFTWARE.                                                   /
////////////////////////////////////////////////////////////////////////////////

package ColorTheme

// Type colorMode is a local boolean of light/dark options.
type ColorMode int

const (
	LightMode ColorMode = iota // LightMode
	DarkMode  ColorMode = iota // DarkMode
)

// Type colorBlind is a local enum list of colorBlind options.
type ColorBlind int

const (
	Normal        ColorBlind = iota
	Achromatopsia ColorBlind = iota
	Protanopia    ColorBlind = iota
	Deuteranopia  ColorBlind = iota
	Tritanopia    ColorBlind = iota
	Protanomaly   ColorBlind = iota
	Deuteranomaly ColorBlind = iota
	Tritanomaly   ColorBlind = iota
)

// Install 'go get -u golang.org/x/tools/...' if stringer is not found.
//go:generate stringer -type=ColorMode,ColorBlind

// Type ColorTheme is the structure for defining our apps color theme
// as selected by user consistent with branding and psychology.
type ColorTheme struct {
	ColorMode  ColorMode
	ColorBlind ColorBlind
}

// Func NewColorTheme returns a pointer to a new ColorTheme with default values.
func NewColorTheme() *ColorTheme {
	return &ColorTheme{LightMode, Normal}
}

// Func String returns a text representation of the ColorTheme.
func (ct *ColorTheme) String() string {
	return "ColorTheme{ColorMode: " + ColorMode.String(ct.ColorMode) + ", ColorBlind: " + ColorBlind.String(ct.ColorBlind) + "}"
}
