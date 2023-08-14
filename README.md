# sto/cbdc know your customer service

[TOC]

## 更新日志

| 时间\|版本            | 更新内容                                                     |
| --------------------- | ------------------------------------------------------------ |
| 2023/08/11\|v1.6.2    | 1、更新 file download 接口，image/text 等内容会直接渲染，而不是下载<br />2、去掉创建 org 时，description 必输的限制 |
| 2023/08/08\|v1.6.1    | 1、更新 server 创建 product 接口，对接合约接口，能够真实创建产品<br />2、更新 server 上线 hidding product 接口，对接合约接口，能够真实创建产品<br />3、更新 server 创建 product 接口，输入字段增加 `symbol`字段，必填 |
| 2023/08/02\|v1.6.0    | 1、增加 server 端创建机构接口 /api/v1/org/admin/org/create<br />2、增加 server 端搜索机构接口 /api/v1/org/admin/org/query<br />3、增加 server 端查看指定 id 机构信息接口 /api/v1/org/admin/org/detail<br />4、增加 server 端创建 product 接口 /api/v1/org/admin/product/create<br />5、增加 server 端上线 hidding product 接口 /api/v1/org/admin/product/setonline<br />6、增加 server 端搜索 product 接口 /api/v1/org/admin/product/query<br />7、增加 server 端查看指定 id 的 product 接口 /api/v1/org/admin/product/detail<br />8、增加 client 端搜索机构接口 /api/v1/org/query<br />9、增加 client 查看指定 id 机构接口 /api/v1/org/detail<br />10、增加 client 端搜索 product 接口 /api/v1/product/query<br />11、增加 client 端查看指定 id 的 product 接口 /api/v1/product/detail<br />12、修改 server 端 level3 接口的查询参数 /api/v1/kyc/admin/level3/query |
| 2023/07/31\|v1.5.0    | 1、增加 server 端对 level3 数据的 confirm 接口 /api/v1/kyc/admin/level3/section/confirm<br />2、增加 server 端对 level3 数据的 reqmoreinfo 接口 /api/v1/kyc/admin/level3/section/reqmoreinfo<br />3、增加 server 端对 level3 数据的 send vc 接口 /api/v1/kyc/admin/level3/section/sendvc<br />4、增加 server 端对 level3 数据的 risk mark 接口 /api/v1/kyc/admin/level3/section/markrisk<br />5、增加 server 端对 send vc 列表查看接口 /api/v1/kyc/admin/level3/sendvc/query |
| 2023/07/17 \| v1.4.0  | 1、增加 client 查看用户自身信息的接口 /api/v1/kyc/user/info<br />2、增加 client 上传文件的接口 /api/v1/kyc/file/upload<br />3、增加 client 下载指定 id 的文件的接口 /api/v1/kyc/file/download<br />4、增加 server 端下载指定 id 的文件的接口 /api/v1/kyc/admin/file/download<br />5、增加 【数据字典】中对 user 信息中 level 字段的解释 |
| 2023/07/14 \| v1.3.0  | 1、增加 server 端查看 user 信息列表接口 /api/v1/kyc/admin/user/query<br />2、增加 server 端查看指定 id 的 user 信息的接口 /api/v1/kyc/admin/user/detail<br />3、调整 client 端提交 level3 数据的输入参数结构，具体见 /api/v1/kyc/level3/create<br />4、增加【数据字典】模块，说明 level3 中的相关数据的 status 值意义 |
| 2023/07/03 \| v1.2.0  | 1、增加 JWT token 验证功能<br />2、增加 level3 的 client api<br />3、增加 level3 的 server 管理 api |
| 2023/06/29 \| v1.1.0  | 1、调整 /api/v1/kyc/passwd/codeverify 输入参数，增加 receiver 字段<br />2、调整 /api/v1/kyc/keypair/recovery/codeverify 输入参数，增加 receiver 字段<br />3、调整 /api/v1/kyc/debug/token 输入及返回数据 |
| 2023/06/29  \| v1.0.1 | 1、获取 token 的 debug 接口，能够测试签名是否正确<br />2、调整 用户 key pai 恢复/找回接口的路径，增加了 keypair 在路径上 |



## api 约定

### 接口请求数据形式与是否必输

- 如果需要提交数据给 server，除非明确指明，一般都是 http request body 的形式传递数据，数据以 json 格式编码。
- 如果没有特别指明，那么字段就是必输，只有非必输字段会特别标记。
- 请求参数包含如下形式
  - url query 参数，比如 /api/abc?arg1=aaa&arg2=bbb，这里 arg1 和 arg2 就是 url query 参数
  - url path 参数，比如 /api/order/:id，这里的 `:id`就是 url path 参数
  - http request body
- 上述的参数可能会同时包含，比如 url query/path 参数与 http request body 一起使用等。



---

### 返回数据

只要是 api 返回的数据，返回的结构都是固定的，其中 data 字段的值，根据情况，可能有值，可能无值。

- code 是一个整数，大多数情况与 http status code 值一致，只有在需要 client 特别处理的情况下，才会有自定义的值。
- msg 是字符串，当有错误时，给出错误的信息提示
- data 是一个对象，可能有值，可能为空

```json
{
    "code": 200,
    "msg": "OK",
    "data": {
        "id": "id1121222"
    }
}
```



---

###  数据字典

#### user info 中的 level 字段的值

| 值   | 意义                 |
| ---- | -------------------- |
| 2    | 用户处于 level2 阶段 |
| 3    | 用户处于 level3 阶段 |



---

#### level3 数据的 status 值

| 值        | 意义             |
| --------- | ---------------- |
| CREATED   | 数据被初始化提交 |
| PROGRESS  | 数据被审核处理中 |
| COMPLETED | 数据处理完毕     |



---

#### level3 数据中各个 section 的 status 值

| 值              | 意义                                                         |
| --------------- | ------------------------------------------------------------ |
| INIT            | 数据未处理的初始化状态，可以被 confirm 或者 request more info |
| CHECKED         | 数据被 confirmed                                             |
| REQUESTMOREINFO | request more info                                            |
| SENDVC          | 已发送给 vc                                                  |

---

#### product.status 的意义

| 值      | 意义                                                         |
| ------- | ------------------------------------------------------------ |
| DRAFT   | 当页面选择了 hidding 时，数据就是这个状态                    |
| ONLINE  | 当页面选择了 showing，或者原来是 hidding，但是选择上线，创建 product 到合约成功后的状态 |
| OFFLINE | 把 ONLINE 的 product 进行下线操作                            |



---

### http status code 及 http response.body.code 

绝大多数情况下，response.body.code 于 http status code 值一致，需要 client 特殊处理的，才会有自定义值，下面分别描述。

**http status code**

使用的是标准的 http status code，关于 status code 的更多参考信息见：[HTTP response status codes - HTTP | MDN](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status) 。

这里仅列出来一些常见的：

- 200/201 请求成功
- 301/302 跳转到新的 url
- 401 用户 token 错误、失效等
- 403 用户 token 正确，但是没有权限访问指定的内容
- 500 系统内部错误，程序有 bug 了
- 501 这个在开发过程中可能会返回，因为预先定义了接口，但是还没有来得及实现
- 502/504 这个一般是反向代理返回的，可能是配置错误，或者 api 挂了



**自定义的 code**

todo



---

### request headers 与 token

token 分为两种：

- kyc token。由 client 使用 private key 进行签名生成的 token 数据，传给 api 时，使用 X-KYC-TOKEN 作为 headers 的 key。
- jwt token。由 server  web 使用账号/密码登录后获取到的 jwt token，传给 api 时，使用 Authorization 作为 headers 的 key，使用 `Bearer + 空格 + token` 作为值传递给 api。

**除非特别标注的外，其余所有的请求 kyc 的 api 中**，必须在 request headers 中增加一个 `X-KYC-TOKEN`  或者 `Authorization` 的字段。选择使用哪个 name 作为 key，规则如下：

- 如果是使用签名生成的，使用 `X-KYC-TOKEN`。
- 如果使用账户密码登录后获取的，使用 `Authorization`。

关于 X-KYC-TOKEN 这个字段的值是按照一定规则生成的。生成规则如下所描述。

#### X-KYC-TOKEN 生成规则

**背景：**因为我们并没有通常应用的那种账户密码的服务器托管账户的交互形式。用户的 key pair 由用户控制，用户的身份认证靠的是用户的签名机制。所以，我们设计的验证请求是否合法，主要依赖于用户的签名。在用户有签名信息的基础上，提前在 server 端保存了用户的 address 地址。进而能够判断是否是一个合法（白名单）的用户请求。

