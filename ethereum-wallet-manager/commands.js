const web3 = require('./ethereum');
const redis = require('./redis');
const queue = require('./queue');

async function listen_to_commands() {
    const queue_consumer = queue.consumer('eth.wallet.manager.commands', ['command'])
    // process messages.
    queue_consumer.on('message', async function(topic_message) {
        try {
            const message = JSON.parse(topic_message.value);
            // create the new address with some reply metadata to match the response to the request
            const resp = await create_address(message.meta);
            if (resp) {
                await queue.send('address.created', [resp]);
            }
        } catch(err) {
            console.log(topic_message, err);
            queue.send('errors', [{ type: 'command', request: topic_message, error_code: err.code, error_message: err.message, err_stack: err.stack}]);
        }
    });

    return queue_consumer;
}

function create_address(meta = {}) {
    // generate the address
    const account = web3.eth.accounts.create()
    
    // address
    address = account.address.toLowerCase();

    // Store the public key
    await redis.setAsync(`eth:address:public:${address}`, JSON.stringify({}));

    // Store the private key
    await redis.setAsync(`eth:address:private:${address}`, account.privateKey);

    return Object.assign({}, meta, {address: account.address});
}

module.exports.listen_to_commands = listen_to_commands;