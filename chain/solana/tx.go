package solana

import (
	"errors"
	"fmt"

	xc "github.com/jumpcrypto/crosschain"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/programs/vote"
	"github.com/gagliardetto/solana-go/rpc"
)

// Tx for Solana, encapsulating a solana.Transaction and other info
type Tx struct {
	SolTx                  *solana.Transaction
	ParsedSolTx            *rpc.ParsedTransaction // similar, but different type
	associatedTokenAccount *token.Account
	parsedTransfer         interface{}
}

var _ xc.Tx = Tx{}

// Hash returns the tx hash or id, for Solana it's signature
func (tx Tx) Hash() xc.TxHash {
	if tx.SolTx != nil && len(tx.SolTx.Signatures) > 0 {
		sig := tx.SolTx.Signatures[0]
		return xc.TxHash(sig.String())
	}
	return xc.TxHash("")
}

// Sighashes returns the tx payload to sign, aka sighashes
func (tx Tx) Sighashes() ([]xc.TxDataToSign, error) {
	if tx.SolTx == nil {
		return nil, errors.New("transaction not initialized")
	}
	messageContent, err := tx.SolTx.Message.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("unable to encode message for signing: %w", err)
	}
	return []xc.TxDataToSign{messageContent}, nil
}

// AddSignatures adds a signature to Tx
func (tx Tx) AddSignatures(signatures ...xc.TxSignature) error {
	if tx.SolTx == nil {
		return errors.New("transaction not initialized")
	}
	solSignatures := make([]solana.Signature, len(signatures))
	for i, signature := range signatures {
		if len(signature) != solana.SignatureLength {
			return fmt.Errorf("invalid signature (%d): %x", len(signature), signature)
		}
		copy(solSignatures[i][:], signature)
	}
	tx.SolTx.Signatures = solSignatures
	return nil
}

// ParseTransfer parses a tx and extracts higher-level transfer information
func (tx *Tx) ParseTransfer() {
	transfer, _ := tx.getSystemTransfer()
	if transfer != nil {
		tx.parsedTransfer = transfer
		return
	}
	tokenTC, _ := tx.getTokenTransferChecked()
	if tokenTC != nil {
		tx.parsedTransfer = tokenTC
		return
	}
	tokenT, _ := tx.getTokenTransfer()
	if tokenT != nil {
		tx.parsedTransfer = tokenT
		return
	}
	voteW, _ := tx.getVoteWithdraw()
	if voteW != nil {
		tx.parsedTransfer = voteW
		return
	}
}

// SetAssociatedTokenAccount sets the associated token account
func (tx *Tx) SetAssociatedTokenAccount(ata *token.Account) {
	tx.associatedTokenAccount = ata
}

// From is the sender of a transfer
func (tx Tx) From() xc.Address {
	switch tf := tx.parsedTransfer.(type) {
	case *system.Transfer:
		from := tf.GetFundingAccount().PublicKey.String()
		return xc.Address(from)
	case *token.TransferChecked:
		from := tf.GetOwnerAccount().PublicKey.String()
		return xc.Address(from)
	case *token.Transfer:
		from := tf.GetOwnerAccount().PublicKey.String()
		return xc.Address(from)
	case *vote.Withdraw:
		// https://docs.rs/solana-vote-program/latest/solana_vote_program/vote_instruction/enum.VoteInstruction.html#variant.Withdraw
		from := tf.GetAccounts()[2].PublicKey.String()
		return xc.Address(from)
	}
	return xc.Address("")
}

// To is the account receiving a transfer
func (tx Tx) To() xc.Address {
	// if ATA is set, return owner
	if tx.associatedTokenAccount != nil {
		return xc.Address(tx.associatedTokenAccount.Owner.String())
	}

	switch tf := tx.parsedTransfer.(type) {
	case *system.Transfer:
		to := tf.GetRecipientAccount().PublicKey.String()
		return xc.Address(to)
	case *vote.Withdraw:
		// https://docs.rs/solana-vote-program/latest/solana_vote_program/vote_instruction/enum.VoteInstruction.html#variant.Withdraw
		to := tf.GetAccounts()[1].PublicKey.String()
		return xc.Address(to)
	}

	// tokens are not available, need to set tx.associatedTokenAccount
	return xc.Address("")
}

// ToAlt returns an alternative recipient, for Solana the Associated Token Account
func (tx Tx) ToAlt() xc.Address {
	// only for tokens
	switch tf := tx.parsedTransfer.(type) {
	case *token.TransferChecked:
		dst := tf.GetDestinationAccount().PublicKey.String()
		return xc.Address(dst)
	case *token.Transfer:
		dst := tf.GetDestinationAccount().PublicKey.String()
		return xc.Address(dst)
	}
	return xc.Address("")
}

