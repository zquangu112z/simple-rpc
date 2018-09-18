package shared

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith interface {
	Multiply(args Args, reply *int) error
	Divide(args Args, quo *Quotient) error
	Greet(guessname string, greeting *string) error
}
