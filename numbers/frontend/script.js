let gameId = localStorage.getItem('gameId');
if (gameId == null) {
    gameId = randomIdString();
    localStorage.setItem('gameId', gameId);
}
let gameState = {}
let gameIdL = document.getElementById('gameId');
let resultScreenL = document.getElementById('resultScreen');
let newGameB = document.getElementById('newGameButton');
let playGameB = document.getElementById('playButton');

newGameB.addEventListener('click', async function () {
    gameId = randomIdString();
    localStorage.setItem('gameId', gameId);
    gameIdL.innerHTML = gameId;
    response = await fetch('/api/game/' + gameId)
    gameState = await response.json()
    resultScreenL.innerHTML = "Your total is: " + gameState.total + " last round was: " + gameState.lastIncrement;
})

playGameB.addEventListener('click', async function () {
    response = await fetch('/api/game/' + gameId, {
        method: "POST",
    })
    gameState = await response.json()
    resultScreenL.innerHTML = "Your total is: " + gameState.total + " last round was: " + gameState.lastIncrement;
})

function randomIdString() {
    return Math.random().toString(36).substring(2, 5);
}

(async function () {
    gameIdL.innerHTML = gameId;
    response = await fetch('/api/game/' + gameId)
    gameState = await response.json()
    resultScreenL.innerHTML = "Your total is: " + gameState.total + " last round was: " + gameState.lastIncrement;
})()