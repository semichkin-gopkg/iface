package iface

// AssertImpl checks source implements Interface.
//
//	var _ = AssertImpl[SomeInterface](&SomeImplementation{})
func AssertImpl[Interface any](source Interface) any { return struct{}{} }
