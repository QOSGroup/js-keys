package main

import (
	"github.com/QOSGroup/js-keys/keys"
	"github.com/gopherjs/gopherjs/js"
)

func main() {

	js.Global.Set("qosKeys", map[string]interface{}{
		//"Bech32ifyQOSAccPubkeyFromBase64PubKey":  keys.Bech32ifyQOSAccPubkeyFromBase64PubKey,
		//"Bech32ifyQOSAccAddressFromBase64PubKey": keys.Bech32ifyQOSAccAddressFromBase64PubKey,
		//"Bech32ifyQOSAccAddressFromPubKey":       keys.Bech32ifyQOSAccAddressFromPubKey,
		//"Bech32ifyQOSAccAddress":                 keys.Bech32ifyQOSAccAddress,
		//"Bech32ifyQOSAccPubKey":                  keys.Bech32ifyQOSAccPubKey,
		"DecodeBase64":                           keys.DecodeBase64,
		"EncodeBase64":                           keys.EncodeBase64,
		"DeriveQOSKey":                           keys.DeriveQOSKey,
		//"AddressFromPubKey":                      keys.AddressFromPubKey,
		"Sign": 								  keys.Sign,
		"SignBase64Message":                      keys.SignBase64Message,
		"VerifyBech32String":                     keys.VerifyBech32String,
		"RecoverFromPrivateKey":                  keys.RecoverFromPrivateKey,
	})
}


