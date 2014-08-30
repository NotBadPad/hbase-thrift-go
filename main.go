package main

import (
	"encoding/json"
	"fmt"
	// "git.apache.org/thrift.git/lib/go/thrift"
	"github.com/sdming/goh"
	"github.com/sdming/goh/Hbase"
	"time"
	// "hbase-thrift/Hbase"
	// "net"
)

// /**
//  * 获取CDSAPI服务
//  */
// func GetHbaseApiClient() (client *Hbase.HbaseClient, err error) {
// 	var trans thrift.TTransport
// 	trans, err = thrift.NewTSocket(net.JoinHostPort("192.168.8.100", "9090"))
// 	if err != nil {
// 		fmt.Printf("Error marshal msg: %s\n", err.Error())
// 	}
// 	trans = thrift.NewTFramedTransport(trans)
// 	protocolFactory := thrift.NewTJSONProtocolFactory()
// 	client = Hbase.NewHbaseClientFactory(trans, protocolFactory)
// 	return
// }

// func test() {
// 	defer func() {
// 		if x := recover(); x != nil {
// 			fmt.Println("error: ", x)
// 		}
// 	}()

// 	client, err := GetHbaseApiClient()

// 	if err != nil {
// 		return
// 	}

// 	defer client.Transport.Close()
// 	if err = client.Transport.Open(); err != nil {
// 		fmt.Printf("Error Open: %s\n", err.Error())
// 	}
// 	fmt.Println(client)

// 	r, _, err := client.GetTableNames()
// 	if err != nil {
// 		fmt.Printf("Error: %s\n", err.Error())
// 	}
// 	fmt.Println("aaa", len(r))
// 	return
// }

func test2() {
	address := "192.168.8.100:9090"

	client, err := goh.NewTcpClient(address, goh.TBinaryProtocol, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = client.Open(); err != nil {
		fmt.Println(err)
		return
	}

	defer client.Close()

	scan := &goh.TScan{
		StartRow: []byte("org|32263_79859470786588|2131175_79859470786588_664766856"),
		StopRow:  []byte("org|32263_79859470794091|2131108_79859470794091_664695117"),
		Caching:  100,
		//FilterString: "substring:value",
	}
	if data, err := client.ScannerOpenWithScan("vm_money", scan, nil); err != nil {
		fmt.Println(err)
	} else {
		dump(data)
		scanId := data
		fmt.Println(scanId)

		if scanId > 0 {
			if data, err := client.ScannerGetList(5, 10); err != nil {
				fmt.Println(err)
			} else {
				printRows(data)
			}
		}
	}
}

func dump(data interface{}) {
	fmt.Println(data)
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json.Marshal error:", err)
		return
	}
	fmt.Println(string(b))
}

func printRows(data []*Hbase.TRowResult) {
	if data == nil {
		fmt.Println("<nil>")
	}

	l := len(data)
	fmt.Println("[]*Hbase.TRowResult len:", l)
	for i, x := range data {
		fmt.Println(i, string(x.Row), "\n[")
		for k, v := range x.Columns {
			fmt.Println("\t", k, ":", string(v.Value), v.Timestamp)
		}
		fmt.Println("]")
	}

}

func main() {
	start := time.Now().UnixNano()
	test2()
	end := time.Now().UnixNano()
	fmt.Println(end - start)
}
