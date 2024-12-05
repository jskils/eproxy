# EProxy

---

## 描述

EProxy ：可利用公有云弹性实例+弹性IP，实现一个私有代理IP池。



## 调用示例

接口调用应在query参数或自定义header携带password

```
curl --location 'localhost:4141/oapi/run' \
--header 'password: eproxy' \
--header 'Content-Type: application/json' \
--data '{
    "amount":1
}'
```

## 接口列表

1. `/oapi/run`

   - 用于创建指定数量的可更换IP的实例，该接口需等待资源创建，比较耗时，建议调用超时时间大于30s，并发数建议小于50，即同时在运行的实例不大于50。

   - 类型：POST

   - 参数：amount(uint)：数量

     ```json
     {
         "amount": 1
     }
     ```

   - 响应：

     ```json
     {
         "data": {
             "instanceId": "i-6wecb11ffkvrtmdwww11",
             "allocationId": "v-6wecb11ffkvrtmdwww11",
             "eip": "127.0.0.1",
             "proxy": "127.0.0.1:7890"
         }
     }
     ```

2. `/oapi/change`

   - 用于为正在运行中的实例更换IP，更换IP最小间隔为1分钟。

   - 类型：POST

   - 参数：instanceId(string)：实例ID，allocationId(string)：弹性IP的ID

     ```
     {
          "instanceId": "i-6wecb11ffkvrtmdwwwz",
          "allocationId": "eip-6we82jwbd4afcvhb38nv"
     }
     ```

     

   - 响应：

     ```
     {
         "data": {
             "instanceId": "i-6wecb11ffkvrtmdwww11",
             "allocationId": "v-6wecb11ffkvrtmdwww11",
             "eip": "127.0.0.1",
             "proxy": "127.0.0.1:7890"
         }
     }
     ```

3. `/oapi/release`

   - 用于释放实例及当前绑定的IP

   - 类型：POST

   - 参数：instanceId(string)：实例ID，allocationId(string)：弹性IP的ID

   - 响应：

     ```
     {
         "data": {
             "instanceId": "i-6wecb11ffkvrtmdwww11",
             "allocationId": "v-6wecb11ffkvrtmdwww11"
         }
     }
     ```

## 配置

```
# 接口密码
PASSWORD=eproxy
# 服务端口
PORT=4141
# 代理端口
PROXY_PORT=4141
# 阿里云KEY_ID
ALIBABA_CLOUD_ACCESS_KEY_ID=
# 阿里云KEY_SECRET
ALIBABA_CLOUD_ACCESS_KEY_SECRET=
# 区域ID
REGION_ID=ap-northeast-1
# 阿里云ECS端点
ECS_ENDPOINT=ecs.ap-northeast-1.aliyuncs.com
# 阿里云启动模板名称
ECS_TEMPLATE_NAME=EPROXY
# 阿里云VPC端点
VPC_ENDPOINT=vpc.ap-northeast-1.aliyuncs.com
```