# js-keys
operation QOS keys


## Methods

### 解析公私钥及签名 

*  DeriveQOSKey: 使用助记词恢复QOS账户公私钥. QOS账户体系中HDPATH为: "44'/389'/0'/0/0"

```typescript
function DeriveQOSKey(mnemonic: string): [Uint8Array,Uint8Array,Error]
```

example: 

```typescript
 let mnemonic = "fury flavor subway start spare hospital tag chief word start pencil borrow town mandate detect pencil cook bridge right scout remain this differ leader";
 let [privKey,pubKey,err] = qosKeys.DeriveQOSKey(mnemonic);

 if(err != null){
   console.log(privKey); // Uint8Array(64) [190, 232, 245, 97, 214, 160, 158, 227, 219, 179, 193, 244, 186, 80, 91, 51, 191, 22, 255, 171, 154, 156, 13, 75, 49, 39, 98, 248, 28, 151, 88, 118, 148, 58, 38, 12, 177, 30, 250, 139, 180, 214, 78, 5, 176, 210, 201, 57, 213, 53, 210, 184, 101, 164, 28, 85, 65, 31, 129, 3, 4, 169, 83, 55]
   console.log(pubKey); //  Uint8Array(32) [148, 58, 38, 12, 177, 30, 250, 139, 180, 214, 78, 5, 176, 210, 201, 57, 213, 53, 210, 184, 101, 164, 28, 85, 65, 31, 129, 3, 4, 169, 83, 55]   
 }  

```

*  DeriveKey: 通用方法, 使用助记词恢复Tendermint账户公私钥

```typescript
function DeriveKey(mnemonic: string, hdpath: string): [Uint8Array,Uint8Array,Error]
```

* Sign: 使用私钥对数据进行签名

```typescript
function Sign(privateKey: Uint8Array, data: Uint8Array): Uint8Array
```

example: 

```typescript
let privateKey  = Uint8Array.from([190, 232, 245, 97, 214, 160, 158, 227, 219, 179, 193, 244, 186, 80, 91, 51, 191, 22, 255, 171, 154, 156, 13, 75, 49, 39, 98, 248, 28, 151, 88, 118, 148, 58, 38, 12, 177, 30, 250, 139, 180, 214, 78, 5, 176, 210, 201, 57, 213, 53, 210, 184, 101, 164, 28, 85, 65, 31, 129, 3, 4, 169, 83, 55]);
let data = Uint8Array.from([10,23,219, 179, 193, 244, 186, 80, 91, 51, 191, 22, 255]);
let signedData = qosKeys.Sign(privateKey, data);
console.log(signedData); //Uint8Array(64) [67, 193, 52, 209, 75, 190, 230, 235, 107, 63, 32, 124, 6, 25, 160, 57, 205, 21, 255, 102, 28, 63, 9, 172, 136, 71, 68, 184, 84, 116, 38, 145, 91, 224, 114, 207, 188, 15, 147, 192, 42, 16, 194, 133, 75, 9, 179, 11, 248, 89, 67, 209, 41, 241, 172, 165, 51, 215, 4, 131, 106, 150, 194, 4]
```

### 公钥转换

* Bech32ifyQOSAccPubkeyFromBase64PubKey: 将Base64格式的公钥字符串转换为Bech32格式的公钥字符串

```typescript
function Bech32ifyQOSAccPubkeyFromBase64PubKey(base64PubKey: string): [string,Error]
```

example: 

```typescript
let base64Pubkey = "lDomDLEe+ou01k4FsNLJOdU10rhlpBxVQR+BAwSpUzc=";
let [bech32Pubkey, err] = qosKeys.Bech32ifyQOSAccPubkeyFromBase64PubKey(base64Pubkey);
if( err != null){
 console.log(bech32Pubkey); //qosaccpub1jsazvr93rmaghdxkfczmp5kf882nt54cvkjpc42pr7qsxp9f2vms2evj9l
}   
```

* Bech32ifyQOSAccPubKey: 将Uint8Array格式的公钥转换为Bech32格式的公钥字符串

```typescript
function Bech32ifyQOSAccPubKey(pubkey:Uint8Array): [string, Error]
```

example: 

```typescript
let pubkey = Uint8Array.from([148, 58, 38, 12, 177, 30, 250, 139, 180, 214, 78, 5, 176, 210, 201, 57, 213, 53, 210, 184, 101, 164, 28, 85, 65, 31, 129, 3, 4, 169, 83, 55]);
let [bech32Address, err] = qosKeys.Bech32ifyQOSAccPubKey(pubkey);   
if(err != null){
 console.log(bech32Address); //qosaccpub1jsazvr93rmaghdxkfczmp5kf882nt54cvkjpc42pr7qsxp9f2vms2evj9l
}
```