所以实现过程是：

- 用户在 register 时，提交了他的 ethereum address，server 端保存了此地址
- 用户在调用其他 api 时，用自己的 private key 签名一个按规则生成的数据，把数据和签名提交给 server
  - server 验证签名是合法的
  - server 验证此签名对应的 address 是提前注册了的（即白名单机制）

**算法：**使用 ethereum 的签名算法，对指定格式的数据进行签名，把原始数据和签名一起 hex 编码后作为 X-KYC-TOKEN 的 value 。

指定格式的数据： `platform + | + timestamp`。platform 可选值是 `ios/android/web`。`timestamp`是一个精确到**秒**的整数。

比如`ios|1686811276`就是原始消息，使用用户的 **private key** 对这个消息进行签名，得到了签名后，把整个签名的和数据一起组装成一个 json 结构的数据，然后对此 json 数据进行 hex 编码，得到的 hex 数据就是 X-KYC-TOKEN 的 value。

比如下面就是一个例子，在这个例子中，展示了各个步骤的数据的结果。

```shell
# 注意，下面的数据是真实的。
# 使用 private key 为 ac178d1ef86b0b32dfc3e3dce5a386e23d407deeadb6fe3014ec0e336712f3dd
# 对应的 address 是 0xCb34c649dAb213A395164474e943C21B0E2126Af
# 对应的 public key 是 049c83fc52d46fbe20ca91d6c2ed28a41bb3dfc0cf8915e9cc50b491cf92cc78426c20f37d19be82225c2f0279b1f1411ff3b85c8f36f5c14783f0751dea033d19

# 使用上面的 private key 对下面的 msg 字段签名，就可以得到 signature 字段的值。

# 消息结构包含两个字段，分别是 msg 和 signature
{"msg":"ios|1687965938","signature":"29512c3a255b233253e34b334c8ee4a459c829cd5658e9f1eee6bdac60d657ac0c1487071ba87564c7864af3b2fea4e16e79904386e43011236a3a5969210b6500"}

# 对上面的 json 数据进行序列化，并对序列化后的数据进行 hex 编码，就得到了 X-KYC-TOKEN
7b226d7367223a22696f737c31363837393635393338222c227369676e6174757265223a2232393531326333613235356232333332353365333462333334633865653461343539633832396364353635386539663165656536626461633630643635376163306331343837303731626138373536346337383634616633623266656134653136653739393034333836653433303131323336613361353936393231306236353030227d

# 下面的 python 代码例子，是一个关于序列化与反序列化的过程的示例。
```



比如下面这个 python 例子：

```python
# 假设有数据如下所示
# 下面的数据，包含两部分数据，分别是签名数据和原始消息
# 注意，这里是例子，与真实的签名数据结构完全不同，仅作为算法的演示，稍候会更新为真实的例子。
# 再注意，为了分步骤说明，代码会显得特别诡异，请勿直接使用。
a = {'kkk1': 'vvv1', 'kkk2': 'vvv2', 'raw': 'ios|1223445'}

# 先对 a 做 json 编码，序列化为字符串
import json
b = json.dumps(a)
c = b.encode('utf-8') # 这一步是 python 3.x 的要求，先把 str 转换为 bytes

# 对 c 进行 hex 编码
import binascii
d = binasicc.b2a_hex(c)
print(d)

# 上面的 d 就是我们的 X-KYC-TOKEN 的 value
# 比如 d 的值 b'7b226b6b6b31223a202276767631222c20226b6b6b32223a202276767632222c2022726177223a2022696f737c31323233343435227d'

# 下面是把 hex 数据转换回原始结构
# server 端收到了 X-KYC-TOKEN 的 d 后
e = binascii.a2b_hex(d)
# e 的可能输出是字符串形式的 b'{"kkk1": "vvv1", "kkk2": "vvv2", "raw": "ios|1223445"}'

# json 反序列化
f = json.loads(e)

# f 就还原回了 a 的样子

# 后续就是签名验证了
# 签名验证首先验证签名的有效性，然后再根据签名反推回得到的 public key，进而得到了 address
# 如果签名有效，且，address 在 server 端记录为正常，就可以正常使用。说明 X-KYC-TOKEN 是 valid。
# 否则返回 401 错误
```



---

### 用户密码与 private key 及 mnemonic words 加密方法

需要注意的是，用户在界面上输入的密码，与用来加密 private key 及 mneomnic words 的密码不是同一个密码。互相之间没有任何关系。

用户密码的处理流程是：

1. 用户输入密码，密码被 sha256 算法进行 hash 并导出为 hex 编码的字符串，client 存储此 hash 值
2. 用户在验证 email/sms code 时，提交此 hash 给 server



private key 及 mnemonic words 也是加密后保存的。

默认的 private key 的 keystore 加密方式是 aes-128-ctr，可以根据情况看能否修改为 aes-256-gcm。

这里先描述密钥的选取步骤，具体的 aes 算法可以稍候再确定。

因为这里的密钥与用户密码没有任何关系，且有两个要求：

- 在使用时，能够解密加密后的 private key，且用户修改了密码后，不影响这里的加密/解密过程；
- 在恢复数据时，能够从 server 端恢复回去数据，server 端具备解密数据的能力。

基于以上两个要求，我们的密钥选取就必须是一个“确定性”的算法生成的数据。

所以，我们的操作步骤是

1. 在生成 master key pair 后，使用 address 的后 20 个字符串作为基础密钥
2. 在这个字符串之前添加固定的 salt，salt 值是 `EMALICBDC`
3. 对上述字符串进行 md5 计算，计算结果导出为 hex 形式的值，且全部是小写形式
4. 上述 3 中最终的得到的 hex 形式的、小写的 hash 值即为 aes 的 key 。

更具体的例子，见 api `验证注册 address 到 server 的验证码的接口`中的 aes 生成算法的描述。



---

## api 列表

### kyc token debug 接口

这个接口的目的是为了调试/获取 X-KYC-TOKEN，同时也是为了方便调用者验证自己的算法是否正确。

- 这个接口无须任何权限验证
- 这个接口仅用户开发测试阶段
- 这个接口会返回 private key、msg、signature、token 等字段

| HTTP METHOD | GET                          |
| ----------- | ---------------------------- |
| URI         | /api/v1/kyc/debug/token      |
| 请求数据    | 可选参数，也可以什么都不输入 |
| HEADERS     | 无                           |

**输入参数**

输入参数是可选的，如果用户不输入，就用程序内置的，如果输入，就用输入的。

两个参数（都是可以不用输入的）：

- words - 这个参数是助记词的用竖线分割的字符串，比如 `wolf|juice|proud|gown|wool|unfair|wall|cliff|insect|more|detail|hub`
- privateKeyHex - private key 的 hex 表示形式

注意，虽然这个接口是 GET 的，但是也是支持 request body 的。



**返回例子**

```json
{
	"code": 200,
	"msg": "OK",
	"data": {
		"address": "0xCb34c649dAb213A395164474e943C21B0E2126Af",
		"aesKey": "0ca65127da42f5125e0fe3016277b21e",
		"data": "{\"msg\":\"ios|1688033513\",\"signature\":\"7ce7ad0f4b40ba3d86f2cd1eb13fa91a2771edb4efdf836e0d0887298d065a0c4ec09fc5b8357d007d3cfb26c48764787ee431c4def3f8ae5175e07fbe2bd60401\"}",
		"encryptedWords": "6d48310d500989f280214af45f22f60caa92393138cf874b0d7f65c5b4400ecca22ee856d411ba8818dfc6691626401a9cbfd763e8a56268556a96b0e712361c42b20fb288e72a901b48b43fcb2b549fd063b61d42c7e522f409119b5f24f9",
		"msg": "ios|1688033513",
		"privateKeyHex": "ac178d1ef86b0b32dfc3e3dce5a386e23d407deeadb6fe3014ec0e336712f3dd",
		"publicKeyHex": "049c83fc52d46fbe20ca91d6c2ed28a41bb3dfc0cf8915e9cc50b491cf92cc78426c20f37d19be82225c2f0279b1f1411ff3b85c8f36f5c14783f0751dea033d19",
		"rawWords": "wolf|juice|proud|gown|wool|unfair|wall|cliff|insect|more|detail|hub",
		"signature": "7ce7ad0f4b40ba3d86f2cd1eb13fa91a2771edb4efdf836e0d0887298d065a0c4ec09fc5b8357d007d3cfb26c48764787ee431c4def3f8ae5175e07fbe2bd60401",
		"token": "7b226d7367223a22696f737c31363838303333353133222c227369676e6174757265223a2237636537616430663462343062613364383666326364316562313366613931613237373165646234656664663833366530643038383732393864303635613063346563303966633562383335376430303764336366623236633438373634373837656534333163346465663366386165353137356530376662653262643630343031227d"
	}
}
```



