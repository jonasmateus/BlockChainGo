package BlockChain

import (
  "time"
  "fmt"
  "crypto/sha256"
  "encoding/json"
  "strings"
)

const DIFFICULTY = 5

type Chain struct {
  chain   []Block
  MenPoll []Transaction
}

//Construtor 
func New() (*Chain, *Block) {

  blockChain := &Chain{}
 
  blockChain.chain = []Block{}
  genesis := blockChain.createGenesisBlock()
  
  return blockChain, genesis 
}

func (blockChain *Chain) createGenesisBlock() *Block {

  genesisBlock := NewBlock()

  genesisBlock.Header.Index = len(blockChain.chain)
  genesisBlock.Header.PreviousHash = [32]byte{}
  genesisBlock.Header.Nonce = 0
  genesisBlock.Header.TimeStamp = time.Now().Unix()
  
  blockChain.chain = append(blockChain.chain, *genesisBlock)

  return genesisBlock
}

func (blockChain *Chain) CreateBlock() *Block {

  newBlock := NewBlock()

  newBlock.Header.Index = len(blockChain.chain)
  newBlock.Header.PreviousHash = GenerateHash(&blockChain.chain[len(blockChain.chain) - 1])
  newBlock.Header.TimeStamp = time.Now().Unix()


  return newBlock
}

func GenerateHash(block *Block) [32]byte {
  blockHeaderBuffer, _ := json.Marshal(block.Header)
  firstStringHash  := fmt.Sprintf("%x", sha256.Sum256(blockHeaderBuffer))
  secondHash := sha256.Sum256([]byte(firstStringHash))

  return secondHash
}

func GetBlockID(block *Block) [32]byte {

  return GenerateHash(block)
}

func isValidProof(block *Block, nonce *int32) bool {

  prefix := strings.Repeat("0", DIFFICULTY)
  block.Header.Nonce = *nonce
  hash := GenerateHash(block)
  encodedStrHash := fmt.Sprintf("%x", hash)

  if(strings.HasPrefix(encodedStrHash, prefix)) {
    return true
  }else{
    *nonce++
    return isValidProof(block, nonce)
  }
}

func (blockChain *Chain) MineProofOfWork(block *Block) int32 {

  nonce := int32(0)

  if(isValidProof(block, &nonce)){
    blockChain.chain = append(blockChain.chain, *block)
    return nonce
  }else {
    return nonce
  }
}

func (blockChain *Chain) PrintChain() {
  
  template := `
+ ---------------------------------------------------------------- +
| %x |
| ---------------------------------------------------------------- | 
| Índice:         Timestamp:              Nonce:                   |
| %x               %v                  %v                    |
|                                                                  |
| Merkle Root:                                                     |	
| %x |                	
|                                                                  |
| Transações:                                                      |
| %v                                                                |
|                                                                  |
| Hash do último bloco:                                            |
| %x |                	
+ ---------------------------------------------------------------- +	
  `
  for i := 0; i < len(blockChain.chain); i++ {
    fmt.Printf(template, GetBlockID(&blockChain.chain[i]), blockChain.chain[i].Header.Index,  blockChain.chain[i].Header.TimeStamp, blockChain.chain[i].Header.Nonce, blockChain.chain[i].Header.MerkleRoot, len(blockChain.chain[i].TransActions), blockChain.chain[i].Header.PreviousHash)
  }
}
