package api

import (
	ports "Arithmetic/internal/ports"
)

type Adapter struct {
	db         ports.DbPort
	arithmetic ports.ArithmeticPort
}

func NewAdapter(db ports.DbPort, arithmetic ports.ArithmeticPort) *Adapter {
	return &Adapter{db: db, arithmetic: arithmetic}
}

func (apia Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := apia.arithmetic.Addition(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(answer, "addition")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (apia Adapter) GetSubtraction(a, b int32) (int32, error) {
	answer, err := apia.arithmetic.Subtraction(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(answer, "subtraction")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (apia Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := apia.arithmetic.Multiplication(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(answer, "multiplication")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (apia Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := apia.arithmetic.Division(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(answer, "division")
	if err != nil {
		return 0, err
	}

	return answer, nil
}