---

### 用户查看自己的信息

| HTTP METHOD | GET                   |
| ----------- | --------------------- |
| URI         | /api/v1/kyc/user/info |
| 输入参数    | 无                    |
| HEADERS     | X-KYC-TOKEN           |

这个接口返回的是调用者自己的基本信息。

在返回的数据中，有一个 `level`字段，标记了当前用户所处的 level 阶段。

- 2 - 标示处于 level2 阶段
- 3 - 标示处于 level3 阶段
- level3Status  结构中的 status 值标记了当前 level3 数据的审核进度



**返回例子**

```json
{
	"code": 200,
	"msg": "OK",
	"data": {
		"id": "64c76ee73b0a8a4622b30aae",
		"phone": "+85212345678",
		"email": "",
		"publicKey": "039c83fc52d46fbe20ca91d6c2ed28a41bb3dfc0cf8915e9cc50b491cf92cc7842",
		"address": "0xcb34c649dab213a395164474e943c21b0e2126af",
		"eip55Address": "0xCb34c649dAb213A395164474e943C21B0E2126Af",
		"extraInfo": null,
		"level": 3,
		"level3Status": {
			"mark": 1,
			"vcSend": 3,
			"status": "PROGRESS"
		},
		"created": 1690791655,
		"updated": 1690791815
	}
}
```



如果用户没有注册，返回如下所示：

1. http status code 是 400
2. response body 结果如下所示

```json
{
	"code": 4000,
	"msg": "address hasn't registered",
	"data": ""
}
```



---

### 用户注册 address 到 server 端

client 在本地创建 key pair 是一个完全 local 的操作，无须与 server 端交互。此时 client 处于 level1 状态

但是，如果 client 想要进行后续的交易，就需要进行 email/sms 验证。验证通过，client 处于 level2 状态。

email/sms 的验证是一个提交 client address 和 mnemonic words 到 server 的操作。

请求流程是：

1. 请求 server 下发验证码
2. 提交收到的验证码给 server，同时提交加密后的 mnemonic words 等信息。



下面是涉及到的两个 api 定义。

#### 注册 address 到 server 的接口 / 验证码发送请求接口

本接口接收 client 提交的数据，根据 client 选择的验证码发送方式，发送验证码给 client。

| HTTP METHOD | POST                              |
| ----------- | --------------------------------- |
| URI         | /api/v1/kyc/addr/register/request |
| 请求数据    | request json body                 |
| HEADERS     | X-KYC-TOKEN 必填                  |

**输入参数**

| field    | value                       | remark                                                       |
| -------- | --------------------------- | ------------------------------------------------------------ |
| type     | 固定值，可选 email 或者 sms |                                                              |
| receiver | 邮箱或者手机号              | 根据 type 不同，此处可能是邮箱，有可能是手机号。<br />如果是手机号，需要带上国际区号前缀，比如 +852 |

*例子*

```json
{
    "type": "sms",
    "receiver": "+8521234567"
}
```



**返回数据**

如果请求成功，会给出请求成功的提示。

如果失败，会给出失败的原因。比如签名验证不通过，缺少必输字段等等信息。



---

#### 验证注册 address 到 server 的验证码的接口

本接口是接收验证码，注册用户 ethereum address 及 mnemonic words 等信息到 server 端，后续可以通过找回的方法来找回 mnemonic words。

| HTTP METHOD | POST                             |
| ----------- | -------------------------------- |
| URI         | /api/v1/kyc/addr/register/verify |
| 请求数据    | request json body                |
| HEADERS     | X-KYC-TOKEN 必填                 |

**输入参数**

| field    | value                                         | remark                       |
| -------- | --------------------------------------------- | ---------------------------- |
| receiver | 邮箱或者手机号                                | 与请求 request 接口一致      |
| code     | 验证码                                        |                              |
| words    | 被 aes-256-gcm 加密后的 mnemonic words 的数据 | 关于数据生成方法，见下面描述 |
| hpasswd  | sha256 hex 编码的用户的密码，无 salt。        |                              |

**words 生成方法**

words 参数是 hd wallet 的 mnemonic words 的加密后的结果。

因为用户在 client 界面上输入的密码可以被修改，而如果 client 忘记了密码，被 aes 加密的数据就无法被解开了。

所以，我们加密数据使用的密码并不是 client 控制的密码。

与此同时，这个加密数据的密码需要在 server 端`明文`计算出来，才能在 client 端进行数据恢复时，根据密码解密加密的数据，并进行数据恢复。

**aes 密码的生成算法是：**

1. 根据用户的 address 地址，截取最后 20 个字符，注意，address 不区分大小写，但是这里要求全部是**小写字符**。
   1. 比如 `001d3f1ef827552ae111`
2. 在截取出来的字符串，再加上固定的 salt 值，salt 值是 `EMALICBDC`。
3. 对上述两部分字符串进行无添加的拼接，以 salt 值作为开头，比如 `EMALICBDC001d3f1ef827552ae111`。
4. 对 3 中的字符串计算其 md5 值。计算的结果 hex 转换后，以所有字符为小写的形式使用。
   1. 比如对 3 的字符串进行 md5 并小写的的值是：`7386454637d5356f160d3220442209f3`
5. 上面 4 得到的字符串就是 aes 加密中的 `key`。



我们采用的是 12 个 words 的 mnemonic 算法，对于这些 words，使用 `|`分割后，按顺序拼接起来，组合成一个最终个 word。

然后对这个 word 进行 `aes-256-gcm` 加密。

比如 `wolf|juice|proud|gown|wool|unfair|wall|cliff|insect|more|detail|hub`。使用上述的 4 的 key 进行加密，加密后的内容是 `cbf9459d2f20a610c9985e8848990db1d0950eab19ef831baa9e4a893411f9f009cda5591f635de7679b48d28f1d4304cfeceb0b55eb63939e0c4ffb5a89730134be7af59f832c488cc86c420c4dd4c839cc8a081ff9c032533c23996e5d98`。

这个就是 words 的内容。

*例子*

```json
{
    "receiver": "+85212345678",
    "code": "123456",
    "words": "cbf9459d2f20a610c9985e8848990db1d0950eab19ef831baa9e4a893411f9f009cda5591f635de7679b48d28f1d4304cfeceb0b55eb63939e0c4ffb5a89730134be7af59f832c488cc86c420c4dd4c839cc8a081ff9c032533c23996e5d98",
    "hpasswd": "63cb5a4324f549dfe8413d1eae5a552753d1c3d315aacb13e773fbf6996f45de"
}
```



**返回数据**

请求成功，返回注册成功的相关提示。

请求失败，返回失败的原因及相关信息。



---

### 用户密码找回

用户的密码 hash 后保存在用户本地，server 端保存了这个 hash 值，目的是恢复数据（比如换了手机）后，不用重置密码即可使用。

用户密码找回接口的本质是，在验证了用户的 email/sms 后，允许 client 端能够重置 client 本地的用户密码。

这个操作的实质是可以绕过 server 端，在 client 端直接重置的，但是，我们的 client 端程序还是要遵循上面的发验证码-验证验证码的流程来假装进行了密码的重置。

整个流程总结如下：

1. 用户在 client ui 上操作`找回密码`，跳转到确认给指定的 email/phone 发送验证码的界面
2. 用户确认此操作，client 请求 server 某发送验证码接口
3. 用户收到验证码，输入验证码，client 发送此验证码到 server 端，server 给出正确/错误提示
4. 如果正确，client ui 跳转到让用户输入新密码的 ui，用户输入了后
   1. 修改本地已保存的密码的 hash 值
   2. 把这个新的 hash 值传递给 server 端
5. 整个找回密码结束

上述流程涉及到如下的 server 端接口

- 发送验证码接口 /api/v1/kyc/passwd/coderequest
- 验证验证码接口 /api/v1/kyc/passwd/codeverify
- 更新用户密码 hash 值接口 /api/v1/kyc/passwd/updatepwdhash

