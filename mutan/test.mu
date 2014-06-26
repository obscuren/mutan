return compile {
	to := this.data[0]
	from := this.origin()
	value := this.data[1]

	if this.store[from] >= value {
		this.store[from] = this.store[from] - value
			this.store[to] = this.store[to] + value
	}
}

