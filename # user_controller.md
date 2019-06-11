# user_controller  
## 留言板相关接口

### 接口一  登陆注册改动
#### 简要描述：登陆注册改动，成功返回参数改动

#### 请求url：
* 不变

#### 请求方式：
* 不变

#### 请求参数列表
* 不变


#### 返回示例
* 加入过队伍
```
{
    "data": {
        "email": "123",
        "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjQsImV4cCI6MTU2MDE1NTY4OCwiaXNzIjoiZ2luLWJsb2cifQ.ysN0eeivBmYiTB8MkVJ-phPk31Bl5m2i9a8CWwd0lWc",
        "motto": "",
***  "teamname": "team1",
        "username": "test1"
    },
    "message": "success",
    "status": 10000,
    "time": "2019-06-10T13:34:48.4331052+08:00"
```
* 无队伍
```
{
    "data": {
        "email": "123",
        "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjQsImV4cCI6MTU2MDE1NTg1NiwiaXNzIjoiZ2luLWJsb2cifQ.8XurdVVEhMxLpGWJcEzKiE41fnXrjEcLQoVb0qEsDLg",
        "motto": "",
        "teamname": "",
        "username": "test1"
    },
    "message": "success",
    "status": 10000,
    "time": "2019-06-10T13:37:36.0445588+08:00"
}
```
#### 返回参数说明

 * 多加入了 teamname 参数 ，返回所在队伍名

#### 返回异常错误说明
* 不变





### 接口二  密码修改


#### 请求url：
* “/user/password"
#### 请求方式：
* PATCH

#### 请求参数列表

```
type ChangePassword struct {
	OldPassword string `json:"oldpassword" binding:"required"`
	NewPassword string `json:"newpassword" binding:"required"`
}
```

|    参数名    |  类型  |   描述   | 是否必须 |
| :----------: | :----: | :------: | :------: |
|   oldpassword     | string | 旧密码 |   必须   |
| newpassword | string | 新密码 |   必须   |


#### 返回示例
* 成功
```
{
    "message": "success",
    "status": 10000,
    "time": "2019-06-11T18:09:46.8134885+08:00"
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
###### 原密码错误
```
{
    "message": "change password error, old password error",
    "status": 10013,
    "time": "2019-06-11T18:08:37.3813945+08:00"
}
```

### 接口二  通过用户名返回用户信息

#### 请求url：
* “/user/message/get"
#### 请求方式：
* POST

#### 请求参数列表
```
type GetUserMessage struct {
	Username string `json:"username" binding:"required"`
}
````
|    参数名    |  类型  |   描述   | 是否必须 |
| :----------: | :----: | :------: | :------: |
|  username   | string | 用户名 |   必须   |


#### 返回示例
* 成功
```
{
    "data": {
        "email": "123",
        "motto": "",
        "teamname": "test111",
        "username": "test1"
    },
    "message": "success",
    "status": 10000,
    "time": "2019-06-11T18:21:26.2702412+08:00"
}
```
#### 返回参数说明
|    参数名    |  类型  |   描述   |
| :----------: | :----: | :------: |
|  username   | string | 用户名 |
|  motto   | string | 格言 |
|  teamname   | string | 队伍名 |
| email   | string | 邮箱 |

#### 返回异常错误说明

##### 参数错误 
```
{
    "message": "param error",
    "status": 10001,
    "time": "2019-06-10T13:23:07.7096313+08:00"
}
```


### 接口三  修改用户信息

#### 请求url：
* “user/message/change"
#### 请求方式：
* PUT

#### 请求参数列表
```
type GetUserMessage struct {
	Username string `json:"username" binding:"required"`
}
````
|    参数名    |  类型  |   描述   | 是否必须 |
| :----------: | :----: | :------: | :------: |
|  username   | string | 用户名 |   必须   |


#### 返回示例
* 成功
```
{
    "data": {
        "email": "123456789",
        "motto": "",
        "teamname": "test111",
        "username": "test111111"
    },
    "message": "success",
    "status": 10000,
    "time": "2019-06-11T18:31:28.4058906+08:00"
}
```
#### 返回参数说明
* 无

#### 返回异常错误说明
|    参数名    |  类型  |   描述   |
| :----------: | :----: | :------: |
|  username   | string | 用户名 |
|  motto   | string | 格言 |
|  teamname   | string | 队伍名 |
| email   | string | 邮箱 |
##### 参数错误 
```
{
    "message": "param error",
    "status": 10001,
    "time": "2019-06-10T13:23:07.7096313+08:00"
}
```
##### 更改的姓名或邮箱重复错误 
```
{
    "message": "username exist or or email is used",
    "status": 10012,
    "time": "2019-06-11T18:35:54.4201066+08:00"
}
```
##### 原密码错误
```
{
    "message": "change password error, old password error",
    "status": 10013,
    "time": "2019-06-11T18:36:59.53623+08:00"
}
```