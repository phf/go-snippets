# blah: silly net/http debug server

A silly net/http server that sometimes helps me debug things.

## Summary

All `blah` does is accept http requests on port 8080 and respond
with a `github.com/kr/pretty` rendition of the `http.Request` and
`http.ResponseWriter` objects passed to the handler function. At
times this is useful for figuring out what another program sends
to a web server. You can use `-p` to change the default port.
