package main

import "fmt"

type repo struct {
	pos Position
}

//New create new Services
func New(p Position) (Services, error) {
	if p.Direction == "" {
		return nil, fmt.Errorf("Wrong Format")
	}
	return &repo{
		pos: p,
	}, nil
}

func (r *repo) Left(pos *Position) error {
	switch pos.Direction {
	case "N":
		pos.Direction = "W"
	case "W":
		pos.Direction = "S"
	case "S":
		pos.Direction = "E"
	case "E":
		pos.Direction = "N"
	default:
		return fmt.Errorf("Wrong Format")
	}
	return nil
}

func (r *repo) Right(pos *Position) error {
	switch pos.Direction {
	case "N":
		pos.Direction = "E"
	case "W":
		pos.Direction = "N"
	case "S":
		pos.Direction = "W"
	case "E":
		pos.Direction = "S"
	default:
		return fmt.Errorf("Wrong Format")
	}
	return nil
}

func (r *repo) Forward(pos *Position) error {
	switch pos.Direction {
	case "N":
		pos.Y += 1
	case "W":
		pos.X -= 1
	case "S":
		pos.Y -= 1
	case "E":
		pos.X += 1
	default:
		return fmt.Errorf("Wrong Format")
	}
	return nil
}
