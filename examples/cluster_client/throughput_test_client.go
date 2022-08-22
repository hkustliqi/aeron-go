package main

import (
	"bytes"
	"fmt"
	"github.com/corymonroe-coinbase/aeron-go/examples/starbase_trading_api"
	"os"
	"strconv"
	"time"

	"github.com/corymonroe-coinbase/aeron-go/aeron"
	"github.com/corymonroe-coinbase/aeron-go/aeron/atomic"
	"github.com/corymonroe-coinbase/aeron-go/aeron/idlestrategy"
	"github.com/corymonroe-coinbase/aeron-go/aeron/logbuffer"
	"github.com/corymonroe-coinbase/aeron-go/cluster/client"
)

type TestContext struct {
	ac                    *client.AeronCluster
	messageCount          int
	latencies             []int64
	nextSendKeepAliveTime int64
}

func (ctx *TestContext) OnConnect(ac *client.AeronCluster) {
	fmt.Printf("OnConnect - sessionId=%d leaderMemberId=%d leadershipTermId=%d\n",
		ac.ClusterSessionId(), ac.LeaderMemberId(), ac.LeadershipTermId())
	ctx.ac = ac
	ctx.nextSendKeepAliveTime = time.Now().UnixMilli() + time.Second.Milliseconds()
}

func (ctx *TestContext) OnDisconnect(cluster *client.AeronCluster, details string) {
	fmt.Printf("OnDisconnect - sessionId=%d (%s)\n", cluster.ClusterSessionId(), details)
	ctx.ac = nil
}

func (ctx *TestContext) OnMessage(cluster *client.AeronCluster, timestamp int64,
	buffer *atomic.Buffer, offset int32, length int32, header *logbuffer.Header) {

	buf := toByteBuffer(buffer, offset, length)
	m := starbase_trading_api.NewSbeGoMarshaller()

	//blockLength := buffer.GetUInt16(offset)
	//templateId := buffer.GetUInt16(offset + 2)
	////schemaId := buffer.GetUInt16(offset + 4)
	//version := buffer.GetUInt16(offset + 6)
	//
	//corelationId := buffer.GetInt64(offset + 8)
	//corelationId += 1
	//evenType := buffer.GetInt8(offset + 16)
	//evenType += 1
	//rejectReason := buffer.GetInt16(offset + 17)
	//rejectReason += 1

	var hdr starbase_trading_api.SbeGoMessageHeader
	if err := hdr.Decode(m, buf); err != nil {
		fmt.Println("Failed to decode message header", err)
	}

	var conf starbase_trading_api.OmsMoonbaseOrderRequestConf

	if hdr.TemplateId != conf.SbeTemplateId() {
		fmt.Println("resp conf template ID mismatch!")
	}

	if err := conf.Decode(m, buf, hdr.Version, hdr.BlockLength, true); err != nil {
		fmt.Println("decoder err: %v", err)
	}

	if conf.RejectReason == starbase_trading_api.NewOrderRejectReason.NullValue {
		fmt.Println("order was accepted")
	}

	ctx.messageCount++
}

func toByteBuffer(buffer *atomic.Buffer, offset int32, length int32) *bytes.Buffer {
	buf := &bytes.Buffer{}
	buffer.WriteBytes(buf, offset, length)
	return buf
}

func (ctx *TestContext) OnNewLeader(cluster *client.AeronCluster, leadershipTermId int64, leaderMemberId int32) {
	fmt.Printf("OnNewLeader - sessionId=%d leaderMemberId=%d leadershipTermId=%d\n",
		cluster.ClusterSessionId(), leaderMemberId, leadershipTermId)
}

func (ctx *TestContext) OnError(cluster *client.AeronCluster, details string) {
	fmt.Printf("OnError - sessionId=%d: %s\n", cluster.ClusterSessionId(), details)
}

func (ctx *TestContext) sendKeepAliveIfNecessary() {
	if now := time.Now().UnixMilli(); now > ctx.nextSendKeepAliveTime && ctx.ac != nil && ctx.ac.SendKeepAlive() {
		ctx.nextSendKeepAliveTime += time.Second.Milliseconds()
	}
}

