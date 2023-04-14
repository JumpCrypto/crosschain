package bitcoin

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/cosmos/btcutil"
	xc "github.com/jumpcrypto/crosschain"
)

// AddressBuilder for Bitcoin
type AddressBuilder struct {
	params        *chaincfg.Params
	UseLegacy     bool
	UseScriptHash bool
}

var _ xc.AddressBuilder = &AddressBuilder{}

// NewAddressBuilder creates a new Bitcoin AddressBuilder
func NewAddressBuilder(asset xc.ITask) (xc.AddressBuilder, error) {
	params, err := GetParams(asset.GetAssetConfig())
	if err != nil {
		return AddressBuilder{}, err
	}
	return AddressBuilder{
		UseLegacy:     true,
		UseScriptHash: false,
		params:        params,
	}, nil
}

// GetAddressFromPublicKey returns an Address given a public key
func (ab AddressBuilder) GetAddressFromPublicKey(publicKeyBytes []byte) (xc.Address, error) {
	// hack to support Taproot until btcutil is bumped
	if len(publicKeyBytes) == 32 {
		publicKeyBytes = append([]byte{0x02}, publicKeyBytes...)
	}
	var address string
	if ab.UseLegacy {
		// legacy
		addressPubKey, err := btcutil.NewAddressPubKey(publicKeyBytes, ab.params)
		if err != nil {
			return "", err
		}
		address = addressPubKey.EncodeAddress()
	} else if ab.UseScriptHash {
		// segwith
		scriptHash := btcutil.Hash160(publicKeyBytes)
		addressPubKey, err := btcutil.NewAddressScriptHashFromHash(scriptHash, ab.params)
		if err != nil {
			return "", err
		}
		address = addressPubKey.EncodeAddress()
	} else {
		// segwith-bech32
		witnessProg := btcutil.Hash160(publicKeyBytes)
		addressPubKey, err := btcutil.NewAddressWitnessPubKeyHash(witnessProg, ab.params)
		if err != nil {
			return "", err
		}
		address = addressPubKey.EncodeAddress()
	}
	return xc.Address(address), nil
}

// GetAllPossibleAddressesFromPublicKey returns all PossubleAddress(es) given a public key
func (ab AddressBuilder) GetAllPossibleAddressesFromPublicKey(publicKeyBytes []byte) ([]xc.PossibleAddress, error) {

	possibles := []xc.PossibleAddress{}
	addressPubKey, err := btcutil.NewAddressPubKey(publicKeyBytes, ab.params)
	if err != nil {
		return possibles, err
	}
	possibles = append(possibles, xc.PossibleAddress{
		Address: xc.Address(addressPubKey.EncodeAddress()),
		Type:    xc.AddressTypeP2PKH,
	})

	witnessProg := btcutil.Hash160(publicKeyBytes)
	addressWitness, err := btcutil.NewAddressWitnessPubKeyHash(witnessProg, ab.params)
	if err != nil {
		return possibles, err
	}
	possibles = append(possibles,
		xc.PossibleAddress{
			Address: xc.Address(addressWitness.EncodeAddress()),
			Type:    xc.AddressTypeP2WPKH,
		},
	)

	return possibles, nil
}
