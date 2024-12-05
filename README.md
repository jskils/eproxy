# EProxy

---

## 描述

EProxy ：可利用公有云弹性实例+弹性IP，实现一个私有代理IP池。



## 接口列表



---

## 接口列表

1. `/oapi/run`

   - 用于创建指定数量的可更换IP的实例，该接口需等待资源创建，比较耗时，建议调用超时时间大于30s

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

