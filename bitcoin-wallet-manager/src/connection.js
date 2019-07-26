const Client = require('bitcoin-core');
const config = require('./config');

const client = new Client({
    host: config.bitcoin_rpc_host,
    port: config.bitcoin_rpc_port,
    username: config.bitcoin_rpc_user,
    password: config.bitcoin_rpc_pass,
});

module.exports = client;