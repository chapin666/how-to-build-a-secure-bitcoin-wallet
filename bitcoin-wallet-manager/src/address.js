const bitcoin = require('bitcoinjs-lib');
const network = bitcoin.networks.bitcoin;

function publicKey() {
    const keyPair = bitcoin.ECPair.makeRandom({
        network: network
    });
    
    const publicKey = bitcoin.payments.p2pkh({
        network: network,
        pubkey: keyPair.publicKey
    }).address
    
    return publicKey
}