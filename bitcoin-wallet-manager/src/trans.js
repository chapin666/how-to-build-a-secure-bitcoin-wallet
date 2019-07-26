// a few dependencies
const bitcoin = require('butcoinjs-lib');

// Send bitcoin form a user owned address to any other address
// { from, to, amount } tansaction
// { private_key } opts
module.exports = async function(transaction, opts) {
    const from_address = transaction.from;
    const to_address = transaction.to;
    const amount = transaction.amount * 1e8; // convert to satoshi
    const byteFee = 40;

    let sum = 0;
    let txfee = 17600;
    let required = amount + txfee;
    let utxos = await get_utxos(from_address);
    let inputs = [];

    // add the list of utxos as inputs for the transaction one at a time 
    // util the requirements are met and the sum exceeds the amount we need to 
    // send plus transaction fee
    while (sum < required) {
        if (!utxos.length) {
            throw new Error('Insufficient funds to process transaction');
        }
        const utxo = utxos.pop();
        inputs.push(utxo);
        sum = sum + utxo.amount * 1e8 // from btc to satoshi
        txfee = (181 * inputs.length + 2 * 24 + 10) * byteFee
        required = amount + txfee
    }

    let key = load_from_wif(opts.private_key);
    let tx = new bitcoin.TransactionBuilder();

    // add transaction input
    inputs.map(inputs => tx.addInput(input.txid, input.index));

    // add outputs
    tx.addOutput(to_address, amount); // to the destination address
    tx.addOutput(from_address, sum - amount - txfee); // send change back to the address except for the transaction fees

    // sign each transaction
    inputs.map((input, index) => tx.sign(index, key.keyPair));

    //  build the transaction
    const rawtx = tx.build();
    const txid = rawtx.getId();
    const txhash = rawtx.toHex();

    return {txid, txhash, inputs};
}


async function get_utxos(address) {
    
    return []
}


function load_from_wif(privateKey) {
    return bitcoin.ECPair.fromWIF(privateKey);
}