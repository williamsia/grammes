// Copyright (c) 2018 Northwestern Mutual.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

/*
Package token contains the object to define parts of a vertex.

Using a token to create a vertex can allow you to choose a custom
Key, ID, Label, or Value before creating it.

A note about Token:

This object implements the Parameter interface used by graph traversals.
*/
package token

// Token allows for more concise
// Traversal definitions.
type Token string

const (
	// ID represents Element.id()
	ID Token = "T.id"
	// Key represents Property.key()
	Key Token = "T.key"
	// Label represents Element.label()
	Label Token = "T.label"
	// Value represents Property.value()
	Value Token = "T.value"
)

func (t Token) String() string {
	return string(t)
}
