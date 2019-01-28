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
|address|0x572856549d51f68ebcc8f15a2749d65874131a25|gnx address|
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
| 10001   | gnx address 格式错误|
| 10003   | gnx address 已存在|

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

###### 6 修改密码

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


###### 7 查询增加保证金的审核列表

url: 
127.0.0.1:8080/deposit/getdepositlist

post 请求参数：

| 参数 | 实例 |描述|
| --------------- | ------------------- |------------------- |
| token   |  dikcggoeqBdELKIL08I3nS5TrpMcrF3OyPMumM5vsn70JgJBqs      |token|
| page   |  1      |page|
| pageSize   |    100    |pageSize|

返回结果：
State 1 待审核 3 审核通过   5 审核拒绝

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


###### 8 增加保证金接口

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
| 20011   | 保证金为0|


###### 9 审核保证金

url: 
127.0.0.1:8080/deposit/depositreview

post 请求参数：

| 参数 | 实例 |描述|
| --------------- | ------------------- |------------------- |
|depositId|3|需要的审核id
|reason:||审核理由|
|token|3aGFEB1Imer3qL1fra2pWi6vST5zDjLesDFH0iIPy1kYDKTGNx|token|
|password|123456|密码|
|states|3|审核状态（3 审核通过 5 审核拒绝 ）|

返回结果：

````json
{
    "code": 200,
    "data": "审核成功"
}
````

错误码

| code | 描述|
| --------------- | ------------------- |
| 20024   | 参数错误|
| 20026   | 参数错误|
| 20016   | token 无效|
| 20018  | 无权限操作|
| 20020   | 密码错误|
| 20022   | 操作错误|
|20024|审核失败|


###### 10 申请提取保证金

url: 
127.0.0.1:8080/deposit/extractdeposit

post 请求参数：

| 参数 | 实例 |描述|
| --------------- | ------------------- |------------------- |
|token|3aGFEB1Imer3qL1fra2pWi6vST5zDjLesDFH0iIPy1kYDKTGNx|token|
|password|123456|密码|
|value|1|提取金额|

返回结果：

````json
{
    "code": 200,
    "data": "申请提取保证金成功"
}
````

错误码

| code | 描述|
| --------------- | ------------------- |
| 20026   | 获取地址失败|
| 20028   | 密码错误|
| 20030   | 提取金额错误|
| 20032  | 保证金余额不足|
| 20034   | 申请提取保证金失败|


###### 11 审核提取保证金

url: 
127.0.0.1:8080/deposit/extractdepositreview

post 请求参数：

| 参数 | 实例 |描述|
| --------------- | ------------------- |------------------- |
|extractDepositId|3|需要审核提取保证金id
|reason:||审核理由|
|token|3aGFEB1Imer3qL1fra2pWi6vST5zDjLesDFH0iIPy1kYDKTGNx|token|
|password|123456|密码|
|states|3|审核状态（3 审核通过 5 审核拒绝 ）|

返回结果：

````json
{
    "code": 200,
    "data": "审核成功"
}
````

错误码

| code | 描述|
| --------------- | ------------------- |
| 20024   | 参数错误|
| 20026   | 参数错误|
| 20016   | token 无效|
| 20018  | 无权限操作|
| 20020   | 密码错误|
| 20022   | 操作错误|
|20024|审核失败|


###### 12 查询审核提取保证金的审核列表

url: 
127.0.0.1:8080/deposit/getextractdepositlist

post 请求参数：

| 参数 | 实例 |描述|
| --------------- | ------------------- |------------------- |
| token   |  dikcggoeqBdELKIL08I3nS5TrpMcrF3OyPMumM5vsn70JgJBqs      |token|
| page   |  1      |page|
| pageSize   |    100    |pageSize|

返回结果：
State 1 待审核 3 审核通过   5 审核拒绝

