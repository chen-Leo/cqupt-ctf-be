# user_controller  
## 留言板相关接口

### 接口 登陆注册改动
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



