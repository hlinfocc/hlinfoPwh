#!/bin/bash

go mod download github.com/astaxie/beego
go get github.com/shiena/ansicolor
go get github.com/astaxie/beego@v1.12.1
go get github.com/astaxie/beego/context@v1.12.1
go get github.com/mattn/go-sqlite3
go get github.com/lib/pq

go mod download github.com/mattn/go-sqlite3
go mod download github.com/lib/pq
