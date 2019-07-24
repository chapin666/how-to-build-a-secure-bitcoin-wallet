'use strict';

// load configuration
const config = require('config');
const redis = require('redis');
const bluebird = require('bluebird');

// promisify the redis client using bluebird
bluebird.promisifyAll(redis.RedisClient.prototype);
bluebird.promisifyAll(redis.Multi.prototype);

// create a new redis client
const client = redis.createClient(config.redis_port, config.redis_host);

// if an error occurs, print it to the console
client.on("error", function(err) {
    console.log("[REDIS] Error encounterd", err);
});

module.exports = client;

