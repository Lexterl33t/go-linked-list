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
    <td>nil</td>
    <td>Init new linked list</td>
  </tr>
  <tr>
    <td>List * Append()</td>
    <td>Value : any</td>
    <td>nil</td>
    <td>insert value at the beginning</td>
  </tr>
  <tr>
    <td>List * Show()</td>
    <td>nil</td>
    <td>nil</td>
    <td>Show formatted linked list</td>
  </tr>
  <tr>
    <td>List * Get()</td>
    <td>index : int</td>
    <td>nil</td>
    <td>Find chunk by index</td>
  </tr>
  <tr>
    <td>List * Pop()</td>
    <td>nil</td>
    <td>nil</td>
    <td>Remove last value of linked list</td>
  </tr>
  <tr>
    <td>List * Delete()</td>
    <td>n : int</td>
    <td>nil</td>
    <td>Remove value of index</td>
  </tr>
</table>

## Init new list

```go
list := linked.NewList(1337)
```
