### note 1 2023-7-19

- ### block

```json
{
  "type":"geter",
  "index":0,
  "cmd":""
}
```

### block have command types

- goto
- geter
- seter
- event
  - click
  - write
  - keypress

- #### goto json format

  ``` json
    {
      "url":"http://example.com"
    }
  ```

- #### getter json format

  ``` json
  {
    "target":"#element path",
    "to":"set to variable(virtual var)"
  }
  ```

- #### setter json format

  ``` json
    {
      "form":"",
      "target":"#element path",
    }
  ```

- #### click event

  ``` json
    {
      "target":"#element path",
    }
  ```

- #### write event

  ``` json
    {
      "target":"#element path",
      "text":"Hello"
    }
  ```
