package main

import (
	"fmt"
	"crypto"
	"crypto/rsa"
    "crypto/rand"
    "log"
    "strconv"
    "crypto/sha256"
   )

func main(){
//Asymmetric Encription
privKey, err := rsa.GenerateKey(rand.Reader, 2048)
if err != nil {
	log.Fatalf("generating random key: %v", err)
	}

plainText := []byte("We're going to keep on gaining power and getting better!!")

cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, &privKey.PublicKey, plainText)
if err != nil {
	log.Fatalf("could not encrypt data: %v", err)
  }
fmt.Printf("%s\n\n", strconv.Quote(string(cipherText)))

decryptedText, err := rsa.DecryptPKCS1v15(nil, privKey, cipherText)
if err != nil {
	log.Fatalf("error decrypting cipher text: %v", err)
}
fmt.Printf("%s\n\n", decryptedText)

// Digital Signatures
hash := sha256.Sum256(plainText)
fmt.Printf("The hash of my message is: %#x\n\n", hash)
// generate a signature using the private key
signature, err := rsa.SignPKCS1v15(rand.Reader, privKey, crypto.SHA256, hash[:])
if err != nil {
	log.Fatal("error creating signature: %v", err)
	}

fmt.Printf("Digital Signature: %s\n", signature)
// use a public key to verify the signature for a message that was created by the private key
verify := func(pub *rsa.PublicKey, msg, signature []byte) error {
	hash := sha256.Sum256(msg)
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hash[:], signature)
	}
fmt.Printf("A bad signature: ")
fmt.Println(verify(&privKey.PublicKey, plainText, []byte("a bad signature")))
fmt.Printf("Different plain text: ")
fmt.Println(verify(&privKey.PublicKey, []byte("crap aint worth a thing"), signature))
fmt.Printf("Correct pubkey/text/signature: ")
fmt.Println(verify(&privKey.PublicKey, plainText, signature))
}
