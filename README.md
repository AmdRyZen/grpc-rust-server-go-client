# grpc-rust-server-go-client
rust grpc 服务端 &amp;&amp; 客户端    go客户端

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


## Mix gRPC Skeleton

- Start
~~~
go build -o bin/go_build_main_go main.go
bin/go_build_main_go grpc:server
bin/go_build_main_go grpc:client
~~~

##如果一切顺利，运行到最后你将看到rust-Client 如下的输出：
```
huzhichaodeMacBook-Pro:mix-go-grpc huzhichao$ bin/go_build_main_go grpc:client
-------------------------------
READY
&{0xc0003fd180}
woods
谦让的七里台加图索
薯条三兄弟
Rafael
魔方范财神
朴实的虹梅南路贝隆
阳光下的你爹
孤独的壮志路罗本
清秀的黄渡镇内马尔
1278461
三不一没有
圆月弯刀2009
妩媚的马栏坡加图索
忠厚的邓家窑贡多齐
沉着的宜兴阿隆索
阳光的浦口拉基蒂奇
善良的中关村寄诚庸
兴奋的天河区戈麦斯
热心的浦口阿尔巴
清秀的圆明园郝海东
--------------------------------
{12901 646 1 20 0 [{1 woods uploads/2019-10/e22466b1-dbb3-340c-b858-944b232f3e9c.png} {2 谦让的七里台加图索 imgMember.png} {3 薯条三兄弟 imgMember.png} {4 Rafael imgMember.imgMember.png} {6 朴实的虹梅南路贝隆 imgMember.png} {7 阳光下的你爹 imgMember.png} {8 孤独的壮志路罗本 imgMember.png} {9 清秀的黄渡镇内马尔 imgMember.png} {10 1278461 uploa53-e52bd4423bec.jpg} {11 三不一没有 imgMember.png} {12 圆月弯刀2009 imgMember.png} {13 妩媚的马栏坡加图索 imgMember.png} {14 忠厚的邓家窑贡多齐 imgMember.png} {15 沉着的宜兴 阳光的浦口拉基蒂奇 imgMember.png} {17 善良的中关村寄诚庸 imgMember.png} {18 兴奋的天河区戈麦斯 imgMember.png} {19 热心的浦口阿尔巴 imgMember.png} {20 清秀的圆明园郝海东 imgMember.png}]}
