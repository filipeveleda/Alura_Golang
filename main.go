package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente!"
	}
}

func (c *ContaCorrente) Depositar(ValorDoDeposito float64) (string, float64) {
	podeDepositar := ValorDoDeposito > 0
	if podeDepositar {
		c.saldo += ValorDoDeposito
		return "Deposito realizado com sucesso! Valor atual: R$", c.saldo
	} else {
		return "Precisa ser um valor positivo! Valor Atual: R$", c.saldo
	}

}

func (c *ContaCorrente) Traferir(ValorTransferencia float64, contaDestino *ContaCorrente) bool {
	if ValorTransferencia <= c.saldo {
		c.saldo -= ValorTransferencia
		contaDestino.Depositar(ValorTransferencia)
		return true
	} else {
		return false
	}
}

func main() {
	contaDoFilipe := ContaCorrente{titular: "Filipe", saldo: 600}
	contaDoLuan := ContaCorrente{titular: "Luan", saldo: 400}

	contaDoFilipe.Traferir(300, &contaDoLuan)

	fmt.Println(contaDoFilipe)
	fmt.Println(contaDoLuan)

}
