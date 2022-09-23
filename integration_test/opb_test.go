package integration

import (
	"encoding/hex"
	"fmt"
	. "github.com/kaifei-bianjie/common-parser/codec"
	"github.com/kaifei-bianjie/common-parser/utils"
)

func (s IntegrationTestSuite) TestOpb() {
	cases := []SubTest{
		{
			"Mint",
			Mint,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func Mint(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)

	txBytes, err := hex.DecodeString("0A720A700A122F69726974612E6F70622E4D73674D696E74125A0801122A696161313864766475746177733837337A773333397433386A377530757238793637636A6D6D6D6663611A2A696161313864766475746177733837337A773333397433386A377530757238793637636A6D6D6D66636112630A4B0A400A192F636F736D6F732E63727970746F2E736D322E5075624B657912230A21031F019DB4CB2D93A06B39BBE048D8330AC56531C6FEEF3633DDBE88638D8B215712040A02080118A62612140A0E0A0475676173120632303030303010C09A0C1A40168F57420C5CD2E43975985534A450B7B2BCEE6CB2F72148E8A2169CF785565E6F569B8BD7E873E7F541F732F759FA348399A5F5825F56DB25005E83EC25996B")
	if err != nil {
		s.T().Log(err.Error())
		panic(err)
	}

	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		s.T().Log(err.Error())
		panic(err)
	}

	for _, msg := range authTx.GetMsgs() {
		if mtDoc, ok := s.Opb.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(mtDoc))
		}
	}
}
