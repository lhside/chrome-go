# Chrome native messaging host in go

This is a package for communicate with Chrome extension using native messaging api.

## Install

```bash
go install github.com/lhside/chrome-go
```

## Example

Read message from standard input.

```go
msg, err := chrome.Recieve(os.Stdin)
```

Post message to standard output
```go
err := chrome.Post(msg, os.Stdout)
```
