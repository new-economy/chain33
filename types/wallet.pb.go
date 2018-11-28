// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wallet.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 钱包模块存贮的tx交易详细信息
// 	 tx : tx交易信息
// 	 receipt :交易收据信息
// 	 height :交易所在的区块高度
// 	 index :交易所在区块中的索引
// 	 blocktime :交易所在区块的时标
// 	 amount :交易量
// 	 fromaddr :交易打出地址
// 	 txhash : 交易对应的哈希值
// 	 actionName  :交易对应的函数调用
//   payload: 保存额外的一些信息，主要是给插件使用
type WalletTxDetail struct {
	Tx         *Transaction `protobuf:"bytes,1,opt,name=tx" json:"tx,omitempty"`
	Receipt    *ReceiptData `protobuf:"bytes,2,opt,name=receipt" json:"receipt,omitempty"`
	Height     int64        `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	Index      int64        `protobuf:"varint,4,opt,name=index" json:"index,omitempty"`
	Blocktime  int64        `protobuf:"varint,5,opt,name=blocktime" json:"blocktime,omitempty"`
	Amount     int64        `protobuf:"varint,6,opt,name=amount" json:"amount,omitempty"`
	Fromaddr   string       `protobuf:"bytes,7,opt,name=fromaddr" json:"fromaddr,omitempty"`
	Txhash     []byte       `protobuf:"bytes,8,opt,name=txhash,proto3" json:"txhash,omitempty"`
	ActionName string       `protobuf:"bytes,9,opt,name=actionName" json:"actionName,omitempty"`
	Payload    []byte       `protobuf:"bytes,10,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *WalletTxDetail) Reset()                    { *m = WalletTxDetail{} }
func (m *WalletTxDetail) String() string            { return proto.CompactTextString(m) }
func (*WalletTxDetail) ProtoMessage()               {}
func (*WalletTxDetail) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *WalletTxDetail) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *WalletTxDetail) GetReceipt() *ReceiptData {
	if m != nil {
		return m.Receipt
	}
	return nil
}

func (m *WalletTxDetail) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *WalletTxDetail) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *WalletTxDetail) GetBlocktime() int64 {
	if m != nil {
		return m.Blocktime
	}
	return 0
}

func (m *WalletTxDetail) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *WalletTxDetail) GetFromaddr() string {
	if m != nil {
		return m.Fromaddr
	}
	return ""
}

func (m *WalletTxDetail) GetTxhash() []byte {
	if m != nil {
		return m.Txhash
	}
	return nil
}

func (m *WalletTxDetail) GetActionName() string {
	if m != nil {
		return m.ActionName
	}
	return ""
}

func (m *WalletTxDetail) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type WalletTxDetails struct {
	TxDetails []*WalletTxDetail `protobuf:"bytes,1,rep,name=txDetails" json:"txDetails,omitempty"`
}

func (m *WalletTxDetails) Reset()                    { *m = WalletTxDetails{} }
func (m *WalletTxDetails) String() string            { return proto.CompactTextString(m) }
func (*WalletTxDetails) ProtoMessage()               {}
func (*WalletTxDetails) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *WalletTxDetails) GetTxDetails() []*WalletTxDetail {
	if m != nil {
		return m.TxDetails
	}
	return nil
}

// 钱包模块存贮的账户信息
// 	 privkey : 账户地址对应的私钥
// 	 label :账户地址对应的标签
// 	 addr :账户地址
// 	 timeStamp :创建账户时的时标
type WalletAccountStore struct {
	Privkey   string `protobuf:"bytes,1,opt,name=privkey" json:"privkey,omitempty"`
	Label     string `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
	Addr      string `protobuf:"bytes,3,opt,name=addr" json:"addr,omitempty"`
	TimeStamp string `protobuf:"bytes,4,opt,name=timeStamp" json:"timeStamp,omitempty"`
}

func (m *WalletAccountStore) Reset()                    { *m = WalletAccountStore{} }
func (m *WalletAccountStore) String() string            { return proto.CompactTextString(m) }
func (*WalletAccountStore) ProtoMessage()               {}
func (*WalletAccountStore) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

func (m *WalletAccountStore) GetPrivkey() string {
	if m != nil {
		return m.Privkey
	}
	return ""
}

func (m *WalletAccountStore) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *WalletAccountStore) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *WalletAccountStore) GetTimeStamp() string {
	if m != nil {
		return m.TimeStamp
	}
	return ""
}

// 钱包模块通过一个随机值对钱包密码加密
// 	 pwHash : 对钱包密码和一个随机值组合进行哈希计算
// 	 randstr :对钱包密码加密的一个随机值
type WalletPwHash struct {
	PwHash  []byte `protobuf:"bytes,1,opt,name=pwHash,proto3" json:"pwHash,omitempty"`
	Randstr string `protobuf:"bytes,2,opt,name=randstr" json:"randstr,omitempty"`
}

func (m *WalletPwHash) Reset()                    { *m = WalletPwHash{} }
func (m *WalletPwHash) String() string            { return proto.CompactTextString(m) }
func (*WalletPwHash) ProtoMessage()               {}
func (*WalletPwHash) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

func (m *WalletPwHash) GetPwHash() []byte {
	if m != nil {
		return m.PwHash
	}
	return nil
}

func (m *WalletPwHash) GetRandstr() string {
	if m != nil {
		return m.Randstr
	}
	return ""
}

// 钱包当前的状态
// 	 isWalletLock : 钱包是否锁状态，true锁定，false解锁
// 	 isAutoMining :钱包是否开启挖矿功能，true开启挖矿，false关闭挖矿
// 	 isHasSeed : 钱包是否有种子，true已有，false没有
// 	 isTicketLock :钱包挖矿买票锁状态，true锁定，false解锁，只能用于挖矿转账
type WalletStatus struct {
	IsWalletLock bool `protobuf:"varint,1,opt,name=isWalletLock" json:"isWalletLock,omitempty"`
	IsAutoMining bool `protobuf:"varint,2,opt,name=isAutoMining" json:"isAutoMining,omitempty"`
	IsHasSeed    bool `protobuf:"varint,3,opt,name=isHasSeed" json:"isHasSeed,omitempty"`
	IsTicketLock bool `protobuf:"varint,4,opt,name=isTicketLock" json:"isTicketLock,omitempty"`
}

func (m *WalletStatus) Reset()                    { *m = WalletStatus{} }
func (m *WalletStatus) String() string            { return proto.CompactTextString(m) }
func (*WalletStatus) ProtoMessage()               {}
func (*WalletStatus) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{4} }

func (m *WalletStatus) GetIsWalletLock() bool {
	if m != nil {
		return m.IsWalletLock
	}
	return false
}

func (m *WalletStatus) GetIsAutoMining() bool {
	if m != nil {
		return m.IsAutoMining
	}
	return false
}

func (m *WalletStatus) GetIsHasSeed() bool {
	if m != nil {
		return m.IsHasSeed
	}
	return false
}

func (m *WalletStatus) GetIsTicketLock() bool {
	if m != nil {
		return m.IsTicketLock
	}
	return false
}

type WalletAccounts struct {
	Wallets []*WalletAccount `protobuf:"bytes,1,rep,name=wallets" json:"wallets,omitempty"`
}

func (m *WalletAccounts) Reset()                    { *m = WalletAccounts{} }
func (m *WalletAccounts) String() string            { return proto.CompactTextString(m) }
func (*WalletAccounts) ProtoMessage()               {}
func (*WalletAccounts) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{5} }

func (m *WalletAccounts) GetWallets() []*WalletAccount {
	if m != nil {
		return m.Wallets
	}
	return nil
}

type WalletAccount struct {
	Acc   *Account `protobuf:"bytes,1,opt,name=acc" json:"acc,omitempty"`
	Label string   `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
}

