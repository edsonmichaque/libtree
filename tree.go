package pungue

type Tree interface {
	Inserte(Key)
	Traverse(Mode, func(Key))
}
