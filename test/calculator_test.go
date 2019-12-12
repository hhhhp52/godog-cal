package calculator_test

import (
	"fmt"

	"github.com/DATA-DOG/godog"
	calc "github.com/hhhhp52/godog-cal"
)

type CalcSuite struct {
	*godog.Suite
	calc *calc.Cal
}

func (cs *CalcSuite) calculatorIsCleared() error {
	cs.calc.Clear()
	return nil
}

func (cs *CalcSuite) iPress(num int) error {
	cs.calc.Press(num)
	return nil
}

func (cs *CalcSuite) iAdd(num int) error {
	cs.calc.Add(num)
	return nil
}

func (cs *CalcSuite) iSubtracr(num int) error {
	cs.calc.Sub(num)
	return nil
}

func (cs *CalcSuite) theResultShouldBe(num int) error {
	result := cs.calc.Result()
	if result == num {
		return nil
	}
	return fmt.Errorf("%d doesn't match expectation: %d", result, num)
}

func (cs *CalcSuite) iMultiplyBy(num int) error {
	cs.calc.MultipleBy(num)
	return nil
}

func FeatureContext(suite *godog.Suite) {
	s := CalcSuite{
		calc:  new(calc.Cal),
		Suite: suite,
	}
	suite.Step(`^calculator is cleared$`, s.calculatorIsCleared)
	suite.Step(`^i press (\d+)$`, s.iPress)
	suite.Step(`^i add (\d+)$`, s.iAdd)
	suite.Step(`^i subtracr (\d+)$`, s.iSubtracr)
	suite.Step(`^the result should be (\d+)$`, s.theResultShouldBe)
}
