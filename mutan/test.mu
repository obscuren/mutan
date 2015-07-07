to := message.data[0]
from := message.caller()
value := message.data[1]

if contract.storage[from] >= value {
    contract.storage[from] = contract.storage[from] - value
    contract.storage[to] = contract.storage[to] + value
}


t := compile {
    func t(var *a) {
        *a = 10;
    }
}