````json
{
    "code": 200,
    "data": {
        "extract_deposit_list": [
            {
                "ID": 1,
                "CreatedAt": "2019-01-24T10:59:53+08:00",
                "UpdatedAt": "2019-01-24T11:17:15+08:00",
                "DeletedAt": null,
                "State": 3,
                "Email": "2581913653@qq.com",
                "Reason": "okxxx",
                "Value": 5,
                "UpdateUser": 2
            },
            {
                "ID": 2,
                "CreatedAt": "2019-01-24T11:00:18+08:00",
                "UpdatedAt": "2019-01-24T11:18:32+08:00",
                "DeletedAt": null,
                "State": 5,
                "Email": "2581913653@qq.com",
                "Reason": "okxxx",
                "Value": 5,
                "UpdateUser": 2
            }
        ],
        "page": 1,
        "pageSize": 2,
        "total": 6
    }
}
````

错误码

| code | 描述|
| --------------- | ------------------- |
| 20014   | token 无效|



###### 13 申请借币挖矿

url: 

127.0.0.1:8080/mining/loanmining

post 请求参数：

| 参数 | 实例 |描述|
| --------------- | ------------------- |------------------- |
|token|3aGFEB1Imer3qL1fra2pWi6vST5zDjLesDFH0iIPy1kYDKTGNx|token|
|password|123456|密码|
|value|1|抵押金额|

返回结果：

````json
{
    "code": 200,
    "data": "申请借币挖矿成功"
}
````

错误码

| code | 描述|
| --------------- | ------------------- |
| 30000   | 获取地址失败|
| 30002   | 密码错误|
| 30004   | 申请挖币金额错误|
| 30006  | 借贷金额不足50万，无法挖矿|
| 30010   | 申请借币挖矿失败|


###### 14 审核借币挖矿

url: 

127.0.0.1:8080/mining/loanminingreview

post 请求参数：

| 参数 | 实例 |描述|
| --------------- | ------------------- |------------------- |
|loanMiningId|3|需要审核的id
|reason:||审核理由|
|token|3aGFEB1Imer3qL1fra2pWi6vST5zDjLesDFH0iIPy1kYDKTGNx|token|
|password|123456|密码|
|states|3|审核状态（3 审核通过 5 审核拒绝 ）|
|address|0x572856549d51f68ebcc8f15a2749d65874131a25|挖矿地址（第一次审核的时候填写，审核前先判断是否绑定挖矿地址）|

返回结果：

````json
{
    "code": 200,
    "data": "审核成功"
}
````

错误码

| code | 描述|
| --------------- | ------------------- |
| 30024   | 参数错误|
| 30026   | 参数错误|
| 30016   | token 无效|
| 30018  | 无权限操作|
| 30020   | 密码错误|
| 30022   | 操作错误|
|30024|审核失败|


###### 15 判断是否绑定挖矿地址

url: 

127.0.0.1:8080/mining/isbindingminingaddress

post 请求参数：

| 参数 | 实例 |描述|
| --------------- | ------------------- |------------------- |
|loanMiningId|3|需要审核的id
|token|3aGFEB1Imer3qL1fra2pWi6vST5zDjLesDFH0iIPy1kYDKTGNx|token|

返回结果：

````json
{
    "code": 200,
    "data": true
}
````

错误码

| code | 描述|
| --------------- | ------------------- |
| 30032   | 参数错误|
| 30028   | token 无效|
| 30030  | 无权限操作|



###### 16 查询申请借币挖矿的审核列表

url: 

127.0.0.1:8080/mining/getloanmininglist

post 请求参数：

| 参数 | 实例 |描述|
| --------------- | ------------------- |------------------- |
| token   |  dikcggoeqBdELKIL08I3nS5TrpMcrF3OyPMumM5vsn70JgJBqs      |token|
| page   |  1      |page|
| pageSize   |    100    |pageSize|

返回结果：
State 1 待审核 3 审核通过   5 审核拒绝

