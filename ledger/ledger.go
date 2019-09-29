package ledger

import (
	"errors"
)

// Represents a blockchain itself
type Ledger struct {
	Genesis *Block
}

func (l *Ledger) Add(h string) string {
	if l.Genesis == nil {
		// If there's no block on the ledger, then we create our genesis block
		l.Genesis = &Block{Hash: h, Next: nil, Previous: nil}
		return h
	}

	// We start counting from 0, which is the head of the list
	block := l.Genesis

	// We traverse the blocks
	for block.Next != nil {
		block = block.Next
	}

	// We then add the block to the end of the list
	block.Next = &Block{Hash: h, Next: nil, Previous: block}

	// We return it's index so we can reference our newly created block later on
	return h
}

func (l *Ledger) Get(h string) (*Block, error) {
	if l.Genesis == nil {
		return nil, errors.New("The chain is Empty")
	}

	block := l.Genesis
	hash := l.Genesis.Hash

	for block != nil {
		// If hashes match, then we return the current block and no errors
		if hash == h {
			return block, nil
		}
		block = block.Next
		hash = block.Hash
	}

	return nil, errors.New("Block not found")

}
