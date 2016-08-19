package api

import "context"

// Api is a fundamental building block of servers and clients.
// It represents a single RPC method. It is a callback function signature.
//
// Or else, think of it as a schematic/diagram to build a f***ing gun. In order,
// to build a gun, you need to implement the details.
//
// In functional-programming world, we call superCharge. You supercharge with
// http-details, you get a f***ing gun that capable of firing http-bullets.
//
// For OOP pals, you may consider this as a mini-version of interface, rule's
// the same: implement first before actual usage.
type Api func(ctx context.Context, request interface{}) (response interface{}, err error)
