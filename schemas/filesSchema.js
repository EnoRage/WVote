var mongoose = require('mongoose');
var Schema = mongoose.Schema;
var files = new Schema({
    url: {
        type: String,
        default: ""
    },
    file: {
        type: String,
        default: ""
    },
    data: {
        type: String,
        default: ''
    }
}, {
    versionKey: false
});

module.exports = mongoose.model('files', files);