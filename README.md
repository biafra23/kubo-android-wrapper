# Kubo wrapper to run Kubo on Android

## Build the wrapper library

```bash
go mod tidy 
go get golang.org/x/mobile/bind 
go get google.golang.org/genproto@latest
gomobile bind -ldflags "-checklinkname=0" -target=android -o wrapper.aar ./wrapper
```

