const web3 = require('./ethereum');

// Sync blocks and start listening for new blocks
async function sync_blocks(current_block_number, opts) {
    let latest_block_number = await web3.eth.getBlockNumber();
    let synced_block_number = await sync_to_block(current_block_number, latest_block_number, opts);

    // subscribe to new blocks
    web3.eth
        .subscribe('newBlockHeaders', (error, result) => error && console.log(error))
        .on('data', async function(blockHeader) {
            return await process_block(blockHeader.number, opts);
        });

    return synced_block_number;
}


// Traverse all unprocessd blocks between the current index and the latest block number
async function sync_to_block(index, lastest, opts) {
    if (index >= lastest) {
        return index;
    }
    await process_block(index + 1, opts);
    return await sync_to_block(index + 1, lastest, opts);
}

// Load all data about the given block and call the callback if defined
async function process_block(block_hash_or_id, opts) {
    // load block information by id or hash
    const block = await web3.eth.getBlock(block_hash_or_id, true);
    // call the onTransactions callback if defined
    opts.onTransactions ? opts.onTransactions(block.transactions) : null;
    // call the onBlock callback if defined.
    opts.onBlock ? opts.onBlock(block_hash_or_id) : null;
    return block;
}

module.exports = sync_blocks;