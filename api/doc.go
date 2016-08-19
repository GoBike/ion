// Package api defines an abstraction for RPCs.
//
// Conceptually, an API can be generalised as function(input) output. In
// particular,
//
// func api(input) output {
//   return any_function_at_all(input)
// }
//
// APIs are a *fundamental building block* for many Go ion components. APIs are
// implemented by servers, and called by clients.
package api
