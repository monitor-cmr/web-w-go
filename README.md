# web-w-go

## Enviroment Requirement

1. `go version go1.17.3 linux/amd64`
   1. Using go module
   2. `github.com/cosmtrek/air` - Auto reload
2. VS Code
3. Linux OS
4. Go-gin framework
   1. [https://golang.org/doc/tutorial/web-service-gin](https://golang.org/doc/tutorial/web-service-gin)

## Day 01

1. Install Go
   1. wget https://golang.org/dl/go1.17.3.linux-amd64.tar.gz
   2. sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.17.3.linux-amd64.tar.gz
   3. export PATH=$PATH:/usr/local/go/bin
   4. source ~/.bash_profile
   5. go version
2. Install `cosmtrek/air`
   1. curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
   2. sudo mv ./bin/air /usr/local/go/bin/
   3. air -v