````json
{
    "code": 200,
    "data": {
        "loan_mining_list": [
            {
                "ID": 3,
                "CreatedAt": "2019-01-24T18:05:06+08:00",
                "UpdatedAt": "2019-01-24T18:05:19+08:00",
                "DeletedAt": null,
                "State": 3,
                "Email": "2581913653@qq.com",
                "Loan": 600000,
                "Reason": "okxxx",
                "Deposit": 6,
                "UpdateUser": 2
            },
            {
                "ID": 4,
                "CreatedAt": "2019-01-24T18:32:25+08:00",
                "UpdatedAt": "2019-01-24T18:58:59+08:00",
                "DeletedAt": null,
                "State": 5,
                "Email": "2581913653@qq.com",
                "Loan": 100000,
                "Reason": "okxxx122",
                "Deposit": 1,
                "UpdateUser": 2
            }
        ],
        "page": 2,
        "pageSize": 2,
        "total": 9
    }
}
````

错误码

| code | 描述|
| --------------- | ------------------- |
| 30054   | token 无效|



###### 17 申请结束挖矿

url: 

127.0.0.1:8080/mining/extractloanmining

post 请求参数：

| 参数 | 实例 |描述|
| --------------- | ------------------- |------------------- |
|password|123456|密码|
|token|3aGFEB1Imer3qL1fra2pWi6vST5zDjLesDFH0iIPy1kYDKTGNx|token|

返回结果：

````json
{
    "code": 200,
    "data": "申请成功"
}
````

错误码

| code | 描述|
| --------------- | ------------------- |
| 30034   | token 无效|
| 30035   | 密码错误|
| 30036  | 你没有开始挖矿|
| 30038  | 申请失败|



###### 18 查询申请结束挖矿审核列表

url: 

127.0.0.1:8080/mining/getextractloanmininglist

post 请求参数：

| 参数 | 实例 |描述|
| --------------- | ------------------- |------------------- |
| token   |  dikcggoeqBdELKIL08I3nS5TrpMcrF3OyPMumM5vsn70JgJBqs      |token|
| page   |  1      |page|
| pageSize   |    100    |pageSize|

返回结果：
State 1 待审核 3 审核通过   5 审核拒绝

````json
{
    "code": 200,
    "data": {
        "loan_mining_list": [
            {
                "ID": 3,
                "CreatedAt": "2019-01-28T14:19:57+08:00",
                "UpdatedAt": "2019-01-28T15:15:33+08:00",
                "DeletedAt": null,
                "Email": "2581913653@qq.com",
                "Deposit": 1,
                "Loan": 500000,
                "State": 3,
                "UpdateUser": 2,
                "Address": "",
                "Reason": "ok",
                "DepositId": 2
            },
            {
                "ID": 4,
                "CreatedAt": "2019-01-28T14:21:22+08:00",
                "UpdatedAt": "2019-01-28T14:21:22+08:00",
                "DeletedAt": null,
                "Email": "2581913653@qq.com",
                "Deposit": 1,
                "Loan": 500000,
                "State": 1,
                "UpdateUser": 0,
                "Address": "",
                "Reason": "",
                "DepositId": 2
            }
        ],
        "page": 2,
        "pageSize": 2,
        "total": 5
    }
}
````

错误码

| code | 描述|
| --------------- | ------------------- |
| 30056   | token 无效|




###### 19 结束挖矿审核

url: 

127.0.0.1:8080/mining/extractloanminingreview

post 请求参数：

| 参数 | 实例 |描述|
| --------------- | ------------------- |------------------- |
|reviewId|3|需要审核的id
|reason:||审核理由|
|token|3aGFEB1Imer3qL1fra2pWi6vST5zDjLesDFH0iIPy1kYDKTGNx|token|
|password|123456|密码|
|states|3|审核状态（3 审核通过 5 审核拒绝 ）|

返回结果：

````json
{
    "code": 200,
    "data": "审核成功"
}
````

错误码

| code | 描述|
| --------------- | ------------------- |
| 30040   | 参数错误|
| 30042   | 参数错误|
| 30044   | token 无效|
| 30046  | 无权限操作|
| 30048   | 密码错误|
| 30050   | 操作错误|
|30052|审核失败|