#### 发送验证码接口

| HTTP METHOD | POST                           |
| ----------- | ------------------------------ |
| URI         | /api/v1/kyc/passwd/coderequest |
| 输入参数    | http request body              |
| HEADERS     | X-KYC-TOKEN                    |

**输入参数**

| field    | value                              | remark               |
| -------- | ---------------------------------- | -------------------- |
| type     | 固定值，可选值为 `sms`或者 `email` |                      |
| receiver | email 地址或者 phone 号码          | 号码必须包含国际区号 |

**返回数据**

标记请求成功或失败的提示。



---

#### 验证验证码接口

| HTTP METHOD | POST                          |
| ----------- | ----------------------------- |
| URI         | /api/v1/kyc/passwd/codeverify |
| 输入参数    | http request body             |
| HEADERS     | X-KYC-TOKEN                   |

**输入参数**

| field    | value          | remark                  |
| -------- | -------------- | ----------------------- |
| code     | 验证码值       |                         |
| receiver | 邮箱或者手机号 | 与请求 request 接口一致 |

**返回数据**

标记请求成功或失败的提示。



---

#### 更新用户密码 hash 值接口

| HTTP METHOD | POST                             |      |
| ----------- | -------------------------------- | ---- |
| URI         | /api/v1/kyc/passwd/updatepwdhash |      |
| 输入参数    | http request body                |      |
| HEADERS     | X-KYC-TOKEN                      |      |

**输入参数**

| field   | value                          | remark |
| ------- | ------------------------------ | ------ |
| hpasswd | 密码的 sha256 hash值，hex 编码 |        |

**返回数据**

标记请求成功或失败的提示。



---

### 用户密码更新

用户可以在 client 端修改密码，用户修改的是本地的密码，但是因为我们 server 端存储了这个密码的  hash 值，所以，在 client 端修改了密码后，也需要请求 server 端，更新密码的 hash 值。

实现步骤如下：

1. client 端修改密码
2. client 把修改后的密码 hash 值发送给 server 端 api，此 api 与上述的用户密码找回中的`更新用户密码 hash 值接口`一致。



---

### level3 的 hkid 等信息认证接口

todo



---

### 用户 key pair 恢复/找回接口

这个接口能够工作的前提是，server 端保存了用户的 mnemonic words。

而要 server 端保存用户的 mnemonic words，就要用户已被验证了 email/sms 值。

所以，这个找回的接口的操作，也是一个根据验证码来操作的一套接口。

整个流程如下：

1. client 端 ui 展示找回/恢复密钥的 ui，用户`输入` phone number 或者 email，client 请求 server 发送验证码。
   1. 需要注意，这里与找回`用户密码`不是同一个东西。用户密码是本地的操作，这是恢复 key pair。
2. 用户输入收到的验证码，client 提交给 server，server 验证通过，会在这个同一次请求中，下发解密后的 mnemonic words 及其他一些信息；
3. client 收到 mnemonic words 后，按 hd wallet 规则，重新构建生成 key pair。
   1. 根据可能的情况，需要去 blockchain 或者其他什么地方拉取一些信息
   2. 需要根据一些其他什么规则，额外生成一些数据，显示在界面上。
4. 恢复数据结束

下面描述使用到的 api。



#### 发送验证码接口

| HTTP METHOD | POST                                     |
| ----------- | ---------------------------------------- |
| URI         | /api/v1/kyc/keypair/recovery/coderequest |
| 输入参数    | http request body                        |
| HEADERS     | **不需要 X-KYC-TOKEN**                   |

**输入参数**

| field    | value                              | remark               |
| -------- | ---------------------------------- | -------------------- |
| type     | 固定值，可选值为 `sms`或者 `email` |                      |
| receiver | email 地址或者 phone 号码          | 号码必须包含国际区号 |

**返回数据**

标记请求成功或失败的提示。



---

#### 验证验证码接口

| HTTP METHOD | POST                                    |
| ----------- | --------------------------------------- |
| URI         | /api/v1/kyc/keypair/recovery/codeverify |
| 输入参数    | http request body                       |
| HEADERS     | **不需要传递 X-KYC-TOKEN**              |

**输入参数**

| field    | value                     | remark               |
| -------- | ------------------------- | -------------------- |
| code     | 验证码值                  |                      |
| receiver | email 地址或者 phone 号码 | 号码必须包含国际区号 |

**返回数据**

本接口的返回数据，主要包含的是助记词的数据。需要注意的是，接口返回的助记词是明文的，且同时包含了 aes key。

client 收到了返回数据，恢复了 master key pair  后，需要验证自己生成的 aes key 与此接口中的 aes key 值是否一致，

如果不一致，一定是出错了，要么是 client，要么是 server。

如果一致，就按规则加密后保存起来即可。

下面是返回的 data 字段的例子。

```json
{
    "key": "7386454637d5356f160d3220442209f3",
    "words": "wolf|juice|proud|gown|wool|unfair|wall|cliff|insect|more|detail|hub",
    "hpasswd": "sagagtagafg"
}
```

如果有错误，就提示错误信息。



---

### client 提交 level3 数据接口

client 提交 level3 数据，按照约定的格式进行数据的提交。

| HTTP METHOD | POST                      |
| ----------- | ------------------------- |
| URI         | /api/v1/kyc/level3/create |
| 输入参数    | 见下面 json 例子          |
| HEADERS     | 传递 X-KYC-TOKEN          |

因为这个结构很啰唆，根据约定，按照“分区”去区分是哪一块的数据。每一个分区内部，可能还有子分区。

在真实的数据上，就是一个按照`从左到右、从上到下`的数据列表。

```json
{
    "section1": {
        "sub1": [],
        "sub2": [],
        "sub3": []
    },
    "section2": {
        "sub1": []
    },
    "section3": {
        "sub1": [],
        "sub2": [],
        "sub3": [],
        "sub4": [],
        "sub5": [],
        "sub6": []
    },
    "section4": {
        "sub1": [],
        "sub2": []
    },
    "section5": {
        "sub1": []
    },
    "section6": {
        "sub1": [],
        "sub2": []
    }
}
```

返回数据

返回成功或者重复的提示。



---

### client 查看自己提交的 level3 数据接口

| HTTP METHOD | GET                     |
| ----------- | ----------------------- |
| URI         | /api/v1/kyc/level3/info |
| 输入参数    | 无参数                  |
| HEADERS     | 传递 X-KYC-TOKEN        |

返回例子

