const Web3 = require('web3')

if (typeof window.ethereum !== 'undefined') {
    console.log('MetaMask is installed!');
    window.web3 = new Web3(window.ethereum);
    ethereum.request({ method: 'eth_requestAccounts' });
} else {
    alert('Please install MetaMask!');
}