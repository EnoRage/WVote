var mongoose = require('mongoose');
var Schema = mongoose.Schema;
var Organisations = new Schema({
    name: {
        type: String,
        default: ""
    },
    description: {
        type: String,
        default: ""
    }
}, {
    versionKey: false
});

module.exports = mongoose.model('Organisations', Organisations);