// Amount returns the tx amount
func (tx Tx) Amount() xc.AmountBlockchain {
	switch tf := tx.parsedTransfer.(type) {
	case *system.Transfer:
		return xc.NewAmountBlockchainFromUint64(*tf.Lamports)
	case *token.TransferChecked:
		return xc.NewAmountBlockchainFromUint64(*tf.Amount)
	case *token.Transfer:
		return xc.NewAmountBlockchainFromUint64(*tf.Amount)
	case *vote.Withdraw:
		return xc.NewAmountBlockchainFromUint64(*tf.Lamports)
	}
	return xc.NewAmountBlockchainFromUint64(0)
}

// ContractAddress returns the contract address for a token transfer
func (tx Tx) ContractAddress() xc.ContractAddress {
	// if ATA is set, return mint
	if tx.associatedTokenAccount != nil {
		return xc.ContractAddress(tx.associatedTokenAccount.Mint.String())
	}

	// only TransferChecked contains mint addr - Transfer does not
	switch tf := tx.parsedTransfer.(type) {
	case *token.TransferChecked:
		contract := tf.GetMintAccount().PublicKey.String()
		return xc.ContractAddress(contract)
	}

	return xc.ContractAddress("")
}

// RecentBlockhash returns the recent block hash used as a nonce for a Solana tx
func (tx Tx) RecentBlockhash() string {
	if tx.ParsedSolTx != nil {
		return tx.ParsedSolTx.Message.RecentBlockHash
	}
	if tx.SolTx != nil {
		return tx.SolTx.Message.RecentBlockhash.String()
	}
	return ""
}

func (tx Tx) getVoteWithdraw() (*vote.Withdraw, error) {
	if tx.SolTx != nil {
		message := tx.SolTx.Message
		for _, instruction := range message.Instructions {
			program, err := message.ResolveProgramIDIndex(instruction.ProgramIDIndex)
			if err != nil {
				continue
			}
			if !program.Equals(solana.VoteProgramID) {
				continue
			}
			accs, err := instruction.ResolveInstructionAccounts(&message)
			if err != nil {
				continue
			}
			inst, err := vote.DecodeInstruction(accs, instruction.Data)
			if err != nil {
				continue
			}
			castedInst, ok := inst.Impl.(*vote.Withdraw)
			if !ok {
				continue
			}
			return castedInst, nil
		}
	}
	return nil, fmt.Errorf("no tx set")
}

func (tx Tx) getSystemTransfer() (*system.Transfer, error) {
	if tx.SolTx != nil {
		message := tx.SolTx.Message
		for _, instruction := range message.Instructions {
			program, err := message.ResolveProgramIDIndex(instruction.ProgramIDIndex)
			if err != nil {
				continue
			}
			if !program.Equals(solana.SystemProgramID) {
				continue
			}
			accs, err := instruction.ResolveInstructionAccounts(&message)
			if err != nil {
				continue
			}
			inst, err := system.DecodeInstruction(accs, instruction.Data)
			if err != nil {
				continue
			}
			castedInst, ok := inst.Impl.(*system.Transfer)
			if !ok {
				continue
			}
			return castedInst, nil
		}
	}
	return nil, fmt.Errorf("no tx set")
}

func (tx Tx) getTokenTransferChecked() (*token.TransferChecked, error) {
	if tx.SolTx != nil {
		message := tx.SolTx.Message
		for _, instruction := range message.Instructions {
			program, err := message.ResolveProgramIDIndex(instruction.ProgramIDIndex)
			if err != nil {
				continue
			}
			if !program.Equals(solana.TokenProgramID) {
				continue
			}
			accs, err := instruction.ResolveInstructionAccounts(&message)
			if err != nil {
				continue
			}
			inst, err := token.DecodeInstruction(accs, instruction.Data)
			if err != nil {
				continue
			}
			castedInst, ok := inst.Impl.(*token.TransferChecked)
			if !ok {
				continue
			}
			return castedInst, nil
		}
		return nil, fmt.Errorf("no instruction is *token.TransferChecked")
	}
	return nil, fmt.Errorf("no tx set")
}

func (tx Tx) getTokenTransfer() (*token.Transfer, error) {
	if tx.SolTx != nil {
		message := tx.SolTx.Message
		for _, instruction := range message.Instructions {
			program, err := message.ResolveProgramIDIndex(instruction.ProgramIDIndex)
			if err != nil {
				continue
			}
			if !program.Equals(solana.TokenProgramID) {
				continue
			}
			accs, err := instruction.ResolveInstructionAccounts(&message)
			if err != nil {
				continue
			}
			inst, err := token.DecodeInstruction(accs, instruction.Data)
			if err != nil {
				continue
			}
			castedInst, ok := inst.Impl.(*token.Transfer)
			if !ok {
				continue
			}
			return castedInst, nil
		}
		return nil, fmt.Errorf("no instruction is *token.Transfer")
	}
	return nil, fmt.Errorf("no tx set")
}

// Serialize returns the serialized tx
func (tx Tx) Serialize() ([]byte, error) {
	if tx.SolTx == nil {
		return []byte{}, errors.New("transaction not initialized")
	}
	return tx.SolTx.MarshalBinary()
}