```json
{
	"code": 200,
	"msg": "OK",
	"data": {
		"id": "64c76f873b0a8a4622b30ab3",
		"cbdcUserId": "64c76ee73b0a8a4622b30aae",
		"address": "0xcb34c649dab213a395164474e943c21b0e2126af",
		"eip55Address": "0xCb34c649dAb213A395164474e943C21B0E2126Af",
		"riskMark": "67",
		"data": {
			"section1": {
				"id": "64c76f873b0a8a4622b30ab4",
				"level3Id": "64c76f873b0a8a4622b30ab3",
				"section": "section1",
				"status": {
					"status": "SENDVC",
					"created": 1690791815,
					"requestMoreInfo": {
						"question": "need more info",
						"answer": "",
						"created": 1690796936,
						"updated": 1690796936
					}
				},
				"sub1": [
					"John",
					"Smith",
					"张",
					"大大",
					"HK ID",
					"Resident",
					"A123456",
					"City",
					"2023-06-23",
					"City"
				],
				"sub2": [
					"Flat1802,BlockB,NewT",
					"/data/user/0/com.emali.cbdc.wallet/cache/9cf48497-4a12-49b4-b0a2-0977528112fb/IMG_20210718_230823_882.jpg",
					"Flat1802,BlockB,NewT"
				],
				"sub3": [
					"+987 65432",
					"+987 65432",
					"+987 65432",
					"johnsmith@gmail.com"
				]
			},
			"section2": {
				"id": "64c76f873b0a8a4622b30ab5",
				"level3Id": "64c76f873b0a8a4622b30ab3",
				"section": "section2",
				"status": {
					"status": "SENDVC",
					"created": 1690791815
				},
				"sub1": [
					"Hong Kong SAR",
					"",
					""
				]
			},
			"section3": {
				"id": "64c76f873b0a8a4622b30ab6",
				"level3Id": "64c76f873b0a8a4622b30ab3",
				"section": "section3",
				"status": {
					"status": "SENDVC",
					"created": 1690791815
				},
				"sub1": [
					"Yes",
					"Loremlpsum",
					"Loremlpsum",
					"Loremlpsum",
					"Loremlpsum",
					"Loremlpsum"
				],
				"sub2": [
					"Loremlpsum",
					"HK$100000"
				],
				"sub3": [
					"Loremlpsum",
					"Loremlpsum",
					"Loremlpsum",
					"Loremlpsum",
					"Loremlpsum"
				],
				"sub4": [
					"Sold to Third Patrty",
					"Loremlpsum",
					"2013-06-23",
					"Loremlpsum"
				],
				"sub5": [
					"Loremlpsum",
					"Loremlpsum",
					"Loremlpsum",
					"2003-06-23",
					"2013-06-23",
					"Loremlpsum"
				],
				"sub6": [
					"Loremlpsum",
					"Loremlpsum"
				]
			},
			"section4": {
				"id": "64c76f873b0a8a4622b30ab7",
				"level3Id": "64c76f873b0a8a4622b30ab3",
				"section": "section4",
				"status": {
					"status": "CHECKED",
					"created": 1690791815
				},
				"sub1": [
					"Single",
					"Professor or doctor",
					"<$500,000",
					"> $1,000,000",
					"Owned"
				],
				"sub2": [
					"Less than 1 year",
					"Capital preservation and regular income",
					"Less than 10%",
					"Do not know how to react",
					"No capital losses are acceptable",
					"Always the possible losses"
				]
			},
			"section5": {
				"id": "64c76f873b0a8a4622b30ab8",
				"level3Id": "64c76f873b0a8a4622b30ab3",
				"section": "section5",
				"status": {
					"status": "CHECKED",
					"created": 1690791815
				},
				"sub1": [
					"No",
					"No",
					"No",
					"No",
					"No",
					"No"
				]
			},
			"section6": {
				"id": "64c76f873b0a8a4622b30ab9",
				"level3Id": "64c76f873b0a8a4622b30ab3",
				"section": "section6",
				"status": {
					"status": "CHECKED",
					"created": 1690791815
				},
				"sub1": [
					"I have read and understood the Client Account Terms and Conditions, including in particular the above Risk Statements, and understand that may contact Arta GlobalN Markets Limited (the contact details of which can be found in the\"Contact Us\", link below) and/or take independent advice if I have any questions.",
					"I have obtained independent tax and legal advice from my home/country of permanent residence confirming that the structure is in compliance with the applicable [Country(ies)] <- populate the country (ies) listed in CRS form tax laws and, where applicable, I will be fully responsible for any tax obligations in my home/permanent residence jurisdiction.",
					"I am currently solvent.",
					"All funds or other assets being transferred to the Company are cleared assets and of a non-criminal origin, and none of my assets, net worth, investments, or income now or in the future have/will come from money laundering, drug trafficking or any other illegal activity at home, in Hong Kong or any other jurisdiction.",
					"I confirm that all the above statements are understood fully and are true and correct as of the date of this declaration.",
					"I warrant that the information I provide in this investor risk profiling questionnaire is true and correct.",
					"I understand that the risk tolerance level result is valid for 12 months only from the date of this assessment. If my risk tolerance level result is expired. I may not be able to purchase certain products. If I believe my risk tolerance level result within the past 12 months is no longer valid. I shall complete a new questionnaire for reassessment purposes.",
					"I confirm that I have been reminded and am aware that I should have adequate liquid funds to meet foreseen and unforeseen events.",
					"I confirm and undertake that I will update you immediately of any changes in the formation I provided."
				],
				"sub2": [
					"Self"
				]
			}
		},
		"status": {
			"mark": 1,
			"vcSend": 3,
			"status": "PROGRESS"
		},
		"created": 1690791815,
		"updated": 1690799273
	}
}
```



---

### client 文件管理接口

#### client 端文件上传接口

需要注意的事，文件上传时，会被计算它的 sha256 值。如果有相同 hash 值的上传，server 端不会保存两份数据，而是把之前的数据的 metadata 返回回来，并标记 `repeated=true`。

| HTTP METHOD | POST                    |
| ----------- | ----------------------- |
| URI         | /api/v1/kyc/file/upload |
| 输入参数    | multiform 形式          |
| HEADERS     | X-KYC-TOKEN             |

文件上传使用 `multipart` form 的形式上传文件，form 的 key 名字是 `file`，值是指定的文件路径。

见下面截图所示。

