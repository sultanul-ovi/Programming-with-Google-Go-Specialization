# Request for Comments

- definitions of internet protocols and formats
- examples
  - HTML - hypertext markup language, 1866
  - URI - uniform resource identifier, 3986
  - HTTP - hypertext transfer protocal, 2616

## Protocol Packages

- golang provides packages for important RFCs
- functions which encode and decode protocol format
- example:
  - `net/http`
    - web communication protocol
    - `http.Get(www.uci.edu)` - get a web page
  - `net`
    - TCP/IP and socket programming
    - `net.Dial("tcp","uci.edu:80)` - connect to a remote host

## JSON

- javascript object notation
- RFC 7159
- format to represent structured information
- attribute-value pairs
  - struct or map
- basic value types
  - bool, int, string, array, object
  - example golang to json:

```golang
p1 := Person(
  name:"joe",
  addr:"a st.",
  phone: "123"
)
```

```json
{
  "name": "joe",
  "addr": "a st",
  "phone": "123"
}
```