func (m *WalletAccount) Reset()                    { *m = WalletAccount{} }
func (m *WalletAccount) String() string            { return proto.CompactTextString(m) }
func (*WalletAccount) ProtoMessage()               {}
func (*WalletAccount) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{6} }

func (m *WalletAccount) GetAcc() *Account {
	if m != nil {
		return m.Acc
	}
	return nil
}

func (m *WalletAccount) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

// 钱包解锁
// 	 passwd : 钱包密码
// 	 timeout :钱包解锁时间，0，一直解锁，非0值，超时之后继续锁定
// 	 walletOrTicket :解锁整个钱包还是只解锁挖矿买票功能，1只解锁挖矿买票，0解锁整个钱包
type WalletUnLock struct {
	Passwd         string `protobuf:"bytes,1,opt,name=passwd" json:"passwd,omitempty"`
	Timeout        int64  `protobuf:"varint,2,opt,name=timeout" json:"timeout,omitempty"`
	WalletOrTicket bool   `protobuf:"varint,3,opt,name=walletOrTicket" json:"walletOrTicket,omitempty"`
}

func (m *WalletUnLock) Reset()                    { *m = WalletUnLock{} }
func (m *WalletUnLock) String() string            { return proto.CompactTextString(m) }
func (*WalletUnLock) ProtoMessage()               {}
func (*WalletUnLock) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{7} }

func (m *WalletUnLock) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

func (m *WalletUnLock) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *WalletUnLock) GetWalletOrTicket() bool {
	if m != nil {
		return m.WalletOrTicket
	}
	return false
}

type GenSeedLang struct {
	Lang int32 `protobuf:"varint,1,opt,name=lang" json:"lang,omitempty"`
}

func (m *GenSeedLang) Reset()                    { *m = GenSeedLang{} }
func (m *GenSeedLang) String() string            { return proto.CompactTextString(m) }
func (*GenSeedLang) ProtoMessage()               {}
func (*GenSeedLang) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{8} }

func (m *GenSeedLang) GetLang() int32 {
	if m != nil {
		return m.Lang
	}
	return 0
}

type GetSeedByPw struct {
	Passwd string `protobuf:"bytes,1,opt,name=passwd" json:"passwd,omitempty"`
}

func (m *GetSeedByPw) Reset()                    { *m = GetSeedByPw{} }
func (m *GetSeedByPw) String() string            { return proto.CompactTextString(m) }
func (*GetSeedByPw) ProtoMessage()               {}
func (*GetSeedByPw) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{9} }

func (m *GetSeedByPw) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

// 存储钱包的种子
// 	 seed : 钱包种子
// 	 passwd :钱包密码
type SaveSeedByPw struct {
	Seed   string `protobuf:"bytes,1,opt,name=seed" json:"seed,omitempty"`
	Passwd string `protobuf:"bytes,2,opt,name=passwd" json:"passwd,omitempty"`
}

func (m *SaveSeedByPw) Reset()                    { *m = SaveSeedByPw{} }
func (m *SaveSeedByPw) String() string            { return proto.CompactTextString(m) }
func (*SaveSeedByPw) ProtoMessage()               {}
func (*SaveSeedByPw) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{10} }

func (m *SaveSeedByPw) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

func (m *SaveSeedByPw) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

type ReplySeed struct {
	Seed string `protobuf:"bytes,1,opt,name=seed" json:"seed,omitempty"`
}

func (m *ReplySeed) Reset()                    { *m = ReplySeed{} }
func (m *ReplySeed) String() string            { return proto.CompactTextString(m) }
func (*ReplySeed) ProtoMessage()               {}
func (*ReplySeed) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{11} }

func (m *ReplySeed) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

type ReqWalletSetPasswd struct {
	OldPass string `protobuf:"bytes,1,opt,name=oldPass" json:"oldPass,omitempty"`
	NewPass string `protobuf:"bytes,2,opt,name=newPass" json:"newPass,omitempty"`
}

func (m *ReqWalletSetPasswd) Reset()                    { *m = ReqWalletSetPasswd{} }
func (m *ReqWalletSetPasswd) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletSetPasswd) ProtoMessage()               {}
func (*ReqWalletSetPasswd) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{12} }

func (m *ReqWalletSetPasswd) GetOldPass() string {
	if m != nil {
		return m.OldPass
	}
	return ""
}

