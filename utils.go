package main

func first_arg[T, U any](val T, _ U) T {
	return val
}
