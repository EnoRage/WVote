var mongoose = require('mongoose');
var Schema = mongoose.Schema;

var VotersSchema = new Schema({
    voteID: {
        type: Schema.Types.ObjectId
    }, 
    address: {
        type: String
    },
    vote: {
        type: Boolean
    }
});

var UserSchema = new Schema({
    sell: [VotersSchema]
});

var Votes = new Schema({
    organisationID: {
        type: Schema.Types.ObjectId
    },
    description: {
        type: String,
        default: ""
    }
}, {
    versionKey: false
});

module.exports.votes = mongoose.model('Votes', Votes);
module.exports.voters = mongoose.model('VotersSchema', VotersSchema);