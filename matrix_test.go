package main

import (
	"reflect"
	"testing"
)

func TestInverse(t *testing.T) {
	m := Matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expRes := Matrix{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}

	res := m.inverse()

	if !reflect.DeepEqual(res, expRes) {
		t.Errorf("Inversed matrix is incorrect.")
	}
}

func TestFlatten(t *testing.T) {
	m := Matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	res := m.flatten()

	if !reflect.DeepEqual(res, expArr) {
		t.Errorf("Flattened matrix is incorrect.")
	}
}

func TestAggregateSum(t *testing.T) {
	m := Matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expSum := 45
	expMult := 362880

	res := m.aggregate(sum, 0)

	if res != expSum {
		t.Errorf("Matrix sum is incorrect.")
	}

	res = m.aggregate(mult, 1)

	if res != expMult {
		t.Errorf("Matrix product is incorrect.")
	}
}

func TestValidate(t *testing.T) {
	mEmpty := Matrix{}
	mNotSquare := Matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}

	res := mEmpty.validate()

	if res != matrixErrorEmpty {
		t.Errorf("Matrix validation should have returned an empty matrix error.")
	}

	res = mNotSquare.validate()

	if res != matrixErrorNotSquare {
		t.Errorf("Matrix validation should have returned a non-square matrix error.")
	}
}

func TestIsSquare(t *testing.T) {
	m := Matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	mNotSquare := Matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}

	res := m.IsSquare()

	if !res {
		t.Errorf("Matrix is not evaluated as a square but it is one.")
	}

	res = mNotSquare.IsSquare()

	if res {
		t.Errorf("Matrix is evaluated as a square but it isn't one.")
	}
}
