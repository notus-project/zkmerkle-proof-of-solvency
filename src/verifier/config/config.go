package config

import (
	"math/big"

	"github.com/notus-project/zkmerkle-proof-of-solvency/src/utils"
)

type Config struct {
	ProofTable    string
	ZkKeyName     string
	CexAssetsInfo []utils.CexAssetInfo
}

type UserConfig struct {
	AccountIndex  uint32
	AccountIdHash string
	TotalEquity   big.Int
	TotalDebt     big.Int
	Root          string
	Assets        []utils.AccountAsset
	Proof         []string
}
