type Ttt struct {
    board [3][3]string
    turn string
    a string
    b string
    blank string
}

func tictactoe(moves [][]int) string {
    if len(moves) < 5 {
        return "Pending"
    }
    
    t := Ttt {a: "A", b: "B"}
    t.init("-")
    
    for _, move := range moves {
        res := t.set(move[0], move[1])
        if res == t.a || res == t.b {
            return res
        }
    }
    
    if t.isDraw() {
        return "Draw"
    }
    return "Pending"
}

func (t *Ttt) init(defaut string) {
    t.board = [3][3]string{}
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            t.board[i][j] = defaut
        }
    }
    t.turn = t.a
    t.blank = defaut
}

func (t *Ttt) set(i int, j int) string {
    t.board[i][j] = t.turn
    
    if t.isWin() {
        return t.turn
    }
    
    if t.turn == t.a {
        t.turn = t.b
    } else {
        t.turn = t.a
    }
    
    return ""
}

func (t *Ttt) isWin() bool {
    if t.board[1][1] == t.turn {
        if t.board[0][0] == t.turn && t.board[2][2] == t.turn {
            return true
        }
        if t.board[0][2] == t.turn && t.board[2][0] == t.turn {
            return true
        }
    }
    
    for i := 0; i < 3; i++ {
        if t.board[i][0] == t.turn && t.board[i][1] == t.turn && t.board[i][2] == t.turn {
            return true
        }
        
        if t.board[0][i] == t.turn && t.board[1][i] == t.turn && t.board[2][i] == t.turn {
            return true
        }
    }
    
    return false
}

func (t *Ttt) isDraw() bool {
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if t.board[i][j] == t.blank {
                return false
            }
        }
    }
    return true
}
