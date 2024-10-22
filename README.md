# mini-im
简单的go写的im服务，流程简单清晰，采用golang编写，分为login、msg、msg-push、online等服务。可以自己扩充其他协议和服务。

## 流程图
![image](https://github.com/user-attachments/assets/e2c9d082-3f4b-4cd5-baf6-56597ef6e684)

### 流程图MMD文档
```
sequenceDiagram
    participant A as 用户A
    participant B as 用户B
    participant ServerLogin as Login服务
    participant ServerMsg as Msg服务
    participant ServerOnline as Online服务
    participant kafka
    participant ServerMsgPush as MsgPush服务
    participant consul

    A->>ServerLogin: 1、发送消息
    ServerLogin->>ServerMsg: 2、rpc调用创建消息
    ServerMsg->>ServerMsg: 3、查看是否有a和b会话没有就创建；创建a和b的session；存放消息
    ServerMsg-->>ServerLogin: 4、发送消息成功ack
    ServerLogin-->>A: 5、消息发送成功ack
    ServerMsg->>kafka: 6、消息进入kafka
    kafka->>ServerMsgPush: 7、收到消息，准备推送给用户b
    ServerMsgPush->>ServerOnline: 8、rpc查询b是否在线
    ServerOnline-->>ServerMsgPush: 9、返回用户在线状态
    alt 用户在线
        ServerMsgPush->>ServerMsgPush: 内存查看是否有链接到login的rpc连接，没有就连接（consul获取地址）
        alt 内存没有rpc login服务的连接
            ServerMsgPush->>consul: 获取B所在的login服务地址
            consul-->>ServerMsgPush: 返回地址
            ServerMsgPush->>ServerMsgPush: 连接和存放到内存map中
        end
        ServerMsgPush->>ServerLogin: 发送msgsync通知消息
        ServerLogin-->>B: 发送msgsync通知消息
    else 用户不在线
        ServerMsgPush->>ServerMsgPush: 进行离线推送
    end

```
