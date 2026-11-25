//go:build !windows

package main

func enterGUI(bool) (func(), error) { return func() {}, nil }
