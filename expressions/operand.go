package expressions

import (
	"fmt"
)

type Operand struct {
	IsPriority bool
	Data       string
}

func (o *Operand) GetBody() string {
	if o.IsPriority {
		return fmt.Sprintf("(%s)", o.Data)
	}

	return o.Data
}
