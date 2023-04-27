# API LIST

#### 新增
```
[POST] /insert
必填 name price weight (status 0, 1 若選 1 remark 必填)
參數 {"name":"bbbb","price":100,"weight":20,"tiket":0,"fare":0,"people":1,"status":0,"profit":10, "note":"note", "remark": "remark"}
```

#### Response
```bash
{
    "result": "ok"
}

{
     "error_code": 100001
}
```

#### 查詢
```
[GET] /quote?name=&offset=1
參數 name offset
```

#### Response
```bash
{
    "data": {
        "comput": [
            {
                "cost": 506,
                "id": 3,
                "name": "bbbb",
                "profit": 46,
                "quote": 552
            },
            {
                "cost": 536,
                "id": 12,
                "name": "bbbb",
                "profit": 62,
                "quote": 598
            }
        ],
        "row": [
            {
                "id": 3,
                "name": "bbbb",
                "price": 2200,
                "weight": 20,
                "freight": 20,
                "exchange_rate": 0.23,
                "profit": 20,
                "created_at": "2023-04-26T02:21:12Z",
                "updated_at": "2023-04-26T02:25:54Z"
            },
            {
                "id": 12,
                "name": "bbbb",
                "price": 2200,
                "weight": 30,
                "freight": 30,
                "people": 1,
                "exchange_rate": 0.23,
                "profit": 30,
                "status": 1,
                "created_at": "2023-04-27T06:30:01Z",
                "updated_at": "2023-04-27T06:34:23Z"
            }
        ]
    }
}


{
     "error_code": 100001
}
```

#### 刪除
```
[DELETE] /:id
參數 id
```

#### Response
```bash
{
    "result": "ok"
}

{
     "error_code": 100001
}
```

#### 更新
```
[PUT] /update
必填 id (status 0, 1 若選 1 remark 必填)
參數 {"id": 1,"name":"bbbb","price":100,"weight":20,"tiket":0,"fare":0,"people":1,"status":0, "exchange_rate": 0.23,"profit":10, "note":"note", "remark": "remark"}
```

#### Response
```bash
{
    "result": "ok"
}

{
     "error_code": 100001
}
```