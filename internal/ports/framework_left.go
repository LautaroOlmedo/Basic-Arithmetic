package ports

import (
	_ "context"
)

type GRPCPort interface {
	RUn()
	GetAddition()
	GetSubtraction()
	GetMultiplication()
	GetDivision()
}
