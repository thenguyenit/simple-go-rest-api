Mình tìm hiểu Golang một cách rất tình cờ và những điều tò mò được giải toả đã nảy sinh những điều thú vị từ Golang, và đây là bài tập nhỏ thứ 2 mà mình thử dùng với Golang.

Mình sẽ xây dụng microservice Restful API đơn giản gồm 2 endpoint: Authorization Server và Product Application Server

## Authorization Server và JWT (JSON Web Tokens) 
Authorization Server sẽ sử dụng một phương thức xác thực đơn giản là JWT với thuật toán hash HS256 để xác thực việc giao tiếp giữa Product Application Server với Client.

```
     
    +----------------+                                  +---------------+
    |   Application  |--(1)- Authorization Request ->   | Authorization |
    |    (Client)    |                                  |     Server    |
    |                |<-(2)-- Return Access Token ---   |               |
    +----------------+                                  +---------------+
            |
            |                       +---------------+
            |                       | Application   |
            --(3)-Use Access Token->|     Server    |
                                    |               |
                                    +---------------+

```

Package:
```
go get github.com/dgrijalva/jwt-go
```

### JSON Web Tokens

Mục đích của Token là để kiểm tra xem một request từ Client tới Application Server đã được xác thực hay chưa (hay đã đăng nhập với username/password). 
Client sẽ xác thực với Authorization Server, sau đó Authorization Server sẽ gửi cho Client một Token, kể từ đó việc giao tiếp của Client đến những Application Server sẽ gửi kèm theo Token này nhằm để Application Server xác thực.

Cấu trúc của một JSON Web Token sẽ bao gồm 3 thứ header, payload và signature
**Header** 
```json
{
    "typ": "JWT",
    "alg": "HS256"
}
```
Giá trị của Header JSON bao gồm `typ` dùng để xác định loại token là JWT và `alg` viết tắc của algorithm là thuật toán hash

Có hai thuật toán hash hay sử dụng mà JWT có support là HS256 và RS256

**HS256** thuộc loại thuật toán đối xứng (symmetric algorithm), có nghĩa là ở đây chỉ có một secret key được dùng cho cả bước tạo signature và validate
HS256 là mặc định của JWT và cũng được recommended vì nó khá nhỏ và việc tính toán trên thuật toán này nhanh hơn.
```
HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  your-256-bit-secret
)
```

Ngược lại, **RS256** thuộc loại thuật toán bất đối xứng (asymmatric algorithm), có nghĩa là nếu bạn cần tạo ra một signature thì phải cần một cặp public/private key.
```
RSASHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDdlatRjRjogo3WojgGHFHYLugdUWAY9iR3fyarWNA1KoS8kVw33cJibXr8bvwUAUparCwlvdbH6dvEOfou0/gCFQsHUfQrSDv+MuSUMAe8jzKE4qW+jK+xQU9a03GUnKHkkle+Q0pX/g6jXZ7r1/xAK5Do2kQ+X5xK9cipRgEKwIDAQAB
-----END PUBLIC KEY-----
,
  
-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQDdlatRjRjogo3WojgGHFHYLugdUWAY9iR3fy4arWNA1KoS8kVw33cJibXr8bvwUAUparCwlvdbH6dvEOfou0/gCFQsHUfQrSDv+MuSUMAe8jzKE4qW+jK+xQU9a03GUnKHkkle+Q0pX/g6jXZ7r1/xAK5Do2kQ+X5xK9cipRgEKwIDAQABAoGAD+onAtVye4ic7VR7V50DF9bOnwRwNXrARcDhq9LWNRrRGElESYYTQ6EbatXS3MCyjjX2eMhu/aF5YhXBwkppwxg+EOmXeh+MzL7Zh284OuPbkglAaGhV9bb6/5CpuGb1esyPbYW+Ty2PC0GSZfIXkXs76jXAu9TOBvD0ybc2YlkCQQDywg2R/7t3Q2OE2+yo382CLJdrlSLVROWKwb4tb2PjhY4XAwV8d1vy0RenxTB+K5Mu57uVSTHtrMK0GAtFr833AkEA6avx20OHo61Yela/4k5kQDtjEf1N0LfI+BcWZtxsS3jDM3i1Hp0KSu5rsCPb8acJo5RO26gGVrfAsDcIXKC+bQJAZZ2XIpsitLyPpuiMOvBbzPavd4gY6Z8KWrfYzJoI/Q9FuBo6rKwl4BFoToD7WIUS+hpkagwWiz+6zLoX1dbOZwJACmH5fSSjAkLRi54PKJ8TFUeOP15h9sQzydI8zJU+upvDEKZsZc/UhT/SySDOxQ4G/523Y0sz/OZtSWcol/UMgQJALesy++GdvoIDLfJX5GBQpuFgFenRiRDabxrE9MNUZ2aPFaFp+DyAe+b4nDwuJaW2LURbr8AEZga7oQj0uYxcYw==
-----END RSA PRIVATE KEY-----
)
```

**Payload**
```json
{
    "userId": 12321,
    "userName": "user1",
    "gender": "male",
    "role": "admin"
}
```
Payload là dữ liệu tự định nghĩa được bao gồm trong một token, nó còn hay được gọi là `claims` trong JWT

**Signature**
Signature là một chữ ký số dùng để xác thực và tin tưởng.

Công thức để tạo ra một signature với thuật toán HMAC và một secret key
```
// signature algorithm
data = base64urlEncode( header ) + “.” + base64urlEncode( payload )
hashedData = hash( data, secret )
signature = base64urlEncode( hashedData )
```

Chúng ta cũng có thể tạo ra một signature dùng public key/private key với thuật toán RSA như ở trên.

**JSON Web Token sẽ nối chúng lại với nhau**
Và cuối cùng một token sẽ là một chuỗi nối giữa header, payload và signature
```
token = base64urlEncode( header ) + “.” + base64urlEncode( payload ) + “.” + signature
```

Chú ý: như đã thấy ở trên header và payload chỉ sử dụng `base64urlEncode` chứ không được mã hoá nhé.

## Product Application Server
Product Application Server chứa các endpoint:
- Get all product
- Get product detail
- Update product
- Delete product

## Packages:

**Libraries to handle network routing**
```
$ go get "github.com/gorilla/mux"
```

**mgo library for handling MongoDB**
```
go get "gopkg.in/mgo.v2"
```