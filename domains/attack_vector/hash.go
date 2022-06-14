package attack_vector

import (
	"fmt"
	"strings"

	"github.com/s9rA16Bf4/go-evil/utility/algorithm/hash"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

type hash_t struct {
	hash_func string // The hash we will utilize
}

var curr_hash hash_t

func Set_hash(hash_to_use string) {
	hash_to_use = strings.ReplaceAll(hash_to_use, "\"", "")
	var available_hashes = []string{"md5", "sha1", "sha224", "sha256", "sha384", "sha512", "sha3_224", "sha3_256", "sha3_384", "sha3_512",
		"blake2s_256", "blake2b_256", "blake2b_384", "blake2b_512"}

	found := false
	for _, hash := range available_hashes {
		if hash == hash_to_use {
			found = true
			break
		}
	}
	if !found {
		notify.Error(fmt.Sprintf("Unknown hash %s", hash_to_use), "attack_vector.Set_hash()")
	}
	curr_hash.hash_func = hash_to_use
}

func Hash(msg string) string {
	msg = strings.ReplaceAll(msg, "\"", "")
	var toReturn string
	switch curr_hash.hash_func {
	case "md5":
		toReturn = hash.Hashing_md5(msg)
	case "sha1":
		toReturn = hash.Hashing_sha1(msg)
	case "sha224":
		toReturn = hash.Hashing_sha224(msg)
	case "sha256":
		toReturn = hash.Hashing_sha256(msg)
	case "sha384":
		toReturn = hash.Hashing_sha384(msg)
	case "sha512":
		toReturn = hash.Hashing_sha512(msg)
	case "sha3_224":
		toReturn = hash.Hashing_sha3_224(msg)
	case "sha3_256":
		toReturn = hash.Hashing_sha3_256(msg)
	case "sha3_384":
		toReturn = hash.Hashing_sha3_384(msg)
	case "sha3_512":
		toReturn = hash.Hashing_sha3_512(msg)
	case "blake2s_256":
		toReturn = hash.Hashing_blake2s_256(msg)
	case "blake2b_256":
		toReturn = hash.Hashing_blake2b_256(msg)
	case "blake2b_384":
		toReturn = hash.Hashing_blake2b_384(msg)
	case "blake2b_512":
		toReturn = hash.Hashing_blake2b_512(msg)
	default:
		toReturn = "NULL"
	}
	return toReturn
}
