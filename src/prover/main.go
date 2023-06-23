package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/notus-project/zkmerkle-proof-of-solvency/src/prover/config"
	"github.com/notus-project/zkmerkle-proof-of-solvency/src/prover/prover"
	"github.com/notus-project/zkmerkle-proof-of-solvency/src/utils"
)

func main() {
	proverConfig := &config.Config{}
	content, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(content, proverConfig)
	if err != nil {
		panic(err.Error())
	}
	remotePasswdConfig := flag.String("remote_password_config", "", "fetch password from aws secretsmanager")
	rerun := flag.Bool("rerun", false, "flag which indicates rerun proof generation")
	flag.Parse()
	if *remotePasswdConfig != "" {
		s, err := utils.GetMysqlSource(proverConfig.MysqlDataSource, *remotePasswdConfig)
		if err != nil {
			panic(err.Error())
		}
		proverConfig.MysqlDataSource = s
	}
	prover := prover.NewProver(proverConfig)
	startingTime := time.Now().UTC()
	prover.Run(*rerun)
	duration := time.Now().UTC().Sub(startingTime)
	fmt.Printf("Running Prover after loading the key Takes [%.3f] Seconds \n", duration.Seconds())
}
