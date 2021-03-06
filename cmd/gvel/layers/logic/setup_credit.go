package logic

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/velo-protocol/DRSv1/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv1/cmd/gvel/utils/crypto"
	"github.com/velo-protocol/DRSv1/cmd/gvel/utils/parser"
	"github.com/velo-protocol/DRSv1/libs/convert"
	"github.com/velo-protocol/DRSv1/libs/txnbuild"
)

func (lo *logic) SetupCredit(input *entity.SetupCreditInput) (*entity.SetupCreditOutput, error) {
	defaultAccount := lo.AppConfig.GetDefaultAccount()

	accountBytes, err := lo.DB.Get([]byte(defaultAccount))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get account from db")
	}

	var account entity.StellarAccount
	err = json.Unmarshal(accountBytes, &account)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal account")
	}

	seedBytes, err := crypto.Decrypt(account.EncryptedSeed, input.Passphrase)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to decrypt the seed of %s with given passphrase", defaultAccount)
	}

	keyPair, err := vconvert.SecretKeyToKeyPair(string(seedBytes))
	if err != nil {
		return nil, errors.Wrap(err, "failed to derive keypair from seed")
	}

	result, err := lo.Velo.Client(keyPair).SetupCredit(context.Background(), vtxnbuild.SetupCredit{
		PeggedValue:    input.PeggedValue,
		PeggedCurrency: input.PeggedCurrency,
		AssetCode:      input.AssetCode,
	})
	if err != nil {
		err = parser.ParseHorizonError(err, lo.AppConfig.GetHorizonUrl(), lo.AppConfig.GetNetworkPassphrase())
		return nil, errors.Wrap(err, "failed to setup credit")
	}

	return &entity.SetupCreditOutput{
		PeggedValue:    input.PeggedValue,
		PeggedCurrency: input.PeggedCurrency,
		AssetCode:      input.AssetCode,
		SourceAddress:  defaultAccount,
		TxResult:       result.HorizonResult,
	}, nil
}
