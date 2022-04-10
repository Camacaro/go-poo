package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)

	if total != 10 {
		t.Errorf("Sum(5, 5) failed. Got %d, expected %d", total, 10)
		t.Error("El test fallo")
	}

	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{3, 4, 7},
	}

	for _, table := range tables {
		total := Sum(table.a, table.b)

		if total != table.n {
			t.Errorf("Sum(%d, %d) failed. Got %d, expected %d", table.a, table.b, total, table.n)
		}
	}
}

func TestGetMax(t *testing.T) {
	tables := []struct {
		firstValue     int
		secondValue    int
		resultExpected int
	}{
		{1, 2, 2},
		{2, 3, 3},
		{5, 4, 5},
	}

	for _, table := range tables {
		result := GetMax(table.firstValue, table.secondValue)

		if result != table.resultExpected {
			t.Errorf("GetMax(%d, %d) failed. Got %d, expected %d", table.firstValue, table.secondValue, result, table.resultExpected)
		}
	}
}

func TestFib(t *testing.T) {
	tables := []struct {
		firstValue     int
		resultExpected int
	}{
		{1, 1},
		{8, 21},
		// {50, 12586269025},
	}

	for _, table := range tables {
		result := Fibonacci(table.firstValue)

		if result != table.resultExpected {
			t.Errorf("Fibonacci(%d) failed. Got %d, expected %d", table.firstValue, result, table.resultExpected)
		}
	}
}
