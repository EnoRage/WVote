const mongoose = require("mongoose");

mongoose.Promise = global.Promise;

const user = 'erage';
const password = 'doBH8993nnjdoBH8993nnj';
const uri = 'mongodb://'+user+':'+password+'@51.144.89.99:27017/VoteDB?authSource=admin';

const options = {
    autoIndex: false,
    reconnectTries: Number.MAX_VALUE,
    reconnectInterval: 500,
    poolSize: 1000,
    bufferMaxEntries: 0
};
const db = mongoose.connect(uri).then(console.log('Mongo DB works fine'));