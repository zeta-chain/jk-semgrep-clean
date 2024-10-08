// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	btcjson "github.com/btcsuite/btcd/btcjson"
	btcutil "github.com/btcsuite/btcd/btcutil"

	chainhash "github.com/btcsuite/btcd/chaincfg/chainhash"

	mock "github.com/stretchr/testify/mock"

	rpcclient "github.com/btcsuite/btcd/rpcclient"

	wire "github.com/btcsuite/btcd/wire"
)

// BTCRPCClient is an autogenerated mock type for the BTCRPCClient type
type BTCRPCClient struct {
	mock.Mock
}

// CreateWallet provides a mock function with given fields: name, opts
func (_m *BTCRPCClient) CreateWallet(name string, opts ...rpcclient.CreateWalletOpt) (*btcjson.CreateWalletResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, name)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateWallet")
	}

	var r0 *btcjson.CreateWalletResult
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...rpcclient.CreateWalletOpt) (*btcjson.CreateWalletResult, error)); ok {
		return rf(name, opts...)
	}
	if rf, ok := ret.Get(0).(func(string, ...rpcclient.CreateWalletOpt) *btcjson.CreateWalletResult); ok {
		r0 = rf(name, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*btcjson.CreateWalletResult)
		}
	}

	if rf, ok := ret.Get(1).(func(string, ...rpcclient.CreateWalletOpt) error); ok {
		r1 = rf(name, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EstimateSmartFee provides a mock function with given fields: confTarget, mode
func (_m *BTCRPCClient) EstimateSmartFee(confTarget int64, mode *btcjson.EstimateSmartFeeMode) (*btcjson.EstimateSmartFeeResult, error) {
	ret := _m.Called(confTarget, mode)

	if len(ret) == 0 {
		panic("no return value specified for EstimateSmartFee")
	}

	var r0 *btcjson.EstimateSmartFeeResult
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, *btcjson.EstimateSmartFeeMode) (*btcjson.EstimateSmartFeeResult, error)); ok {
		return rf(confTarget, mode)
	}
	if rf, ok := ret.Get(0).(func(int64, *btcjson.EstimateSmartFeeMode) *btcjson.EstimateSmartFeeResult); ok {
		r0 = rf(confTarget, mode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*btcjson.EstimateSmartFeeResult)
		}
	}

	if rf, ok := ret.Get(1).(func(int64, *btcjson.EstimateSmartFeeMode) error); ok {
		r1 = rf(confTarget, mode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateToAddress provides a mock function with given fields: numBlocks, address, maxTries
func (_m *BTCRPCClient) GenerateToAddress(numBlocks int64, address btcutil.Address, maxTries *int64) ([]*chainhash.Hash, error) {
	ret := _m.Called(numBlocks, address, maxTries)

	if len(ret) == 0 {
		panic("no return value specified for GenerateToAddress")
	}

	var r0 []*chainhash.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, btcutil.Address, *int64) ([]*chainhash.Hash, error)); ok {
		return rf(numBlocks, address, maxTries)
	}
	if rf, ok := ret.Get(0).(func(int64, btcutil.Address, *int64) []*chainhash.Hash); ok {
		r0 = rf(numBlocks, address, maxTries)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*chainhash.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(int64, btcutil.Address, *int64) error); ok {
		r1 = rf(numBlocks, address, maxTries)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBalance provides a mock function with given fields: account
func (_m *BTCRPCClient) GetBalance(account string) (btcutil.Amount, error) {
	ret := _m.Called(account)

	if len(ret) == 0 {
		panic("no return value specified for GetBalance")
	}

	var r0 btcutil.Amount
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (btcutil.Amount, error)); ok {
		return rf(account)
	}
	if rf, ok := ret.Get(0).(func(string) btcutil.Amount); ok {
		r0 = rf(account)
	} else {
		r0 = ret.Get(0).(btcutil.Amount)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockCount provides a mock function with given fields:
func (_m *BTCRPCClient) GetBlockCount() (int64, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetBlockCount")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func() (int64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockHash provides a mock function with given fields: blockHeight
func (_m *BTCRPCClient) GetBlockHash(blockHeight int64) (*chainhash.Hash, error) {
	ret := _m.Called(blockHeight)

	if len(ret) == 0 {
		panic("no return value specified for GetBlockHash")
	}

	var r0 *chainhash.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*chainhash.Hash, error)); ok {
		return rf(blockHeight)
	}
	if rf, ok := ret.Get(0).(func(int64) *chainhash.Hash); ok {
		r0 = rf(blockHeight)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chainhash.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(blockHeight)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockHeader provides a mock function with given fields: blockHash
func (_m *BTCRPCClient) GetBlockHeader(blockHash *chainhash.Hash) (*wire.BlockHeader, error) {
	ret := _m.Called(blockHash)

	if len(ret) == 0 {
		panic("no return value specified for GetBlockHeader")
	}

	var r0 *wire.BlockHeader
	var r1 error
	if rf, ok := ret.Get(0).(func(*chainhash.Hash) (*wire.BlockHeader, error)); ok {
		return rf(blockHash)
	}
	if rf, ok := ret.Get(0).(func(*chainhash.Hash) *wire.BlockHeader); ok {
		r0 = rf(blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*wire.BlockHeader)
		}
	}

	if rf, ok := ret.Get(1).(func(*chainhash.Hash) error); ok {
		r1 = rf(blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockVerbose provides a mock function with given fields: blockHash
func (_m *BTCRPCClient) GetBlockVerbose(blockHash *chainhash.Hash) (*btcjson.GetBlockVerboseResult, error) {
	ret := _m.Called(blockHash)

	if len(ret) == 0 {
		panic("no return value specified for GetBlockVerbose")
	}

	var r0 *btcjson.GetBlockVerboseResult
	var r1 error
	if rf, ok := ret.Get(0).(func(*chainhash.Hash) (*btcjson.GetBlockVerboseResult, error)); ok {
		return rf(blockHash)
	}
	if rf, ok := ret.Get(0).(func(*chainhash.Hash) *btcjson.GetBlockVerboseResult); ok {
		r0 = rf(blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*btcjson.GetBlockVerboseResult)
		}
	}

	if rf, ok := ret.Get(1).(func(*chainhash.Hash) error); ok {
		r1 = rf(blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockVerboseTx provides a mock function with given fields: blockHash
func (_m *BTCRPCClient) GetBlockVerboseTx(blockHash *chainhash.Hash) (*btcjson.GetBlockVerboseTxResult, error) {
	ret := _m.Called(blockHash)

	if len(ret) == 0 {
		panic("no return value specified for GetBlockVerboseTx")
	}

	var r0 *btcjson.GetBlockVerboseTxResult
	var r1 error
	if rf, ok := ret.Get(0).(func(*chainhash.Hash) (*btcjson.GetBlockVerboseTxResult, error)); ok {
		return rf(blockHash)
	}
	if rf, ok := ret.Get(0).(func(*chainhash.Hash) *btcjson.GetBlockVerboseTxResult); ok {
		r0 = rf(blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*btcjson.GetBlockVerboseTxResult)
		}
	}

	if rf, ok := ret.Get(1).(func(*chainhash.Hash) error); ok {
		r1 = rf(blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNetworkInfo provides a mock function with given fields:
func (_m *BTCRPCClient) GetNetworkInfo() (*btcjson.GetNetworkInfoResult, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetNetworkInfo")
	}

	var r0 *btcjson.GetNetworkInfoResult
	var r1 error
	if rf, ok := ret.Get(0).(func() (*btcjson.GetNetworkInfoResult, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *btcjson.GetNetworkInfoResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*btcjson.GetNetworkInfoResult)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNewAddress provides a mock function with given fields: account
func (_m *BTCRPCClient) GetNewAddress(account string) (btcutil.Address, error) {
	ret := _m.Called(account)

	if len(ret) == 0 {
		panic("no return value specified for GetNewAddress")
	}

	var r0 btcutil.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (btcutil.Address, error)); ok {
		return rf(account)
	}
	if rf, ok := ret.Get(0).(func(string) btcutil.Address); ok {
		r0 = rf(account)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(btcutil.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRawTransaction provides a mock function with given fields: txHash
func (_m *BTCRPCClient) GetRawTransaction(txHash *chainhash.Hash) (*btcutil.Tx, error) {
	ret := _m.Called(txHash)

	if len(ret) == 0 {
		panic("no return value specified for GetRawTransaction")
	}

	var r0 *btcutil.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(*chainhash.Hash) (*btcutil.Tx, error)); ok {
		return rf(txHash)
	}
	if rf, ok := ret.Get(0).(func(*chainhash.Hash) *btcutil.Tx); ok {
		r0 = rf(txHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*btcutil.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(*chainhash.Hash) error); ok {
		r1 = rf(txHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRawTransactionVerbose provides a mock function with given fields: txHash
func (_m *BTCRPCClient) GetRawTransactionVerbose(txHash *chainhash.Hash) (*btcjson.TxRawResult, error) {
	ret := _m.Called(txHash)

	if len(ret) == 0 {
		panic("no return value specified for GetRawTransactionVerbose")
	}

	var r0 *btcjson.TxRawResult
	var r1 error
	if rf, ok := ret.Get(0).(func(*chainhash.Hash) (*btcjson.TxRawResult, error)); ok {
		return rf(txHash)
	}
	if rf, ok := ret.Get(0).(func(*chainhash.Hash) *btcjson.TxRawResult); ok {
		r0 = rf(txHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*btcjson.TxRawResult)
		}
	}

	if rf, ok := ret.Get(1).(func(*chainhash.Hash) error); ok {
		r1 = rf(txHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransaction provides a mock function with given fields: txHash
func (_m *BTCRPCClient) GetTransaction(txHash *chainhash.Hash) (*btcjson.GetTransactionResult, error) {
	ret := _m.Called(txHash)

	if len(ret) == 0 {
		panic("no return value specified for GetTransaction")
	}

	var r0 *btcjson.GetTransactionResult
	var r1 error
	if rf, ok := ret.Get(0).(func(*chainhash.Hash) (*btcjson.GetTransactionResult, error)); ok {
		return rf(txHash)
	}
	if rf, ok := ret.Get(0).(func(*chainhash.Hash) *btcjson.GetTransactionResult); ok {
		r0 = rf(txHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*btcjson.GetTransactionResult)
		}
	}

	if rf, ok := ret.Get(1).(func(*chainhash.Hash) error); ok {
		r1 = rf(txHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListUnspent provides a mock function with given fields:
func (_m *BTCRPCClient) ListUnspent() ([]btcjson.ListUnspentResult, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ListUnspent")
	}

	var r0 []btcjson.ListUnspentResult
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]btcjson.ListUnspentResult, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []btcjson.ListUnspentResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]btcjson.ListUnspentResult)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListUnspentMinMaxAddresses provides a mock function with given fields: minConf, maxConf, addrs
func (_m *BTCRPCClient) ListUnspentMinMaxAddresses(minConf int, maxConf int, addrs []btcutil.Address) ([]btcjson.ListUnspentResult, error) {
	ret := _m.Called(minConf, maxConf, addrs)

	if len(ret) == 0 {
		panic("no return value specified for ListUnspentMinMaxAddresses")
	}

	var r0 []btcjson.ListUnspentResult
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, []btcutil.Address) ([]btcjson.ListUnspentResult, error)); ok {
		return rf(minConf, maxConf, addrs)
	}
	if rf, ok := ret.Get(0).(func(int, int, []btcutil.Address) []btcjson.ListUnspentResult); ok {
		r0 = rf(minConf, maxConf, addrs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]btcjson.ListUnspentResult)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, []btcutil.Address) error); ok {
		r1 = rf(minConf, maxConf, addrs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendRawTransaction provides a mock function with given fields: tx, allowHighFees
func (_m *BTCRPCClient) SendRawTransaction(tx *wire.MsgTx, allowHighFees bool) (*chainhash.Hash, error) {
	ret := _m.Called(tx, allowHighFees)

	if len(ret) == 0 {
		panic("no return value specified for SendRawTransaction")
	}

	var r0 *chainhash.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(*wire.MsgTx, bool) (*chainhash.Hash, error)); ok {
		return rf(tx, allowHighFees)
	}
	if rf, ok := ret.Get(0).(func(*wire.MsgTx, bool) *chainhash.Hash); ok {
		r0 = rf(tx, allowHighFees)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chainhash.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(*wire.MsgTx, bool) error); ok {
		r1 = rf(tx, allowHighFees)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBTCRPCClient creates a new instance of BTCRPCClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBTCRPCClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *BTCRPCClient {
	mock := &BTCRPCClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
