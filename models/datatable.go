package models

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
)

type DtListResponse struct {
	Draw            int64       `json:"draw"`
	RecordsTotal    int64       `json:"recordsTotal"`
	RecordsFiltered int64       `json:"recordsFiltered"`
	Data            interface{} `json:"data"`
	Eer             string      `json:"error"`
}

type DtListOrder struct {
	Column int64
	Dir    string
}

type DtListRequest struct {
	Start  int
	Length int
	Order  DtListOrder
}

type DtList struct {
	Data       string
	Name       string
	Searchable bool
	Orderable  bool
	Search     struct {
		Value string
		Regex string
	}
	Index string
}

type DtListWrapper struct {
	Array []DtList
}

func (dtlist *DtListWrapper) GenericQuery(selectedColumn []string, where []string, dtlr *DtListRequest, tableName string) (string, string) {
	var query string
	var filteredCount string
	query = query + `SELECT ` + strings.Join(selectedColumn, ", ") + ` FROM "` + tableName + `" AS "table1"`
	filteredCount = filteredCount + `SELECT COUNT(*) as filtered FROM "` + tableName + `"`
	if len(where) > 0 {
		query = query + ` WHERE ` + strings.Join(where, " AND ")
		filteredCount = filteredCount + ` WHERE ` + strings.Join(where, " AND ")
	}
	query = query + `GROUP BY "table1"."id" ORDER BY ` + dtlist.Array[dtlr.Order.Column].Data + ` ` + strings.ToUpper(dtlr.Order.Dir) + ` LIMIT ? OFFSET ? `

	return query, filteredCount
}

func (dtlist *DtListWrapper) CheckIfFieldInStruct(v reflect.Value, dtlistData string) (bool, string) {
	exist := false
	notExistColumn := ""
	for i := 0; i < v.NumField(); i++ {
		fmt.Println((v.Type().Field(i).Tag.Get("dt")))
		if strings.ToLower(dtlistData) == (v.Type().Field(i).Tag.Get("dt")) {
			exist = true
		} else {
			notExistColumn = dtlistData
		}
	}
	return exist, notExistColumn
}

func (dtlist *DtListWrapper) IterateValues(v reflect.Value, dtlr *DtListRequest) ([]interface{}, []string, []interface{}, []string, error) {
	selectedColumn := []string{}
	var values []interface{}
	var where []string
	var whereValues []interface{}
	for i := 0; i < len(dtlist.Array); i++ {
		exist, notExistColumn := dtlist.CheckIfFieldInStruct(v, dtlist.Array[i].Data)
		if !exist {
			return nil, nil, nil, nil, errors.New("Requested Column doesnt exists " + notExistColumn)
		}
		selectedColumn = append(selectedColumn, `"`+dtlist.Array[i].Data+`"`)

		if dtlist.Array[i].Search.Value != "" {
			where = append(where, fmt.Sprintf(`"%s" LIKE ?`, dtlist.Array[i].Data))
			whereValues = append(whereValues, "%"+dtlist.Array[i].Search.Value+"%")
		}

	}
	where = append(where, ` show = true `)
	values = append(values, whereValues...)
	values = append(values, dtlr.Length)
	values = append(values, dtlr.Start)
	return values, where, whereValues, selectedColumn, nil
}

func (dtlist *DtListWrapper) Create(r *http.Request) (DtListRequest, error) {
	size, err := strconv.Atoi(chi.URLParam(r, "total"))
	if err != nil {
		return DtListRequest{}, nil
	}
	var dtlr DtListRequest

	start, _ := strconv.Atoi(r.URL.Query()["start"][0])
	length, _ := strconv.Atoi(r.URL.Query()["length"][0])
	for i := 0; i < size; i++ {
		columnString := "columns[" + strconv.Itoa(i) + "]"
		searchable := r.URL.Query()[columnString+".searchable"][0]
		orderable := r.URL.Query()[columnString+".orderable"][0]
		fmt.Println(searchable)
		fmt.Println(orderable)
		dt := DtList{
			Data:       r.URL.Query()[columnString+".data"][0],
			Name:       r.URL.Query()[columnString+".name"][0],
			Searchable: false,
			Orderable:  false,
			Search: struct {
				Value string
				Regex string
			}{
				Value: r.URL.Query()[columnString+".search.value"][0],
				Regex: r.URL.Query()[columnString+".search.regex"][0],
			},
			Index: r.URL.Query()[columnString+".index"][0],
		}
		fmt.Println(r.URL.Query()[columnString+".data"][0])
		dtlist.Array = append(dtlist.Array, dt)

	}
	columnInt, cerr := strconv.Atoi(r.URL.Query()["order[0].column"][0])
	if cerr != nil {
		return DtListRequest{}, nil

	}
	order := DtListOrder{
		Column: int64(columnInt),
		Dir:    r.URL.Query()["order[0].dir"][0],
	}

	dtlr = DtListRequest{
		Start:  start,
		Length: length,
		Order:  order,
	}

	return dtlr, nil

}
