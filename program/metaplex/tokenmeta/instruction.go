package tokenmeta

import (
	"github.com/creamy-cpu/solana-go-sdk/common"
	"github.com/creamy-cpu/solana-go-sdk/types"
	"github.com/near/borsh-go"
)

type Instruction uint8

const (
	InstructionCreateMetadataAccount Instruction = iota
)

func CreateMetadataAccount(metadata, mint, mintAuthority, payer, updateAuthority common.PublicKey, updateAuthorityIsSigner, isMutable bool, mintData Data) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Data        Data
		IsMutable   bool
	}{
		Instruction: InstructionCreateMetadataAccount,
		Data:        mintData,
		IsMutable:   isMutable,
	})

	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     mint,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     mintAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     updateAuthority,
				IsSigner:   updateAuthorityIsSigner,
				IsWritable: false,
			},
			{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SysVarRentPubkey,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}
