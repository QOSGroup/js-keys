package keys

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"golang.org/x/crypto/ed25519"
	"math/big"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil/bech32"
	"github.com/tyler-smith/go-bip39"
)

func DeriveQOSKey(mnemonic string) ([]byte, []byte, error) {
	return DeriveKey(mnemonic, "44'/389'/0'/0/0")
}

func DeriveKey(mnemonic, hdpath string) (priKeyBz []byte, pubKeyBz []byte, err error) {
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		return nil, nil, err
	}

	secret, chaincode := computeMastersFromSeed(seed)
	privateKeySeed, err := derivePrivateKeyForPath(secret, chaincode, hdpath)
	if err != nil {
		return nil, nil, err
	}

	hasher := sha256.New()
	hasher.Write(privateKeySeed[:])
	hash256Seed := hasher.Sum(nil)

	priKey := ed25519.NewKeyFromSeed(hash256Seed)
	return priKey, priKey[32:], nil
}


func Sign(privKey, message []byte) []byte {
	pk := ed25519.PrivateKey(privKey)
	return ed25519.Sign(pk, message)
}

func Bech32ifyQOSAccPubkeyFromBase64PubKey(base64pubkey string) (string, error) {
	bz, err := base64.StdEncoding.DecodeString(base64pubkey)
	if err != nil {
		return "", err
	}
	return Bech32ifyQOSAccPubKey(bz)
}

func Bech32ifyQOSAccPubKey(pubkey []byte) (string, error) {
	return ConvertAndEncode("qosaccpub", pubkey)
}

func Bech32ifyQOSAccAddressFromBase64PubKey(base64pubkey string) (string, error) {
	bz, err := base64.StdEncoding.DecodeString(base64pubkey)
	if err != nil {
		return "", err
	}
	return Bech32ifyQOSAccAddressFromPubKey(bz)
}

func Bech32ifyQOSAccAddressFromPubKey(pubkey []byte) (string, error) {
	return Bech32ifyQOSAccAddress(AddressFromPubKey(pubkey))
}

func Bech32ifyQOSAccAddress(addr []byte) (string, error) {
	return ConvertAndEncode("qosacc", addr)
}

func DecodeBase64(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}

func EncodeBase64(bz []byte) string {
	return base64.StdEncoding.EncodeToString(bz)
}

func ConvertAndEncode(hrp string, data []byte) (string, error) {
	converted, err := bech32.ConvertBits(data, 8, 5, true)
	if err != nil {
		return "", errors.New("encoding bech32 failed")
	}
	return bech32.Encode(hrp, converted)

}

func AddressFromPubKey(bz []byte) []byte {
	hash := sha256.Sum256(bz)
	return hash[:20]
}

func computeMastersFromSeed(seed []byte) (secret [32]byte, chainCode [32]byte) {

	masterSecret := []byte("Bitcoin seed")
	secret, chainCode = i64(masterSecret, seed)

	return
}

func derivePrivateKeyForPath(privKeyBytes [32]byte, chainCode [32]byte, path string) ([32]byte, error) {
	data := privKeyBytes
	parts := strings.Split(path, "/")
	for _, part := range parts {
		harden := part[len(part)-1:] == "'"
		if harden {
			part = part[:len(part)-1]
		}
		idx, err := strconv.Atoi(part)
		if err != nil {
			return [32]byte{}, fmt.Errorf("invalid BIP 32 path: %s", err)
		}
		if idx < 0 {
			return [32]byte{}, errors.New("invalid BIP 32 path: index negative ot too large")
		}
		data, chainCode = derivePrivateKey(data, chainCode, uint32(idx), harden)
	}
	var derivedKey [32]byte
	n := copy(derivedKey[:], data[:])
	if n != 32 || len(data) != 32 {
		return [32]byte{}, fmt.Errorf("expected a (secp256k1) key of length 32, got length: %v", len(data))
	}

	return derivedKey, nil
}

func derivePrivateKey(privKeyBytes [32]byte, chainCode [32]byte, index uint32, harden bool) ([32]byte, [32]byte) {
	var data []byte
	if harden {
		index = index | 0x80000000
		data = append([]byte{byte(0)}, privKeyBytes[:]...)
	} else {
		_, ecPub := btcec.PrivKeyFromBytes(btcec.S256(), privKeyBytes[:])
		pubkeyBytes := ecPub.SerializeCompressed()
		data = pubkeyBytes

	}
	data = append(data, uint32ToBytes(index)...)
	data2, chainCode2 := i64(chainCode[:], data)
	x := addScalars(privKeyBytes[:], data2[:])
	return x, chainCode2
}

// modular big endian addition
func addScalars(a []byte, b []byte) [32]byte {
	aInt := new(big.Int).SetBytes(a)
	bInt := new(big.Int).SetBytes(b)
	sInt := new(big.Int).Add(aInt, bInt)
	x := sInt.Mod(sInt, btcec.S256().N).Bytes()
	x2 := [32]byte{}
	copy(x2[32-len(x):], x)
	return x2
}

func uint32ToBytes(i uint32) []byte {
	b := [4]byte{}
	binary.BigEndian.PutUint32(b[:], i)
	return b[:]
}

// i64 returns the two halfs of the SHA512 HMAC of key and data.
func i64(key []byte, data []byte) (IL [32]byte, IR [32]byte) {
	mac := hmac.New(sha512.New, key)
	// sha512 does not err
	_, _ = mac.Write(data)
	I := mac.Sum(nil)
	copy(IL[:], I[:32])
	copy(IR[:], I[32:])
	return
}
