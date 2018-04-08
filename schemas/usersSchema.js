var mongoose = require('mongoose');
var Schema = mongoose.Schema;
var Users = new Schema({
    userID: {
        type: String,
        default: ""
    },
    name: {
        type: String,
        default: ""
    },
    encryptedSeed: {
        type: String,
        default: ""
    },
    address: {
        type: String,
        default: ""
    },
    sessionKey: {
        type: String,
        default: ""
    }
}, {
    versionKey: false
});

module.exports = mongoose.model('Users', Users);