func (m *ReqWalletSetPasswd) GetNewPass() string {
	if m != nil {
		return m.NewPass
	}
	return ""
}

type ReqNewAccount struct {
	Label string `protobuf:"bytes,1,opt,name=label" json:"label,omitempty"`
}

func (m *ReqNewAccount) Reset()                    { *m = ReqNewAccount{} }
func (m *ReqNewAccount) String() string            { return proto.CompactTextString(m) }
func (*ReqNewAccount) ProtoMessage()               {}
func (*ReqNewAccount) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{13} }

func (m *ReqNewAccount) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

// 获取钱包交易的详细信息
// 	 fromTx : []byte( Sprintf("%018d", height*100000 + index)，
// 				表示从高度 height 中的 index 开始获取交易列表；
// 			    第一次传参为空，获取最新的交易。)
// 	 count :获取交易列表的个数。
// 	 direction :查找方式；0，上一页；1，下一页。
type ReqWalletTransactionList struct {
	FromTx    []byte `protobuf:"bytes,1,opt,name=fromTx,proto3" json:"fromTx,omitempty"`
	Count     int32  `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	Direction int32  `protobuf:"varint,3,opt,name=direction" json:"direction,omitempty"`
}

func (m *ReqWalletTransactionList) Reset()                    { *m = ReqWalletTransactionList{} }
func (m *ReqWalletTransactionList) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletTransactionList) ProtoMessage()               {}
func (*ReqWalletTransactionList) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{14} }

func (m *ReqWalletTransactionList) GetFromTx() []byte {
	if m != nil {
		return m.FromTx
	}
	return nil
}

func (m *ReqWalletTransactionList) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReqWalletTransactionList) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

type ReqWalletImportPrivkey struct {
	// bitcoin 的私钥格式
	Privkey string `protobuf:"bytes,1,opt,name=privkey" json:"privkey,omitempty"`
	Label   string `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
}

func (m *ReqWalletImportPrivkey) Reset()                    { *m = ReqWalletImportPrivkey{} }
func (m *ReqWalletImportPrivkey) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletImportPrivkey) ProtoMessage()               {}
func (*ReqWalletImportPrivkey) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{15} }

func (m *ReqWalletImportPrivkey) GetPrivkey() string {
	if m != nil {
		return m.Privkey
	}
	return ""
}

func (m *ReqWalletImportPrivkey) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

// 发送交易
// 	 from : 打出地址
// 	 to :接受地址
// 	 amount : 转账额度
// 	 note :转账备注
type ReqWalletSendToAddress struct {
	From        string `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	To          string `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
	Amount      int64  `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
	Note        []byte `protobuf:"bytes,4,opt,name=note,proto3" json:"note,omitempty"`
	IsToken     bool   `protobuf:"varint,5,opt,name=isToken" json:"isToken,omitempty"`
	TokenSymbol string `protobuf:"bytes,6,opt,name=tokenSymbol" json:"tokenSymbol,omitempty"`
}

func (m *ReqWalletSendToAddress) Reset()                    { *m = ReqWalletSendToAddress{} }
func (m *ReqWalletSendToAddress) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletSendToAddress) ProtoMessage()               {}
func (*ReqWalletSendToAddress) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{16} }

func (m *ReqWalletSendToAddress) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *ReqWalletSendToAddress) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *ReqWalletSendToAddress) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ReqWalletSendToAddress) GetNote() []byte {
	if m != nil {
		return m.Note
	}
	return nil
}

func (m *ReqWalletSendToAddress) GetIsToken() bool {
	if m != nil {
		return m.IsToken
	}
	return false
}

func (m *ReqWalletSendToAddress) GetTokenSymbol() string {
	if m != nil {
		return m.TokenSymbol
	}
	return ""
}

type ReqWalletSetFee struct {
	Amount int64 `protobuf:"varint,1,opt,name=amount" json:"amount,omitempty"`
}

func (m *ReqWalletSetFee) Reset()                    { *m = ReqWalletSetFee{} }
func (m *ReqWalletSetFee) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletSetFee) ProtoMessage()               {}
func (*ReqWalletSetFee) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{17} }

func (m *ReqWalletSetFee) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type ReqWalletSetLabel struct {
	Addr  string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
}

func (m *ReqWalletSetLabel) Reset()                    { *m = ReqWalletSetLabel{} }
func (m *ReqWalletSetLabel) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletSetLabel) ProtoMessage()               {}
func (*ReqWalletSetLabel) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{18} }

func (m *ReqWalletSetLabel) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *ReqWalletSetLabel) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type ReqWalletMergeBalance struct {
	To string `protobuf:"bytes,1,opt,name=to" json:"to,omitempty"`
}

func (m *ReqWalletMergeBalance) Reset()                    { *m = ReqWalletMergeBalance{} }
func (m *ReqWalletMergeBalance) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletMergeBalance) ProtoMessage()               {}
func (*ReqWalletMergeBalance) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{19} }

func (m *ReqWalletMergeBalance) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

type ReqTokenPreCreate struct {
	CreatorAddr  string `protobuf:"bytes,1,opt,name=creator_addr,json=creatorAddr" json:"creator_addr,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Symbol       string `protobuf:"bytes,3,opt,name=symbol" json:"symbol,omitempty"`
	Introduction string `protobuf:"bytes,4,opt,name=introduction" json:"introduction,omitempty"`
	OwnerAddr    string `protobuf:"bytes,5,opt,name=owner_addr,json=ownerAddr" json:"owner_addr,omitempty"`
	Total        int64  `protobuf:"varint,6,opt,name=total" json:"total,omitempty"`
	Price        int64  `protobuf:"varint,7,opt,name=price" json:"price,omitempty"`
}

func (m *ReqTokenPreCreate) Reset()                    { *m = ReqTokenPreCreate{} }
func (m *ReqTokenPreCreate) String() string            { return proto.CompactTextString(m) }
func (*ReqTokenPreCreate) ProtoMessage()               {}
func (*ReqTokenPreCreate) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{20} }

func (m *ReqTokenPreCreate) GetCreatorAddr() string {
	if m != nil {
		return m.CreatorAddr
	}
	return ""
}

