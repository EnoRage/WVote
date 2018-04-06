const express = require('express'),
    url = require("url"),
    bodyParser = require('body-parser'),
    Waves = require('./waves.js'),
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
    
});

// Получаем баланс аккаунта
app.post('/getBalance', (req, res) => {
    let data = req.body;
    
});

app.listen(3000);