const rp = require('request-promise');

// Получение курса по 7 валютам
function getCourse(rubOrUsd, callback) {
        var options = {
            uri: 'https://min-api.cryptocompare.com/data/price?fsym=' + rubOrUsd + '&tsyms=WAVES,BTC,ETH,ZEC,LTC,USD,EUR',
            json: true
        };

       rp(options)
        .then(
            (courses) => {
                callback(courses);
            }
        )
        .catch(
            (err) => {
                callback('400');
            }
        )
}

module.exports.getCourse = getCourse;
