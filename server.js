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
    let encryptedSeed = Waves.createSeed(data.userID);
    res.send(encryptedSeed[0]);
});

// Получаем адрес из seed
app.post('/getAddress', (req, res) => {
    let data = req.body;
    let address = Waves.getAddress(data.userID, data.encryptedSeed);
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
    let balance = Waves.getBalance(data.address, data.currency, (balance) => {
        if (balance == false) res.send('400');
        res.send(balance);
    });
});

app.listen(3000);