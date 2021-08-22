## RabbitMQ消息ACK确认机制

#### 消息ACK机制

1. 生产者produce发送消息到mq时， mq会发送ack给produce告知消息是否成功
   - 发送过程
     - produce将消息发送到broker, 即发送到exchange交换机
     - 消息通过交换机exchange被路由到队列queue (消息只有被正确投递到队列queue中，才算发送成功)
     
    - 失败情况
   
      - producter连接mq失败，消息没有发送到mq
   
        在发送消息时可以通过捕捉AmqpException异常，将消息保存db中后续进行重发处理。
   
      - producter连接mq成功，但是发送到exchange失败
   
         通过实现ConfirmCallback接口，对发送结果进行处理。
   
      - 消息发送到exchange成功，但是路由到queue失败；
   
         通过实现ReturnCallback接口，对回退消息进行重发处理。
   
        

2. 消费者consumer接收消息后，consumer会发送ack给mq告知消息是否处理成功

   - 手动ACK。 处理异常时，消息会储存在MQ的Unacked消息里，不会丢失

     1. 如果一个消费者在处理消息出现了网络不稳定、服务器异常等现象，那么就不会有ACK反馈，RabbitMQ会认为这个消息没有正常消费，会将消息重新放入队列中。
     2. 如果在集群的情况下，RabbitMQ会立即将这个消息推送给这个在线的其他消费者。这种机制保证了在消费者服务端故障的时候，不丢失任何消息和任务。
     3. 消息永远不会从RabbitMQ中删除，只有当消费者正确发送ACK反馈，RabbitMQ确认收到后，消息才会从RabbitMQ服务器的数据中删除。
     
   - 可以设置ACK的重试次数 (达到重试次数后，消息会被丢弃)
   
     ```
     # 开启重试
     spring.rabbitmq.listener.simple.retry.enabled=true
     # 重试次数,默认为3次
     spring.rabbitmq.listener.simple.retry.max-attempts=5
     ```



生产者发送消息时，对应交换器未找到绑定的队列，消息会默认丢失掉，可以使用备份交换器（建议类型为fanout，如果为其他的类型，rountingKey不匹配同样会丢失）进行绑定，这样未被路由的消息会存储到备份交换器绑定的队列上（在声明消息发送交换器时，增加参数alternate-exchange值为备份交换器名来实现）



AMQP 提供的是“**至少一次交付**”（at-least-once delivery），异常情况下，消息会被重复消费，此时业务要实现幂等性（重复消息处理）



#### 参考

https://www.cnblogs.com/biehongli/p/11789098.html

https://juejin.cn/post/6844903906074427400
