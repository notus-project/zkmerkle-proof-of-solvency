package circuit

import (
	"github.com/consensys/gnark/frontend"
	"github.com/notus-project/zkmerkle-proof-of-solvency/src/utils"
)

type (
	Variable = frontend.Variable
	API      = frontend.API
)

type CexAssetInfo struct {
	TotalEquity Variable
	TotalDebt   Variable
	BasePrice   Variable
}

type UserAssetInfo struct {
	Equity Variable
	Debt   Variable
}

type CreateUserOperation struct {
	BeforeAccountTreeRoot Variable
	AfterAccountTreeRoot  Variable
	Assets                []UserAssetInfo
	AccountIndex          Variable
	AccountIdHash         Variable
	AccountProof          [utils.AccountTreeDepth]Variable
}
