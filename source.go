package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Block structure
type Block struct {
	Timestamp string `json:"timestamp"`
	Data      string `json:"data"`
}

// Global variable to store blocks
var blocks []Block

// DisplayAllBlocks function to display all blocks
func DisplayAllBlocks() {
	if len(blocks) == 0 {
		fmt.Println("No blocks available.")
	} else {
		fmt.Println("All Blocks:")
		for _, block := range blocks {
			fmt.Printf("Timestamp: %s\nData: %s\n\n", block.Timestamp, block.Data)
		}
	}
}

// NewBlock function to create a new block
func NewBlock(data string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	block := Block{
		Timestamp: timestamp,
		Data:      data,
	}
	blocks = append(blocks, block)
	fmt.Println("New block created successfully.")
}

// ModifyBlock function to modify a block
func ModifyBlock(index int, newData string) {
	if index >= 0 && index < len(blocks) {
		blocks[index].Data = newData
		blocks[index].Timestamp = time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("Block at index %d modified successfully.\n", index)
	} else {
		fmt.Println("Invalid index. Block modification failed.")
	}
}

func main() {
	// Create a new block
	NewBlock("Initial data")

	// Display all blocks
	DisplayAllBlocks()

	// Modify the first block
	ModifyBlock(0, "Updated data")

	// Display all blocks again
	DisplayAllBlocks()
}
