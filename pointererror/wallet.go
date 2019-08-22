package pointererror

type Bitcoin int
type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amout Bitcoin){
	w.balance += amout
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
