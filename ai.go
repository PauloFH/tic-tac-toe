package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const MCTS_ITERATIONS = 10000

func checkWinner(board [9]string, player string) bool {
	wins := [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 6}}
	for _, l := range wins {
		if board[l[0]] == player && board[l[1]] == player && board[l[2]] == player {
			return true
		}
	}
	return false
}

func isFull(board [9]string) bool {
	for _, v := range board {
		if v == VAZIO {
			return false
		}
	}
	return true
}

func getLegalMoves(board [9]string) []int {
	moves := []int{}
	for i, val := range board {
		if val == VAZIO {
			moves = append(moves, i)
		}
	}
	return moves
}

func isTerminal(board [9]string) bool {
	return checkWinner(board, IA) || checkWinner(board, HUMANO) || isFull(board)
}

type Node struct {
	board         [9]string
	parent        *Node
	children      []*Node
	move          int
	currentPlayer string
	visits        float64
	wins          float64
	untried       []int
}

func NewNode(board [9]string, parent *Node, move int, currentPlayer string) *Node {
	return &Node{
		board:         board,
		parent:        parent,
		move:          move,
		currentPlayer: currentPlayer,
		untried:       getLegalMoves(board),
		children:      []*Node{},
	}
}

func (n *Node) SelectChild() *Node {
	if len(n.children) == 0 {
		return nil
	}
	bestScore := -math.MaxFloat64
	var bestNode *Node
	c := 1.41

	for _, child := range n.children {
		if child.visits == 0 {
			return child
		}

		ucb1 := (child.wins / child.visits) +
			c*math.Sqrt(math.Log(n.visits)/child.visits)

		if ucb1 > bestScore {
			bestScore = ucb1
			bestNode = child
		}
	}
	if bestNode == nil && len(n.children) > 0 {
		return n.children[0]
	}
	return bestNode
}

func (n *Node) Expand() *Node {
	idx := rand.Intn(len(n.untried))
	move := n.untried[idx]
	n.untried = append(n.untried[:idx], n.untried[idx+1:]...)

	newBoard := n.board
	newBoard[move] = n.currentPlayer
	next := HUMANO
	if n.currentPlayer == HUMANO {
		next = IA
	}

	child := NewNode(newBoard, n, move, next)
	n.children = append(n.children, child)
	return child
}

func (n *Node) Simulate() float64 {
	currentBoard := n.board
	currentPlayer := n.currentPlayer

	for !isTerminal(currentBoard) {
		moves := getLegalMoves(currentBoard)
		randomMove := moves[rand.Intn(len(moves))]

		currentBoard[randomMove] = currentPlayer
		if currentPlayer == HUMANO {
			currentPlayer = IA
		} else {
			currentPlayer = HUMANO
		}
	}

	if checkWinner(currentBoard, IA) {
		return 1.0
	} else if checkWinner(currentBoard, HUMANO) {
		return 0.0
	}
	return 0.5
}

func (n *Node) Backpropagate(result float64) {
	node := n
	for node != nil {
		node.visits++
		if node.currentPlayer == HUMANO {
			node.wins += result
		} else {
			node.wins += (1.0 - result)
		}
		node = node.parent
	}
}

func bestMove(board [9]string) int {
	root := NewNode(board, nil, -1, IA)

	start := time.Now()

	for i := 0; i < MCTS_ITERATIONS; i++ {
		node := root
		for len(node.untried) == 0 && len(node.children) > 0 {
			next := node.SelectChild()
			if next == nil {
				break
			}
			node = next
		}
		if len(node.untried) > 0 {
			node = node.Expand()
		}
		result := node.Simulate()
		node.Backpropagate(result)
	}

	elapsed := time.Since(start)
	fmt.Printf("MCTS (%d ops) em %s\n", MCTS_ITERATIONS, elapsed)

	bestMove := -1
	bestVisits := -1.0

	for _, c := range root.children {
		if c.visits > bestVisits {
			bestVisits = c.visits
			bestMove = c.move
		}
	}

	if bestMove == -1 {
		fmt.Println("MCTS não encontrou nó filho válido! Usando movimento aleatório.")
		legal := getLegalMoves(board)
		if len(legal) > 0 {
			return legal[0]
		}
	}

	return bestMove
}
