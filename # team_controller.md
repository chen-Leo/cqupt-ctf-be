# team_controller  
## 队伍相关接口

### 接口一 :  创建队伍

#### 请求url：
* ”/team/create“

#### 请求方式：
* POST

#### 请求参数列表
```
type CreateTeam struct {
	Name         string `json:"name" binding:"required"`
	Introduction string `json:"introduction" `
}
```

|   参数名    |  类型  |    描述    |是否必须|
| :---------: | :----: | :--------: | :--------: |
|   name   | string |  队伍姓名  |必须 |
|    introduction  | string |  队伍简介 | 必须|



#### 成功返回示例

##### 返回 
```
{
    "data": {
        "application": 1,
        "introduction": "nothing to say",
        "name": "test_team114",
        "score": 0
    },
    "message": "success",
    "status": 10000,
    "time": "2019-06-11T11:13:24.6049473+08:00"
}
```
#### 返回参数说明
|   参数名    |  类型  |    描述    |
| :---------: | :----: | :--------: |
|  name | string |  队伍名  |
|  score   | int| 分数 |
|  application   | int| 是否同意申请 1表示同意，-1不同意|
|introduction|string|简介|

#### 返回异常错误说明
##### 参数错误 
```
{
    "message": "param error",
    "status": 10001,
    "time": "2019-06-10T13:23:07.7096313+08:00"
}
```
##### 队名重复错误 
```
{
    "message": "team name exist",
    "status": 10042,
    "time": "2019-06-10T16:41:58.0257982+08:00"
}
```
##### 已加入过队伍错误 
```
{
    "message": "error,you already join a team, you can not join or create other team",
    "status": 10041,
    "time": "2019-06-10T16:42:09.7050956+08:00"
}
```
### 接口二 ：解散或退出队伍
*  队长解散自己队伍，不是队长退出该队
#### 请求url：
* "/team/exit"

#### 请求方式：
* DELETE

#### 请求参数列表
* 无
* 后端自己获取当前登陆用户的id



#### 成功返回示例
##### 返回 
```
{
    "message": "success",
    "status": 10000,
    "time": "2019-06-10T16:52:26.4677114+08:00"
}
```
#### 返回参数说明
* 无


#### 返回异常错误说明
##### 参数错误 （删除失败)
* 这个应该是不会出现的,出现了应该是数据库的锅
```
{
    "message": "param error",
    "status": 10001,
    "time": "2019-06-10T13:23:07.7096313+08:00"
}
```

##### 无队伍错误 
```
{
    "message": "team name exist",
    "status": 10042,
    "time": "2019-06-10T16:41:58.0257982+08:00"
}
```


### 接口三 :  同意他人加入队伍

#### 请求url：
* ”/team/application/agree“

#### 请求方式：
* POST

#### 请求参数列表
```
type TeamApplication struct {
	NewUserName string `json:"newusername" binding:"required"`
	AgreeOrNot  int    `json:"agreeornot" binding:"required"`
}
```

|   参数名    |  类型  |    描述    |是否必须|
| :---------: | :----: | :--------: | :--------: |
|  newusername | string |  待加入者姓名  |必须 |
|  agreeornot   | int| 是否同意 1表示同意，-1不同意（除了1都不同意） | 必须|



#### 成功返回示例

##### 返回 
 * 同意不同意都成功返回这个
```
{
    "message": "success",
    "status": 10000,
    "time": "2019-06-10T16:40:37.9734017+08:00"
}
```
#### 返回参数说明
* 无

#### 返回异常错误说明
##### 参数错误 
```
{
    "message": "param error",
    "status": 10001,
    "time": "2019-06-10T13:23:07.7096313+08:00"
}
```
##### 不是队长错误 
```
{
    "message": "you are not the team leader ",
    "status": 10043,
    "time": "2019-06-10T21:24:36.7964837+08:00"
}
```

##### 加入队伍申请不存在错误 
```
{
    "message": "the team application do not exist ",
    "status": 10047,
    "time": "2019-06-10T21:30:34.0675919+08:00"
}
```
##### 申请者已加队伍错误 
```
{
    "message": "error,you already join a team, you can not join or create other team",
    "status": 10041,
    "time": "2019-06-11T11:44:18.2110606+08:00"
}
```


