return compile {
	to := this.data[0]
	from := this.caller()
	value := this.data[1]

	if contract.storage[from] >= value {
		contract.storage[from] = contract.storage[from] - value
		contract.storage[to] = contract.storage[to] + value
	}
}

