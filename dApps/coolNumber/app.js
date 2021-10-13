async function loadWeb3() {
    if (window.ethereum) {
        window.web3 = new Web3(window.ethereum);
        window.ethereum.enable();
    }
}

async function loadContract() {
    return await new window.web3.eth.Contract([
        {
            "inputs": [],
            "name": "coolNumber",
            "outputs": [
                {
                    "internalType": "uint256",
                    "name": "",
                    "type": "uint256"
                }
            ],
            "stateMutability": "view",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "uint256",
                    "name": "_coolNumber",
                    "type": "uint256"
                }
            ],
            "name": "setCoolNumber",
            "outputs": [],
            "stateMutability": "nonpayable",
            "type": "function"
        }
    ], '0xB202CCCd3b66F63f567D40d6e22C1C1BAB627824');
}


async function printCoolNumber() {
    updateStatus('fetching Cool Number...');
    const coolNumber = await window.contract.methods.coolNumber().call();
    updateStatus(`coolNumber: ${coolNumber}`);
}

async function getCurrentAccount() {
    const accounts = await window.web3.eth.getAccounts();
    return accounts[0];
}

async function changeCoolNumber() {
    const value = Math.floor(Math.random() * 100);
    updateStatus(`Updating coolNumber with ${value}`);
    const account = await getCurrentAccount();
    const coolNumber = await window.contract.methods.setCoolNumber(value).send({ from: account });
    updateStatus('Updated.');
}

async function load() {
    await loadWeb3();
    window.contract = await loadContract();
    updateStatus('Ready!');
}

function updateStatus(status) {
    const statusEl = document.getElementById('status');
    statusEl.innerHTML = status;
    console.log(status);
}

load();
