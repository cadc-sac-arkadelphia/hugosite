////////////////////////////////////////////////////////////////////////////////
// ColorMode.go                                                                /
//                                                                             /
// by david                                                                    /
// on 12/13/19, 8:24 AM                                                        /
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

// Install 'go get -u golang.org/x/tools/...' if stringer is not found.
//go:generate stringer -type=ColorMode
