FROM golang:1.19-bullseye
RUN go install github.com/cweill/gotests/gotests@v1.6.0 
RUN go install github.com/fatih/gomodifytags@v1.16.0
RUN go install github.com/josharian/impl@v1.1.0
RUN go install github.com/haya14busa/goplay/cmd/goplay@v1.0.0 
RUN go install github.com/go-delve/delve/cmd/dlv@latest 
RUN go install honnef.co/go/tools/cmd/staticcheck@v0.3.3 
RUN go install golang.org/x/tools/gopls@latest