package expressions

import (
	"fmt"
)

type Binary struct {
	IsPriority bool
	Left       Expression
	Right      Expression
	Operator   string
}

func (b *Binary) GetBody() string {
	result := fmt.Sprintf("%s %s %s", b.Left.GetBody(), b.Operator, b.Right.GetBody())

	if b.IsPriority {
		return fmt.Sprintf("(%s)", result)
	}

	return result
}
