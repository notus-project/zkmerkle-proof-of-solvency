package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/notus-project/zkmerkle-proof-of-solvency/src/utils"
	"github.com/notus-project/zkmerkle-proof-of-solvency/src/witness/config"
	"github.com/notus-project/zkmerkle-proof-of-solvency/src/witness/witness"
)

func main() {
	remotePasswdConfig := flag.String("remote_password_config", "", "fetch password from aws secretsmanager")
	flag.Parse()
	witnessConfig := &config.Config{}
	content, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(content, witnessConfig)
	if err != nil {
		panic(err.Error())
	}
	if *remotePasswdConfig != "" {
		s, err := utils.GetMysqlSource(witnessConfig.MysqlDataSource, *remotePasswdConfig)
		if err != nil {
			panic(err.Error())
		}
		witnessConfig.MysqlDataSource = s
	}

	accounts, cexAssetsInfo, err := utils.ParseUserDataSet(witnessConfig.UserDataFile)
	fmt.Println("account counts", len(accounts))
	if err != nil {
		panic(err.Error())
	}
	accountTree, err := utils.NewAccountTree(witnessConfig.TreeDB.Driver, witnessConfig.TreeDB.Option.Addr)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("account tree init height is ", accountTree.LatestVersion())
	fmt.Printf("account tree root is %x\n", accountTree.Root())

	witnessService := witness.NewWitness(accountTree, uint32(len(accounts)), accounts, cexAssetsInfo, witnessConfig)
	startingTime := time.Now().UTC()
	witnessService.Run()
	duration := time.Now().UTC().Sub(startingTime)
	fmt.Printf("Running Prover after loading the key Takes [%.3f] Seconds \n", duration.Seconds())
	fmt.Println("witness service run finished...")
}
