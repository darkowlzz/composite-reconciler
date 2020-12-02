package v1

import (
	"fmt"
	"testing"
	"time"

	"github.com/darkowlzz/composite-reconciler/operate/v1/operand"
)

func TestCompositeOperatorOrder(t *testing.T) {
	secretName := "secret"
	configmapName := "configmap"
	daemonsetName := "daemonset"
	deploymentAName := "deploymentA"
	deploymentBName := "deploymentB"

	operands := []*operand.Operand{
		&operand.Operand{
			Name: secretName,
		},
		&operand.Operand{
			Name: configmapName,
		},
		&operand.Operand{
			Name:     daemonsetName,
			Requires: []string{secretName, configmapName},
		},
		&operand.Operand{
			Name:     deploymentAName,
			Requires: []string{secretName},
		},
		&operand.Operand{
			Name:     deploymentBName,
			Requires: []string{deploymentAName},
		},
	}

	wantOrder := `[
  0: [ configmap secret ]
  1: [ daemonset deploymentA ]
  2: [ deploymentB ]
]`

	co, err := NewCompositeOperator(operands...)
	if err != nil {
		t.Errorf("unexpected error while creating a composite operator: %v", err)
	}
	od := co.Order()
	if od.String() != wantOrder {
		t.Errorf("unexpected operator order:\n\t(WNT) %q\n\t(GOT) %q", wantOrder, od.String())
	}
}

func TestCompositeOperatorEnsure(t *testing.T) {
	operandA := &operand.Operand{
		Name: "opA",
		Ensure: func() error {
			fmt.Println("RUNNING opA")
			return nil
			// return errors.New("some error for opA")
		},
		Delete: func() error {
			fmt.Println("DELETING opA")
			return nil
		},
	}

	operandB := &operand.Operand{
		Name: "opB",
		Ensure: func() error {
			time.Sleep(1 * time.Second)
			fmt.Println("RUNNING opB")
			return nil
			// return errors.New("some error for opB")
		},
		Delete: func() error {
			fmt.Println("DELETING opB")
			return nil
		},
	}

	operandC := &operand.Operand{
		Name:     "opC",
		Requires: []string{operandA.Name},
		Ensure: func() error {
			fmt.Println("RUNNING opC")
			return nil
		},
		Delete: func() error {
			fmt.Println("DELETING opC")
			return nil
		},
	}

	// tests := []struct {
	//     name string
	//     operands []*operand.Operand
	//     wantRequeue bool
	// }{
	//     {

	//     },
	// }

	co, err := NewCompositeOperator(operandA, operandB, operandC)
	if err != nil {
		t.Errorf("unexpected error while creating a composite operator: %v", err)
	}

	res, eve, eerr := co.Ensure()
	fmt.Println("EERR:", eerr)
	fmt.Println("RES:", res)
	fmt.Println("EVENT:", eve)

	cres, ceve, ceerr := co.Cleanup()
	fmt.Println("EERR:", ceerr)
	fmt.Println("RES:", cres)
	fmt.Println("EVENT:", ceve)

}