# compete_controller  
## 比赛相关接口

### 接口一
#### 简要描述：返回所有的比赛列表

#### 请求url：****
* ”/compete/get“

  #### 请求方式：
* GET

  #### 请求参数列表
* 空

#### 返回示例
##### 空值返回 
```
{
    "data": {
        "competes": []
    },
    "message": "success",
    "status": 10000,
    "time": "2019-06-10T00:37:45.5537962+08:00"
}
```

##### 非空返回   结束时间在前面的先返回，然后按创建时间逆序
```
{
    "data": {
        "competes": [
            {
                "Name": "比赛2",
                "Introduction": "这是比赛二",
                "CreateTime": "2019-05-28 00:40:03",
                "Type": "团队赛",
                "EndTime": "2019-06-10 22:07:09"
            },
            {
                "Name": "比赛1",
                "Introduction": "这是比赛一",
                "CreateTime": "2019-06-09 00:39:43",
                "Type": "个人赛",
                "EndTime": "2019-06-11 22:07:05"
            },
            {
                "Name": "比赛3",
                "Introduction": "这是比赛三",
                "CreateTime": "2019-06-04 00:40:19",
                "Type": "个人赛",
                "EndTime": "2019-06-11 22:07:12"
            }
        ]
    },
    "message": "success",
    "status": 10000,
    "time": "2019-06-11T22:08:41.1256874+08:00"
}
```
#### 返回参数说明
|    参数名    |  类型  |   描述   |
| :----------: | :----: | :------: |
|     Name     | string | 比赛名字 |
| Introduction | string |   简介   |
|  CreateTime  | string | 创建时间 |
|  Type | string | 比赛类型 |
|EndTime|string|结束时间|
#### 返回异常错误说明
 * 无