##### genaro xpool api document

code 200 表示成功

code 401 token无效/未登录

###### 1 根据邮箱获取验证码

url: 
127.0.0.1:8080/user/getverificationcode

post 请求参数：

| 参数 | 实例 |描述|
| --------------- | ------------------- |------------------- |
| email   |   2581913653@qq.com      |邮箱|

返回结果：

````json
{
    "code": 200,
    "data": "邮件发送成功"
}
````

错误码

| code | 描述|
| --------------- | ------------------- |
| 10008   | email 格式错误|
| 10010   | 邮件发送失败|
| 10012   | 邮件发送失败|


###### 2 注册帐号

url: 
127.0.0.1:8080/user/createuser

post 请求参数：

| 参数 | 实例 |描述|
| --------------- | ------------------- |------------------- |
| email   |   2581913653@qq.com      |邮箱|
|password|123456|密码|
|code|WTPcQ|验证码|
返回结果：

````json
{
    "code": 200,
    "data": "注册成功"
}
````

错误码

| code | 描述|
| --------------- | ------------------- |
| 10000   | email 格式错误|
| 10002   | email 已存在|
| 10004   | 验证码错误或已过期|
| 10006   | password 长度应大于5位|
| 10004   | 验证码错误或已过期|
| 10004   | 验证码错误或已过期|


###### 3 帐号登录

url: 
127.0.0.1:8080/user/login

post 请求参数：

| 参数 | 实例 |描述|
| --------------- | ------------------- |------------------- |
| email   |   2581913653@qq.com      |邮箱|
|password|123456|密码|

返回结果：

````json
{
    "code": 200,
    "data": {
        "ID": 1,
        "CreatedAt": "2019-01-22T10:48:52.055027436+08:00",
        "UpdatedAt": "2019-01-22T10:48:52.055027436+08:00",
        "DeletedAt": null,
        "email": "2581913653@qq.com",
        "TokenRes": "dikcggoeqBdELKIL08I3nS5TrpMcrF3OyPMumM5vsn70JgJBqs",
        "timestamp": 1548125332
    }
}
````

错误码

| code | 描述|
| --------------- | ------------------- |
| 10014   | 登录失败|
| 10016   | 登录失败|



###### 4 根据token获取用户信息

url: 
127.0.0.1:8080/user/getuserbytoken

post 请求参数：

| 参数 | 实例 |描述|
| --------------- | ------------------- |------------------- |
| token   |  dikcggoeqBdELKIL08I3nS5TrpMcrF3OyPMumM5vsn70JgJBqs      |token|

返回结果：

````json
{
    "code": 200,
    "data": {
        "email": "2581913653@qq.com"
    }
}
````

错误码

| code | 描述|
| --------------- | ------------------- |
| 10018   | 获取用户信息失败|
| 10020   | 获取用户信息失败|



###### 5 找回密码

url: 
127.0.0.1:8080/user/forgetpassword

post 请求参数：

| 参数 | 实例 |描述|
| --------------- | ------------------- |------------------- |
| email   |   2581913653@qq.com      |邮箱|
|password|123456|密码|
|code|WTPcQ|验证码|

返回结果：

````json
{
    "code": 200,
    "data": "密码找回成功"
}
````

错误码

| code | 描述|
| --------------- | ------------------- |
| 10022   | email 格式错误|
| 10024   | 密码找回失败|
| 10026   | 验证码错误或已过期|
| 10028   | password 长度应大于5位|

###### 5 修改密码

url: 
127.0.0.1:8080/user/resetpassword

post 请求参数：

| 参数 | 实例 |描述|
| --------------- | ------------------- |------------------- |
| token   |  dikcggoeqBdELKIL08I3nS5TrpMcrF3OyPMumM5vsn70JgJBqs      |token|
|password|123456|原始密码|
|newPassword|123456|新密码|

返回结果：

````json
{
    "code": 200,
    "data": "重置成功"
}
````

错误码

| code | 描述|
| --------------- | ------------------- |
| 10030   | newPassword 长度应大于5位|
| 10032   | token 无效|
| 10033   | 原始密码错误|


###### 6 查询增加保证金的审核列表

url: 
127.0.0.1:8080/deposit/getdepositlist

post 请求参数：

| 参数 | 实例 |描述|
| --------------- | ------------------- |------------------- |
| token   |  dikcggoeqBdELKIL08I3nS5TrpMcrF3OyPMumM5vsn70JgJBqs      |token|

返回结果：

````json
{
    "code": 200,
    "data": {
        "depositList": [
            {
                "ID": 1,
                "CreatedAt": "2019-01-22T18:53:56+08:00",
                "UpdatedAt": "2019-01-22T18:53:56+08:00",
                "DeletedAt": null,
                "State": 1,
                "Email": "2581913653@qq.com",
                "Hash": "0x1e50dae433a66f4f8cd2a8e9f572661efb9bcec2737e437cfc3e82b451d73bb0",
                "Reason": "",
                "Value": 9.223
            },
            {
                "ID": 3,
                "CreatedAt": "2019-01-22T18:55:54+08:00",
                "UpdatedAt": "2019-01-22T18:55:54+08:00",
                "DeletedAt": null,
                "State": 1,
                "Email": "2581913653@qq.com",
                "Hash": "0x65717a8adf52eaf949236457feb3d1f1f65ae6b6c9cb73bde780599b22024e35",
                "Reason": "",
                "Value": 1
            }
        ],
        "page": 1,
        "pageSize": 100,
        "total": 2
    }
}
````

错误码

| code | 描述|
| --------------- | ------------------- |
| 20014   | token 无效|


###### 7 增加保证金接口

url: 
127.0.0.1:8080/deposit/adddeposit

post 请求参数：

| 参数 | 实例 |描述|
| --------------- | ------------------- |------------------- |
| token   |  dikcggoeqBdELKIL08I3nS5TrpMcrF3OyPMumM5vsn70JgJBqs      |token|
| hash   |  0x65717a8adf52eaf949236457feb3d1f1f65ae6b6c9cb73bde780599b22024e35      |交易hash|
| password   |  1234560     |password|

返回结果：

````json
{
    "code": 200,
    "data": "增加保证金成功"
}
````

错误码

| code | 描述|
| --------------- | ------------------- |
| 20000   | 获取地址失败|
| 20002   | 密码错误|
| 20010   | hash 错误|
| 20004  | 获取交易失败|
| 20006   | 获取交易失败|
| 20008   | 获取交易失败|
| 20012   | 增加保证金Hash已存在|
| 20010   | 保证金为0|