package contains

import "testing"

func Test_StartsWith(t *testing.T) {
	result := StartsWith("123meJason", []string{"123"})
	if !result {
		t.Log("Result should be true! Got false")
		t.Fail()
	}

	result = StartsWith("meJason123", []string{"123"})
	if result {
		t.Log("Result should be false! Got true")
		t.Fail()
	}
}

func Test_EndsWith(t *testing.T) {
	result := EndsWith("123meJason", []string{"123"})
	if result {
		t.Log("Result should be false! Got true")
		t.Fail()
	}

	result = EndsWith("meJason123", []string{"123"})
	if !result {
		t.Log("Result should be true! Got false")
		t.Fail()
	}
}

func Test_Contains(t *testing.T) {
	result := Contains("123meJason", "123")
	if !result {
		t.Log("Result should be true! Got false")
		t.Fail()
	}
	result = Contains("meJason123", "123")
	if !result {
		t.Log("Result should be true! Got false")
		t.Fail()
	}
	result = Contains("me123Jason", "123")
	if !result {
		t.Log("Result should be true! Got false")
		t.Fail()
	}
	result = Contains("meJason123", "meJason")
	if !result {
		t.Log("Result should be true! Got false")
		t.Fail()
	}

	result = Contains("meJason", "123")
	if result {
		t.Log("Result should be false! Got true")
		t.Fail()
	}

	result = Contains("meJason", "mek")
	if result {
		t.Log("Result should be false! Got true")
		t.Fail()
	}
}
