package main

import (
  "fmt"
  "BlockChain"
)

func main () {
  
  blockChain, _ := BlockChain.New()

  nonce := blockChain.MineProofOfWork(blockChain.CreateBlock())

  nonce = blockChain.MineProofOfWork(blockChain.CreateBlock())

  nonce = blockChain.MineProofOfWork(blockChain.CreateBlock())

  fmt.Println(nonce)

  blockChain.PrintChain()

}