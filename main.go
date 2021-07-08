package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/DataDog/kafka-kit/v3/kafkazk"
)

func main() {
	// Handle flags.
	file := flag.String("file", "", "input map file")
	brokersToDemote := flag.String("brokers-to-demote", "", "brokers to remote")
	flag.Parse()

	// Broker ID set from CSV string.
	brokers := strings.Split(*brokersToDemote, ",")
	brokerSet := map[int]struct{}{}
	for _, id := range brokers {
		idInt, err := strconv.Atoi(id)
		exitOnErr(err)
		brokerSet[idInt] = struct{}{}
	}

	// Deserialize the input map.
	f, err := ioutil.ReadFile(*file)
	exitOnErr(err)

	pm1, err := kafkazk.PartitionMapFromString(string(f))
	exitOnErr(err)

	pm2 := pm1.Copy()

	// Swap positions for all leaders in the broker set.
	for _, partn := range pm2.Partitions {
		// Skip out of range errors.
		if len(partn.Replicas) < 2 {
			continue
		}
		leader := partn.Replicas[0]
		// If the leader is in the demotion set, swap its position with the
		// next replica.
		if _, toDemote := brokerSet[leader]; toDemote {
			partn.Replicas[0], partn.Replicas[1] = partn.Replicas[1], partn.Replicas[0]
		}
	}

	// Report changes.
	changes(pm1, pm2)

	// Write the output.
	_, err = json.Marshal(pm2)
	exitOnErr(err)

	// fmt.Println(string(out))
}

func changes(pm1, pm2 *kafkazk.PartitionMap) {
	fmt.Println("\nPartition map changes:")
	for i := range pm1.Partitions {
		change := whatChanged(pm1.Partitions[i].Replicas,
			pm2.Partitions[i].Replicas)

		fmt.Printf("%s p%d: %v -> %v %s\n",
			pm1.Partitions[i].Topic,
			pm1.Partitions[i].Partition,
			pm1.Partitions[i].Replicas,
			pm2.Partitions[i].Replicas,
			change)
	}
}

func exitOnErr(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}
