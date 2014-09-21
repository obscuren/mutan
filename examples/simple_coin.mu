// Premined coin
contract.storage[tx.sender()] = 10**12

return compile {
	value := message.data[0]
	to := message.data[1]
	from := tx.sender()

	if contract.storage[from] >= value {
		contract.storage[from] = contract.storage[from] - value
		contract.storage[to] = contract.storage[to] + value
	}
}

