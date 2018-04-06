const WavesAPI = require('waves-api');

const Waves = WavesAPI.create(WavesAPI.TESTNET_CONFIG);

const newConfig = {
    networkByte: Waves.constants.MAINNET_BYTE,
    nodeAddress: 'https://nodes.wavesnodes.com',
    matcherAddress: 'https://nodes.wavesnodes.com/matcher',
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
    let restoredPhrase = Waves.Seed.decryptSeedPhrase(encryptedSeed, user_id);
    let seed = Waves.Seed.fromExistingPhrase(restoredPhrase);
    return seed;
}

// Получаем адрес из зашифрованного Seed
function getAddress(userID, encryptedSeed) {
    let seed = decryptSeed(userID, encryptedSeed);
    let address = seed.address;
    return address;
}



module.exports.returnedWaves = Waves;
module.exports.createSeed = createSeed;
module.exports.decryptSeed = decryptSeed;
module.exports.getAddress = getAddress;