# 集合列表

## 用户集合
```js
{
	"account":"账号",
    "password":"密码",
	"nickname":"昵称",
	"sex":1,	// 0-未知，1-男，2-女
	"email":"邮箱",
	"avater":"头像",
	"created_at":1,//创建时间
	"updated_at":1,//更新时间
}
```

## 消息集合
```js
{
    "user_identity":"用户的唯一表示",
	"room_identity":"房间的唯一表示",
	"data":"发送的数据",
	"created_at":1,
	"updated_at":1,
}
```

## 房间集合
```js
{
    "number":"房间号",
	"name":"房间的名称",
	"info":"房间简介",
	"user_identity":"房间创建者的唯一标识",
	"created_at":1,
	"updated_at" :1,
}
```

## 用户房间集合
```js
{
    "user_identity":"用户的唯一标识",
		"room_ideneity":"房间的唯一标识",
		"message_identity":"消息的唯一标识",
		"created_at":1,
		"updated_at":1,
}
```