// Auto-generated by avdl-compiler v1.3.13 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/block.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type GetBlockRes struct {
	BlockKey string `codec:"blockKey" json:"blockKey"`
	Buf      []byte `codec:"buf" json:"buf"`
}

type BlockRefNonce [8]byte
type BlockReference struct {
	Bid       BlockIdCombo  `codec:"bid" json:"bid"`
	Nonce     BlockRefNonce `codec:"nonce" json:"nonce"`
	ChargedTo UID           `codec:"chargedTo" json:"chargedTo"`
}

type BlockReferenceCount struct {
	Ref       BlockReference `codec:"ref" json:"ref"`
	LiveCount int            `codec:"liveCount" json:"liveCount"`
}

type DowngradeReferenceRes struct {
	Completed []BlockReferenceCount `codec:"completed" json:"completed"`
	Failed    BlockReference        `codec:"failed" json:"failed"`
}

type BlockPingResponse struct {
}

type GetSessionChallengeArg struct {
}

type AuthenticateSessionArg struct {
	Signature string `codec:"signature" json:"signature"`
}

type PutBlockArg struct {
	Bid      BlockIdCombo `codec:"bid" json:"bid"`
	Folder   string       `codec:"folder" json:"folder"`
	BlockKey string       `codec:"blockKey" json:"blockKey"`
	Buf      []byte       `codec:"buf" json:"buf"`
}

type GetBlockArg struct {
	Bid    BlockIdCombo `codec:"bid" json:"bid"`
	Folder string       `codec:"folder" json:"folder"`
}

type AddReferenceArg struct {
	Folder string         `codec:"folder" json:"folder"`
	Ref    BlockReference `codec:"ref" json:"ref"`
}

type DelReferenceArg struct {
	Folder string         `codec:"folder" json:"folder"`
	Ref    BlockReference `codec:"ref" json:"ref"`
}

type ArchiveReferenceArg struct {
	Folder string           `codec:"folder" json:"folder"`
	Refs   []BlockReference `codec:"refs" json:"refs"`
}

type DelReferenceWithCountArg struct {
	Folder string           `codec:"folder" json:"folder"`
	Refs   []BlockReference `codec:"refs" json:"refs"`
}

type ArchiveReferenceWithCountArg struct {
	Folder string           `codec:"folder" json:"folder"`
	Refs   []BlockReference `codec:"refs" json:"refs"`
}

type GetUserQuotaInfoArg struct {
}

type BlockPingArg struct {
}

type BlockInterface interface {
	GetSessionChallenge(context.Context) (ChallengeInfo, error)
	AuthenticateSession(context.Context, string) error
	PutBlock(context.Context, PutBlockArg) error
	GetBlock(context.Context, GetBlockArg) (GetBlockRes, error)
	AddReference(context.Context, AddReferenceArg) error
	DelReference(context.Context, DelReferenceArg) error
	ArchiveReference(context.Context, ArchiveReferenceArg) ([]BlockReference, error)
	DelReferenceWithCount(context.Context, DelReferenceWithCountArg) (DowngradeReferenceRes, error)
	ArchiveReferenceWithCount(context.Context, ArchiveReferenceWithCountArg) (DowngradeReferenceRes, error)
	GetUserQuotaInfo(context.Context) ([]byte, error)
	BlockPing(context.Context) (BlockPingResponse, error)
}

