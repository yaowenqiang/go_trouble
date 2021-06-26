# trouble shooting go application development
> export PATH=$PATH:$GOPATH/bin
> https://github.com/andlabs/ui
> https://github.com/therecipe/qt
> https://go-proverbs.github.io


> nl -ba main.go

+ Don't panic
+ Don't ignore errors

> docker run -d -rm -e POSTGRES_USER=gopher -e POSTGERS_PASSWORD=rdbmsftw --name db -p 5432:5432 postgres:10.3-alphine
> docker exec -it db psql -U gopher -w
> select * from contacts;

> wtch -p -n 1 free -m
> from time import sleep # python
> time sleep 1 # bash

> date +%I:%M:%S

> from datetime import datetime
> print datetime.now().strftime("%I:%M:%S")
> print datetime.now().strftime("%H:%M:%S")
