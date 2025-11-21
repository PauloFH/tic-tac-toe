package main

import "math"

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

func minimax(board [9]string, depth int, isMaximizing bool) int {
	if checkWinner(board, IA) {
		return 10 - depth
	}
	if checkWinner(board, HUMANO) {
		return depth - 10
	}
	if isFull(board) {
		return 0
	}

	if isMaximizing {
		best := math.MinInt32
		for i := 0; i < 9; i++ {
			if board[i] == VAZIO {
				board[i] = IA
				score := minimax(board, depth+1, false)
				board[i] = VAZIO
				if score > best {
					best = score
				}
			}
		}
		return best
	} else {
		best := math.MaxInt32
		for i := 0; i < 9; i++ {
			if board[i] == VAZIO {
				board[i] = HUMANO
				score := minimax(board, depth+1, true)
				board[i] = VAZIO
				if score < best {
					best = score
				}
			}
		}
		return best
	}
}

func bestMove(board [9]string) int {
	bestVal := math.MinInt32
	move := -1
	for i := 0; i < 9; i++ {
		if board[i] == VAZIO {
			board[i] = IA
			score := minimax(board, 0, false)
			board[i] = VAZIO
			if score > bestVal {
				bestVal = score
				move = i
			}
		}
	}
	return move
}
