package main

import (
	
	"testing"
	"unicode/utf8"
)

func TestSum(t *testing.T) {
got := Sum(10, 20)

expected := 30

if got != expected {
	t.Fatalf("ne werni otwet: %d want %d", got, expected)
}


}

//---------------------------------

type Test struct {
	in  int
	out string
}

var tests = []Test{
	{-1, "negative"},
	{5, "short"},
	{0, "zero"},
	{101, "very long"},
	{55, "long"},
}


func TestLength(t *testing.T) {
	for i, test := range tests {
		size := Length(test.in)
		if size != test.out {
			t.Errorf("#%d: Size(%d)=%s; want %s", i, test.in, size, test.out)
		}
	}
}


//-----------------------------------------------------

func TestMultiply(t *testing.T){
	got := Multiply(10, 20)

	expected := 200
	
	if got != expected {
		t.Fatalf("ne werni otwet: %d want %d", got, expected)
	}
	
}

//--------------------------------------------------------

func TestDeleteVowels(t *testing.T) {
	in := DeleteVowels("fytaeuiolnbcf")

	out := "fytlnbcf"

	if in != out {
		t.Fatalf("ne werni otwet: %s want %s", in, out)
	}
}

//_______________________________________________________


type TestUTF struct{
	slbyte []byte;
	out int;
	outErr error

}



func TestGetUTFLength(t *testing.T) {
	testes := []TestUTF{
   {[]byte{48,49,50,51,52}, 5, nil},
   {[]byte{51,52}, 2, nil},
   {[]byte{0xff, 0xff}, 0, ErrInvalidUTF8},
   }
   
   for _, v := range testes {
	   res, err := GetUTFLength(v.slbyte)
   
   if res != v.out{
	   t.Fatalf("ne werni otwet: %d want %d",utf8.RuneCount(v.slbyte),v.out)
   }
   if err != v.outErr{
	   t.Fatalf("errors %s", err)
   }
   
   }
   
   }