const express = require('express'),
    url = require("url"),
    hbs = require("hbs"),
    bodyParser = require('body-parser'),
    Cookies = require('cookies').express,
    Waves = require('./waves.js'),
    Course = require('./course.js'),
    db = require('./db.js');


const app = express();

// Чтоб считывать статические файлы
app.use(express.static(__dirname + '/views'));
app.use(Cookies('секрет_хех'));
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

// Для выхода. 
app.get('/logout', (req, res) => {
    res.clearCookie('session');
    res.redirect('/login');
});

// Вход в личный кабинет
app.post('/success_login', (req, res) => {
    if (Object.keys(req.body).length == 2) {
        // Распарсенные данные
        var data = req.body;
        // рандомная строка для сессии
        const string = db.randomString();
        // Проводим аутентификацию и присваиваем строку сессии пользователю
        db.authentication(data.encrSeed, data.password, string, (isTrue) => {
            if (isTrue == true) {
                // Добавляем захешированную строку в куки
                res.cookies.set('session', md5(string).toString());
                res.redirect('/');
            } else {
                // Неправильно ввёл или такого юзера не существует
                res.redirect('/login');
            }
        });
    } else {
        res.redirect('/login');
    }
});

// Регистрация пользователя в личном кабинете
app.post('/success_registration', (req, res, next) => {
    if (Object.keys(req.body).length == 4) {
        var data = req.body;
        db.addUser(data.password, "No Name", data.encrSeed, data.address);
        res.status(200);
        // Тут баг небольшой. Надо перенаправлять на логин после нажатия кнопки ok
        res.redirect('/login');
    } else {
        res.redirect('/login');
    }
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

// Создать голосование
app.post('/createVote', (req, res) => {
    let data = req.body;
    db.createVote(data.userID, data.description, data.endTime);
    res.send('200');
});

// Проголосовать
app.post('/vote', (req, res) => {
    let data = req.body;
    console.log(data)
    // Тут в блокчейн заносится
    db.takePartInVote(Number(data.voteNum), data.address, data.vote);
    res.send('200');
});

// Подсчитать голоса
app.post('/totalVotes', (req, res) => {
    // единица - за, нуль - против
    let data = req.body;
    if (data.whatVote == '0') {
        db.voteNo(data.num, (count) => {
            res.send(count);
        });
    } else {
        db.voteYes(data.num, (count) => {
            res.send(count);
        });
    }
});

// Найти все голосования
app.post('/findAllVotes', (req, res) => {
    let data = req.body;
    db.findAllVotes((votesArray) => {
        res.send(votesArray);
    });
});

// Найти всех голосующих
app.post('/findAllVotes', (req, res) => {
    let data = req.body;
    db.findAllVoters((votesArray) => {
        res.send(votesArray);
    });
});

hbs.registerPartials(__dirname + '/views/templates', () => {
    app.set("view engine", "hbs");


    var pathname = ''; // Переменная для хранения текущего url

    // Обработчик запросов
    app.use((req, res) => {
        // Определяем текущий url
        pathname = url.parse(req.url).pathname;

        db.findUserBySessionKey(res.cookies.get('session'), (user) => {
            db.findFileByURL(pathname, (page) => {
                if (page != false) {
                    if (user != false) {
                        if (pathname == '/') {
                            var arr = '';
                            db.findAllVotes((votes) => {

                                for (let i in votes) {
                                    console.log(votes[i].num)
                                    db.findVotersByNum(votes[i].num, (voters) => {
                                        console.log(voters)
                                        var votersB = "";
                                        for (let b in voters) {
                                            var isTrue;
                                            if (voters[b].vote == true) {
                                                isTrue = 'За';
                                            } else {
                                                isTrue = 'Против';
                                            }
                                            votersB += `<tr><td>${voters[b].address}</td><td>${isTrue}</td></tr>`
                                        }

                                        arr += `<div class="col-lg-12 col-md-12"><div class="card"><div class="card-header card-header-tabs card-header-primary"><div class="nav-tabs-navigation"><div class="nav-tabs-wrapper"><span class="nav-tabs-title"> <p class="crd-p" style="font-weight: 500">Голосование ${votes[i].num}: ${votes[i].description}</p> </span><ul class="nav nav-tabs justify-content-end" data-tabs="tabs"><li class="nav-item"><a class="nav-link " href="#"><i class="material-icons">update</i></a> </li> </ul> </div> </div> </div><div class = "card-body table-responsive" ><table class = "table table-hover" thead class = "text-primary" ><th> Address </th> <th> Тип </th> </thead>  <tbody> ${votersB} </tbody> </table> </div> <div class = "card-footer" ><div class = "stats" ><i class = "material-icons"> more </i> <a href = "./votingplus.html"> Подробнее </a> </div> </div> </div> </div>`

                                        if (votes.length - 1 == i) {
                                            console.log(arr)
                                            res.render(page.file, {
                                                pages: arr

                                            });
                                        }

                                    })

                                }

                            })
                        }
                    } else if (pathname == '/register') {
                        res.render(page.file);
                    } else {
                        res.render('login.hbs');
                    }
                } else {
                    res.status(404).render('404.hbs'); // Иначе, отправить ошибку 404
                }
            })
        })
    })
})


app.listen(3001);