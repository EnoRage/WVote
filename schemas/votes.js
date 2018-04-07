import { Bool } from 'waves-api/raw/src/classes/ByteProcessor';

var mongoose = require('mongoose');
var Schema = mongoose.Schema;

var Voters = new Schema({
    num: {
        type: Number
    }, 
    address: {
        type: String
    },
    vote: {
        type: Boolean
    }
});

var Votes = new Schema({
    num: {
        type: Number
    },
    description: {
        type: String,
        default: ""
    },
    approvedAddresses: {
        type: Array,
        default: ""
    },
    startTime: {
        type: Date,
        default: Date.now()
    },
    endTime: {
        type: Date,
        default: Date.now() + 99999
    },
    end: {
        type: Boolean,
        default: false
    }
}, {
    versionKey: false
});

module.exports.votes = mongoose.model('Votes', Votes);
module.exports.voters = mongoose.model('Voters', Voters);