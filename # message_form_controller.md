# message_form_controller  
## 留言板相关接口

### 接口一
#### 简要描述：获取所有的留言

#### 请求url：
* ”/messageform/get“

#### 请求方式：
* GET

#### 请求参数列表
* 空

#### 返回示例
##### 空值返回 
```
{
    "data": {
        "message": null
    },
    "message": "success",
    "status": 10000,
    "time": "2019-06-10T01:13:34.8948758+08:00"
}
```

##### 非空返回 时间逆序（最新的在前面)
```
{
    "data": {
        "message": [
            {
                "Id": 1,
                "Content": "留言1",
                "Username": "test1",
                "Time": "2019-06-11 00:54:29",
                "OthersMessageForm": [
                    {
                        "Id": 5,
                        "Content": "留言1的评论5",
                        "Username": "test7",
                        "Time": "2019-06-10 00:58:30",
                        "OthersMessageForm": [
                            {
                                "Id": 6,
                                "Content": "评论5的评论（回复）6",
                                "Username": "test1",
                                "Time": "2019-06-10 00:59:21",
                                "OthersMessageForm": []
                            }
                        ]
                    }
                ]
            },
            {
                "Id": 4,
                "Content": "留言4",
                "Username": "test5",
                "Time": "2019-06-10 00:57:13",
                "OthersMessageForm": []
            },
            {
                "Id": 3,
                "Content": "留言3",
                "Username": "test7",
                "Time": "2019-06-09 00:56:01",
                "OthersMessageForm": [
                    {
                        "Id": 10,
                        "Content": "留言3的评论7",
                        "Username": "test5",
                        "Time": "2019-06-13 01:02:25",
                        "OthersMessageForm": []
                    },
                    {
                        "Id": 7,
                        "Content": "留言3的评论7",
                        "Username": "test2",
                        "Time": "2019-06-10 01:00:25",
                        "OthersMessageForm": [
                            {
                                "Id": 9,
                                "Content": "评论7的评论（回复）9",
                                "Username": "test2",
                                "Time": "2019-06-10 01:01:34",
                                "OthersMessageForm": []
                            },
                            {
                                "Id": 8,
                                "Content": "评论7的评论（回复）8",
                                "Username": "test9",
                                "Time": "2019-06-10 01:01:05",
                                "OthersMessageForm": []
                            }
                        ]
                    }
                ]
            },
            {
                "Id": 2,
                "Content": "留言2",
                "Username": "test2",
                "Time": "2019-06-07 00:55:00",
                "OthersMessageForm": []
            }
        ]
    },
    "message": "success",
    "status": 10000,
    "time": "2019-06-10T01:12:26.745386+08:00"
}
```
#### 返回示例
#### 返回参数说明
| 参数名 | 类型 | 描述 |
| :----: | :----:  | :----: |
| Id | uid | 留言或评论的id |
| Content |string | 留言评论内容 |
| Username|string |用户名|
| Time |string | 创建时间 |

#### 返回异常错误说明
 * 无



### 接口二
#### 简要描述： 添加留言或者评论

#### 请求url：
* ”/messageform/add“

#### 请求方式：
* POST

#### 请求参数列表
```
  type MessageLeave struct {
  Pid     uint   `json:"pid"`
  Content string `json:"content" binding:"required"`
  }
```

| 参数名 | 类型 | 描述 |是否必须|
| :----: | :----:  | :----: || :----: |
|  pid | uint | 回复或评论的id | 非必须|
| content |string | 留言或评论内容 | 必须|
 *  注 : 后端会根据当前登陆用户获取uid
#### 成功返回示例
```
{
    "message": "success",
    "status": 10000,
    "time": "2019-06-10T02:09:25.982462+08:00"
}
```

#### 错误返回示例

###### 错误一 ： 传入参数错误
```
  {
    "message": "param error",
    "status": 10001,
    "time": "2019-06-10T02:10:57.5013593+08:00"
```
###### 错误一 ：pid 所对应的表在数据库中不存在



```
{
    "message": "leave message error,the message doesn't exist.",
    "status": 10061,
    "time": "2019-06-10T02:12:07.2748182+08:00"
}
```

