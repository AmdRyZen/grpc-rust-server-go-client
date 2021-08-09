<p prefix="center">
<img src="https://picb.zhimg.com/v2-cb1db68b184ed26bc6e2ff0b3108a827_1440w.jpg?source=172ae18b" alt="Rust">
</p>

# Examples

Set of examples that show off the features provided by `tonic`.

## Helloworld

### Client

```bash
$ cargo run --bin helloworld-client
```

### Server

```bash
$ cargo run --bin helloworld-server
```

##如果一切顺利，运行到最后你将看到rust-Client 如下的输出：
```
huzhichaodeMacBook-Pro:tonic-grpc huzhichao$ sudo cargo run --bin helloworld-client
    Finished dev [unoptimized + debuginfo] target(s) in 0.32s
     Running `target/debug/helloworld-client`
RESPONSE=Response { metadata: MetadataMap { headers: {"content-type": "application/grpc", "date": "Mon, 09 Aug 2021 15:58:07 GMT", "grpc-status": "0"} }, message: HelloReply { message: "{\"page_no\":1,\"page_size\":20,\"pages\":646,\"records\":[{\"avatar\":\"uploads/2019-10/e22466b1-dbb3-340c-b858-944b232f3e9c.png\",\"id\":1,\"name\":\"woods\"},{\"avatar\":\"imgMember.png\",\"id\":2,\"name\":\"谦让的七里台加图索\"},{\"avatar\":\"imgMember.png\",\"id\":3,\"name\":\"薯条三兄弟\"},{\"avatar\":\"imgMember.png\",\"i:\"Rafael\"},{\"avatar\":\"imgMember.png\",\"id\":5,\"name\":\"魔方范财神\"},{\"avatar\":\"imgMember.png\",\"id\":6,\"name\":\"朴实的虹梅南路贝隆\"},{\"avatar\":\"imgMember7,\"name\":\"阳光下的你爹\"},{\"avatar\":\"imgMember.png\",\"id\":8,\"name\":\"孤独的壮志路罗本\"},{\"avatar\":\"imgMember.png\",\"id\":9,\"name\":\"清秀的黄渡镇内马尔\"},{19-12/1c0a7d29-f855-3b6e-9353-e52bd4423bec.jpg\",\"id\":10,\"name\":\"1278461\"},{\"avatar\":\"imgMember.png\",\"id\":11,\"name\":\"三不一没有\"},{\"avatar\":\"imgMember.pn"id\":12,\"name\":\"圆月弯刀2009\"},{\"avatar\":\"imgMember.png\",\"id\":13,\"name\":\"妩媚的马栏坡加图索\"},{\"avatar\":\"imgMember.png\",\"id\":14,\"name\":\"忠厚的邓家窑\"imgMember.png\",\"id\":15,\"name\":\"沉着的宜兴阿隆索\"},{\"avatar\":\"imgMember.png\",\"id\":16,\"name\":\"阳光的浦口拉基蒂奇\"},{\"avatar\":\"imgMember.png\",\"id\":17,寄诚庸\"},{\"avatar\":\"imgMember.png\",\"id\":18,\"name\":\"兴奋的天河区戈麦斯\"},{\"avatar\":\"imgMember.png\",\"id\":19,\"name\":\"热心的浦口阿尔巴\"},{\"avatar\":\"imgM0,\"name\":\"清秀的圆明园郝海东\"}],\"search_count\":true,\"total\":12901}" }, extensions: Extensions }
```

## RouteGuide

### Client

```bash
$ cargo run --bin routeguide-client
```

### Server

```bash
$ cargo run --bin routeguide-server
```

## Authentication

### Client

```bash
$ cargo run --bin authentication-client
```

### Server

```bash
$ cargo run --bin authentication-server
```

## Load Balance

### Client

```bash
$ cargo run --bin load-balance-client
```

### Server

```bash
$ cargo run --bin load-balance-server
```

## Dynamic Load Balance

### Client

```bash
$ cargo run --bin dynamic-load-balance-client
```

### Server

```bash
$ cargo run --bin dynamic-load-balance-server
```

## TLS (rustls)

### Client

```bash
$ cargo run --bin tls-client
```

### Server

```bash
$ cargo run --bin tls-server
```

## Health Checking

### Server

```bash
$ cargo run --bin health-server
```

## Server Reflection

### Server
```bash
$ cargo run --bin reflection-server
```

## Tower Middleware

### Server

```bash
$ cargo run --bin tower-server
```

## Autoreloading Server

### Server
```bash
systemfd --no-pid -s http::[::1]:50051 -- cargo watch -x 'run --bin autoreload-server'
```

### Notes:

If you are using the `codegen` feature, then the following dependencies are
**required**:

* [bytes](https://crates.io/crates/bytes)
* [prost](https://crates.io/crates/prost)
* [prost-derive](https://crates.io/crates/prost-derive)

The autoload example requires the following crates installed globally:

* [systemfd](https://crates.io/crates/systemfd)
* [cargo-watch](https://crates.io/crates/cargo-watch)