func (m *ReqTokenPreCreate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqTokenPreCreate) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ReqTokenPreCreate) GetIntroduction() string {
	if m != nil {
		return m.Introduction
	}
	return ""
}

func (m *ReqTokenPreCreate) GetOwnerAddr() string {
	if m != nil {
		return m.OwnerAddr
	}
	return ""
}

func (m *ReqTokenPreCreate) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReqTokenPreCreate) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

type ReqTokenFinishCreate struct {
	FinisherAddr string `protobuf:"bytes,1,opt,name=finisher_addr,json=finisherAddr" json:"finisher_addr,omitempty"`
	Symbol       string `protobuf:"bytes,2,opt,name=symbol" json:"symbol,omitempty"`
	OwnerAddr    string `protobuf:"bytes,3,opt,name=owner_addr,json=ownerAddr" json:"owner_addr,omitempty"`
}

func (m *ReqTokenFinishCreate) Reset()                    { *m = ReqTokenFinishCreate{} }
func (m *ReqTokenFinishCreate) String() string            { return proto.CompactTextString(m) }
func (*ReqTokenFinishCreate) ProtoMessage()               {}
func (*ReqTokenFinishCreate) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{21} }

func (m *ReqTokenFinishCreate) GetFinisherAddr() string {
	if m != nil {
		return m.FinisherAddr
	}
	return ""
}

func (m *ReqTokenFinishCreate) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ReqTokenFinishCreate) GetOwnerAddr() string {
	if m != nil {
		return m.OwnerAddr
	}
	return ""
}

type ReqTokenRevokeCreate struct {
	RevokerAddr string `protobuf:"bytes,1,opt,name=revoker_addr,json=revokerAddr" json:"revoker_addr,omitempty"`
	Symbol      string `protobuf:"bytes,2,opt,name=symbol" json:"symbol,omitempty"`
	OwnerAddr   string `protobuf:"bytes,3,opt,name=owner_addr,json=ownerAddr" json:"owner_addr,omitempty"`
}

func (m *ReqTokenRevokeCreate) Reset()                    { *m = ReqTokenRevokeCreate{} }
func (m *ReqTokenRevokeCreate) String() string            { return proto.CompactTextString(m) }
func (*ReqTokenRevokeCreate) ProtoMessage()               {}
func (*ReqTokenRevokeCreate) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{22} }

func (m *ReqTokenRevokeCreate) GetRevokerAddr() string {
	if m != nil {
		return m.RevokerAddr
	}
	return ""
}

func (m *ReqTokenRevokeCreate) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ReqTokenRevokeCreate) GetOwnerAddr() string {
	if m != nil {
		return m.OwnerAddr
	}
	return ""
}

type ReqModifyConfig struct {
	Key      string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Op       string `protobuf:"bytes,2,opt,name=op" json:"op,omitempty"`
	Value    string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	Modifier string `protobuf:"bytes,4,opt,name=modifier" json:"modifier,omitempty"`
}

func (m *ReqModifyConfig) Reset()                    { *m = ReqModifyConfig{} }
func (m *ReqModifyConfig) String() string            { return proto.CompactTextString(m) }
func (*ReqModifyConfig) ProtoMessage()               {}
func (*ReqModifyConfig) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{23} }

func (m *ReqModifyConfig) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReqModifyConfig) GetOp() string {
	if m != nil {
		return m.Op
	}
	return ""
}

func (m *ReqModifyConfig) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ReqModifyConfig) GetModifier() string {
	if m != nil {
		return m.Modifier
	}
	return ""
}

type ReqSignRawTx struct {
	Addr    string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Privkey string `protobuf:"bytes,2,opt,name=privkey" json:"privkey,omitempty"`
	TxHex   string `protobuf:"bytes,3,opt,name=txHex" json:"txHex,omitempty"`
	Expire  string `protobuf:"bytes,4,opt,name=expire" json:"expire,omitempty"`
	Index   int32  `protobuf:"varint,5,opt,name=index" json:"index,omitempty"`
	// 签名的模式类型
	// 0：普通交易
	// 1：隐私交易
	// int32  mode  = 6;
	Token string `protobuf:"bytes,7,opt,name=token" json:"token,omitempty"`
}

func (m *ReqSignRawTx) Reset()                    { *m = ReqSignRawTx{} }
func (m *ReqSignRawTx) String() string            { return proto.CompactTextString(m) }
func (*ReqSignRawTx) ProtoMessage()               {}
func (*ReqSignRawTx) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{24} }

func (m *ReqSignRawTx) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *ReqSignRawTx) GetPrivkey() string {
	if m != nil {
		return m.Privkey
	}
	return ""
}

func (m *ReqSignRawTx) GetTxHex() string {
	if m != nil {
		return m.TxHex
	}
	return ""
}

func (m *ReqSignRawTx) GetExpire() string {
	if m != nil {
		return m.Expire
	}
	return ""
}

func (m *ReqSignRawTx) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ReqSignRawTx) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ReplySignRawTx struct {
	TxHex string `protobuf:"bytes,1,opt,name=txHex" json:"txHex,omitempty"`
}

func (m *ReplySignRawTx) Reset()                    { *m = ReplySignRawTx{} }
func (m *ReplySignRawTx) String() string            { return proto.CompactTextString(m) }
func (*ReplySignRawTx) ProtoMessage()               {}
func (*ReplySignRawTx) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{25} }

func (m *ReplySignRawTx) GetTxHex() string {
	if m != nil {
		return m.TxHex
	}
	return ""
}

type ReportErrEvent struct {
	Frommodule string `protobuf:"bytes,1,opt,name=frommodule" json:"frommodule,omitempty"`
	Tomodule   string `protobuf:"bytes,2,opt,name=tomodule" json:"tomodule,omitempty"`
	Error      string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *ReportErrEvent) Reset()                    { *m = ReportErrEvent{} }
func (m *ReportErrEvent) String() string            { return proto.CompactTextString(m) }
func (*ReportErrEvent) ProtoMessage()               {}
func (*ReportErrEvent) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{26} }

