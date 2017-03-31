package block

type Chain struct {
	magic    int32    //Magic value indicating message origin network, and used to seek to next message when stream state is unknown
	command  [48]byte //ASCII string identifying the packet content, NULL padded (non-NULL padding results in packet rejected)
	length   int32    //Length of payload in number of bytes
	checksum int32    //First 4 bytes of sha256(sha256(payload))
	payload  []byte   //The actual data
}
type Block struct {
	version     int32 //Block version information (note, this is signed)
	prev_block  [32]byte
	merkle_root [32]byte
	timestamp   int32
	bits        int32
	nonce       int32
	txn_count   int32
}

func NewBlock(version int32, prev_block [32]byte, merkle_root [32]byte, timestamp int32, bits int32, nonce int32, txn_count int32) *Block {
	return &Block{
		version:     version,
		prev_block:  prev_block,
		merkle_root: merkle_root,
		timestamp:   timestamp,
		bits:        bits,
		nonce:       nonce,
		txn_count:   txn_count,
	}
}
