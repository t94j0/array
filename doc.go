// Package array does mathematical array-related operations
//
// Why
//
// I find myself doing things like checking if a value is in an array a lot.
// Since Go has the ability to do reflection very easily, I made this project
// to only write array utilities once.
//
// The inspiration for this project comes from SQL databases. I like the union,
// intersect, and except operators, so this package will implement them.
//
// Usage
//
// See the examples for usage guides. Always remmeber to unpack your arrays
// so that you don't try to do array work with interfaces! The Go compiler
// should throw errors if you try.
package array
