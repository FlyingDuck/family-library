package endpoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
	"strconv"
)

type BookDO struct {
	Id         int64;
	Title      string;
	SubTitle   string `sql:"sub_title"`;
	Author     string;
	Press      string;
	Pages      int;
	Desc       string;
	KeyWords   string `sql:"key_words"`;
	CreateTime int64  `sql:"create_time"`;
}

func BookListEndpoint(context *gin.Context) {
	db, err := sql.Open("mysql", "lib:familylib@tcp(127.0.0.1:3306)/family_library")
	if nil != err {
		log.Fatal(err)
	}
	err = db.Ping()
	if nil != err {
		log.Fatal(err)
	}
	defer db.Close()

	//var title string;
	//db.QueryRow("SELECT * FROM book LIMIT 1").Scan(&title)

	rows, err := db.Query("SELECT * FROM book")
	if err != nil {
		log.Fatal(err)
	}

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var books []BookDO = make([]BookDO, 0)
	//index := 0
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal(err)
		}

		var book BookDO
		var value string
		for i, colValue := range values {
			if colValue == nil {
				value = "NULL"
			} else {
				value = string(colValue)
				switch columns[i] {
				case "id":
					book.Id, err = strconv.ParseInt(value, 10, 64)
					break
				case "title":
					book.Title = value
					break
				case "sub_title":
					book.SubTitle = value
					break
				case "press":
					book.Press = value
					break
				case "pages":
					book.Pages, err = strconv.Atoi(value)
					break
				case "desc":
					book.Desc = value
					break
				case "author":
					book.Author = value
					break
				case "key_words":
					book.KeyWords = value
					break
				case "create_time":
					book.CreateTime, err = strconv.ParseInt(value, 10, 64)
					break
				}
			}
			//fmt.Println(columns[i], ":", value)
		}
		books = append(books, book)
		fmt.Println(book)
		fmt.Println("-----------------------------------")
		//index++
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	context.HTML(http.StatusOK, "book-list.html", gin.H{"books": books})
}

func BookDetailEndpoint(context *gin.Context) {
	context.HTML(http.StatusOK, "book-detail.html", gin.H{"title": "重构", "subTitle": "改善既有代码的设计"})
}
