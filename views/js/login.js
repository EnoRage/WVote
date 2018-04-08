
<script src="../../node_modules/waves-api/dist/waves-api.min.js"></script>

const HACKNET_CONFIG = {
    minimumSeedLength: 1,
    networkByte: 'U'.charCodeAt(0),
    nodeAddress: 'http://nodes.unblock.wavesnodes.com:6869',
    matcherAddress: 'http://nodes.unblock.wavesnodes.com/matcher',

}

function postLoginDataToServer() {
    var seed = $("#exampleInputPassword1").val()
    var password = $("#exampleInputPassword2").val()
    let seed1 = Waves.Seed.fromExistingPhrase(seed);
    const Waves = WavesAPI.create(HACKNET_CONFIG);

    const encrypted = seed1.encrypt(password);

    $.post(
        "/success_login", {
            "encrSeed": encrypted,
            "passwor": password
        },
        (res) => {
            
        }
    )
}
