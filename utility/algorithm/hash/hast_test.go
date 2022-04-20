package hash

import (
	"fmt"
	"testing"
)

func Test_Hashing_md5(t *testing.T) {
	result := Hashing_md5("hello world")
	if result != "5eb63bbbe01eeed093cb22bb8f5acdc3" {
		t.Log(fmt.Sprintf("Expected the hash '5eb63bbbe01eeed093cb22bb8f5acdc3' got '%s'", result))
		t.Fail()
	}
}

func Test_Hashing_sha1(t *testing.T) {
	result := Hashing_sha1("hello world")
	if result != "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed" {
		t.Log(fmt.Sprintf("Expected the hash '2aae6c35c94fcfb415dbe95f408b9ce91ee846ed' got '%s'", result))
		t.Fail()
	}
}

func Test_Hashing_sha224(t *testing.T) {
	result := Hashing_sha224("hello world")
	if result != "2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b" {
		t.Log(fmt.Sprintf("Expected the hash '2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b' got '%s'", result))
		t.Fail()
	}
}

func Test_Hashing_sha256(t *testing.T) {
	result := Hashing_sha256("hello world")
	if result != "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9" {
		t.Log(fmt.Sprintf("Expected the hash 'b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9' got '%s'", result))
		t.Fail()
	}
}

func Test_Hashing_sha384(t *testing.T) {
	result := Hashing_sha384("hello world")
	if result != "fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd" {
		t.Log(fmt.Sprintf("Expected the hash 'fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd' got '%s'", result))
		t.Fail()
	}
}

func Test_Hashing_sha512(t *testing.T) {
	result := Hashing_sha512("hello world")
	if result != "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f" {
		t.Log(fmt.Sprintf("Expected the hash '309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f' got '%s'", result))
		t.Fail()
	}
}

func Test_Hashing_sha3_224(t *testing.T) {
	result := Hashing_sha3_224("hello world")
	if result != "dfb7f18c77e928bb56faeb2da27291bd790bc1045cde45f3210bb6c5" {
		t.Log(fmt.Sprintf("Expected the hash 'dfb7f18c77e928bb56faeb2da27291bd790bc1045cde45f3210bb6c5' got '%s'", result))
		t.Fail()
	}
}
func Test_Hashing_sha3_256(t *testing.T) {
	result := Hashing_sha3_256("hello world")
	if result != "644bcc7e564373040999aac89e7622f3ca71fba1d972fd94a31c3bfbf24e3938" {
		t.Log(fmt.Sprintf("Expected the hash '644bcc7e564373040999aac89e7622f3ca71fba1d972fd94a31c3bfbf24e3938' got '%s'", result))
		t.Fail()
	}
}
func Test_Hashing_sha3_384(t *testing.T) {
	result := Hashing_sha3_384("hello world")
	if result != "83bff28dde1b1bf5810071c6643c08e5b05bdb836effd70b403ea8ea0a634dc4997eb1053aa3593f590f9c63630dd90b" {
		t.Log(fmt.Sprintf("Expected the hash '83bff28dde1b1bf5810071c6643c08e5b05bdb836effd70b403ea8ea0a634dc4997eb1053aa3593f590f9c63630dd90b' got '%s'", result))
		t.Fail()
	}
}

func Test_Hashing_sha3_512(t *testing.T) {
	result := Hashing_sha3_512("hello world")
	if result != "840006653e9ac9e95117a15c915caab81662918e925de9e004f774ff82d7079a40d4d27b1b372657c61d46d470304c88c788b3a4527ad074d1dccbee5dbaa99a" {
		t.Log(fmt.Sprintf("Expected the hash '840006653e9ac9e95117a15c915caab81662918e925de9e004f774ff82d7079a40d4d27b1b372657c61d46d470304c88c788b3a4527ad074d1dccbee5dbaa99a' got '%s'", result))
		t.Fail()
	}
}

// The tests for blake2s and blake2b's will always fail as we utilize a random generated key
// Therefore I've choosen to opt out out testing them and instead hope that they are implemented correctly
