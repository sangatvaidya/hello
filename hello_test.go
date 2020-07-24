package hello

import "testing"

func TestHello(t *testing.T) {
    expected := "Hello, world!!!"
	if ret := Hello(); ret != expected {
	  t.Errorf("Hello()=%q, want %q",ret,expected)
	}  
	
	}