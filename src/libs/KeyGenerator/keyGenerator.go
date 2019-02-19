package KeyGenerator

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func GeneratePrivate(size int) *rsa.PrivateKey {
	privateK, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		fmt.Errorf("Can't generate private key:\n%+v", err)
		os.Exit(1)
	}

	return privateK
}

func GeneratePEMFiles(key *rsa.PrivateKey, name string) {
	WritePrivateFile(key, name+".key")
	public := &key.PublicKey
	WritePublicFile(public, name+".crt")
}

func WritePrivateFile(key *rsa.PrivateKey, path string) {
	keyFile, keyErr := os.Create(path)

	if keyErr != nil {
		fmt.Println(keyErr)
		os.Exit(1)
	}

	var pemBlock = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err := pem.Encode(keyFile, pemBlock)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	keyFile.Close()
}

func WritePublicFile(publicKey *rsa.PublicKey, path string) {
	crtFile, crtErr := os.Create(path)

	if crtErr != nil {
		fmt.Println(crtErr)
		os.Exit(1)
	}

	pemBlock := &pem.Block{
		Type:  "RSA Public Key",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	}

	pem.Encode(crtFile, pemBlock)

	crtFile.Close()
}

func ReadFilePrivate(fileName string) *rsa.PrivateKey {
	privateKeyFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)
	data, _ := pem.Decode([]byte(pembytes))
	privateKeyFile.Close()

	privateKey, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return privateKey
}

func ReadFilePublic(fileName string) *rsa.PublicKey {
	publicKeyFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pemfileinfo, _ := publicKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pembytes)
	data, _ := pem.Decode([]byte(pembytes))
	publicKeyFile.Close()

	publicKey, err := x509.ParsePKCS1PublicKey(data.Bytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return publicKey
}
