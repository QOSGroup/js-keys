# js-keys
operation QOS keys


### Methods

*  DeriveQOSKey
 
 使用助记词恢复QOS账户公私钥. QOS账户体系中HDPATH为: "44'/389'/0'/0/0"

```
 let [privKey,pubKey,err] = qosKeys.DeriveQOSKey("助记词...")
```

*  DeriveKey

通用方法, 使用助记词恢复Tendermint账户公私钥


```
 let [privKey,pubKey,err] = qosKeys.DeriveKey("助记词...", "hdpath")

```

* Sign

使用私钥对数据进行签名
```
    qosKeys.Sign(privateKeyBuffer, messageDataBuffer)
```

* Bech32ifyQOSAccPubkeyFromBase64PubKey
* Bech32ifyQOSAccPubKey

以QOS Bech32格式显示账户公钥

* Bech32ifyQOSAccAddressFromBase64PubKey
* Bech32ifyQOSAccAddressFromPubKey
* Bech32ifyQOSAccAddress

以QOS Bech32格式显示账户地址

* AddressFromPubKey

通过账户公钥获取账户地址

