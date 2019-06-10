# team_controller  
## 队伍相关接口

### 接口一 :  创建队伍

#### 请求url：
* ”/news/get“

#### 请求方式：
* GET

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
* "/team/exite"

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
* ”/team/agreeadd“

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
