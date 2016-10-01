package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/hirokazumiyaji/goflake"
)

const version = "0.0.1"

type server struct {
	worker *goflake.IDWorker
	retry  int
}

func (s *server) GetID(ctx context.Context, in *GetIDRequest) (*GetIDResponse, error) {
	id, err := getID(s.worker, "GRPC", s.retry)
	if err != nil {
		return nil, err
	}
	return &GetIDResponse{ID: strconv.FormatUint(id, 10)}, nil
}

func main() {
	var (
		port         int
		datacenterID int
		workerID     int
		retry        int
		startTime    string
		err          error
	)

	defaultPort := 50051
	if v := os.Getenv("PORT"); v != "" {
		defaultPort, err = strconv.Atoi(v)
		if err != nil {
			log.Fatalf("port value type not integer. %v\n", err)
		}
	}
	flag.IntVar(&port, "port", defaultPort, "port to use (default 50051)")
	flag.IntVar(&port, "p", defaultPort, "port to use (default 50051)")

	defaultStartTime := os.Getenv("START_TIME")
	if defaultStartTime == "" {
		defaultStartTime = "2016-01-01 00:00:00 +0000"
	}
	flag.StringVar(&startTime, "start-time", defaultStartTime, "id generate start time (default '2016-01-01 00:00:00 +0000')")
	flag.StringVar(&startTime, "s", defaultStartTime, "id generate start time (short)")

	defaultDatacenterID := 1
	if v := os.Getenv("DATACENTER_ID"); v != "" {
		defaultDatacenterID, err = strconv.Atoi(v)
		if err != nil {
			log.Fatalf("datacenter id value type not integer. %v\n", err)
		}
	}
	flag.IntVar(&datacenterID, "datacenter-id", defaultDatacenterID, "datacenter id (default 1)")
	flag.IntVar(&datacenterID, "d", defaultDatacenterID, "datacenter id (short)")

	defaultWorkerID := 1
	if v := os.Getenv("WORKER_ID"); v != "" {
		defaultWorkerID, err = strconv.Atoi(v)
		if err != nil {
			log.Fatalf("worker id value type not integer. %v\n", err)
		}
	}
	flag.IntVar(&workerID, "worker-id", defaultWorkerID, "worker id (default 1)")
	flag.IntVar(&workerID, "w", defaultWorkerID, "worker id (short)")

	defaultRetry := 5
	if v := os.Getenv("RETRY"); v != "" {
		defaultRetry, err = strconv.Atoi(v)
		if err != nil {
			log.Fatalf("retry value type not integer. %v\n", err)
		}
	}
	flag.IntVar(&retry, "retry", defaultRetry, "generate id retry count (default 5)")
	flag.IntVar(&retry, "r", defaultRetry, "generate id retry count (short)")

	flag.Parse()

	t, err := time.Parse("2006-01-02 15:04:05 -0700", startTime)
	if err != nil {
		log.Fatalf("start time parse error. %v\n", err)
	}

	w, err := goflake.NewIDWorker(uint16(datacenterID), uint16(workerID), t)
	if err != nil {
		log.Fatalf("could not create id worker. %v\n", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	RegisterGoflakeServer(s, &server{worker: w, retry: retry})
	log.Fatal(s.Serve(lis))
}

func getID(w *goflake.IDWorker, u string, r int) (uint64, error) {
	var (
		id  uint64
		err error
	)
	for i := 0; i < r; i++ {
		id, err = w.GetID(u)
		if err == nil {
			break
		}
	}
	return id, err
}
