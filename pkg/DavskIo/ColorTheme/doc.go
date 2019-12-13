////////////////////////////////////////////////////////////////////////////////
// doc.go                                                                      /
//                                                                             /
// by david                                                                    /
// on 12/13/19, 8:56 AM                                                        /
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

// Package ColorTheme defines standard options for DavskApps for console, gui, print, etc.
//
// Overview
//
// When an app begins, a new default ColorTheme will be created,
//
// ColorTheme options may then be loaded from local or cloud profile
// or the user may select the options using DavskIo, from the console, gui, web or other interface,
// with the option of saving changes to local and or cloud.
//
// ColorTheme will then be used by DavskIo interface to provide appropriate colors
// for specific uses that accommodate the options selected.
//
// Best Practices
//
// ColorTheme should only be used by the implementation of DavskIo interface
// and should be altered directly within the app.
//
// The programmer should avoid specifying colors directly, but should specify specific color sets within the theme
// when using the DavskIo interface. So a PostItNote may mean Yellow, ErrorMsg may mean Red, InfoNote may mean Blue,
// but the programmer does not choose the color as the color needs to accommodate the hardware and user requirements
// and comply with branding.
//
// Copyright © 2019. Davsk Ltd. Co.                                            /
// All Rights Reserved. Licensed under the MIT License.
package ColorTheme