func BlockProtocol(i BlockInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.block",
		Methods: map[string]rpc.ServeHandlerDescription{
			"getSessionChallenge": {
				MakeArg: func() interface{} {
					ret := make([]GetSessionChallengeArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.GetSessionChallenge(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"authenticateSession": {
				MakeArg: func() interface{} {
					ret := make([]AuthenticateSessionArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]AuthenticateSessionArg)
					if !ok {
						err = rpc.NewTypeError((*[]AuthenticateSessionArg)(nil), args)
						return
					}
					err = i.AuthenticateSession(ctx, (*typedArgs)[0].Signature)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"putBlock": {
				MakeArg: func() interface{} {
					ret := make([]PutBlockArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PutBlockArg)
					if !ok {
						err = rpc.NewTypeError((*[]PutBlockArg)(nil), args)
						return
					}
					err = i.PutBlock(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getBlock": {
				MakeArg: func() interface{} {
					ret := make([]GetBlockArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetBlockArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetBlockArg)(nil), args)
						return
					}
					ret, err = i.GetBlock(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"addReference": {
				MakeArg: func() interface{} {
					ret := make([]AddReferenceArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]AddReferenceArg)
					if !ok {
						err = rpc.NewTypeError((*[]AddReferenceArg)(nil), args)
						return
					}
					err = i.AddReference(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"delReference": {
				MakeArg: func() interface{} {
					ret := make([]DelReferenceArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]DelReferenceArg)
					if !ok {
						err = rpc.NewTypeError((*[]DelReferenceArg)(nil), args)
						return
					}
					err = i.DelReference(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"archiveReference": {
				MakeArg: func() interface{} {
					ret := make([]ArchiveReferenceArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ArchiveReferenceArg)
					if !ok {
						err = rpc.NewTypeError((*[]ArchiveReferenceArg)(nil), args)
						return
					}
					ret, err = i.ArchiveReference(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"delReferenceWithCount": {
				MakeArg: func() interface{} {
					ret := make([]DelReferenceWithCountArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]DelReferenceWithCountArg)
					if !ok {
						err = rpc.NewTypeError((*[]DelReferenceWithCountArg)(nil), args)
						return
					}
					ret, err = i.DelReferenceWithCount(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"archiveReferenceWithCount": {
				MakeArg: func() interface{} {
					ret := make([]ArchiveReferenceWithCountArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ArchiveReferenceWithCountArg)
					if !ok {
						err = rpc.NewTypeError((*[]ArchiveReferenceWithCountArg)(nil), args)
						return
					}
					ret, err = i.ArchiveReferenceWithCount(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getUserQuotaInfo": {
				MakeArg: func() interface{} {
					ret := make([]GetUserQuotaInfoArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.GetUserQuotaInfo(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"blockPing": {
				MakeArg: func() interface{} {
					ret := make([]BlockPingArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.BlockPing(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type BlockClient struct {
	Cli rpc.GenericClient
}

func (c BlockClient) GetSessionChallenge(ctx context.Context) (res ChallengeInfo, err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.getSessionChallenge", []interface{}{GetSessionChallengeArg{}}, &res)
	return
}

func (c BlockClient) AuthenticateSession(ctx context.Context, signature string) (err error) {
	__arg := AuthenticateSessionArg{Signature: signature}
	err = c.Cli.Call(ctx, "keybase.1.block.authenticateSession", []interface{}{__arg}, nil)
	return
}

func (c BlockClient) PutBlock(ctx context.Context, __arg PutBlockArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.putBlock", []interface{}{__arg}, nil)
	return
}

func (c BlockClient) GetBlock(ctx context.Context, __arg GetBlockArg) (res GetBlockRes, err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.getBlock", []interface{}{__arg}, &res)
	return
}

func (c BlockClient) AddReference(ctx context.Context, __arg AddReferenceArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.addReference", []interface{}{__arg}, nil)
	return
}

func (c BlockClient) DelReference(ctx context.Context, __arg DelReferenceArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.delReference", []interface{}{__arg}, nil)
	return
}

func (c BlockClient) ArchiveReference(ctx context.Context, __arg ArchiveReferenceArg) (res []BlockReference, err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.archiveReference", []interface{}{__arg}, &res)
	return
}

func (c BlockClient) DelReferenceWithCount(ctx context.Context, __arg DelReferenceWithCountArg) (res DowngradeReferenceRes, err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.delReferenceWithCount", []interface{}{__arg}, &res)
	return
}

func (c BlockClient) ArchiveReferenceWithCount(ctx context.Context, __arg ArchiveReferenceWithCountArg) (res DowngradeReferenceRes, err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.archiveReferenceWithCount", []interface{}{__arg}, &res)
	return
}

func (c BlockClient) GetUserQuotaInfo(ctx context.Context) (res []byte, err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.getUserQuotaInfo", []interface{}{GetUserQuotaInfoArg{}}, &res)
	return
}

func (c BlockClient) BlockPing(ctx context.Context) (res BlockPingResponse, err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.blockPing", []interface{}{BlockPingArg{}}, &res)
	return
}