func (m *ReportErrEvent) GetFrommodule() string {
	if m != nil {
		return m.Frommodule
	}
	return ""
}

func (m *ReportErrEvent) GetTomodule() string {
	if m != nil {
		return m.Tomodule
	}
	return ""
}

func (m *ReportErrEvent) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Int32 struct {
	Data int32 `protobuf:"varint,1,opt,name=data" json:"data,omitempty"`
}

func (m *Int32) Reset()                    { *m = Int32{} }
func (m *Int32) String() string            { return proto.CompactTextString(m) }
func (*Int32) ProtoMessage()               {}
func (*Int32) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{27} }

func (m *Int32) GetData() int32 {
	if m != nil {
		return m.Data
	}
	return 0
}

// 某些交易需要存在一些复杂的算法，所以需要请求服务器协助构建交易，返回对应交易哈希值，后续签名可以根据此哈希值进行处理
type ReqCreateTransaction struct {
	Tokenname string `protobuf:"bytes,1,opt,name=tokenname" json:"tokenname,omitempty"`
	// 构建交易类型
	// 0：普通的交易（暂不支持）
	// 1：隐私交易 公开->隐私
	// 2：隐私交易 隐私->隐私
	// 3：隐私交易 隐私->公开
	Type   int32  `protobuf:"varint,2,opt,name=type" json:"type,omitempty"`
	Amount int64  `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
	Note   []byte `protobuf:"bytes,4,opt,name=note,proto3" json:"note,omitempty"`
	// 普通交易的发送方
	From string `protobuf:"bytes,5,opt,name=from" json:"from,omitempty"`
	// 普通交易的接收方
	To string `protobuf:"bytes,6,opt,name=to" json:"to,omitempty"`
	// 隐私交易，接收方的公钥对
	Pubkeypair string `protobuf:"bytes,10,opt,name=pubkeypair" json:"pubkeypair,omitempty"`
	Mixcount   int32  `protobuf:"varint,11,opt,name=mixcount" json:"mixcount,omitempty"`
	Expire     int64  `protobuf:"varint,12,opt,name=expire" json:"expire,omitempty"`
}

func (m *ReqCreateTransaction) Reset()                    { *m = ReqCreateTransaction{} }
func (m *ReqCreateTransaction) String() string            { return proto.CompactTextString(m) }
func (*ReqCreateTransaction) ProtoMessage()               {}
func (*ReqCreateTransaction) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{28} }

func (m *ReqCreateTransaction) GetTokenname() string {
	if m != nil {
		return m.Tokenname
	}
	return ""
}

func (m *ReqCreateTransaction) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqCreateTransaction) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ReqCreateTransaction) GetNote() []byte {
	if m != nil {
		return m.Note
	}
	return nil
}

func (m *ReqCreateTransaction) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *ReqCreateTransaction) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *ReqCreateTransaction) GetPubkeypair() string {
	if m != nil {
		return m.Pubkeypair
	}
	return ""
}

func (m *ReqCreateTransaction) GetMixcount() int32 {
	if m != nil {
		return m.Mixcount
	}
	return 0
}

func (m *ReqCreateTransaction) GetExpire() int64 {
	if m != nil {
		return m.Expire
	}
	return 0
}

type ReqAccountList struct {
	WithoutBalance bool `protobuf:"varint,1,opt,name=withoutBalance" json:"withoutBalance,omitempty"`
}

func (m *ReqAccountList) Reset()                    { *m = ReqAccountList{} }
func (m *ReqAccountList) String() string            { return proto.CompactTextString(m) }
func (*ReqAccountList) ProtoMessage()               {}
func (*ReqAccountList) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{29} }

func (m *ReqAccountList) GetWithoutBalance() bool {
	if m != nil {
		return m.WithoutBalance
	}
	return false
}

func init() {
	proto.RegisterType((*WalletTxDetail)(nil), "types.WalletTxDetail")
	proto.RegisterType((*WalletTxDetails)(nil), "types.WalletTxDetails")
	proto.RegisterType((*WalletAccountStore)(nil), "types.WalletAccountStore")
	proto.RegisterType((*WalletPwHash)(nil), "types.WalletPwHash")
	proto.RegisterType((*WalletStatus)(nil), "types.WalletStatus")
	proto.RegisterType((*WalletAccounts)(nil), "types.WalletAccounts")
	proto.RegisterType((*WalletAccount)(nil), "types.WalletAccount")
	proto.RegisterType((*WalletUnLock)(nil), "types.WalletUnLock")
	proto.RegisterType((*GenSeedLang)(nil), "types.GenSeedLang")
	proto.RegisterType((*GetSeedByPw)(nil), "types.GetSeedByPw")
	proto.RegisterType((*SaveSeedByPw)(nil), "types.SaveSeedByPw")
	proto.RegisterType((*ReplySeed)(nil), "types.ReplySeed")
	proto.RegisterType((*ReqWalletSetPasswd)(nil), "types.ReqWalletSetPasswd")
	proto.RegisterType((*ReqNewAccount)(nil), "types.ReqNewAccount")
	proto.RegisterType((*ReqWalletTransactionList)(nil), "types.ReqWalletTransactionList")
	proto.RegisterType((*ReqWalletImportPrivkey)(nil), "types.ReqWalletImportPrivkey")
	proto.RegisterType((*ReqWalletSendToAddress)(nil), "types.ReqWalletSendToAddress")
	proto.RegisterType((*ReqWalletSetFee)(nil), "types.ReqWalletSetFee")
	proto.RegisterType((*ReqWalletSetLabel)(nil), "types.ReqWalletSetLabel")
	proto.RegisterType((*ReqWalletMergeBalance)(nil), "types.ReqWalletMergeBalance")
	proto.RegisterType((*ReqTokenPreCreate)(nil), "types.ReqTokenPreCreate")
	proto.RegisterType((*ReqTokenFinishCreate)(nil), "types.ReqTokenFinishCreate")
	proto.RegisterType((*ReqTokenRevokeCreate)(nil), "types.ReqTokenRevokeCreate")
	proto.RegisterType((*ReqModifyConfig)(nil), "types.ReqModifyConfig")
	proto.RegisterType((*ReqSignRawTx)(nil), "types.ReqSignRawTx")
	proto.RegisterType((*ReplySignRawTx)(nil), "types.ReplySignRawTx")
	proto.RegisterType((*ReportErrEvent)(nil), "types.ReportErrEvent")
	proto.RegisterType((*Int32)(nil), "types.Int32")
	proto.RegisterType((*ReqCreateTransaction)(nil), "types.ReqCreateTransaction")
	proto.RegisterType((*ReqAccountList)(nil), "types.ReqAccountList")
}

func init() { proto.RegisterFile("wallet.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 1271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x5d, 0x6e, 0x1b, 0xb7,
	0x13, 0xc7, 0x4a, 0x96, 0x6d, 0xd1, 0xb2, 0x93, 0x10, 0x49, 0xb0, 0xf0, 0xff, 0x9f, 0x44, 0x61,
	0x91, 0xd4, 0x05, 0x0a, 0x07, 0x88, 0x5e, 0x8a, 0x02, 0x05, 0xe2, 0x7c, 0x3a, 0x80, 0x93, 0x1a,
	0x94, 0x8a, 0x02, 0x7d, 0x29, 0xa8, 0x5d, 0x5a, 0x22, 0xb4, 0x5a, 0xae, 0xb9, 0xd4, 0xd7, 0x45,
	0x8a, 0x1e, 0xa0, 0x47, 0xe8, 0x45, 0x7a, 0x8f, 0x1e, 0xa2, 0x98, 0x21, 0xb9, 0xda, 0x4d, 0xdc,
	0x87, 0xa0, 0x6f, 0xf3, 0x1b, 0x0e, 0x67, 0x38, 0xbf, 0x19, 0x0e, 0x49, 0x7a, 0x2b, 0x91, 0x65,
	0xd2, 0x9e, 0x16, 0x46, 0x5b, 0x4d, 0x3b, 0x76, 0x53, 0xc8, 0xf2, 0xf8, 0x8e, 0x35, 0x22, 0x2f,
	0x45, 0x62, 0x95, 0xce, 0xdd, 0xca, 0xf1, 0xed, 0x71, 0xa6, 0x93, 0x59, 0x32, 0x15, 0x2a, 0x68,
	0x0e, 0x45, 0x92, 0xe8, 0x45, 0xee, 0xb7, 0x1e, 0x1f, 0xc9, 0xb5, 0x4c, 0x16, 0x56, 0x1b, 0x87,
	0xd9, 0x9f, 0x2d, 0x72, 0xf4, 0x33, 0xfa, 0x1e, 0xad, 0x5f, 0x4b, 0x2b, 0x54, 0x46, 0x19, 0x69,
	0xd9, 0x75, 0x1c, 0xf5, 0xa3, 0x93, 0x83, 0xe7, 0xf4, 0x14, 0x43, 0x9d, 0x8e, 0xb6, 0x91, 0x78,
	0xcb, 0xae, 0xe9, 0xb7, 0x64, 0xcf, 0xc8, 0x44, 0xaa, 0xc2, 0xc6, 0xad, 0x86, 0x21, 0x77, 0xda,
	0xd7, 0xc2, 0x0a, 0x1e, 0x4c, 0xe8, 0x7d, 0xb2, 0x3b, 0x95, 0x6a, 0x32, 0xb5, 0x71, 0xbb, 0x1f,
	0x9d, 0xb4, 0xb9, 0x47, 0xf4, 0x2e, 0xe9, 0xa8, 0x3c, 0x95, 0xeb, 0x78, 0x07, 0xd5, 0x0e, 0xd0,
	0xff, 0x93, 0x2e, 0x66, 0x61, 0xd5, 0x5c, 0xc6, 0x1d, 0x5c, 0xd9, 0x2a, 0xc0, 0x97, 0x98, 0x43,
	0x42, 0xf1, 0xae, 0xf3, 0xe5, 0x10, 0x3d, 0x26, 0xfb, 0x57, 0x46, 0xcf, 0x45, 0x9a, 0x9a, 0x78,
	0xaf, 0x1f, 0x9d, 0x74, 0x79, 0x85, 0x61, 0x8f, 0x5d, 0x4f, 0x45, 0x39, 0x8d, 0xf7, 0xfb, 0xd1,
	0x49, 0x8f, 0x7b, 0x44, 0x1f, 0x12, 0xe2, 0x72, 0xfa, 0x28, 0xe6, 0x32, 0xee, 0xe2, 0xae, 0x9a,
	0x86, 0xc6, 0x64, 0xaf, 0x10, 0x9b, 0x4c, 0x8b, 0x34, 0x26, 0xb8, 0x31, 0x40, 0xf6, 0x96, 0xdc,
	0x6a, 0xb2, 0x56, 0xd2, 0x01, 0xe9, 0xda, 0x00, 0xe2, 0xa8, 0xdf, 0x3e, 0x39, 0x78, 0x7e, 0xcf,
	0x93, 0xd2, 0x34, 0xe5, 0x5b, 0x3b, 0xb6, 0x24, 0xd4, 0x2d, 0x9e, 0xb9, 0x2a, 0x0d, 0xad, 0x36,
	0x2e, 0xae, 0x51, 0xcb, 0x99, 0xdc, 0x60, 0x19, 0xba, 0x3c, 0x40, 0x60, 0x2c, 0x13, 0x63, 0x99,
	0x21, 0xeb, 0x5d, 0xee, 0x00, 0xa5, 0x64, 0x07, 0xf3, 0x6e, 0xa3, 0x12, 0x65, 0x60, 0x11, 0xf8,
	0x1a, 0x5a, 0x31, 0x2f, 0x90, 0xdf, 0x2e, 0xdf, 0x2a, 0xd8, 0x0b, 0xd2, 0x73, 0x71, 0x2f, 0x57,
	0xe7, 0xc0, 0xc4, 0x7d, 0xb2, 0x5b, 0xa0, 0x84, 0x01, 0x7b, 0xdc, 0x23, 0x38, 0x89, 0x11, 0x79,
	0x5a, 0x5a, 0xe3, 0x23, 0x06, 0xc8, 0x7e, 0x8f, 0x82, 0x8b, 0xa1, 0x15, 0x76, 0x51, 0x52, 0x46,
	0x7a, 0xaa, 0x74, 0x9a, 0x0b, 0x9d, 0xcc, 0xd0, 0xd1, 0x3e, 0x6f, 0xe8, 0x9c, 0xcd, 0xd9, 0xc2,
	0xea, 0x0f, 0x2a, 0x57, 0xf9, 0x04, 0x7d, 0xa2, 0xcd, 0x56, 0x07, 0x07, 0x57, 0xe5, 0xb9, 0x28,
	0x87, 0x52, 0xa6, 0x98, 0xd1, 0x3e, 0xdf, 0x2a, 0x9c, 0x87, 0x91, 0x4a, 0x66, 0x3e, 0xca, 0x4e,
	0xf0, 0xb0, 0xd5, 0xb1, 0x17, 0xa1, 0xa5, 0x3d, 0xa9, 0x25, 0x3d, 0x25, 0x7b, 0xee, 0x02, 0x85,
	0xca, 0xdc, 0x6d, 0x54, 0xc6, 0xdb, 0xf1, 0x60, 0xc4, 0xde, 0x91, 0xc3, 0xc6, 0x0a, 0xed, 0x93,
	0xb6, 0x48, 0x12, 0x7f, 0x29, 0x8e, 0xfc, 0xe6, 0xb0, 0x0d, 0x96, 0x6e, 0xae, 0x0c, 0x9b, 0x06,
	0x92, 0x7e, 0xca, 0x91, 0x00, 0xe0, 0x59, 0x94, 0xe5, 0x2a, 0xf5, 0x85, 0xf5, 0x08, 0x78, 0x86,
	0xe2, 0xe8, 0x85, 0xbb, 0x4f, 0x6d, 0x1e, 0x20, 0x7d, 0x4a, 0x8e, 0xdc, 0xa9, 0x7e, 0x34, 0x2e,
	0x45, 0xcf, 0xc9, 0x27, 0x5a, 0xf6, 0x98, 0x1c, 0xbc, 0x93, 0x39, 0x70, 0x74, 0x21, 0xf2, 0x09,
	0xb4, 0x44, 0x26, 0xf2, 0x09, 0x86, 0xe9, 0x70, 0x94, 0xd9, 0x13, 0x30, 0xb1, 0x60, 0xf2, 0x72,
	0x73, 0xb9, 0xfa, 0xb7, 0xb3, 0xb0, 0xef, 0x49, 0x6f, 0x28, 0x96, 0xb2, 0xb2, 0xa3, 0x64, 0xa7,
	0x84, 0x5a, 0x38, 0x2b, 0x94, 0x6b, 0x7b, 0x5b, 0x8d, 0xbd, 0x8f, 0x48, 0x97, 0xcb, 0x22, 0xdb,
	0x60, 0xad, 0x6e, 0xd8, 0xc8, 0xce, 0x09, 0xe5, 0xf2, 0xda, 0x37, 0x8e, 0xb4, 0x97, 0x55, 0xfa,
	0x3a, 0x4b, 0x01, 0x84, 0x86, 0xf7, 0x10, 0x56, 0x72, 0xb9, 0xc2, 0x15, 0xdf, 0x80, 0x1e, 0xb2,
	0x27, 0xe4, 0x90, 0xcb, 0xeb, 0x8f, 0x72, 0x15, 0x6a, 0x54, 0x55, 0x20, 0xaa, 0x57, 0xe0, 0x8a,
	0xc4, 0x55, 0xc0, 0xda, 0x14, 0xbb, 0x50, 0x25, 0xce, 0x25, 0x98, 0x11, 0xa3, 0x75, 0xe8, 0x7a,
	0x87, 0xc0, 0x13, 0xba, 0xc4, 0x90, 0x1d, 0xee, 0x00, 0x34, 0x66, 0xaa, 0x8c, 0xc4, 0xed, 0x58,
	0x84, 0x0e, 0xdf, 0x2a, 0xd8, 0x39, 0xb9, 0x5f, 0xc5, 0x79, 0x3f, 0x2f, 0xb4, 0xb1, 0x97, 0xfe,
	0xce, 0x7e, 0xe1, 0x6d, 0x66, 0x7f, 0x44, 0x35, 0x57, 0x43, 0x99, 0xa7, 0x23, 0x7d, 0x96, 0xa6,
	0x46, 0x96, 0x25, 0x30, 0x0a, 0x47, 0x0c, 0x8c, 0x82, 0x4c, 0x8f, 0x48, 0xcb, 0x6a, 0xef, 0xa1,
	0x65, 0x75, 0x6d, 0x40, 0xb6, 0x1b, 0x03, 0x92, 0x92, 0x9d, 0x5c, 0x5b, 0x89, 0x37, 0xa6, 0xc7,
	0x51, 0x86, 0xa3, 0xa9, 0x72, 0xa4, 0x67, 0x32, 0xc7, 0x41, 0xbb, 0xcf, 0x03, 0xa4, 0x7d, 0x72,
	0x60, 0x41, 0x18, 0x6e, 0xe6, 0x63, 0x9d, 0xe1, 0xac, 0xed, 0xf2, 0xba, 0x8a, 0x7d, 0x43, 0x6e,
	0xd5, 0x2b, 0xf9, 0x56, 0xd6, 0x67, 0x73, 0x54, 0x0f, 0xcd, 0x7e, 0x20, 0x77, 0xea, 0xa6, 0x17,
	0x8d, 0xa1, 0x15, 0xd5, 0x86, 0xd6, 0xcd, 0x84, 0x7c, 0x4d, 0xee, 0x55, 0xdb, 0x3f, 0x48, 0x33,
	0x91, 0x2f, 0x45, 0x26, 0xf2, 0x44, 0xfa, 0xd4, 0xa3, 0x90, 0x3a, 0xfb, 0x2b, 0xc2, 0x40, 0x98,
	0xc1, 0xa5, 0x91, 0xaf, 0x8c, 0x14, 0x56, 0xd2, 0xc7, 0xa4, 0x97, 0x80, 0xa4, 0xcd, 0xaf, 0xb5,
	0x80, 0x07, 0x5e, 0x07, 0xd4, 0x22, 0x37, 0xf0, 0x04, 0xb8, 0xb0, 0x28, 0x43, 0x32, 0xa5, 0x4b,
	0xde, 0x8d, 0x55, 0x8f, 0x70, 0x02, 0xe5, 0xd6, 0xe8, 0x74, 0xe1, 0x3a, 0xc1, 0xcd, 0xd6, 0x86,
	0x8e, 0x3e, 0x20, 0x44, 0xaf, 0x72, 0xe9, 0x03, 0x76, 0xdc, 0xf4, 0x45, 0xcd, 0x99, 0x4f, 0xd3,
	0x6a, 0x2b, 0x32, 0xff, 0x84, 0x39, 0x00, 0xda, 0xc2, 0xa8, 0x44, 0xe2, 0xf3, 0xd5, 0xe6, 0x0e,
	0x30, 0x43, 0xee, 0x86, 0x94, 0xde, 0xaa, 0x5c, 0x95, 0x53, 0x9f, 0xd5, 0x57, 0xe4, 0xf0, 0x0a,
	0xb1, 0x6c, 0xa4, 0xd5, 0x0b, 0xca, 0x33, 0xff, 0xf0, 0xf9, 0x1c, 0x5a, 0x8d, 0x1c, 0x9a, 0xe7,
	0x6b, 0x7f, 0x72, 0x3e, 0x56, 0x6c, 0x63, 0x72, 0xb9, 0xd4, 0xb3, 0x1a, 0x93, 0x06, 0x71, 0x93,
	0x49, 0xaf, 0xfb, 0x2f, 0x11, 0x25, 0x36, 0xd3, 0x07, 0x9d, 0xaa, 0xab, 0xcd, 0x2b, 0x9d, 0x5f,
	0xa9, 0x09, 0xbd, 0x4d, 0xda, 0xdb, 0x2b, 0x03, 0x22, 0x94, 0x5b, 0x17, 0xa1, 0xd3, 0x75, 0x01,
	0x84, 0x2d, 0x45, 0xb6, 0x90, 0xde, 0x9d, 0x03, 0xf0, 0x11, 0x98, 0x83, 0x1f, 0x25, 0x8d, 0xaf,
	0x4d, 0x85, 0xd9, 0x6f, 0x11, 0xe9, 0x71, 0x79, 0x3d, 0x54, 0x93, 0x9c, 0x8b, 0xd5, 0x68, 0x7d,
	0x63, 0x13, 0xd6, 0xee, 0x6b, 0xeb, 0xb3, 0xfb, 0x6a, 0xd7, 0xe7, 0x72, 0x1d, 0x02, 0x22, 0x80,
	0x94, 0xe5, 0xba, 0x50, 0x46, 0xfa, 0x70, 0x1e, 0x6d, 0x7f, 0x37, 0x1d, 0x37, 0x45, 0xdc, 0xef,
	0x06, 0x6b, 0x0f, 0x17, 0x6e, 0xcf, 0xfb, 0x00, 0xc0, 0x9e, 0x92, 0x23, 0x37, 0x37, 0xab, 0x93,
	0x55, 0xb1, 0xa2, 0x5a, 0x2c, 0x36, 0x46, 0x3b, 0x6d, 0xec, 0x1b, 0x63, 0xde, 0x2c, 0x65, 0x6e,
	0xe1, 0x0f, 0x03, 0x63, 0x60, 0xae, 0xd3, 0x45, 0x26, 0xbd, 0x71, 0x4d, 0x03, 0x74, 0x58, 0xed,
	0x57, 0x5d, 0x3a, 0x15, 0x86, 0x18, 0xd2, 0x18, 0x1d, 0xea, 0xe1, 0x00, 0xfb, 0x1f, 0xe9, 0xbc,
	0xcf, 0xed, 0xe0, 0x39, 0x90, 0x93, 0x0a, 0x2b, 0xc2, 0x1b, 0x02, 0x32, 0xfb, 0x3b, 0xc2, 0xde,
	0x70, 0x0d, 0x51, 0x9b, 0xa7, 0xf8, 0xdf, 0x80, 0x54, 0xf0, 0x1e, 0x45, 0xfe, 0xbf, 0x11, 0x14,
	0xe0, 0x0a, 0xde, 0x4c, 0x3f, 0x50, 0x51, 0xfe, 0xa2, 0x41, 0x15, 0x06, 0x5f, 0xe7, 0xb3, 0xc1,
	0xb7, 0x5b, 0x0d, 0xbe, 0x87, 0x84, 0x14, 0x8b, 0xf1, 0x4c, 0x6e, 0x0a, 0xa1, 0x0c, 0x7e, 0xd8,
	0xba, 0xbc, 0xa6, 0xc1, 0xc6, 0x50, 0x6b, 0x37, 0xd8, 0x0f, 0xf0, 0x1c, 0x15, 0xae, 0xd5, 0xb0,
	0xe7, 0xce, 0xe2, 0x10, 0xfb, 0x0e, 0xf8, 0xbe, 0xf6, 0x2f, 0x0c, 0xbe, 0x19, 0xf0, 0x1e, 0x2b,
	0x3b, 0xd5, 0x0b, 0xeb, 0xa7, 0x90, 0xff, 0xe8, 0x7c, 0xa2, 0x7d, 0xf9, 0xe8, 0x97, 0x07, 0x13,
	0x65, 0xa7, 0x8b, 0xf1, 0x69, 0xa2, 0xe7, 0xcf, 0x06, 0x83, 0x24, 0x7f, 0x86, 0xdf, 0xf2, 0xc1,
	0xe0, 0x19, 0xfe, 0x1e, 0xc6, 0xbb, 0xf8, 0x01, 0x1f, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0xe6,
	0x69, 0x17, 0x4b, 0xdb, 0x0b, 0x00, 0x00,
}
