const mongoose = require("mongoose"),
    ObjectId = require('mongoose').Types.ObjectId;
Users = require('./schemas/usersSchema.js');
Organizations = require('./schemas/organisationsSchema.js'),
    Votes = require('./schemas/votes.js').votes;
Voters = require('./schemas/votes.js').voters;

mongoose.Promise = global.Promise;

const user = 'erage';
const password = 'doBH8993nnjdoBH8993nnj';
const uri = 'mongodb://' + user + ':' + password + '@51.144.89.99:27017/unblock?authSource=admin';

const options = {
    autoIndex: false,
    reconnectTries: Number.MAX_VALUE,
    reconnectInterval: 500,
    poolSize: 1000,
    bufferMaxEntries: 0
};
const db = mongoose.connect(uri).then(console.log('Mongo DB works fine'));

// Добавление организации в базу данных
function addOrganisation(name, descriptions) {
    Organizations.create({
        name: name,
        description: description
    }, (err, doc) => {

    });
}

// Поиск доступных адресов для голсования конкретной организации
function findApprovedAddresses(organisationID, callback) {
    Votes.find({
        _id: new ObjectId(organisationID)
    }, (err, doc) => {
        callback(doc[0].approvedAddresses);
    });
}

// Поиск доступных голосований
function findApprovedAddressesInAll(address, callback) {
    var organisationsID = [];

    Votes.find({}, (err, vote) => {
        for (let i in vote) {
            for (let j in vote[i].approvedAddresses) {
                if (vote[i].approvedAddresses[j] == address) {
                    organisationsID.push(organisations[i].organisationID);
                }
            }
        }

        callback(organisationsID);
    });
}

// Создание нового пользователя
function addUser(userID, name, encryptedSeed, address) {
    Users.create({
        userID: userID,
        name: name,
        encryptedSeed: encryptedSeed,
        address: address
    }, (err, doc) => {

    });
}

// Поиск пользователя по userID
function findUserByUserID(userID, callback) {
    Users.find({
        userID: userID
    }, (err, doc) => {
        callback(doc[0]);
    });
}

// Поиск пользователя по адресу
function findUserByAddress(address, callback) {
    Users.find({
        address: address
    }, (err, doc) => {
        callback(doc[0]);
    })
}

// Поиск всех пользователей
function findAllUsers(address, callback) {
    Users.find({}, (err, doc) => {
        callback(doc);
    });
}


// Создание голосования
function createVote(organizationID, description) {
    Votes.create({
        organizationID: new ObjectId(organizationID),
        description: description
    }, (err, doc) => {

    });
}
// addUser('213142131', 'Kirill', 'wefrgfhgnhfrewfgfngdse', '23453rfewdwefdgrr')
// Принять участие в голосовании
function takePartInVote(voteID, description) {
    Votes.create({
        voteID: new ObjectId(voteID),
        address: address,
        vote: vote
    }, (err, doc) => {

    });
}

module.exports.addOrganisation = addOrganisation;
module.exports.findApprovedAddresses = findApprovedAddresses;
module.exports.findApprovedAddressesInAll = findApprovedAddressesInAll;
module.exports.addUser = addUser;
module.exports.findUserByUserID = findUserByUserID;
module.exports.findUserByAddress = findUserByAddress;
module.exports.findAllUsers = findAllUsers;
module.exports.createVote = createVote;
module.exports.takePartInVote = takePartInVote;