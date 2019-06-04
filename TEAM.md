### 1.login  及signup

url：不变

传参不变


###### 信息返回格式改变 加了个teamid

![1559633793168](https:\\github.com\chen-Leo\cqupt-ctf-be\edit\test\typora-user-images\1559633793168.png)



### 2.创建队伍

> urL：/team/create POST
>
> 传参：

```
Name           名称    string `json:"name" binding:"required"`
Introduction   简介  string `json:"introduction" `
```





###### 成功 返回

![1559632977594](typora-user-images\1559632977594.png)



###### 参数错误返回

![1559633020144](typora-user-images\1559633020144.png)

1.已有队伍没退出返回

![1559633129769](typora-user-images\1559633129769.png)





### 3.申请某队伍

> urL：/team/add  POST
>
> 传参：

```
TeamId uint `json:"teamId" binding:"required"` 想加的队伍
```



###### 成功返回

![1559633260816](typora-user-images\1559633260816.png)

###### 参数错误返回

![1559633020144](typora-user-images\1559633020144.png)



###### 1.加入或创建过其他队伍

![1559633839351](typora-user-images\1559633839351.png)

###### 2.申请过该队伍重复申请

![1559633908479](typora-user-images\1559633908479.png)

###### 3.加入的队伍未开放申请

![1559633953769](typora-user-images\1559633953769.png)



### 4.退出队伍

> urL：/team/exite DELETE
>
> 传参：无，根据jwt后端自己获取uid

成功返回

![1559633260816](\typora-user-images\1559633260816.png)

参数错误返回

![1559633020144](typora-user-images\1559633020144.png)

1.没有队伍错误返回

![1559634093776](typora-user-images\1559634093776.png)

### 5.解散队伍

> urL：/team/break DELETE
>
> 传参： 无，根据jwt后端自己获取uid
>
> 返回与4退出队伍相同



### 6.同意申请

urL：/team/agreeadd POST

传参：   申请人用户名          

```
NewUserName string `json:"newusername" binding:"required"`   申请人用户名       
AgreeOrNot    int `json:"agreeornot" binding:"required"` 是否同意  1-表示同意 （-1不同意，其实除了1都不同意） 
```

返回



###### 成功返回

![1559633260816](typora-user-images\1559633260816.png)

###### 参数错误返回

![1559633020144](typora-user-images\1559633020144.png)



###### 1.不是加入的本队的队伍（防止恶意构造表单)

![1559634892633](typora-user-images\1559634892633.png)

###### 2.不是队长

![1559634967785](typora-user-images\1559634967785.png)



### 7.踢人

> urL：/team/kickpeople DELETE
>
> 传参：  踢掉的人 id uint `json:"pooruid" binding:"required"`
>
> 返回

###### 成功返回

![1559633260816](typora-user-images\1559633260816.png)

###### 参数错误返回

![1559633020144](typora-user-images\1559633020144.png)

###### 1.不是队长

![1559634967785](typora-user-images\1559634967785.png)

### 8.获取登陆者队伍信息

> urL：/team/getmessage GET
>
> 传参：无 后端自动获得
>
> 返回

成功返回

![1559635351724](typora-user-images\1559635351724.png)

```
Name:             team.Name,
Score:            team.Score,        
LeaderName:       leaderName,        //队长
Introduction:     team.Introduction, //队伍简介 
Application:      team.Application,  //是否接受申请 1->接受，-1->不接受 取
LsLeader:         isLeader,          //是否是队长，1->是，-1->不是
ApplicationUsers: applicationUsers,  //申请人名字列表
```

无数据返回

”team“：”“

```
response.OkWithData(c, gin.H{"team": teamMessageAll})
```

### 9.队伍信息修改

> urL：/team/changemessage POST
>
> 传参： 

```
 `Name             string   json:"name"    //姓名`
`Introduction     string   json:"introduction" //简介`
`Application        int      `json:"application" //是否同意申请``
```

返回

###### 成功返回

![1559633260816](\typora-user-images\1559633260816.png)

###### 参数错误返回

![1559633020144](C:\Users\hello\AppData\Roaming\Typora\typora-user-images\1559633020144.png)

###### 1.不是队长

![1559634967785](typora-user-images\1559634967785.png)
