package main
import (
    "database/sql"
    "log"
    "fmt"
    _ "github.com/lib/pq"
)

func main() {
    db, err := sql.Open("postgers", "postgres://gopher:rdbmsftw@localhost:5432/gopher?sslmode=disable")
    if err != nil {
        log.Fatalf("Failed to make connection: %s\n", err)
    }

    err  = db.Ping()
    if err != nil {
        log.Fatalf("Failed to ping: %s\n", err)
    }

    _, err = db.Exec(
        `
        CREATE TABLE contacts (
            id serial primary key,
             first varchar(100),
             last varchar(100)
        )

        `
    )

    if err != nil {
        log.Fatalf("Failed to create table: %s\n", err)
    }

    type person struct {
        first string
        last string

    }

    for _, p := range []person(
        person{first:"Rob", last: "Pike"},
        person{first:"Ken", last: "thompson"},
        person{first:"Robert", last: "Griesemer"},
    ){
        _, err = db.Exec(
            `
            insert into contacts (first, last)
            values ($1, $2)
            `,
            p.first, p.last
        )

        if err != nil {
            log.Fatalf("Failed to insert record: %s\n", err)
        }
    }

    fmt.Println("done")

}

