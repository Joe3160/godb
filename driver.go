package godb

type Driver string

func (d Driver) String() string {
	return string(d)
}

