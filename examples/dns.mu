#define CONFADDR 0x661005d2720d855f1d9976f88bb10c1a3398c77f

var a = 0
var nameRegAddr = 0

call(CONFADDR, 0, 1000, a, nameRegAddr)

var[2] nameRegArgs
nameRegArgs[0] = "register"
nameRegArgs[1] = "DnsReg"

call(nameRegAddr, 0, 1000, nameRegArgs, nil)

contract.storage["owner"] = tx.sender()

return compile {
    if message.data[0] == "register" {
        if contract.storage[message.data[1]] == 0 {
            // Deregister old
            if contract.storage[tx.sender()] != 0 {
                contract.storage[contract.storage[tx.sender()]] = 0
            }

            // Ex. jeff.eth : "127.0.0.1"
            // [0]: "jeff.eth"
            // [1]: 0xac000001
            contract.storage[message.data[1]] = message.data[2]
            contract.storage[tx.sender()] = message.data[1]
        }
    } else if message.data[0] == "deregister" {
        if contract.storage[tx.sender()] == message.data[1] {
            contract.storage[tx.sender()] = 0
            contract.storage[message.data[1]] = 0
        }
    } else if message.data[0] == "kill" {
        if contract.storage["owner"] == tx.sender() {
            suicide(tx.sender())
        }
    }
}