### 接口四 :  申请加入某队

#### 请求url：
* ”/team/application/add“

#### 请求方式：
* POST

#### 请求参数列表
```
type TeamName struct {
	TeamName string `json:"teamname" binding:"required"`
}
```

|   参数名    |  类型  |    描述    |是否必须|
| :---------: | :----: | :--------: | :--------: |
|   teamname   | string |  想要申请加入的队伍名字  |必须 |




#### 成功返回示例

##### 返回 
```
{
    "message": "success",
    "status": 10000,
    "time": "2019-06-11T11:31:04.3568873+08:00"
}
```
#### 返回参数说明
* 无

#### 返回异常错误说明
##### 参数错误 
```
{
    "message": "param error",
    "status": 10001,
    "time": "2019-06-10T13:23:07.7096313+08:00"
}
```

##### 队伍未开放申请或队伍不存在（已注销)重复错误 
```
{
    "message": "the team is not open the application or team is not exit",
    "status": 10044,
    "time": "2019-06-11T11:28:11.0506256+08:00"
}
```
##### 重复申请错误 
```
{
    "message": "Application Already before",
    "status": 10046,
    "time": "2019-06-11T11:31:40.9616975+08:00"
}
```

##### 已有队伍错误
```
{
    "message": "error,you already join a team, you can not join or create other team",
    "status": 10041,
    "time": "2019-06-11T11:33:07.388652+08:00"
}
```



### 接口五 :返回当前登陆者队伍信息

#### 请求url：
* ”/team/message“

#### 请求方式：
* GET

#### 请求参数列表
无
后端自己获取登陆者id

#### 成功返回示例

##### 有队伍返回 
{
    "data": {
        "team": {
            "Name": "test111",
            "Score": 0,
            "LeaderName": "test1",
            "Introduction": "nothing to say",
            "Application": 1,
            "LsLeader": 1,
            "ApplicationUsers": [
                "test7"
            ],
            "Members": [
                "test1",
                "test9",
                "test5"
            ]
        }
    },
    "message": "success",
    "status": 10000,
    "time": "2019-06-11T15:50:48.0485103+08:00"
}

##### 无队伍返回 
{
    "data": {
        "team": ""
    },
    "message": "success",
    "status": 10000,
    "time": "2019-06-11T15:54:07.8908638+08:00"
}

#### 返回参数说明
|   参数名    |  类型  |    描述    |
| :---------: | :----: | :--------: |
|   Name   | string |  队伍名  |
|  Score   | int | 分数 |
|   Introduction|     string   |队伍简介 |
|   Application | int |  是否接受申请，1接受，-1不接受|
|   isleader | int|是否队长  1是，-1不是 |
|ApplicationUsers|  []string|申请者姓名          |
|Members|  []string|队员姓名           |



#### 返回异常错误说明
* 无

### 接口六 :根据队伍姓名返回队伍信息

#### 请求url：
* ”/team/message“

#### 请求方式：
* POST

#### 请求参数列表
```
type TeamName struct {
	TeamName string `json:"teamname" binding:"required"`
}
```



#### 请求参数说明
|   参数名    |  类型  |    描述    | 是否必须|
| :---------: | :----: | :--------: |:--------: |
|   teamname   | string |  队伍名  |必须|


#### 成功返回示例

##### 无对应队伍返回 
{
    "data": {
        "team": {
            "Name": "",
            "Score": 0,
            "Introduction": "",
            "Application": 0,
            "Members": null
        }
    },
    "message": "success",
    "status": 10000,
    "time": "2019-06-11T16:54:38.4068193+08:00"
}

##### 有队伍返回 
{
    "data": {
        "team": {
            "Name": "test111",
            "Score": 0,
            "Introduction": "nothing to say",
            "Application": 1,
            "Members": [
                "test9",
                "test5"
            ]
        }
    },
    "message": "success",
    "status": 10000,
    "time": "2019-06-11T16:57:17.8066317+08:00"
}

