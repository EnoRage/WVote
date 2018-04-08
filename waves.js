const WavesAPI = require('waves-api'),
    WavesData = require('./waves-data/test.js'),
    bs58 = require('bs58'),
    // WavesUtils = require('./node_modules/waves-api/src/utils/request.ts'),
    // WavesTx = require("./node_modules/  waves-api/src/classes/Transactions.ts"),
    SafeMath = require('./safeMath.js'),
    Objects = require('./objects.js');

const HACKNET_CONFIG = {
    minimumSeedLength: 1,
    networkByte: 'U'.charCodeAt(0),
    nodeAddress: 'http://nodes.unblock.wavesnodes.com:6869',
    matcherAddress: 'http://nodes.unblock.wavesnodes.com/matcher',

}

const Waves = WavesAPI.create(HACKNET_CONFIG);
// Waves.API.Node.v1.transactions.
// const newConfig = {
//     // networkByte: Waves.constants.MAINNET_BYTE,
//     minimumSeedLength: 50
// };

// Waves.config.set(newConfig);

// Создаём Seed пользователю и шифруем его уникальным идентификатора телеграма
function createSeed(userID) {
    let seed = Waves.Seed.create();
    let encrypted = seed.encrypt(userID);
    return [encrypted, seed];
}

// var seed = createSeed('ffer4tgr');
// console.log(seed[1])

// Расшифровываем seed с помощью уникального идентификатора телеграма
function decryptSeed(userID, encryptedSeed) {
    let restoredPhrase = Waves.Seed.decryptSeedPhrase(encryptedSeed, userID);
    let seed = Waves.Seed.fromExistingPhrase(restoredPhrase);
    return seed;
}

// Получаем адрес из Seed
function getAddress(userID, _seed) {
    let seed = Waves.Seed.fromExistingPhrase(_seed);
    let address = seed.address;
    return address;
}

// Получаем баланс токенов или Waves
function getBalance(address, currency, callback) {
    Waves.API.Node.v1.assets.balance(String(address), String(Objects.currency[currency].assetID))
        .then(
            (balance) => {
                callback(SafeMath.powMinus8(balance.balance));
            })
        .catch(
            (err) => {
                console.log(err)
                callback(false);
            });
}

// getBalance("3NLSgTvf71NeUAuVtbTrBf8GPr52Kbup7W2", "Waves", (balance) => {
//     console.log(balance)
// })


function sendDataTx(userID, encryptedSeed, voteNum, vote) {
    const seed = decryptSeed(userID, encryptedSeed)
    var _vote;
    if (vote == "0") {
        _vote = false;
    } else {
        _vote = true;
    }
    WavesData.sendDataToWavesBlockchain(seed, Number(voteNum), _vote);
}

function sendTx(address, currency, amount, userID, encryptedSeed, validatorEncrSeed, _attachment, callback) {
    let seed = decryptSeed(userID, encryptedSeed);

    const transferData = {
        recipient: address,
        assetId: Objects.currency[currency].assetID,
        amount: SafeMath.pow8(amount),
        feeAssetId: 'WAVES',
        fee: 100000,
        attachment: Number(_attachment),
        timestamp: Date.now()
    };

    console.log(transferData)

    Waves.API.Node.v1.assets.transfer(transferData, seed.keyPair).then(
            (responseData) => {
                // console.log(responseData);
                var attachment = responseData.attachment;
                var decodeAttechment = bs58.decode(attachment).toString();

                var vote = decodeAttechment.substr(1, 1);
                var voteNum = decodeAttechment.substring(2);
                sendDataTx('validator',validatorEncrSeed, voteNum, vote)
                callback('200');
            })
        .catch(
            (err) => {
                console.log(err);
                callback('400');
            });
}

function sendAttachmentToValidator (userID, encryptedSeed, _vote, voteNum, validatorEncrSeed, validatorAddress) {

    var attachment = "1"+_vote + voteNum
    console.log(attachment)
    sendTx(validatorAddress, 'Waves', 0.005, userID, encryptedSeed, validatorEncrSeed, attachment, () => {

    })
}

const to_b58 = function (B, A) {
    var d = [], s = "", i, j, c, n;
    for (i in B) {
        j = 0, c = B[i];
        s += c || s.length ^ i ? "" : 1;
        while (j in d || c) {
            n = d[j];
            n = n ? n * 256 + c : c;
            c = n / 58 | 0;
            d[j] = n % 58;
            j++
        }
    }
    while (j--) s += A[d[j]];
    return s
};


// sendDataTx()

module.exports.returnedWaves = Waves;
module.exports.createSeed = createSeed;
module.exports.decryptSeed = decryptSeed;
module.exports.getAddress = getAddress;
module.exports.getBalance = getBalance;
module.exports.sendTx = sendTx;
module.exports.sendDataTx = sendDataTx;
module.exports.sendAttachmentToValidator = sendAttachmentToValidator;