const express = require('express'),
    url = require("url"),
    bodyParser = require('body-parser'),
    Waves = require('./waves.js'),
    Course = require('./course.js'),
    db = require('./db.js');


const app = express();

app.use(bodyParser.urlencoded({
    extended: false
}))

// Создаём seed (отправляем зашифрованный seed)
app.post('/createSeed', (req, res) => {
    let data = req.body;
    let seed = Waves.createSeed(data.userID);
    let address = Waves.getAddress(data.userID, seed[1].phrase);
    db.addUser(data.userID, data.name, seed[0], address)
    res.send(seed[1].phrase);
});

// Получаем адрес из seed
app.post('/decryptSeed', (req, res) => {
    let data = req.body;
    let seed = Waves.decryptSeed(data.userID, data.encryptedSeed);
    res.send(seed.phrase);
});

// Получаем адрес из seed
app.post('/getAddress', (req, res) => {
    let data = req.body;
    let address = Waves.getAddress(data.userID, data.seed);
    res.send(address);
});

// Отправляем транзакцию в блокчейн
app.post('/sendTx', (req, res) => {
    let data = req.body;
    Waves.sendTx(data.address, data.currency, data.amount, data.userID, data.encryptedSeed, (transactionStatus) => {
        if (transactionStatus == '200') {
            res.send('200');
        } else {
            res.send('400');
        }
    })
});

// Получаем баланс токенов или Waves
app.post('/getBalance', (req, res) => {
    let data = req.body;
    console.log(data)
    let balance = Waves.getBalance(data.address, data.currency, (balance) => {
        // if (balance == false) {res.send('400');}
        res.send(String(balance));
    });
});

// Создать нового пользователя
app.post('/addUser', (req, res) => {
    let data = req.body;
    db.addUser(data.userID, data.name, data.encryptedSeed, data.address);
    res.send('200');
});

app.listen(3000);