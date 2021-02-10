# Solution of the task on the comment tree (Golang)

## Task:

Insert comments in tree, parallel 

**Input:**
```
{
  "id": 1,
  "replies": [
    {
      "id": 2,
      "replies": []
    },
    {
      "id": 3,
      "replies": [
        {
          "id": 4,
          "replies": []
        },
        {
          "id": 5,
          "replies": []
        }
      ]
    }
  ]
}
```
**Output**
```
{
  "id": 1,
  "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita...",
  "replies": [
    {
      "id": 2,
      "body": "est rerum tempore vitae...",
      "replies": []
    },
    {
      "id": 3,
      "body": "et iusto sed quo iure...",
      "replies": [
        {
          "id": 4,
          "body": "ullam et saepe reiciendis voluptatem adipisci...",
          "replies": []
        },
        {
          "id": 5,
          "body": "repudiandae veniam quaerat sunt sed...",
          "replies": []
        }
      ]
    }
  ]
}
```