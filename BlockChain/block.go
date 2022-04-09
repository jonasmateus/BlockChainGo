package BlockChain

type Block struct {
  Header BlockHeader
  Txs    []Tx
}

func NewBlock() *Block {
  return &Block{}
}
