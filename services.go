package main

//Services interface
type Services interface {
	Left(pos *Position) error
	Right(pos *Position) error
	Forward(pos *Position) error
}
