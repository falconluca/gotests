package hbase_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
	"log"
	"strconv"
	"testing"
)

// create 'bi:datasource', '100', 'meta_info', 'biz_info'

func TestHbaseClient(t *testing.T) {
	tt := assert.New(t)

	client := gohbase.NewClient("localhost")
	defer client.Close()

	tt.NotNil(client)
}

func TestInsertCellAndGetEntireRow(t *testing.T) {
	tt := assert.New(t)

	values := map[string]map[string][]byte{
		"meta_info": {
			"name":  []byte("Luca"),
			"ds_id": []byte("98123123432"),
			"state": []byte(strconv.Itoa(1)),
			"code":  []byte(strconv.Itoa(9897)),
		},
		"biz_info": {
			"platform":  []byte("data works"),
			"biz_code":  []byte("1998"),
			"biz_state": []byte(strconv.Itoa(1)),
		},
	}
	table, rowKey := "bi:datasource", "100"
	putRequest, err := hrpc.NewPutStr(context.Background(), table, rowKey, values)
	if err != nil {
		tt.Error(err)
	}

	client := gohbase.NewClient("localhost")
	defer client.Close()

	_, err = client.Put(putRequest)
	if err != nil {
		tt.Error(err)
	}

	// 根据 rowKey 获取整行
	getRequest, err := hrpc.NewGetStr(context.Background(), table, rowKey)
	if err != nil {
		tt.Error(err)
	}
	resp, err := client.Get(getRequest)
	if err != nil {
		tt.Error(err)
	}

	for _, cell := range resp.Cells {
		log.Printf("%s:%s %s\n", cell.Family, cell.Qualifier, cell.Value)
	}
}

func TestGetSpecificCellOfRow(t *testing.T) {
	tt := assert.New(t)

	families := map[string][]string{
		"meta_info": {
			"state",
		},
		"biz_info": {
			"biz_state",
		},
	}
	table, rowKey := "bi:datasource", "100"
	getReq, err := hrpc.NewGetStr(context.Background(), table, rowKey, hrpc.Families(families))
	if err != nil {
		tt.Error(err)
	}

	client := gohbase.NewClient("localhost")
	defer client.Close()

	resp, err := client.Get(getReq)
	if err != nil {
		tt.Error(err)
	}

	for _, cell := range resp.Cells {
		log.Printf("%s:%s %s\n", cell.Family, cell.Qualifier, cell.Value)
	}
}
