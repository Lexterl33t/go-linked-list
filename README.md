# go-linked-list
little library to use linked list

## Installation
```sh
go get -u github.com/Lexterl33t/go-linked-list
```

```go
import "github.com/Lexterl33t/go-linked-list"
```

<table>
  <tr>
    <th>Protoype</th>
    <th>Arg1</th>
    <th>Arg2</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>NewList()</td>
    <td>DefaultValue : Integer</td>
    <td>NULL</td>
    <td>Init new linked list</td>
  </tr>
  <tr>
    <td>List * Append()</td>
    <td>Value : any</td>
    <td>NULL</td>
    <td>insert value at the beginning</td>
  </tr>
</table>

## Init new list

```go
list := linked.NewList(1337)
```
