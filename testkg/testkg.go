// Command gokey is a vaultless password and key manager.
package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/cloudflare/gokey"
)

var (
	pass, keyType, seedPath, realm, output string
	unsafe                                 bool
	seedSkipCount                          int
)

var keyTypes = map[string]gokey.KeyType{
	"ec256":   gokey.EC256,
	"ec521":   gokey.EC521,
	"rsa2048": gokey.RSA2048,
	"rsa4096": gokey.RSA4096,
}

func genSeed(w io.Writer) {
	seed, err := gokey.GenerateEncryptedKeySeed(pass)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = w.Write(seed)
	if err != nil {
		log.Fatalln(err)
	}
}

func genPass(seed []byte) string {
	password, err := gokey.GetPass(pass, realm, seed, &gokey.PasswordSpec{16, 2, 3, 2, 0, ""})
	if err != nil {
		log.Fatalln(err)
	}
	return password
}

func genKey(seed []byte, w io.Writer) {
	key, err := gokey.GetKey(pass, realm, seed, keyTypes[keyType], unsafe)
	if err != nil {
		log.Fatalln(err)
	}

	err = gokey.EncodeToPem(key, w)
	if err != nil {
		log.Fatalln(err)
	}
}

// TODO: parametrize size
// generates raw 32 bytes
func genRaw(seed []byte, w io.Writer) {
	raw, err := gokey.GetRaw(pass, realm, seed, unsafe)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = io.CopyN(w, raw, 32)
	if err != nil {
		log.Fatalln(err)
	}
}

func logFatal(format string, args ...interface{}) {
	log.Printf(format, args...)
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {

	//  var pass = "qwerty"
	//	var realm = "llwyd-consulting.cymru"
	var seed []byte

	var key = genPass(seed)
	println(key)
}
