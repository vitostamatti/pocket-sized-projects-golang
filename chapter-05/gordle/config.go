package gordle

import (
	"bufio"
	"io"
)

type ConfigFunc func(g *Game) error

func WithReader(reader io.Reader) ConfigFunc {
	return func(g *Game) error {
		g.reader = bufio.NewReader(reader)
		return nil
	}
}

func WithSolution(solution []rune) ConfigFunc {
	return func(g *Game) error {
		g.solution = solution
		return nil
	}
}

func WithMaxAttempts(maxAttempts int) ConfigFunc {
	return func(g *Game) error {
		g.maxAttempts = maxAttempts
		return nil
	}
}
