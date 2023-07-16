package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime by definition!"},
		{"one", 1, false, "1 is not prime by definition!"},
		{"negative", -11, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}
		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}
		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}

	}
}

func Test_prompt(t *testing.T) {
	//save a copy of os.Stdout
	oldOut := os.Stdout

	// Create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	prompt()

	// Close our writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	//perform test
	if string(out) != "-> " {
		t.Errorf("incorrect prompt expected -> but got %s", string(out))
	}
}

func Test_intro(t *testing.T) {
	//save a copy of os.Stdout
	oldOut := os.Stdout

	// Create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	intro()

	// Close our writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	//perform test
	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("intro test not correct: got %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a whole number!"},
		{name: "zero", input: "0", expected: "0 is not prime by definition!"},
		{name: "one", input: "1", expected: "1 is not prime by definition!"},
		{name: "negative", input: "-111", expected: "Negative numbers are not prime, by definition!"},
		{name: "eight", input: "8", expected: "8 is not a prime number because it is divisible by 2!"},
		{name: "seven", input: "7", expected: "7 is a prime number!"},
		{name: "typed", input: "eight", expected: "Please enter a whole number!"},
		{name: "decimal", input: "3.3", expected: "Please enter a whole number!"},
		{name: "quit", input: "q", expected: ""},
		{name: "QUIT", input: "Q", expected: ""},
	}

	for _, e := range tests {

		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: expected %s, but got %s", e.name, e.expected, res)
		}
	}
}

func Test_readUserInput(t *testing.T) {
	doneChan := make(chan bool)

	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)

}
