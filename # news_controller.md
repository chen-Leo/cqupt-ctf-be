# news_controller  
## 留言板相关接口

### 接口一
#### 简要描述：获取所有的留言

#### 请求url：
* ”/team/create“

#### 请求方式：
* post

#### 请求参数列表
* page 页数
* page 为空  或 小于0 返回 第一页
* page 大于最大页数 返回最后一页



#### 返回示例
** 注：数据库中有一加权时间，置顶时更改置顶时间，返回置顶时间最新的，然后返回没有置顶的（按时间最新的排前面)

##### 空值返回 
```
{
    "data": [],
    "message": "success",
    "status": 10000,
    "time": "2019-06-10T13:16:48.911516+08:00"
}
```
##### 非空返回 
```
{
    "data": [
        {
            "content": "新闻2",
            "currentPage ": 1,
            "number": 1,
            "time": "2019-06-04 15:13:40",
            "title": "2",
            "totalPage": 2
        },
        {
            "content": "新闻7",
            "currentPage ": 1,
            "number": 2,
            "time": "2019-06-04 15:13:40",
            "title": "7",
            "totalPage": 2
        },
        {
            "content": "新闻1",
            "currentPage ": 1,
            "number": 3,
            "time": "2019-06-05 15:13:34",
            "title": "1",
            "totalPage": 2
        },
        {
            "content": "新闻6",
            "currentPage ": 1,
            "number": 4,
            "time": "2019-06-05 15:13:34",
            "title": "6",
            "totalPage": 2
        },
        {
            "content": "新闻5",
            "currentPage ": 1,
            "number": 5,
            "time": "2019-06-13 15:14:45",
            "title": "5",
            "totalPage": 2
        }
    ],
    "message": "success",
    "status": 10000,
    "time": "2019-06-10T12:43:09.951439+08:00"
}
```
##### 超出最大页数例子
* 请求 "http://localhost:8888/news/get?page=414141"
{
    "data": [
        {
            "content": "新闻8",
            "currentPage ": 3,
            "number": 11,
            "time": "2019-06-02 15:13:44",
            "title": "8",
            "totalPage": 2
        }
    ],
    "message": "success",
    "status": 10000,
    "time": "2019-06-10T13:13:09.2329212+08:00"
}


#### 返回参数说明
|  参数名  |  类型  |      描述      |
| :------: | :----: | :------------: |
|  title  |  string  | 公告题目 |
| Content  | string |  公告内容  |
|  time   | string |    时间    |
|  number  |int |  第几个公告   |
|currentPage|int|目前页数|
| totalPage | int |   总页数   |

#### 返回异常错误说明
##### page参数错误 
```
{
    "message": "param error",
    "status": 10001,
    "time": "2019-06-10T13:23:07.7096313+08:00"
}
```
