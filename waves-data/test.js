const axlsign = require('./dep/axlsign').default;
const base58 = require('./dep/base58').default;
const sec = require('./dep/secure-random').default;
const fetch = require('node-fetch');

const DATA_ENTRY_TYPES = {
	INTEGER: 'integer',
	BOOLEAN: 'boolean',
	BINARY_ARRAY: 'binary'
};


function buildTransactionSignature(dataBytes, privateKey) {
	if (!dataBytes || !(dataBytes instanceof Uint8Array)) {
		throw new Error('Missing or invalid data');
	}
	
	if (!privateKey || typeof privateKey !== 'string') {
		throw new Error('Missing or invalid private key');
	}
	
	const privateKeyBytes = Uint8Array.from(base58.decode(privateKey));
	
	const signature = axlsign.sign(privateKeyBytes, dataBytes, sec.randomUint8Array(64));
	return base58.encode(signature);
	
}

const byteProcessors = {
	byte: function (value) {
		if (typeof value !== 'number') throw new Error('You should pass a number to Byte constructor');
		if (value < 0 || value > 255) throw new Error('Byte value must fit between 0 and 255');
		return [value];
	},
	bool: function (value) {
		if (typeof value !== 'boolean') {
			throw new Error('Boolean input is expected');
		}
		
		const bytes = value ? [1] : [0];
		return bytes
		
	},
	long: function (input) {
		if (typeof input !== 'number') {
			throw new Error('Numeric input is expected');
		}
		
		const bytes = new Array(7);
		for (let k = 7; k >= 0; k--) {
			bytes[k] = input & (255);
			input = input / 256;
		}
		
		return bytes;
		
	},
	short: function (input) {
		if (typeof input !== 'number') {
			throw new Error('Numeric input is expected');
		}
		
		const bytes = new Array(1);
		for (let k = 1; k >= 0; k--) {
			bytes[k] = input & (255);
			input = input / 256;
		}
		
		return bytes;
	},
	string: function (str) {
		str = unescape(encodeURIComponent(str));
		
		let bytes = new Array(str.length);
		for (let i = 0; i < str.length; ++i)
			bytes[i] = str.charCodeAt(i);
		
		return bytes;
		
	}
};

function getSigDataBytes(tx) {
	const getTypeByte = (type) => {
		switch (type) {
			case DATA_ENTRY_TYPES.INTEGER: {
				return 0;
			}
			case DATA_ENTRY_TYPES.BOOLEAN: {
				return 1;
			}
			case DATA_ENTRY_TYPES.BINARY_ARRAY: {
				return 2;
			}
		}
	};
	
	const getValueBytes = (type, value) => {
		switch (type) {
			case DATA_ENTRY_TYPES.INTEGER: {
				return byteProcessors.long(value);
			}
			case DATA_ENTRY_TYPES.BOOLEAN: {
				return byteProcessors.bool(!!value);
			}
			case DATA_ENTRY_TYPES.BINARY_ARRAY: {
				return byteProcessors.long(value);
			}
		}
	};
	
	
	let sigBytes = [12]; //tx type
	sigBytes = sigBytes.concat(byteProcessors.byte(1)); //version
	sigBytes = sigBytes.concat(base58.decode(tx.senderPublicKey)); //version
	sigBytes = sigBytes.concat(byteProcessors.short(tx.data.length));
	
	tx.data.forEach((object) => {
		let keyLength = byteProcessors.short(object.key.length);
		let key = byteProcessors.string(object.key);
		let bytes = keyLength.concat(key,
			getTypeByte(object.type),
			getValueBytes(object.type, object.value));
		sigBytes = sigBytes.concat(bytes);
	});
	
	sigBytes = sigBytes.concat(byteProcessors.long(tx.timestamp));
	sigBytes = sigBytes.concat(byteProcessors.long(tx.fee));
	
	return sigBytes;
	
}

function processTx(tx, keyPair) {
	const sigBytesArray = getSigDataBytes(tx);
	const sigData = Uint8Array.from(sigBytesArray);
	const signature = buildTransactionSignature(sigData, keyPair.privateKey);
	
	return {
		version: 1,
		...tx,
		type: 12,
		proofs: [signature],
		signature: signature
	}
}

//SEE BELOW



function sendDataToWavesBlockchain(_seed, voteNum,voteValue) {
	var seed = _seed;
	const dataTx = {
		
		// An arbitrary address; mine, in this example
		sender: seed.address,
		senderPublicKey: seed.keyPair.publicKey,
		data: [
			{ key: 'voteID', value: Number(voteNum), type: DATA_ENTRY_TYPES.INTEGER },
			{ key: 'vote', value: voteValue, type: DATA_ENTRY_TYPES.BOOLEAN }
		],
		fee: 100000,
		timestamp: Date.now()
	};
	
	
	var options = {
		headers: {
			Accept: 'application/json',
			'Content-Type': 'application/json;charset=UTF-8'
		},
		method: 'POST',
		body: JSON.stringify(processTx(dataTx, {
			privateKey: seed.keyPair.privateKey,
			publicKey: seed.keyPair.publicKey
		}))
	};

	fetch('http://nodes.unblock.wavesnodes.com:6869/transactions/broadcast', options)
	.then(res => res.text())
	.then(body => console.log(body));
}

module.exports.sendDataToWavesBlockchain = sendDataToWavesBlockchain;