### 地址转换

* Bech32ifyQOSAccAddressFromBase64PubKey: 将Base64格式的公钥字符串转换为Bech32格式的地址字符串

```typescript
function Bech32ifyQOSAccAddressFromBase64PubKey(base64PubKey: string): [string,Error]
```

example: 

```typescript
let base64Pubkey = "lDomDLEe+ou01k4FsNLJOdU10rhlpBxVQR+BAwSpUzc=";
let [bech32Address, err] = qosKeys.Bech32ifyQOSAccAddressFromBase64PubKey(base64Pubkey);   
if(err != null){
 console.log(bech32Address); //qosacc163m0ww25yld86rrss0vntasn4cs72y5nl9evw3
} 
```

* Bech32ifyQOSAccAddressFromPubKey: 将Uint8Array格式的公钥转换为Bech32格式的地址字符串

```typescript
function Bech32ifyQOSAccAddressFromPubKey(pubkey: Uint8Array): [string,Error]
```

example: 

```typescript
let pubkey = Uint8Array.from([148, 58, 38, 12, 177, 30, 250, 139, 180, 214, 78, 5, 176, 210, 201, 57, 213, 53, 210, 184, 101, 164, 28, 85, 65, 31, 129, 3, 4, 169, 83, 55]);
let [bech32Address, err] = qosKeys.Bech32ifyQOSAccAddressFromPubKey(pubkey);
if(err != null){
 console.log(bech32Address); //qosacc163m0ww25yld86rrss0vntasn4cs72y5nl9evw3
}      
```

* Bech32ifyQOSAccAddress: 将Uint8Array格式的地址转换为Bech32格式的地址字符串

```typescript
function Bech32ifyQOSAccAddress(address: Uint8Array): [string,Error]
```

example: 

```typescript
let address = Uint8Array.from([212, 118, 247, 57, 84, 39, 218, 125, 12, 112, 131, 217, 53, 246, 19, 174, 33, 229, 18, 147]);
let [bech32Address, err] = qosKeys.Bech32ifyQOSAccAddress(address);
if(err != null){
 console.log(bech32Address); //qosacc163m0ww25yld86rrss0vntasn4cs72y5nl9evw3
}   
```


* AddressFromPubKey: 从公钥中获取地址

```typescript
function AddressFromPubKey(pubkey: Uint8Array): Uint8Array
```

example: 

```typescript
let pubkey = Uint8Array.from([148, 58, 38, 12, 177, 30, 250, 139, 180, 214, 78, 5, 176, 210, 201, 57, 213, 53, 210, 184, 101, 164, 28, 85, 65, 31, 129, 3, 4, 169, 83, 55]);
let address = qosKeys.AddressFromPubKey(pubkey);
console.log(address); //Uint8Array(20) [212, 118, 247, 57, 84, 39, 218, 125, 12, 112, 131, 217, 53, 246, 19, 174, 33, 229, 18, 147]
```


### 辅助方法

* EncodeBase64: base64编码

```typescript
function EncodeBase64(buffer: Uint8Array): string
```

example: 

```typescript
let buffer = Uint8Array.from([148, 58, 38, 12, 177, 30, 250, 139, 180, 214, 78, 5, 176, 210, 201, 57, 213, 53, 210, 184, 101, 164, 28, 85, 65, 31, 129, 3, 4, 169, 83, 55]);
let base64Str = qosKeys.EncodeBase64(buffer);
console.log(base64Str); //lDomDLEe+ou01k4FsNLJOdU10rhlpBxVQR+BAwSpUzc=
```

* DecodeBase64: base64解码

```typescript
function DecodeBase64(base64Str: string): [Uint8Array,Error]
```

example: 

```typescript
let base64Str = "lDomDLEe+ou01k4FsNLJOdU10rhlpBxVQR+BAwSpUzc=";
let [buffer, err] = qosKeys.DecodeBase64(base64Str);
if(err != null){
  console.log(buffer); //Uint8Array(32) [148, 58, 38, 12, 177, 30, 250, 139, 180, 214, 78, 5, 176, 210, 201, 57, 213, 53, 210, 184, 101, 164, 28, 85, 65, 31, 129, 3, 4, 169, 83, 55]
}
```

