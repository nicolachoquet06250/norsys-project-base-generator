const tmpAstilectron = { callbacks: {}, counters: {}, onMessageOnce() {}, onMessage() {}, registerCallbacks() {}, sendMessage(message, cb) {} }
astilectron = astilectron ?? tmpAstilectron

/**
 * @param {string} json.channel
 * @param {Record<string, string>|Record<string, string>[]|undefined} json.data
 */
function sendMessage(json) {
    if (json.data === undefined) {
        json.data = {}
    }
    astilectron.sendMessage(json)
}