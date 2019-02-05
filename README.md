# BeeGoREST
My Playground for learning Beego API generator from MySQL database

## Installation
1. Open localhost/PHPMyAdmin or any MySQL database management tool that you are using
2. Create a new database "beegorest"
3. Run the command below
``` bash
# clone repo
git clone https://github.com/HendraPB/BeeGoREST.git

# change your directory to project directory
cd BeeGoREST

# install bee CLI
go get -u github.com/beego/bee

# install package manager and register it on your path
go get -u github.com/Masterminds/glide && export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

# install project dependency
glide install

# running migrations
bee migrate -driver=mysql -conn="<dbUser>:<dbPass>@tcp(<dbHost>:<dbPort>)/<dbName>"
# example
bee migrate -driver=mysql -conn="root:@tcp(localhost:3306)/beegorest"

# generate swagger and run project on localhost:8080/swagger/
bee run -downdoc=true -gendoc=true
```

## Beego API Generator Documentation
https://beego.me/blog/beego_api