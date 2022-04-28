package transport

import (
	"context"
	"fmt"
	"log"
	"monografia/lib/database"
	"os"
	"sync"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/network"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Benchmark() func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	var mutex sync.Mutex
	var cpuValues []float64
	var netValues []uint64

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		metadata, _ := metadata.FromIncomingContext(ctx)

		var testID, reqID string

		val, ok := metadata["x-test"]
		if ok && len(val) > 0 {
			testID = val[0]
		}

		val, ok = metadata["x-req"]
		if ok && len(val) > 0 {
			reqID = val[0]
		}

		ctx = context.WithValue(ctx, "x-test", testID)
		ctx = context.WithValue(ctx, "x-req", reqID)

		if os.Getenv("ENV") != "production" {
			return handler(ctx, req)
		}

		if reqID == "" {
			return nil, fmt.Errorf("missing x-req")
		}
		if testID == "" {
			return nil, fmt.Errorf("missing x-test")
		}

		var cpuUsage float64
		var netUsage uint64

		before, err := cpu.Get()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Duration(1) * time.Second)
		after, err := cpu.Get()
		if err != nil {
			log.Fatal(err)
		}
		total := float64(after.Total - before.Total)
		cpuUser := float64(after.User-before.User) / total * 100
		cpuSystem := float64(after.System-before.System) / total * 100

		cpuUsage = cpuUser + cpuSystem

		mutex.Lock()
		cpuValues = append(cpuValues, cpuUsage)
		var cpuMed float64
		var sum float64
		for _, v := range cpuValues {
			sum += v
		}
		cpuMed = sum / float64(len(cpuValues))
		mutex.Unlock()

		var startBytes, endBytes uint64

		netStats, err := network.Get()
		if err != nil {
			log.Fatal(err)
		}
		for _, n := range netStats {
			if n.Name == "eth0" {
				startBytes = n.RxBytes
				break
			}
		}

		time.Sleep(time.Second)

		netStats, err = network.Get()
		if err != nil {
			log.Fatal(err)
		}
		for _, n := range netStats {
			if n.Name == "eth0" {
				endBytes = n.RxBytes
				break
			}
		}

		netUsage = endBytes - startBytes
		mutex.Lock()
		netValues = append(netValues, netUsage)
		var netMed float64
		var sumNet uint64
		for _, v := range netValues {
			sumNet += v
		}
		netMed = float64(sumNet) / float64(len(netValues))
		mutex.Unlock()

		_, err = db.Exec(`
				INSERT INTO benchmark(test, resource, x, y)
				VALUES (?, ?, ?, ?), (?, ?, ?, ?)
			`,
			testID, "cpu", reqID, cpuMed,
			testID, "net", reqID, netMed/128.0,
		)
		if err != nil {
			log.Fatal(err)
		}

		return handler(ctx, req)
	}
}