func main() {
	ctx := aeron.NewContext()
	//if aeronDir := os.Getenv("AERON_DIR"); aeronDir != "" {
	//	ctx.AeronDir(aeronDir)
	//	fmt.Println("aeron dir: ", aeronDir)
	//} else if _, err := os.Stat("/dev/shm"); err == nil {
	//	path := fmt.Sprintf("/dev/shm/aeron-%s", aeron.UserName)
	//	ctx.AeronDir(path)
	//	fmt.Println("aeron dir: ", path)
	//}
	aeronDir := "/tmp/aeron/"
	ctx.AeronDir(aeronDir)
	fmt.Println("aeron dir: ", aeronDir)

	opts := client.NewOptions()
	if idleStr := os.Getenv("NO_OP_IDLE"); idleStr != "" {
		opts.IdleStrategy = &idlestrategy.Busy{}
	}
	opts.IngressChannel = "aeron:udp?alias=cluster-ingress"
	opts.IngressEndpoints = "0=127.0.0.1:9150"
	opts.IngressStreamId = 200
	//opts.EgressChannel = "aeron:udp?alias=cluster-egress|endpoint=localhost:11111"

	//opts.EgressChannel = "aeron:udp?endpoint=localhost:0"

	listener := &TestContext{
		latencies: make([]int64, 1),
	}
	clusterClient, err := client.NewAeronCluster(ctx, opts, listener)
	if err != nil {
		panic(err)
	}

	for !clusterClient.IsConnected() {
		opts.IdleStrategy.Idle(clusterClient.Poll())
	}

	var buffer = new(bytes.Buffer)

	m := starbase_trading_api.NewSbeGoMarshaller()
	var order starbase_trading_api.OmsMoonbaseNewOrder
	header := starbase_trading_api.SbeGoMessageHeader{
		BlockLength: order.SbeBlockLength(),
		TemplateId:  order.SbeTemplateId(),
		SchemaId:    order.SbeSchemaId(),
		Version:     order.SbeSchemaVersion(),
	}
	if err := header.Encode(m, buffer); err != nil {
		fmt.Println("Encoding header Error", err)
	}

	populateOrder(&order)

	if err := order.Encode(m, buffer, false); err != nil {
		fmt.Println("Encoding Error", err)
	}

	//sendBuf := atomic.MakeBuffer(make([]byte, 16))
	sendBuf := atomic.MakeBuffer(buffer.Bytes())
	fmt.Printf("starting to send \n")
	listener.messageCount = 0
	sentCt := 0
	for i := 1; i <= 1; i++ {
		//sendBuf.PutInt64(0, int64(999))
		//sendBuf.PutInt64(8, int64(1234))
		for {
			if r := clusterClient.Offer(sendBuf, 0, sendBuf.Capacity()); r >= 0 {
				sentCt++
				fmt.Println("sentCt " + strconv.Itoa(sentCt))
				break
			}
			clusterClient.Poll()
			listener.sendKeepAliveIfNecessary()
		}
	}
	for listener.messageCount < sentCt {
		pollCt := clusterClient.Poll()
		if pollCt == 0 {
			listener.sendKeepAliveIfNecessary()
		}
		opts.IdleStrategy.Idle(pollCt)
	}
	clusterClient.Close()
	fmt.Println("done")
	time.Sleep(time.Second)
}

func populateOrder(order *starbase_trading_api.OmsMoonbaseNewOrder) {
	var userName [80]byte
	copy(userName[:], "user1")
	order.UserName = userName

	order.CorrelationId = 9876541

	var symbol [80]byte
	copy(symbol[:], "BTC-USD")
	order.InstrumentSymbol = symbol

	var clientOID [2]int64
	clientOID[0] = 99999
	order.ClientOrderId = clientOID

	var portfolio [80]byte
	copy(portfolio[:], "portfolio1")
	order.PortfolioName = portfolio

	var limitPrice [80]byte
	copy(limitPrice[:], "99999")
	order.LimitPrice = limitPrice

	var quantity [80]byte
	copy(quantity[:], "14")
	order.Quantity = quantity

	var stopPrice [80]byte
	copy(stopPrice[:], "5")
	order.StopPrice = stopPrice

	var funds [80]byte
	copy(funds[:], "0")
	order.Funds = funds

	order.ExpireTime = 1989
	order.OrderFlags = [64]bool{false}
	order.Side = starbase_trading_api.Side.BUY
	order.SelfMatchPreventionMode = starbase_trading_api.SelfMatchPreventionMode.CANCEL_NEWER
	order.TimeInForce = starbase_trading_api.TimeInForce.GOOD_TIL_TIME
}