#### 返回参数说明
|   参数名    |  类型  |    描述    |
| :---------: | :----: | :--------: |
|   Name   | string |  队伍名  |
|  Score   | int | 分数 |
|   Introduction|     string   |队伍简介 |
|   Application | int |  是否接受申请，1接受，-1不接受|
|Members|  []string|队员姓名           |

#### 返回异常错误说明
 ```
  {
    "message": "param error",
    "status": 10001,
    "time": "2019-06-10T13:23:07.7096313+08:00"
  }

 ```


### 接口七 :踢出某人出队伍
#### 请求url：
* ”/team/kick“

#### 请求方式：
* DELETE

#### 请求参数列表
```
type KickName struct {
	PoorName string `json:"poorname" binding:"required"`
}
```



#### 请求参数说明
|   参数名    |  类型  |    描述    | 是否必须|
| :---------: | :----: | :--------: |:--------: |
|   poorname   | string |  被踢的人姓名  |必须|


#### 成功返回示例

##### 成功返回
```
{
    "message": "success",
    "status": 10000,
    "time": "2019-06-11T17:03:17.5752593+08:00"
}
```


#### 返回参数说明
无

#### 返回异常错误说明
##### 参数异常
```
  {
    "message": "param error",
    "status": 10001,
    "time": "2019-06-10T13:23:07.7096313+08:00"
  }

```
###### 踢的用户不存在
```
{
    "message": "username is not exist",
    "status": 10048,
    "time": "2019-06-11T17:08:01.7293018+08:00"
}
```
###### 不是队长
```
{
    "message": "you are not the team leader ",
    "status": 10043,
    "time": "2019-06-11T17:26:43.6452226+08:00"
}
```
##### 不能踢你自己
```
{
    "message": "you can not kick yourself",
    "status": 10049,
    "time": "2019-06-11T17:31:10.950067+08:00"
}

```

###### 不是你队友，你不能踢
```
{
    "message": "not your member",
    "status": 10140,
    "time": "2019-06-11T17:32:03.3956571+08:00"
}
```





### 接口八 :修改队伍是否允许申请

#### 请求url：
* ”/team/application/change“

#### 请求方式：
* PUT

#### 请求参数列表
无
后台自己获取uid



#### 请求参数说明
无

#### 成功返回示例



##### 成功返回 
```
{
    "message": "success",
    "status": 10000,
    "time": "2019-06-11T17:03:17.5752593+08:00"
}

```

#### 返回参数说明
无

#### 返回异常错误说明
##### 参数异常 （在数据库修改application时出现，一般不可能出现)
```
  {
    "message": "param error",
    "status": 10001,
    "time": "2019-06-10T13:23:07.7096313+08:00"
  }

```

###### 不是队长
```
{
    "message": "you are not the team leader ",
    "status": 10043,
    "time": "2019-06-11T17:38:23.9433406+08:00"
}
```



### 接口九 :修改队伍信息
#### 请求url：
* ”/team/message/change“

#### 请求方式：
* put

#### 请求参数列表
```
type NewTeamMessage struct {
	Name         string `json:"name" `
	Introduction string `json:"introduction" `
	Application  int    `json:"application" `
}
```



#### 请求参数说明
|   参数名    |  类型  |    描述    |是否必须|
| :---------: | :----: | :--------: |:--------: |
|   name   | string |  新队伍名  |否|
|   introduction   | string |  简介  |否|
|   application   |int |  是否开发申请 只能填-1，1，否则报错  |否|

#### 成功返回示例

##### 成功返回
```
{
    "data": {
        "team": {
            "Name": "best",
            "Score": 0,
            "Introduction": "ahahahah",
            "Application": -1,
            "Members": [
                "test9",
                "test5",
                "test111111"
            ]
        }
    },
    "message": "success",
    "status": 10000,
    "time": "2019-06-11T18:55:13.4752342+08:00"
}
```


#### 返回参数说明
无

#### 返回异常错误说明
##### 参数异常
```
  {
    "message": "param error",
    "status": 10001,
    "time": "2019-06-10T13:23:07.7096313+08:00"
  }

```

###### 不是队长
```
{
    "message": "you are not the team leader ",
    "status": 10043,
    "time": "2019-06-11T17:26:43.6452226+08:00"
}
```




