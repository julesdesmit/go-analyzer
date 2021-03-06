// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT License was not distributed with this
// file, you can obtain one at https://opensource.org/licenses/MIT.

// Package lints defines all available lints for go-analyzer.
package lints

// Fn defines the standard function signature for a lint.
type Fn func(string) []error
