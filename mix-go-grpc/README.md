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
```


- Install

~~~
go get -u github.com/mix-go/mixcli
~~~

- New project

~~~
mixcli new hello
~~~

~~~
 Use the arrow keys to navigate: ↓ ↑ → ← 
 ? Select project type:
     CLI
     API
     Web (contains the websocket)
   ▸ gRPC
 ~~~
