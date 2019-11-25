# js-keys
operate QOS keys


## Methods

### 解析公私钥及签名 

*  DeriveQOSKey 

使用助记词恢复QOS账户公私钥. QOS账户体系中HDPATH为: "44'/389'/0'/0/0"

私钥为16进制编码的字符串, 公钥为bech32格式的字符串, 用户地址为bech32格式的字符串


```typescript
function DeriveQOSKey(mnemonic: string): [privateKey: string,pubkey: string,accAddress: string, err: Error]
```

example: 

```typescript
 let mnemonic = "fury flavor subway start spare hospital tag chief word start pencil borrow town mandate detect pencil cook bridge right scout remain this differ leader";
 let [hexPrivateKey, bech32PubKey, bech32Address, err] = qosKeys.DeriveQOSKey(mnemonic);

 if(err != null){
   console.log(hexPrivateKey); // BEE8F561D6A09EE3DBB3C1F4BA505B33BF16FFAB9A9C0D4B312762F81C975876943A260CB11EFA8BB4D64E05B0D2C939D535D2B865A41C55411F810304A95337
   console.log(bech32PubKey); //  qosaccpub1jsazvr93rmaghdxkfczmp5kf882nt54cvkjpc42pr7qsxp9f2vms2evj9l
   console.log(bech32Address);//  qosacc163m0ww25yld86rrss0vntasn4cs72y5nl9evw3
 }  

```

* RecoverFromPrivateKey

使用私钥恢复账户公私钥及地址信息

```typescript
function RecoverFromPrivateKey(hexPrivateKey: string): [privateKey: string,pubkey: string,accAddress: string, err: Error]
```

example:

```typescript

let hexPrivateKey = "BEE8F561D6A09EE3DBB3C1F4BA505B33BF16FFAB9A9C0D4B312762F81C975876943A260CB11EFA8BB4D64E05B0D2C939D535D2B865A41C55411F810304A95337";

 let [hexPrivateKey1, bech32PubKey, bech32Address, err] = qosKeys.RecoverFromPrivateKey(hexPrivateKey);

 if(err != null){
   console.log(hexPrivateKey1); // BEE8F561D6A09EE3DBB3C1F4BA505B33BF16FFAB9A9C0D4B312762F81C975876943A260CB11EFA8BB4D64E05B0D2C939D535D2B865A41C55411F810304A95337
   console.log(bech32PubKey); //  qosaccpub1jsazvr93rmaghdxkfczmp5kf882nt54cvkjpc42pr7qsxp9f2vms2evj9l
   console.log(bech32Address);//  qosacc163m0ww25yld86rrss0vntasn4cs72y5nl9evw3
 }  
```

* Sign: 使用私钥对数据进行签名

```typescript
function Sign(privateKey: string, data: string): Uint8Array
```

example: 

```typescript
let privateKey  = "BEE8F561D6A09EE3DBB3C1F4BA505B33BF16FFAB9A9C0D4B312762F81C975876943A260CB11EFA8BB4D64E05B0D2C939D535D2B865A41C55411F810304A95337";
let message = "你好,QOS钱包";
let signedData = qosKeys.Sign(privateKey, message);
console.log(signedData); //Uint8Array(64) [246, 14, 144, 150, 177, 6, 5, 1, 13, 242, 162, 41, 69, 213, 88, 145, 174, 185, 32, 212, 198, 218, 132, 209, 197, 103, 232, 65, 134, 20, 2, 249, 108, 246, 12, 220, 135, 61, 120, 127, 98, 45, 144, 181, 70, 4, 201, 231, 187, 228, 61, 143, 23, 132, 198, 79, 229, 10, 162, 224, 204, 208, 35, 1]
```

* SignBase64Message: 使用私钥对Base64编码的字符串进行签名

```typescript
function SignBase64Message(privateKey: string, base64Message: string): Uint8Array
```

example: 

```typescript
let privateKey  = "BEE8F561D6A09EE3DBB3C1F4BA505B33BF16FFAB9A9C0D4B312762F81C975876943A260CB11EFA8BB4D64E05B0D2C939D535D2B865A41C55411F810304A95337";
let message = "你好,QOS钱包";
let base64Message = qosKeys.EncodeBase64(message);
let signedData = qosKeys.SignBase64Message(privateKey, base64Message);
console.log(signedData); //Uint8Array(64) [246, 14, 144, 150, 177, 6, 5, 1, 13, 242, 162, 41, 69, 213, 88, 145, 174, 185, 32, 212, 198, 218, 132, 209, 197, 103, 232, 65, 134, 20, 2, 249, 108, 246, 12, 220, 135, 61, 120, 127, 98, 45, 144, 181, 70, 4, 201, 231, 187, 228, 61, 143, 23, 132, 198, 79, 229, 10, 162, 224, 204, 208, 35, 1]
```

### 辅助方法

* EncodeBase64: base64编码

```typescript
function EncodeBase64(buffer: Uint8Array | string): string
```

example: 

```typescript
let buffer = Uint8Array.from([148, 58, 38, 12, 177, 30, 250, 139, 180, 214, 78, 5, 176, 210, 201, 57, 213, 53, 210, 184, 101, 164, 28, 85, 65, 31, 129, 3, 4, 169, 83, 55]);
let base64Str = qosKeys.EncodeBase64(buffer);
console.log(base64Str); //lDomDLEe+ou01k4FsNLJOdU10rhlpBxVQR+BAwSpUzc=

let str = "你好,QOS钱包";
console.log(qosKeys.EncodeBase64(str)); //5L2g5aW9LFFPU+mSseWMhQ==


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


* VerifyBech32String: 校验bech32格式的字符串是否合法

```typescript
function VerifyBech32String(bech32Str: string): boolean
```

example:

```typescript
let str = "qosacc163m0ww25yld86rrss0vntasn4cs72y5nl9evw3";
let isValidate = qosKeys.VerifyBech32String(str);
console.log(isValidate)
```


