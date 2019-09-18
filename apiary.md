FORMAT: 1A
HOST: https://api.ancient.hatlonely.com/

# ancient

获取古诗文接口

## 所有作品 [/ancient?offset={offset}&limit={limit}]

- Parameters

    - offset: `0` (number) - 页偏移
    - limit: `10` (number) - 最多结果

### 获取所有古诗文 [GET]

- Response 200 (application/json)

    - Body

            [
                {
                    "id": 123,
                    "title": "静夜思",
                    "author": "李白",
                    "dynasty": "唐",
                },
                {
                    "id": 123,
                    "title": "绝句",
                    "author": "杜甫",
                    "dynasty": "唐",
                },
            ]

- Response 204
- Response 400

## 作品 [/ancient/{id}]

- Parameters

    - id: `123` (required) - ancient id

### 获取一篇古诗文 [GET]

- Response 200 (application/json)

    - Body

            {
                "id": 123,
                "title": "静夜思",
                "author": "李白",
                "dynasty": "唐",
                "content": "床前明月光，疑是地上霜。举头望明月，低头思故乡。",
            }

- Response 204
- Response 400

## 作者 [/author?offset={offset}&limit={limit}]

- Parameters

    - offset: `0` (number) - 页偏移
    - limit: `10` (number) - 最多结果

### 获取所有作者 [GET]

- Response 200 (application/json)

    - Body

            [
                "李白",
                "杜甫"
            ]

- Response 204
- Response 400

## 作者作品 [/author/{author}?offset={offset}&limit={limit}]

- Parameters

    - author: `李白` (string) - 作者名字
    - offset: `0` (number) - 页偏移
    - limit: `10` (number) - 最多结果

### 获取一个作者的作品 [GET]

- Response 200 (application/json)

    - Body

            [
                {
                    "id": 123,
                    "title": "静夜思",
                    "author": "李白",
                },
                {
                    "id": 124,
                    "title": "将进酒",
                    "author": "李白",
                }
            ]

- Response 204
- Response 400

## 朝代 [/dynasty?offset={offset}&limit={limit}]

- Parameters

    - offset: `0` (number) - 页偏移
    - limit: `10` (number) - 最多结果

### 获取所有朝代 [GET]

- Response 200 (application/json)

    - Body

            [
                "宋",
                "唐"
            ]

- Response 204
- Response 400

## 朝代作品 [/dyansty/{dyansty}?offset={offset}&limit={limit}]

- Parameters

    - dyansty: `唐` (string) - 朝代
    - offset: `0` (number) - 页偏移
    - limit: `10` (number) - 最多结果

### 获取一个朝代的作品 [GET]

- Response 200 (application/json)

    - Body

            [
                {
                    "id": 123,
                    "title": "静夜思",
                    "author": "李白",
                },
                {
                    "id": 124,
                    "title": "将进酒",
                    "author": "李白",
                }
            ]

