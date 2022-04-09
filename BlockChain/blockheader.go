package BlockChain

type BlockHeader struct {
  Index         int
  TimeStamp     int64
  Nonce         int32
  MerkleRoot    [32]byte
  PreviousHash  [32]byte
}