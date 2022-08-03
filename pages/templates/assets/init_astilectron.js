const tmpAstilectron = { callbacks: {}, counters: {}, onMessageOnce() {}, onMessage() {}, registerCallbacks() {}, sendMessage(message, cb) {} }
astilectron = astilectron ?? tmpAstilectron

/**
 * @param {string} json.channel
 * @param {Record<string, string>|Record<string, string>[]} json.data
 */
function sendMessage(json) {
    astilectron.sendMessage(JSON.stringify(json))
}