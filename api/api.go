package api

import "context"

// Api is a fundamental building block of servers and clients.
// It represents a single RPC method. It is a callback function signature.
//
// Or else, think of it as a schematic/diagram to build a fucking gun. In order,
// to build a gun, you need to implement the details.
//
// In functional-programming term, we call superCharge. You supercharge with http-details,
// you get a fucking gun that capable of firing http-bullets. Pretty cool, huh. ;)
type Api func(ctx context.Context, request interface{}) (response interface{}, err error)
