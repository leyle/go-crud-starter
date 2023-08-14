package smsapp

import "testing"

func TestGenerateRandomCode(t *testing.T) {
	code := GenerateRandomCode(4)
	t.Log(code)
}

func TestGenerateRandomCode2(t *testing.T) {
	code := GenerateRandomCode(6)
	t.Log(code)
}

func TestGenerateRandomCode3(t *testing.T) {
	code := GenerateRandomCode(3)
	t.Log(code)
}

func TestGenerateRandomCode4(t *testing.T) {
	code := GenerateRandomCode(9)
	t.Log(code)
}
