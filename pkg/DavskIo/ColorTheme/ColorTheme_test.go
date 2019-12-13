////////////////////////////////////////////////////////////////////////////////
// ColorTheme_test.go                                                          /
//                                                                             /
// by david                                                                    /
// on 12/9/19, 6:05 PM                                                         /
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

import (
	"reflect"
	"testing"
)

func TestColorTheme_String(t *testing.T) {
	type fields struct {
		colorMode  ColorMode
		colorBlind ColorBlind
	}
	var tests = []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{name: "Default", fields: fields{colorMode: LightMode, colorBlind: Normal}, want: "Oops."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			co := &ColorTheme{
				ColorMode:  tt.fields.colorMode,
				ColorBlind: tt.fields.colorBlind,
			}
			if got := co.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewColorTheme(t *testing.T) {
	tests := []struct {
		name string
		want *ColorTheme
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewColorTheme(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewColorTheme() = %v, want %v", got, tt.want)
			}
		})
	}
}
