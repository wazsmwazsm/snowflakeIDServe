package main

import (
	"fmt"
	"log"
	"strconv"
	"bytes"
	"net/http"
	"snowflake"
)

func main() {

	node, err := snowflake.NewNode(1)

	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/snowflakeID", func (w http.ResponseWriter, r *http.Request) {
		id := node.Generate().Int64()

		w.Header().Set("Content-Type", "applicasettion/json")

		b := bytes.Buffer{}
		b.WriteString("{\"id\":")
		b.WriteString(strconv.FormatInt(id, 10))
		b.WriteString("}")

		fmt.Fprint(w, b.String()) 
	})

	err = http.ListenAndServe(":2333", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