![image-20230717155229853](https://raw.githubusercontent.com/leyle/picbase/master/gopic/image-20230717155229853.png)



**返回数据**

返回的数据包含了很多的内容，其中最重要的是 `id`字段的值。

这个 `id`就是后续用来下载文件的接口中的 id 参数的值。

如果上传了文件后，需要保存文件相关的内容，这个 id 是必须要存放到相应的位置上的。



---

#### client 端文件下载接口

文件下载只支持明确指定 id 的文件的下载。

| HTTP METHOD | GET                                                        |
| ----------- | ---------------------------------------------------------- |
| URI         | /api/v1/kyc/file/download                                  |
| 输入参数    | 支持两个 url query string 参数，具体见下面描述             |
| HEADERS     | X-KYC-TOKEN (注意，为了在浏览器调试方便，暂时不做权限验证) |

**输入参数**

支持两个 url query string 参数：

- id 参数，必输，指的是要下载的文件的 id，这个值从上传接口返回的数据中的 id 值。
- debug=yes，可选参数，当这个参数存在时，api 返回的不再是原始文件数据，而是 metadata 信息。

下面是两个请求例子

```shell
# 请求文件
GET /api/v1/kyc/file/download?id=64b4d815cebde73b15ce2ec1

# 仅请求文件 metadata，不需要下载文件
GET /api/v1/kyc/file/download?id=64b4d815cebde73b15ce2ec1&debug=yes
```



**返回数据**

返回文件，或者 metadata 的 json 数据。



---

### client 端 market/product 查看

#### client 搜索机构列表

**注意：因为目前只有一个隐藏的机构，所以，在使用此接口时，先按照 page=1&size=1 传递参数，如果有返回值，就使用列表的第一个机构**

| HTTP METHOD | GET                                                          |
| ----------- | ------------------------------------------------------------ |
| URI         | /api/v1/org/query                                            |
| 输入参数    | 支持 url query 参数。目前支持分页参数<br />1、page<br />2、size<br />需要注意的是，默认返回的数据是按照 updated 逆序输出的 |
| HEADERS     | X-KYC-TOKEN                                                  |

输入参数：

url query 参数。

例子：

```shell
/api/v1/org/query?page=1&size=20
```

返回例子：

返回的结构中，在 data 字段，会包含如下结构

- total - 满足条件的数据总数
- page - 用户输入的 page 原样返回
- size - 用户输入的 size 原样返回
- data - 真实的数据列表

```json
{
	"code": 200,
	"msg": "OK",
	"data": {
		"data": [
			{
				"id": "64c8ae1790c07a3515370d1c",
				"name": "FuTu cc",
				"description": "this is description info",
				"image": "img-id",
				"extraInfo": "optional extra info, json string format",
				"created": 1690873367,
				"updated": 1690873367
			},
			{
				"id": "64c8ad758ba67a9d76d44430",
				"name": "FuTu inc",
				"description": "this is description info",
				"image": "",
				"extraInfo": "optional extra info, json string format",
				"created": 1690873205,
				"updated": 1690873205
			}
		],
		"page": 1,
		"size": 10,
		"total": 2
	}
}
```





---

#### client 查看指定 id 的机构信息

| HTTP METHOD | GET                                           |
| ----------- | --------------------------------------------- |
| URI         | /api/v1/org/detail                            |
| 输入参数    | 支持一个 url query 参数 id<br />比如 ?id=xxxx |
| HEADERS     | X-KYC-TOKEN                                   |

输入例子

```shell
/api/v1/org/detail?id=64c8ae1790c07a3515370d1c
```

返回例子

```json
{
	"code": 200,
	"msg": "OK",
	"data": {
		"id": "64c8ae1790c07a3515370d1c",
		"name": "FuTu cc",
		"description": "this is description info",
		"image": "img-id",
		"extraInfo": "optional extra info, json string format",
		"created": 1690873367,
		"updated": 1690873367
	}
}
```



---

#### client 搜索 product 列表

| HTTP METHOD | GET                                                          |
| ----------- | ------------------------------------------------------------ |
| URI         | /api/v1/product/query                                        |
| 输入参数    | 支持 url query 参数。目前支持分页参数<br />1、page<br />2、size<br />3、name - 指的是 product name，模糊匹配<br />4、org - 指的是机构名，模糊匹配<br />需要注意的是，默认返回的数据是按照 updated 逆序输出的 |
| HEADERS     | X-KYC-TOKEN                                                  |

输入参数：

url query 参数。

例子：

```shell
/api/v1/product/query?page=1&size=20
```

返回例子：

返回的结构中，在 data 字段，会包含如下结构

- total - 满足条件的数据总数
- page - 用户输入的 page 原样返回
- size - 用户输入的 size 原样返回
- data - 真实的数据列表

```json
{
	"code": 200,
	"msg": "OK",
	"data": {
		"data": [
			{
				"id": "64c9cbcefa42db7f577041f5",
				"orgId": "64c8ae1790c07a3515370d1c",
				"orgName": "FuTu cc",
				"creatorId": "646448e3ab25112e85fe7dc1",
				"productId": 0,
				"name": "product name 4",
				"description": "description text",
				"images": [
					"img1",
					"img2"
				],
				"extraInfo": "extra info",
				"totalSupplySize": 1000000,
				"minSubSize": 100,
				"maxSubSize": 1000,
				"interestRate": 1.278,
				"lockPeriod": 86400,
				"syncToContract": false,
				"status": "ONLINE",
				"created": 1690946510,
				"updated": 1690947645
			}
		],
		"page": 1,
		"size": 10,
		"total": 1
	}
}
```



---

#### client 查看指定 id 的 product

| HTTP METHOD | GET                                           |
| ----------- | --------------------------------------------- |
| URI         | /api/v1/product/detail                        |
| 输入参数    | 支持一个 url query 参数 id<br />比如 ?id=xxxx |
| HEADERS     | X-KYC-TOKEN                                   |

输入例子

```shell
/api/v1/product/detail?id=64c9cbcefa42db7f577041f5
```

返回例子

```json
{
	"code": 200,
	"msg": "OK",
	"data": {
		"id": "64c9cbcefa42db7f577041f5",
		"orgId": "64c8ae1790c07a3515370d1c",
		"orgName": "FuTu cc",
		"creatorId": "646448e3ab25112e85fe7dc1",
		"productId": 0,
		"name": "product name 4",
		"description": "description text",
		"images": [
			"img1",
			"img2"
		],
		"extraInfo": "extra info",
		"totalSupplySize": 1000000,
		"minSubSize": 100,
		"maxSubSize": 1000,
		"interestRate": 1.278,
		"lockPeriod": 86400,
		"syncToContract": false,
		"status": "ONLINE",
		"created": 1690946510,
		"updated": 1690947645
	}
}
```



---

### server 端管理接口

⚠️需要注意的是，server 端使用的是用户名、密码获取的 token。

所以，在调用 server 端管理接口时，需要在 headers 配置一个 `Authorization` 的 header，它的值是 jwt token 的格式。比如 `Bearer + 空格 +   jwt token`。比如 `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiI2NDY0NDhlM2FiMjUxMTJlODVmZTdkYzEiLCJpYXQiOjE2ODgzNTg1ODUsImV4cCI6MTY4ODM1OTQ4NX0.EH5binIM6hFnxus_bhMB-ZH5V5ZpIyZB6OaPRlCloSg`。

目前 server 端主要是管理 level3 相关的数据，包括查看、审核等等操作。

这里列出来可能有的 api 列表：

- /api/v1/kyc/admin/user/query - 获取已注册的用户列表
- /api/v1/kyc/admin/user/detail - 读取指定 id 的用户信息
- /api/v1/kyc/admin/level3/query   - 获取已提交的 level3 数据列表
- /api/v1/kyc/admin/level3/detail   - 查看指定 id 的 level3 数据详情
- 其他 api



---

#### server 端查看 user 数据列表

| HTTP METHOD | GET                                                          |
| ----------- | ------------------------------------------------------------ |
| URI         | /api/v1/kyc/admin/user/query                                 |
| 输入参数    | 支持 url query 参数。目前支持分页参数<br />1、page<br />2、size<br />需要注意的是，默认返回的数据是按照 updated 逆序输出的 |
| HEADERS     | Authorization                                                |

输入参数：

url query 参数。

例子：

```shell
/api/v1/kyc/admin/user/query?page=1&size=20
```

返回例子：

返回的结构中，在 data 字段，会包含如下结构

- total - 满足条件的数据总数
- page - 用户输入的 page 原样返回
- size - 用户输入的 size 原样返回
- data - 真实的数据列表

```json
{
	"code": 200,
	"msg": "OK",
	"data": {
		"data": [
			{
				"id": "64b1003944a9e7479d3bde84",
				"phone": "+85212345670",
				"email": "",
				"publicKey": "039c83fc52d46fbe20ca91d6c2ed28a41bb3dfc0cf8915e9cc50b491cf92cc7842",
				"address": "0xcb34c649dab213a395164474e943c21b0e2126af",
				"extraInfo": null,
				"level": 3,
				"created": 1689321529,
				"updated": 1689324675
			}
		],
		"page": 1,
		"size": 10,
		"total": 1
	}
}
```



---

#### server 端查看指定 id 的 user 信息

| HTTP METHOD | GET                                           |
| ----------- | --------------------------------------------- |
| URI         | /api/v1/kyc/admin/user/detail                 |
| 输入参数    | 支持一个 url query 参数 id<br />比如 ?id=xxxx |
| HEADERS     | Authorization                                 |

输入例子

```shell
/api/v1/kyc/admin/user/detail?id=64b1003944a9e7479d3bde84
```

返回例子

```json
{
	"code": 200,
	"msg": "OK",
	"data": {
		"id": "64b1003944a9e7479d3bde84",
		"phone": "+85212345670",
		"email": "",
		"publicKey": "039c83fc52d46fbe20ca91d6c2ed28a41bb3dfc0cf8915e9cc50b491cf92cc7842",
		"address": "0xcb34c649dab213a395164474e943c21b0e2126af",
		"extraInfo": null,
		"level": 3,
		"created": 1689321529,
		"updated": 1689324675
	}
}
```



---

#### server 端获取 level3 数据列表

| HTTP METHOD | GET                                                          |
| ----------- | ------------------------------------------------------------ |
| URI         | /api/v1/kyc/admin/level3/query                               |
| 输入参数    | 支持 url query 参数。目前支持分页参数<br />1、page<br />2、size<br />3、addr - 模糊匹配<br />4、status - 默认返回的数据是PENDING  和 PROGRESS 的<br />5、phone - 模糊匹配<br />需要注意的是，默认返回的数据是按照 updated 逆序输出的 |
| HEADERS     | Authorization                                                |

输入参数：

url query 参数。

例子：

```shell
/api/v1/kyc/admin/level3/query?page=1&size=20
```

返回例子：

返回的结构中，在 data 字段，会包含如下结构

- total - 满足条件的数据总数
- page - 用户输入的 page 原样返回
- size - 用户输入的 size 原样返回
- data - 真实的数据列表

```json
{
	"code": 200,
	"msg": "OK",
	"data": {
		"data": [
			{
				"id": "64a24e496fedd46e9f803091",
				"cbdcUserId": "649d57f067f45f78cad9d6bf",
				"address": "0xcb34c649dab213a395164474e943c21b0e2126af",
				"data": {
					"section1": {
						"sub1": [],
						"sub2": [],
						"sub3": []
					},
					"section2": [],
					"section3": {
						"sub1": [],
						"sub2": [],
						"sub3": [],
						"sub4": [],
						"sub5": [],
						"sub6": []
					},
					"section4": {
						"sub1": [],
						"sub2": []
					},
					"section5": [],
					"section6": {
						"sub1": [],
						"sub2": []
					}
				},
				"status": "PENDING",
				"created": 1688358473,
				"updated": 1688358473
			}
		],
		"page": 1,
		"size": 10,
		"total": 1
	}
}
```



---

#### server 端查看指定的 id 的 level3 数据

| HTTP METHOD | GET                                           |
| ----------- | --------------------------------------------- |
| URI         | /api/v1/kyc/admin/level3/detail               |
| 输入参数    | 支持一个 url query 参数 id<br />比如 ?id=xxxx |
| HEADERS     | Authorization                                 |

输入例子

```shell
/api/v1/kyc/admin/level3/detail?id=64a24e496fedd46e9f803091
```

返回例子

```json
{
	"code": 200,
	"msg": "OK",
	"data": {
		"id": "64a24e496fedd46e9f803091",
		"cbdcUserId": "649d57f067f45f78cad9d6bf",
		"address": "0xcb34c649dab213a395164474e943c21b0e2126af",
		"data": {
			"section1": {
				"sub1": [],
				"sub2": [],
				"sub3": []
			},
			"section2": [],
			"section3": {
				"sub1": [],
				"sub2": [],
				"sub3": [],
				"sub4": [],
				"sub5": [],
				"sub6": []
			},
			"section4": {
				"sub1": [],
				"sub2": []
			},
			"section5": [],
			"section6": {
				"sub1": [],
				"sub2": []
			}
		},
		"status": "PENDING",
		"created": 1688358473,
		"updated": 1688358473
	}
}
```



---

#### server 端下载指定 id 的文件

文件下载只支持明确指定 id 的文件的下载。

| HTTP METHOD | GET                                                          |
| ----------- | ------------------------------------------------------------ |
| URI         | /api/v1/kyc/admin/file/download                              |
| 输入参数    | 支持两个 url query string 参数，具体见下面描述               |
| HEADERS     | Authorization (注意，为了在浏览器调试方便，暂时不做权限验证) |

**输入参数**

支持两个 url query string 参数：

- id 参数，必输，指的是要下载的文件的 id，这个值从上传接口返回的数据中的 id 值。
- debug=yes，可选参数，当这个参数存在时，api 返回的不再是原始文件数据，而是 metadata 信息。

下面是两个请求例子

```shell
# 请求文件
GET/api/v1/kyc/admin/file/download?id=64b4d815cebde73b15ce2ec1

# 仅请求文件 metadata，不需要下载文件
GET /api/v1/kyc/admin/file/download?id=64b4d815cebde73b15ce2ec1&debug=yes
```



**返回数据**

返回文件，或者 metadata 的 json 数据。



---

#### server 端 check level3 数据

| HTTP METHOD | POST                                     |
| ----------- | ---------------------------------------- |
| URI         | /api/v1/kyc/admin/level3/section/confirm |
| 输入参数    | http request body，具体见下面例子        |
| HEADERS     | Authorization                            |

**输入参数**

`id`指的是通过 `/api/v1/kyc/admin/level3/detail`接口返回的数据中，各个 `sectionX`中的那个 `id`值，不是 level3Id，需要特别注意。

```json
{
	"id": "64c76f873b0a8a4622b30ab9"
}
```

当请求满足要求时，会返回输入的 id，如果重复请求，仅仅提示请求成功，无额外信息返回。



---

#### server 端对 level3 数据发送 request more info 请求

| HTTP METHOD | POST                                         |
| ----------- | -------------------------------------------- |
| URI         | /api/v1/kyc/admin/level3/section/reqmoreinfo |
| 输入参数    | http request body，具体见下面例子            |
| HEADERS     | Authorization                                |

**输入参数**

`id`指的是通过 `/api/v1/kyc/admin/level3/detail`接口返回的数据中，各个 `sectionX`中的那个 `id`值，不是 level3Id，需要特别注意。

```json
{
	"id": "64c76f873b0a8a4622b30ab4",
	"info": "need more info"
}
```

当请求满足要求时，会返回输入的 id，如果重复请求，仅仅提示请求成功，无额外信息返回。



---

#### server 端对 level3 数据进行 send vc 操作

| HTTP METHOD | POST                                    |
| ----------- | --------------------------------------- |
| URI         | /api/v1/kyc/admin/level3/section/sendvc |
| 输入参数    | http request body，具体见下面例子       |
| HEADERS     | Authorization                           |

**输入参数**

`id`指的是通过 `/api/v1/kyc/admin/level3/detail`接口返回的数据中，各个 `sectionX`中的那个 `id`值，不是 level3Id，需要特别注意。

```json
{
	"id": "64c76f873b0a8a4622b30ab6"
}
```

当请求满足要求时，会返回输入的 id，如果重复请求，仅仅提示请求成功，无额外信息返回。



---

#### server 端对 level3 数据进行 risk mark

| HTTP METHOD | POST                                      |
| ----------- | ----------------------------------------- |
| URI         | /api/v1/kyc/admin/level3/section/markrisk |
| 输入参数    | http request body，具体见下面例子         |
| HEADERS     | Authorization                             |

**输入参数**

`id`指的是**整条 level3 数据的 id**，与上面几个 section 接口的 section id 不同。需要特别注意。

```json
{
	"id": "64c76f873b0a8a4622b30ab3",
	"mark": "67"
}
```

当请求满足要求时，会返回输入的 id，如果重复请求，仅仅提示请求成功，无额外信息返回。



---

#### server 端查看 send vc 的历史记录（列表）

| HTTP METHOD | GET                                                          |
| ----------- | ------------------------------------------------------------ |
| URI         | /api/v1/kyc/admin/level3/sendvc/query                        |
| 输入参数    | 支持 url query 参数。目前支持分页参数<br />1、page<br />2、size<br />需要注意的是，默认返回的数据是按照 updated 逆序输出的 |
| HEADERS     | Authorization                                                |

输入参数：

url query 参数。

例子：

```shell
/api/v1/kyc/admin/level3/sendvc/query?page=1&size=20
```

返回例子：

返回的结构中，在 data 字段，会包含如下结构

- total - 满足条件的数据总数
- page - 用户输入的 page 原样返回
- size - 用户输入的 size 原样返回
- data - 真实的数据列表

```json
{
	"code": 200,
	"msg": "OK",
	"data": {
		"data": [
			{
				"id": "64c78ca99f82732104e5a732",
				"level3Id": "64c76f873b0a8a4622b30ab3",
				"sectionId": "64c76f873b0a8a4622b30ab6",
				"user": {
					"id": "64c76ee73b0a8a4622b30aae",
					"phone": "+85212345678",
					"email": "",
					"publicKey": "039c83fc52d46fbe20ca91d6c2ed28a41bb3dfc0cf8915e9cc50b491cf92cc7842",
					"address": "0xcb34c649dab213a395164474e943c21b0e2126af",
					"eip55Address": "0xCb34c649dAb213A395164474e943C21B0E2126Af",
					"extraInfo": null,
					"level": 3,
					"created": 1690791655,
					"updated": 1690791815
				},
				"created": 1690799273,
				"updated": 1690799273
			},
			{
				"id": "64c78b5f089963176a87b616",
				"level3Id": "64c76f873b0a8a4622b30ab3",
				"sectionId": "64c76f873b0a8a4622b30ab5",
				"user": {
					"id": "64c76ee73b0a8a4622b30aae",
					"phone": "+85212345678",
					"email": "",
					"publicKey": "039c83fc52d46fbe20ca91d6c2ed28a41bb3dfc0cf8915e9cc50b491cf92cc7842",
					"address": "0xcb34c649dab213a395164474e943c21b0e2126af",
					"eip55Address": "0xCb34c649dAb213A395164474e943C21B0E2126Af",
					"extraInfo": null,
					"level": 3,
					"created": 1690791655,
					"updated": 1690791815
				},
				"created": 1690798943,
				"updated": 1690798943
			}
		],
		"page": 1,
		"size": 10,
		"total": 2
	}
}
```



---

#### server 端创建机构

| HTTP METHOD | POST                              |
| ----------- | --------------------------------- |
| URI         | /api/v1/org/admin/org/create      |
| 输入参数    | http request body，具体见下面例子 |
| HEADERS     | Authorization                     |

**输入参数**

必填参数

- name - 必填，且不能重复

其他参数为非必填

```json
{
	"name": "FuTu cc",
	"description": "this is description info",
	"image": "img-id",
	"extraInfo": "optional extra info, json string format"
}
```



---

#### server 端搜索机构列表

| HTTP METHOD | GET                                                          |
| ----------- | ------------------------------------------------------------ |
| URI         | /api/v1/org/admin/org/query                                  |
| 输入参数    | 支持 url query 参数。目前支持分页参数<br />1、page<br />2、size<br />需要注意的是，默认返回的数据是按照 updated 逆序输出的 |
| HEADERS     | Authorization                                                |

输入参数：

url query 参数。

例子：

```shell
/api/v1/org/admin/org/query?page=1&size=20
```

返回例子：

返回的结构中，在 data 字段，会包含如下结构

- total - 满足条件的数据总数
- page - 用户输入的 page 原样返回
- size - 用户输入的 size 原样返回
- data - 真实的数据列表

```json
{
	"code": 200,
	"msg": "OK",
	"data": {
		"data": [
			{
				"id": "64c8ae1790c07a3515370d1c",
				"name": "FuTu cc",
				"description": "this is description info",
				"image": "img-id",
				"extraInfo": "optional extra info, json string format",
				"created": 1690873367,
				"updated": 1690873367
			},
			{
				"id": "64c8ad758ba67a9d76d44430",
				"name": "FuTu inc",
				"description": "this is description info",
				"image": "",
				"extraInfo": "optional extra info, json string format",
				"created": 1690873205,
				"updated": 1690873205
			}
		],
		"page": 1,
		"size": 20,
		"total": 2
	}
}
```



---

#### server 端查看指定 id 的机构信息

| HTTP METHOD | GET                                           |
| ----------- | --------------------------------------------- |
| URI         | /api/v1/org/admin/org/detail                  |
| 输入参数    | 支持一个 url query 参数 id<br />比如 ?id=xxxx |
| HEADERS     | Authorization                                 |

输入例子

```shell
/api/v1/org/admin/org/detail?id=64c8ae1790c07a3515370d1c
```

返回例子

```json
{
	"code": 200,
	"msg": "OK",
	"data": {
		"id": "64c8ae1790c07a3515370d1c",
		"name": "FuTu cc",
		"description": "this is description info",
		"image": "img-id",
		"extraInfo": "optional extra info, json string format",
		"created": 1690873367,
		"updated": 1690873367
	}
}
```



---

#### server 端创建 product

| HTTP METHOD | POST                              |
| ----------- | --------------------------------- |
| URI         | /api/v1/org/admin/product/create  |
| 输入参数    | http request body，具体见下面例子 |
| HEADERS     | Authorization                     |

**输入参数**

必填参数

- orgId -  从之前的 org query 接口得到数据的 id 值
- name - 且不能重复
- symbol - token 的缩写，比如 TPA,  DDD 之类的，必填
- totalSupplySize
- minSubSize
- maxSubSize
- interestRate - 浮点数，目前精确到万一之一，即 x.yz%。比如可以输入 1%|1.2%|1.23%，但是不能输如 1.234%，4 会被截断
- lockPeriod - 秒，比如 24小时就是 86400 秒
- online - 页面上的 showing 就是 true，hidding 就是 false

其他参数为非必填

```json
{
	"orgId": "64c8ae1790c07a3515370d1c",
	"name": "product name 4",
	"description": "description text",
	"images": ["img1", "img2"],
	"extraInfo": "extra info",
	"totalSupplySize": 1000000,
	"minSubSize": 100,
	"maxSubSize": 1000,
	"interestRate": 1.278,
	"lockPeriod": 86400,
	"online": false	
}
```



---

#### server 端上线 hidding 状态的 product

当创建商品时，标记为 hidding 状态，才可以进行此操作。

当商品在创建时被标记为 hidding 状态，在这条 product 数据的 status 的值就是 `DRAFT`。换句话说，只有 product status 值是 `DRAFT`的才能够做这个操作。

| HTTP METHOD | POST                                |
| ----------- | ----------------------------------- |
| URI         | /api/v1/org/admin/product/setonline |
| 输入参数    | http request body，具体见下面例子   |
| HEADERS     | Authorization                       |

**输入参数**

必填参数

- id - 指的是 product 数据的 `id`，注意，商品中还有个 productId, 不是这个 id。

其他参数为非必填

```json
{
	"id": "64c9cbcefa42db7f577041f5"
}
```



---

#### server 端下线 online 状态的 product

todo



---

#### server 端搜索 product 列表

| HTTP METHOD | GET                                                          |
| ----------- | ------------------------------------------------------------ |
| URI         | /api/v1/org/admin/product/query                              |
| 输入参数    | 支持 url query 参数。目前支持分页参数<br />1、page<br />2、size<br />3、name - 商品名，模糊匹配<br />4、status - 支持多个 status 筛选，比如 status=DRAFT&status=ONLINE <br />需要注意的是，默认返回的数据是按照 updated 逆序输出的 |
| HEADERS     | Authorization                                                |

输入参数：

url query 参数。

例子：

```shell
/api/v1/org/admin/product/query?page=1&size=20
```

返回例子：

返回的结构中，在 data 字段，会包含如下结构

- total - 满足条件的数据总数
- page - 用户输入的 page 原样返回
- size - 用户输入的 size 原样返回
- data - 真实的数据列表

```json
{
	"code": 200,
	"msg": "OK",
	"data": {
		"data": [
			{
				"id": "64c9cbcefa42db7f577041f5",
				"orgId": "64c8ae1790c07a3515370d1c",
				"orgName": "FuTu cc",
				"creatorId": "646448e3ab25112e85fe7dc1",
				"productId": 0,
				"name": "product name 4",
				"description": "description text",
				"images": [
					"img1",
					"img2"
				],
				"extraInfo": "extra info",
				"totalSupplySize": 1000000,
				"minSubSize": 100,
				"maxSubSize": 1000,
				"interestRate": 1.278,
				"lockPeriod": 86400,
				"syncToContract": false,
				"status": "ONLINE",
				"created": 1690946510,
				"updated": 1690947645
			},
			{
				"id": "64c9cbadfa42db7f577041f2",
				"orgId": "64c8ae1790c07a3515370d1c",
				"orgName": "FuTu cc",
				"creatorId": "646448e3ab25112e85fe7dc1",
				"productId": 0,
				"name": "product name 3",
				"description": "description text",
				"images": [
					"img1",
					"img2"
				],
				"extraInfo": "extra info",
				"totalSupplySize": 1000000,
				"minSubSize": 100,
				"maxSubSize": 1000,
				"interestRate": 1.27,
				"lockPeriod": 86400,
				"syncToContract": false,
				"status": "DRAFT",
				"created": 1690946477,
				"updated": 1690946477
			},
			{
				"id": "64c9cb8cfa42db7f577041f0",
				"orgId": "64c8ae1790c07a3515370d1c",
				"orgName": "FuTu cc",
				"creatorId": "646448e3ab25112e85fe7dc1",
				"productId": 0,
				"name": "product name 2",
				"description": "description text",
				"images": [
					"img1",
					"img2"
				],
				"extraInfo": "extra info",
				"totalSupplySize": 1000000,
				"minSubSize": 100,
				"maxSubSize": 1000,
				"interestRate": 1.24,
				"lockPeriod": 86400,
				"syncToContract": false,
				"status": "DRAFT",
				"created": 1690946444,
				"updated": 1690946444
			},
			{
				"id": "64c9caeec0262addc0c079e4",
				"orgId": "64c8ae1790c07a3515370d1c",
				"orgName": "FuTu cc",
				"creatorId": "646448e3ab25112e85fe7dc1",
				"productId": 0,
				"name": "product name 1",
				"description": "description text",
				"images": [
					"img1",
					"img2"
				],
				"extraInfo": "extra info",
				"totalSupplySize": 1000000,
				"minSubSize": 100,
				"maxSubSize": 1000,
				"interestRate": 1.2300000190734863,
				"lockPeriod": 86400,
				"syncToContract": false,
				"status": "DRAFT",
				"created": 1690946286,
				"updated": 1690946286
			}
		],
		"page": 1,
		"size": 10,
		"total": 4
	}
}
```



----

#### server 端查看指定 id 的 product

| HTTP METHOD | GET                                           |
| ----------- | --------------------------------------------- |
| URI         | /api/v1/org/admin/product/detail              |
| 输入参数    | 支持一个 url query 参数 id<br />比如 ?id=xxxx |
| HEADERS     | Authorization                                 |

输入例子

```shell
/api/v1/org/admin/product/detail?id=64c9cbcefa42db7f577041f5
```

返回例子

```json
{
	"code": 200,
	"msg": "OK",
	"data": {
		"id": "64c9cbcefa42db7f577041f5",
		"orgId": "64c8ae1790c07a3515370d1c",
		"orgName": "FuTu cc",
		"creatorId": "646448e3ab25112e85fe7dc1",
		"productId": 0,
		"name": "product name 4",
		"description": "description text",
		"images": [
			"img1",
			"img2"
		],
		"extraInfo": "extra info",
		"totalSupplySize": 1000000,
		"minSubSize": 100,
		"maxSubSize": 1000,
		"interestRate": 1.278,
		"lockPeriod": 86400,
		"online": true,
		"status": "ONLINE",
		"created": 1690946510,
		"updated": 1690947645
	}
}
```



---

