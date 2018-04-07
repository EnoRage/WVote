const WavesAPI = require('waves-api'),
    SafeMath = require('./safeMath.js'),
    Objects = require('./objects.js');

const Waves = WavesAPI.create(WavesAPI.TESTNET_CONFIG);

const newConfig = {
    networkByte: Waves.constants.MAINNET_BYTE,
    nodeAddress: 'http://nodes.unblock.wavesnodes.com:6869',
    minimumSeedLength: 50
};

Waves.config.set(newConfig);

// Создаём Seed пользователю и шифруем его уникальным идентификатора телеграма
function createSeed(userID) {
    let seed = Waves.Seed.create();
    let encrypted = seed.encrypt(userID);
    return [encrypted, seed];
}

// Расшифровываем seed с помощью уникального идентификатора телеграма
function decryptSeed(userID, encryptedSeed) {
    let restoredPhrase = Waves.Seed.decryptSeedPhrase(encryptedSeed, userID);
    let seed = Waves.Seed.fromExistingPhrase(restoredPhrase);
    return seed;
}

// Получаем адрес из зашифрованного Seed
function getAddress(userID, _seed) {
    let seed = Waves.Seed.fromExistingPhrase(_seed);
    let address = seed.address;
    return address;
}

// Получаем баланс токенов или Waves
function getBalance(address, currency, callback) {
    Waves.API.Node.v1.assets.balance(address, Objects.currency[currency].assetID)
        .then(
            (balance) => {
                callback(SafeMath.powMinus8(balance.balance));
            })
        .catch(
            (err) => {
                callback(false);
            });
}

function sendTx(address, currency, amount, userID, encryptedSeed, callback) {
    let seed = decryptSeed(userID, encryptedSeed);

    const transferData = {
        recipient: address,
        assetId: Objects.currency['currency'].assetID,
        amount: SafeMath.pow8(amount),
        feeAssetId: 'WAVES',
        fee: 100000,
        attachment: '',
        timestamp: Date.now()
    };

    Waves.API.Node.v1.assets.transfer(transferData, seed.keyPair).then(
            (responseData) => {
                console.log(responseData);
                callback('200');
            })
        .catch(
            (err) => {
                console.log(err);
                callback('400');
            });
}


module.exports.returnedWaves = Waves;
module.exports.createSeed = createSeed;
module.exports.decryptSeed = decryptSeed;
module.exports.getAddress = getAddress;
module.exports.getBalance = getBalance;
module.exports.sendTx = sendTx;