var mongoose = require('mongoose');
var Schema = mongoose.Schema;

var VotersSchema = new Schema({
    voteID: {
        type: Schema.Types.ObjectId
    }, 
    userID: {
        type: String
    },
    vote: {
        type: Boolean
    }
});

var Votes = new Schema({
    organisationID: {
        type: Schema.Types.ObjectId
    },
    description: {
        type: String,
        default: ""
    },
    approvedAddresses: {
        type: Array,
        default: ""
    },
    endTime: {
        type: Date,
        default: Date.now() + 99999999
    }
}, {
    versionKey: false
});

module.exports.votes = mongoose.model('Votes', Votes);
module.exports.voters = mongoose.model('VotersSchema', VotersSchema);