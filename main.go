package main

import (
    "strings"
    "fmt"
    "os"
    "bufio"
    "strconv"
)

var board = [][]string{
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
}

const winningPatterns = [][]struct{
    line, col int
}{
    [{0, 0}, {0, 1}, {0, 2}],
    [{1, 0}, {1, 1}, {1, 2}],
    [{2, 0}, {2, 1}, {2, 2}],
    [{0, 0}, {1, 0}, {2, 0}],
    [{0, 1}, {1, 1}, {2, 1}],
    [{0, 2}, {1, 2}, {2, 2}],
    [{0, 0}, {1, 1}, {2, 2}],
    [{0, 2}, {1, 1}, {2, 0}],
}

var player = "O"

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    printBoard()
    for {
        fmt.Printf(">> ")
        scanner.Scan()
        l := scanner.Text()
        nums := strings.SplitN(l, " ", 2)
        if len(nums) != 2 {
            usage()
            continue
        }
        lineNumber, err := strconv.Atoi(nums[0])
        if err != nil {
            usage()
            continue
        }
        colNumber, err := strconv.Atoi(nums[1])
        if err != nil {
            usage()
            continue
        }
        if lineNumber > 3 || lineNumber <0 || colNumber > 3 || colNumber < 0 {
            usage()
            continue
        }
        board[lineNumber][colNumber] = player
        if player == "O" {
            player = "X"
        } else {
            player = "O"
        }
        printBoard()
        if ok := checkWinner(); ok {
            break
        }
    }
}

func usage() {
    fmt.Println("type like: <line number> <column number>")
}

func printBoard() {
    for i := 0; i < 3; i++ {
        line := strings.Join(board[i], " ")
        fmt.Printf("%s\n", line)
    }
}

func checkWinner () bool {
}
